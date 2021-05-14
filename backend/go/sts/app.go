package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lichv/go"
	"github.com/thinkerou/favicon"
	"net/http"
	"os"
	"path"
	"strconv"
	"testing"
)

type ReadFileForm struct {
	Filepath  string `json:"filepath" form:"filepath"`
}

func main() {
	var outport int
	var docs_path string
	var public_path string
	flag.IntVar(&outport,"o",7654,"输出端口号")
	flag.StringVar(&docs_path,"d","./docs/","文档目录")
	flag.StringVar(&public_path,"s","./public/","前端目录")

	testing.Init()
	if !flag.Parsed() {
		flag.Parse()
	}

	if !lichv.IsExist(docs_path) {
		fmt.Println("文档目录不存在")
		return
	}
	if !lichv.IsExist(public_path) {
		fmt.Println("静态文件目录不存在")
		return
	}

	engine := gin.Default()
	engine.Use(favicon.New(path.Join(public_path,"favicon.ico")))
	engine.Static("/_assets", path.Join(public_path,"/_assets"))
	engine.LoadHTMLFiles(path.Join(public_path,"/index.html"))
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	engine.Any("/api/aliyun/sts", func(c *gin.Context) {
		res,err :=lichv.GetSTSToken("LTAI5tKnra54xozuA3KktFur", "VyIHrtVQZxXeiuuBWUW2oG34qe87dk","acs:ram::1378573870105843:role/osser","osser","cn-hangzhou")
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(200,gin.H{
				"state":4001,
				"msg":"获取sts失败！",
			})
		}else {
			fmt.Print("ok")
			c.JSON(200, gin.H{
				"state":2000,
				"message":"success",
				"data": res,
			})
		}
	})

	engine.Any("/api/markdown/files", func(context *gin.Context) {
		result,_ := lichv.GetFileTree(docs_path)
		context.JSON(200,gin.H{
			"state":2000,
			"message":"success",
			"data":result,
		})
	})

	engine.Any("/api/markdown/read", func(context *gin.Context) {
		var readfile ReadFileForm
		context.Bind(&readfile)
		filepath := readfile.Filepath
		if filepath == ""{
			filepath = context.DefaultPostForm("path","")
		}
		if filepath == "" {
			filepath = context.DefaultQuery("path","")
		}
		filepath = path.Join(docs_path,filepath)
		if lichv.IsDir(filepath) {
			filepath = path.Join(filepath,"index.md")
		}
		if lichv.IsExist(filepath){
			bs,err := os.ReadFile(filepath)
			if err != nil {
				context.JSON(404,gin.H{
					"state":4000,
					"message":"failed",
				})
				return
			}
			context.JSON(200,gin.H{
				"state":2000,
				"message":"success",
				"data":string(bs),
			})
		}else{
			context.JSON(404,gin.H{
				"state":4000,
				"message":"failed",
			})
		}
	})

	outportStr := strconv.Itoa(outport)
	fmt.Println("从浏览器打开：http://localhost:"+outportStr)
	engine.Run(":"+outportStr)
}