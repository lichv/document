#!/usr/bin/env python
#coding=utf-8

from aliyunsdkcore.client import AcsClient
from aliyunsdkcore.acs_exception.exceptions import ClientException
from aliyunsdkcore.acs_exception.exceptions import ServerException
from aliyunsdksts.request.v20150401.AssumeRoleRequest import AssumeRoleRequest

#构建一个阿里云客户端，用于发起请求。
#构建阿里云客户端时需要设置AccessKey ID和AccessKey Secret。
client = AcsClient('LTAI5tKnra54xozuA3KktFur', 'VyIHrtVQZxXeiuuBWUW2oG34qe87dk', 'cn-shanghai')

#构建请求。
request = AssumeRoleRequest()
request.set_accept_format('json')

#设置参数。
request.set_RoleArn("acs:ram::1378573870105843:role/osser")
request.set_RoleSessionName("osser")

#发起请求，并得到响应。
response = client.do_action_with_exception(request)
# python2:  print(response)
print(str(response, encoding='utf-8'))

