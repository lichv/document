from lichv.utils import * 

dirname = './docs/'


filename = './docs/内含子.md'

# text = read(filename)


# print(text)

def getFiles(source='.', extends='.md'):
	source = source.replace('/','\\')
	if not source.endswith('\\'):
		source = source + '\\'
	items = []
	if os.path.isdir(source):
		for parent, dirnames, filenames in os.walk(source, topdown=False,	followlinks=True):
			for filename in filenames:
				if filename.lower().endswith(extends) :
					name = filename[:filename.index(extends)]
					dirname = parent[len(source):]
					print(name)
					print(parent+'\\'+filename)
					items.append({'filename':name,'dirname':dirname,'path':parent+'\\'+filename})
	return items



dirname = os.getcwd()+"../../../frontend"
print(dirname)
items = getFiles(dirname)
print(items)