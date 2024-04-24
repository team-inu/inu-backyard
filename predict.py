import sys
import time
import mysql.connector

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

  db_cursor = mysql_connection.cursor()
  query = "UPDATE prediction SET status='DONE', result='some equation' WHERE id = %s"
  args = (sys.argv[6],)
  db_cursor.execute(query, args)
  mysql_connection.commit()

  sys.exit(0)
