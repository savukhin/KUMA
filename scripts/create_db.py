import psycopg2
import argparse

parser = argparse.ArgumentParser()

#-db DATABASE -u USERNAME -p PASSWORD -size 20000
parser.add_argument("-host", "--hostname", dest="hostname", default="127.0.0.1", help="Server name")
parser.add_argument("-P", "--port", dest="port", default="5432", help="Server port")
parser.add_argument("-db", "--database", dest="db", default="kuma", help="Database name")
parser.add_argument("-u", "--username", dest="username", default="20624880", help="User name")
parser.add_argument("-p", "--password", dest="password", default="admin", help="Password")

parser.add_argument("-d", "--drop", dest="drop", default=False, help="Need to drop database", action=argparse.BooleanOptionalAction)

args = parser.parse_args()

dbname = args.db
username = args.username
password = args.password
hostname = args.hostname
port = args.port
drop = args.drop

with open("./sql/cnc_checker.sql") as f:
    cnc_checker_script = f.read()
with open("./sql/cnc_status_type.sql") as f:
    cnc_status_script = f.read()
with open("./sql/employee.sql") as f:
    employee_script = f.read()

conn = psycopg2.connect(dbname=dbname, user=username, password=password, port=port, host=hostname)
conn.autocommit = True
cur = conn.cursor()


if drop:
    print("Dropping")
    
    cur.execute("DROP TABLE IF EXISTS cnc_checkers")
    cur.execute("DROP TABLE IF EXISTS employees")
    
    cur.execute("DROP TYPE IF EXISTS cnc_status")
    
    print("Dropped")

cur.execute(cnc_status_script)
cur.execute(cnc_checker_script)
cur.execute(employee_script)

conn.commit()

cur.close()
conn.close()

print("Done")
