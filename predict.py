import sys
import time
import mysql.connector
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.preprocessing import OneHotEncoder
from sklearn.model_selection import train_test_split
from sklearn.compose import ColumnTransformer
from sklearn.preprocessing import StandardScaler
from sklearn.metrics import r2_score
from sklearn.linear_model import LinearRegression
from sklearn.decomposition import PCA
from sklearn.feature_selection import r_regression
from scipy.stats import pearsonr
import statsmodels.api as sm
# python3 predict.py <user> <password> <host> <port> <database> <programme_name> <old_gpax> <math_gpa> <eng_gpa> <sci_gpa> <school> <admission>
# python3 predict.py root root mysql 3306 inu_backyard

# TODO: add visibility

if __name__ == '__main__':
  mysql_connection = mysql.connector.connect(
    user=sys.argv[1],
    password=sys.argv[2],
    host=sys.argv[3],
    port=sys.argv[4],
    database=sys.argv[5]
  )

  np.set_printoptions(suppress=True)

  # time.sleep(10)

  db_cursor = mysql_connection.cursor()
  db_cursor.execute("WITH grades AS (SELECT student_id, ROUND(AVG(grade), 2) as gpax FROM `grade` GROUP BY student_id) SELECT programme_name, student.gpax AS old_gpax, math_gpa, eng_gpa, sci_gpa, school, admission, remark, grades.gpax FROM student INNER JOIN grades ON grades.student_id = student.id WHERE student.gpax > 0.1 AND student.math_gpa != 0  AND student.eng_gpa != 0 AND student.sci_gpa != 0;")
  features_name = np.array(["programme", "old_gpax", "math_gpa", "eng_gpa", "sci_gpa", "school, admission"])
  labels_name = np.array(["remark", "gpax"])
  results = db_cursor.fetchall()

  data = np.asarray(results)
  programme = np.asarray(data[:,0])
  old_gpax = np.asarray(data[:,1], dtype=float)
  math_gpa = np.asarray(data[:,2], dtype=float)
  eng_gpa = np.asarray(data[:,3], dtype=float)
  sci_gpa = np.asarray(data[:,4], dtype=float)
  school = np.asarray(data[:,5])
  admission = np.asarray(data[:,6])
  remark = np.asarray(data[:,7])
  gpax = np.asarray(data[:,8], dtype=float)
  data = pd.DataFrame({'programme': programme, 'old_gpax': old_gpax, 'math_gpa': math_gpa, 'eng_gpa': eng_gpa, 'sci_gpa': sci_gpa, 'school': school, 'admission': admission, 'remark': remark, 'gpax': gpax}).to_numpy()

  X = data[:,:7]
  ctx = ColumnTransformer([('school', OneHotEncoder(sparse_output=False,handle_unknown='ignore'), [5]), ('admission', OneHotEncoder(sparse_output=False,handle_unknown='ignore'), [6]), ('programme', OneHotEncoder(handle_unknown='ignore'), [0])], remainder=StandardScaler(), sparse_threshold=0)


  ## Predict GPAX from admission information

  pca = PCA()
  y_gpax = np.array(data[:,[8]], dtype=float)
  Xt_gpax = ctx.fit_transform(X)

  included_list = []
  for i in range(Xt_gpax.shape[1]):
    corr, p_value = pearsonr(Xt_gpax[:,i],y_gpax[:,0])

    # change this value to adjust the significance accepted
    if p_value < 0.95:
      included_list.append(True)
      # print(ctx.get_feature_names_out()[i], corr, p_value)
    else:
      included_list.append(False)

  sig_array = np.asarray(included_list)

  X_gpax_train, X_gpax_test, y_gpax_train, y_gpax_test = train_test_split(Xt_gpax[:, sig_array], y_gpax, test_size=0.25, random_state=0)
  X_gpax_pca = pca.fit_transform(X_gpax_train)

  cumsum = pca.explained_variance_ratio_.cumsum() > 0.9
  pca_index = 0
  for i in range(len(cumsum)):
    if cumsum[i]:
      pca_index = i
      break

  yscaler = StandardScaler().fit(y_gpax_train[:,-1].reshape(-1, 1))
  y_gpax_train = yscaler.transform(y_gpax_train[:,-1].reshape(-1, 1))

  # # plt.figure(figsize=(4,4))
  # # plt.scatter(X_gpax_train[:,-1], y_gpax_train[:,0])
  # # plt.xlabel("old_gpax")
  # # plt.ylabel("gpax")
  # # plt.show()

  modelregr = LinearRegression()

  modelregr.fit(X_gpax_pca[:,:pca_index], y_gpax_train)
  # print(modelregr.score(pca.transform(X_gpax_test)[:,:pca_index], yscaler.transform(y_gpax_test)))
  y_gpax_predict = modelregr.predict(pca.transform(X_gpax_test)[:,:pca_index]).reshape(-1, 1)
  y_gpax_predict_iscaled = yscaler.inverse_transform(y_gpax_predict.reshape(-1, 1))

  target = pd.DataFrame([[sys.argv[6], sys.argv[7], sys.argv[8], sys.argv[9], sys.argv[10], sys.argv[11], sys.argv[12]]]).to_numpy()
  target = ctx.transform(target)
  target = target[:, sig_array]
  # print(target)
  target = pca.transform(target)
  prediction = modelregr.predict(target[:,:pca_index])
  print(round(yscaler.inverse_transform(prediction.reshape(-1, 1))[0,0], 2))

  sys.exit(0)

  # print(r2_score(yscaler.transform(y_gpax_test[:,-1].reshape(-1, 1)), y_gpax_predict))

  ## Predict remark from admission and current GPAX

  # y_remark = data[:,[6]]
  # ohe = OneHotEncoder(handle_unknown='ignore').fit(y_remark.reshape(-1, 1))
  # yt_remark = ohe.transform(y_remark.reshape(-1, 1)).toarray()
  # Xt_remark = np.append(Xt_gpax, np.round(4 * np.random.random_sample((data.shape[0], 1)), 2), axis=1)
  # X_remark_train, X_remark_test, y_remark_train, y_remark_test = train_test_split(Xt_remark, yt_remark, test_size=0.25, random_state=0)

  # modelcls = KNeighborsClassifier(n_neighbors = 5).fit(X_remark_train, y_remark_train)
  # y_remark_predict = modelcls.predict(X_remark_test)

  # print(y_remark_predict)



  # test = ohe.transform(X[:,5].reshape(1, -1))
  # print(X.shape)
  # print([y=='ตกออก'])
  # print(test)
  # print(test[y=='ตกออก'])
  # query = "UPDATE prediction SET status='DONE', result='some equation' WHERE id = %s"
  # args = (sys.argv[6],)
  # db_cursor.execute(query, args)
  # mysql_connection.commit()
  # print(test)
  # ctx = ColumnTransformer([('school', OneHotEncoder(handle_unknown='ignore'), [5]), ('programme', OneHotEncoder(handle_unknown='ignore'), [0])], remainder="passthrough")
  # ctx.sparse_output_ = True

  # cty = ColumnTransformer([('remark', OneHotEncoder(handle_unknown='ignore'), [0])], remainder="passthrough")
  # yt = cty.fit_transform(y)

  # print(ctx.get_feature_names_out())
  # print(Xt[0])

  # print(cty.get_feature_names_out())
  # print(yt[0])
  # print(cty.get_feature_names_out())
  # print(y_train[:, -1])
