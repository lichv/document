from lichv.utils import * 
from copy import deepcopy

import sys  # 导入sys模块
sys.setrecursionlimit(3000)  # 将默认的递归深度修改为3000


dirname = './docs/'


filename = './docs/内含子.md'

# name = read(filename)


# print(name)

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
					path = filename.split('/')
					items.append({'filename':filename,'path':path})
	return items

def makeParent(path,dirname=''):
	if not path or not isinstance(path,list):
		return {}
	length = len(path)
	if length == 1:
		# return {"name":'root',"path":'',"children":[{"name":path[-1],"path":dirname}]}
		return {"name":path[-1],"parent":dirname,"path":os.path.join(dirname,path[-1])}
	elif length > 1:
		# return {"name":'root',"path":'',"children":[{"name":path[-1*length],"path":dirname,"children":[makeParent(path[1:],os.path.join(dirname,path[-1*length]))]}]}
		return {"name":path[-1*length],"parent":dirname,"path":os.path.join(dirname,path[-1*length]),"children":[makeParent(path[1:],os.path.join(dirname,path[-1*length]))]}
	else:
		return {}

def getFilePath(path,dirname=''):
	filepath = makeParent(path,dirname)
	if isinstance(filepath,list):
		return {'name':"root","parent":"","path":"","children":filepath}
	else:
		return {'name':"root","parent":"","path":"","children":[filepath]}

def merge(path1,path2,dirname=''):
	if not path1:
		return path2
	if not path2:
		return path1
	temp = path1
	if 'children' in path1 and 'children' in path2:
		lis = []
		for a in path1['children']:
			lis.append(a)
		for b in path2['children']:
			lis.append(b)
		newitem = None
		print('****************')
		print(lis)
		print('****************')
		for li in lis:
			newitem = merge(newitem,li)
		if isinstance(newitem,list):
			temp['children'] = newitem
		else:
			temp['children'] = [newitem]
		return temp
	else:
		if 'children' in path1:
			temp['children'] = path1['children']
			return temp
		if 'children' in path2:
			temp['children'] = path2['children']
		return temp


# dirname = os.getcwd()+"/docs"
# items = getFiles(dirname)

# for item in items:
# 	print(item)
# temp = None
# for item in items:
# 	print('------------------------------')
# 	file = makeParent(item['path'])
# 	print('temp,file',temp,file)
# 	temp = merge(temp,file)
	
# 	print('------------------------------')
# print(temp)


path1 = ['数据结构','用户模块','用户表','获取用户列表.md']
path2 = ['数据结构','用户模块','用户表','修改用户列表.md']
files1 = getFilePath(path1)
files2 = getFilePath(path2)
print(files1)
print(files2)
print(merge(files1,files2))

# path = ['获取用户列表.md']
# 1, -1
# children 0
# result = {"name":'获取用户列表.md'}
# path = ['数据结构','用户模块','用户表','获取用户列表.md']
# 1, -2
# children 
# result = {"name":"用户表","children":[{"name":'获取用户列表.md'}]}
# path = ['数据结构','用户模块','用户表','获取用户列表.md']
# 1, -3
# result = {"name":'用户模块',"children":[{"name":"用户表","children":[{"name":'获取用户列表.md'}]}]}
# path = ['数据结构','用户模块','用户表','获取用户列表.md']
# 1, -4
# result = {"name":'数据结构',"children":[{"name":'用户模块',"children":[{"name":"用户表","children":[{"name":'获取用户列表.md'}]}]}]}


# file1 = {'name': '开发文档', 'path': '', 'children': [{'name': '价格备案.md', 'path': '开发文档'}, {'name': '接口文档.md', 'path': '开发文档'}]}
# file2 = {'name': '开发文档', 'path': '', 'children': [{'name': '流程总览.md', 'path': '开发文档'}]}
# print(merge(file1,file2))
# print(mergeList(file1,file3))