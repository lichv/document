# 本地网络接口
# http://localhost:8110
import os
import decimal
import json
from flask import Flask,render_template,request
from lichv.utils import * 
from lichv.postgresqldb import PostgresqlDBService

from aliyunsdkcore.client import AcsClient
from aliyunsdkcore.acs_exception.exceptions import ClientException
from aliyunsdkcore.acs_exception.exceptions import ServerException
from aliyunsdksts.request.v20150401.AssumeRoleRequest import AssumeRoleRequest


db = PostgresqlDBService.instance(host='localhost', port=5432, user='postgres', passwd='123456', db='data')

app = Flask(__name__,template_folder=os.getcwd()+"../../../frontend/dist/",static_folder=os.getcwd()+"../../../frontend/dist/_assets")

def getToken():
	client = AcsClient('LTAI5tKnra54xozuA3KktFur', 'VyIHrtVQZxXeiuuBWUW2oG34qe87dk', 'cn-shanghai')
	request = AssumeRoleRequest()
	request.set_accept_format('json')
	request.set_RoleArn("acs:ram::1378573870105843:role/osser")
	request.set_RoleSessionName("osser")
	response = client.do_action_with_exception(request)
	result = json.loads(str(response, encoding='utf-8'))
	return result

def getFiles(source='.', extends='.md'):
	source = source.replace('/','\\')
	if not source.endswith('\\'):
		source = source + '\\'
	items = []
	if os.path.isdir(source):
		for parent, dirnames, filenames in os.walk(source, topdown=False,	followlinks=True):
			for filename in filenames:
				if filename.lower().endswith(extends) :
					filepath = os.path.join( parent,filename )
					filename = filepath[len(source):]
					filename = filename.replace('\\','/')
					items.append({'filename':filename})
	return items

@app.route('/', methods=['POST', 'GET'])
def index():
	return render_template('index.html')
	# return {"state":2000,"msg":"success"}

@app.route('/api/aliyun/sts', methods=['POST', 'GET'])
def getSTSToken():
	result = getToken()
	return {"state":2000,"data":result['Credentials']}

@app.route('/api/markdown/files', methods=['POST', 'GET'])
def getMarkdownFiles():
	files = getFiles('./docs')
	return {"state":2000,"msg":"success","data":files}

@app.route('/api/markdown/read', methods=['POST', 'GET'])
def getMarkdownDetail():
	text = ''
	if request.method == "GET":
		path = request.args.get("path")
	if request.method == "POST":
		if request.content_type.startswith('application/json'): 
			path = request.json.get('path')
		elif request.content_type.startswith('multipart/form-data'):
			path = request.form.get('path')
		else:
			path = request.values.get("path")
	print("第一次："+path)
	path = ".\\docs\\"+path
	print("第二次："+path)
	if os.path.isdir(path):
		path = os.path.join( path,'index.md' )
	if os.path.exists( path ):
		text = read(path)
	else:
		print('path不存在：'+path)
	print(text)
	return {"state":2000,"msg":"success","data":text}

if __name__ == '__main__':
    app.run(host="0.0.0.0",port=8110)
