package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//文件目录树形结构节点
type fileNode struct {
	name  string `json:"name"`
	child []fileNode `json:"child"`
}

//递归遍历文件目录
func getFilesTree(pathName string) (fileNode, error) {
	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		log.Fatalf("Read dir '%s' failed: %v", pathName, err)
	}
	var tree, childNode fileNode
	tree.name = pathName
	var name, fullName string
	for _, fileDir := range rd {
		name = fileDir.Name()
		fullName = pathName + "/" + name
		if fileDir.IsDir() {
			childNode, err = getFilesTree(fullName)
			if err != nil {
				log.Fatalf("Read dir '%s' failed: %v", fullName, err)
			}
		} else {
			childNode.name = name
			childNode.child = nil
		}
		tree.child = append(tree.child, childNode)
	}
	return tree, nil
}

//递归打印文件目录
func printDirTree(tree fileNode, prefix string) {
	fmt.Println(prefix + tree.name)
	if len(tree.child) > 0 {
		prefix += "----"
		for _, childNode := range tree.child {
			printDirTree(childNode, prefix)
		}
	}
}

//func main() {
//	dirName := "./docs"
//	tree, err := getFilesTree(dirName)
//	if err != nil {
//		log.Fatalln("read dir '%s' fail: %v", dirName, err)
//	}
//	printDirTree(tree, "")
//	s,err:=json.Marshal(tree)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(s)
//
//}