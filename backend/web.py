# 本地网络接口
# http://localhost:8110
from flask import Flask,render_template
import os
from lichv.postgresqldb import PostgresqlDBService
import decimal

db = PostgresqlDBService.instance(host='localhost', port=5432, user='postgres', passwd='123456', db='data')

def list2tree(items,steps,key='code',children='children'):
	mapdata = {}
	length = sum(steps)
	count = len(steps)
	for item in items:
		temp = item
		if len(item[key]) < length:
			temp[children] = []
		mapdata[item[key]] = item

	for i in range(1,count):
		current = sum(steps[0:count-i+1])
		parent = sum(steps[0:count-i])

		for item in items:
			if len(item[key])== current:
				mapdata[item[key][:parent]][children].append(item)
	items = {}
	for item in items:
		if len(item[key]) == steps[0]:
			items[item[key]] = mapdata[item[key]]
	return items;
		

			
	
app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def index():
	return {"state":2000,"msg":"success"}

@app.route('/api/area', methods=['POST', 'GET'])
def area():
	items = {}
	areas = set()
	lu = db.getList('zone',{})
	items = list2tree(lu,(2,2,2))
	return {"state":2000,"data":items}

@app.route('/api/<path:table>', methods=['POST', 'GET'])
def route(table='user'):
	items = []
	data = db.getList(table,{},'*',[('id',1)])
	for item in data:
		for ke in item:
			if isinstance(item[ke], decimal.Decimal):
				item[ke] = float(item[ke])
		items.append(item)
	return {"state":2000,"data":items}

if __name__ == '__main__':
    app.run(host="0.0.0.0",port=8110)