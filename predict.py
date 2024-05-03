import sys
import time
import mysql.connector
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.preprocessing import OneHotEncoder
from sklearn.model_selection import train_test_split
from sklearn.neighbors import KNeighborsRegressor
from sklearn.neighbors import KNeighborsClassifier
from sklearn.multioutput import MultiOutputRegressor
from sklearn.compose import ColumnTransformer
from sklearn.preprocessing import StandardScaler
from sklearn.impute import SimpleImputer
from sklearn.metrics import r2_score
from sklearn.neural_network import MLPRegressor
from sklearn.linear_model import LinearRegression
from sklearn.svm import SVR
from sklearn.linear_model import Ridge
from sklearn.ensemble import RandomForestRegressor
from sklearn.decomposition import PCA
from sklearn.feature_selection import r_regression
# python3 predict.py <user> <password> <host> <port> <database> <prediction_id>
# python3 predict.py root root mysql 3306 inu_backyard 01HW8F6EAJ1JM46DH98QTMKN15

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
  # print(r_regression(admission, y_gpax.reshape(-1, 1)))
  X_gpax_train, X_gpax_test, y_gpax_train, y_gpax_test = train_test_split(Xt_gpax, y_gpax, test_size=0.25, random_state=0)
  X_gpax_pca = pca.fit_transform(X_gpax_train)
  # print(y_gpax[0][0])
  # print(type(y_gpax[0][0]))
  # print(y_gpax[:,0])
  # print(X_gpax_train[0])
  # print(y_gpax_train[0])
  # print (pca.explained_variance_ratio_.cumsum())
  yscaler = StandardScaler().fit(y_gpax_train[:,-1].reshape(-1, 1))

  y_gpax_train = yscaler.transform(y_gpax_train[:,-1].reshape(-1, 1))

  # plt.figure(figsize=(4,4))
  # plt.scatter(X_gpax_train[:,-1], y_gpax_train[:,0])
  # plt.xlabel("old_gpax")
  # plt.ylabel("gpax")
  # plt.show()

  # modelregr = MLPRegressor(hidden_layer_sizes=(5,15,5), max_iter=5000)
  # modelregr = RandomForestRegressor(criterion="squared_error", max_depth=5, n_estimators=1000)
  # modelregr = SVR()
  modelregr = LinearRegression()
  # modelregr = KNeighborsRegressor(n_neighbors=30)

  modelregr.fit(X_gpax_pca[:,:42], y_gpax_train)
  y_gpax_predict = modelregr.predict(pca.transform(X_gpax_test)[:,:42]).reshape(-1, 1)
  y_gpax_predict_iscaled = yscaler.inverse_transform(y_gpax_predict.reshape(-1, 1))


  # test = pd.DataFrame([['regular', 3.99, 4,3.96, 4, 'โรงเรียนจำลอง', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 1, 1,1, 1, 'หมีน้อย', 'เรียนดี']]).to_numpy()
  test = pd.DataFrame([['regular', 1, 1,1, 1, 'เตรียมอุดมศึกษาน้อมเกล้า', 'เรียนดี']]).to_numpy()
  test = pd.DataFrame([['regular', 4, 4, 4, 4, 'เตรียมอุดมศึกษาน้อมเกล้า', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 3.99, 4,3.96, 4, 'อิสลามวิทยาลัยแห่งประเทศไทย', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 1, 1,1, 1, 'อิสลามวิทยาลัยแห่งประเทศไทย', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 4, 4, 4, 4, 'อิสลามวิทยาลัยแห่งประเทศไทย', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 3.8, 3.45,4, 3.7, 'มหิดลวิทยานุสรณ์', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 1, 1,1, 1, 'มหิดลวิทยานุสรณ์', 'เรียนดี']]).to_numpy()
  # test = pd.DataFrame([['regular', 4, 4, 4, 4, 'มหิดลวิทยานุสรณ์', 'เรียนดี']]).to_numpy()
  test = ctx.transform(test)
  test = pca.transform(test)
  prediction = modelregr.predict(test[:,:42])
  print(yscaler.inverse_transform(prediction.reshape(-1, 1)))

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







  sys.exit(0)
