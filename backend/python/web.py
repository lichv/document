# 本地网络接口
# http://localhost:8110
import os
import decimal
import json
from flask import Flask,render_template
from lichv.postgresqldb import PostgresqlDBService

from aliyunsdkcore.client import AcsClient
from aliyunsdkcore.acs_exception.exceptions import ClientException
from aliyunsdkcore.acs_exception.exceptions import ServerException
from aliyunsdksts.request.v20150401.AssumeRoleRequest import AssumeRoleRequest


db = PostgresqlDBService.instance(host='localhost', port=5432, user='postgres', passwd='123456', db='data')

app = Flask(__name__)

def getToken():
	client = AcsClient('LTAI5tKnra54xozuA3KktFur', 'VyIHrtVQZxXeiuuBWUW2oG34qe87dk', 'cn-shanghai')
	request = AssumeRoleRequest()
	request.set_accept_format('json')
	request.set_RoleArn("acs:ram::1378573870105843:role/osser")
	request.set_RoleSessionName("osser")
	response = client.do_action_with_exception(request)
	result = json.loads(str(response, encoding='utf-8'))
	return result

@app.route('/', methods=['POST', 'GET'])
def index():
	return {"state":2000,"msg":"success"}

@app.route('/api/aliyun/sts', methods=['POST', 'GET'])
def getSTSToken():
	result = getToken()
	return {"state":2000,"data":result}

if __name__ == '__main__':
    app.run(host="0.0.0.0",port=8110)