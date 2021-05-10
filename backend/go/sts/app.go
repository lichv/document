package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type ReadFileForm struct {
	Filepath  string `json:"filepath" form:"filepath"`
}

func main() {
	var outport int
	flag.IntVar(&outport,"o",7654,"输出端口号")
	flag.Parse()

	engine := gin.Default()
	engine.Use(favicon.New("./public/favicon.ico"))
	engine.Static("/_assets", "./public/_assets")
	engine.LoadHTMLFiles("public/index.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	engine.Any("/api/aliyun/sts", func(c *gin.Context) {
		res,err :=GetSTSToken("LTAI5tKnra54xozuA3KktFur", "VyIHrtVQZxXeiuuBWUW2oG34qe87dk","acs:ram::1378573870105843:role/osser","osser")
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
	engine.Any("/api/msg/sendsms", func(context *gin.Context) {
		mobile := context.DefaultPostForm("mobile","")
		if mobile == "" {
			mobile = context.DefaultQuery("mobile","")
		}
		content := context.DefaultPostForm("content","")
		if content == "" {
			content = context.DefaultQuery("content","")
		}
		if mobile == "" || content == "" {
			context.JSON(200,gin.H{
				"state":3000,
				"message":"参数错误",
			})
			return
		}
		result, _ := Send("N1211363", "uEwY0cngbOa6e6", "易之科技", mobile, content)
		context.JSON(200,gin.H{
			"state":2000,
			"data":*result,
			"message":"success",
		})
	})

	engine.Any("/api/markdown/files", func(context *gin.Context) {
		dirName := "./docs"
		items := GetALLFIles_walk(dirName)
		result := []map[string]interface{}{}
		for _,item := range items {
			s := item[len(dirName)-1:]
			tmp := map[string]interface{}{}
			tmp["filename"] = strings.Replace(s,"\\","/",-1)
			result = append(result, tmp)
		}

		
		context.JSON(200,gin.H{
			"state":2000,
			"message":"success",
			"data":result,
		})
	})

	engine.Any("/api/markdown/read", func(context *gin.Context) {
		var readfile ReadFileForm
		context.Bind(&readfile)
		fmt.Println(readfile.Filepath)
		//fmt.Println(context)
		filepath := readfile.Filepath
		if filepath == ""{
			filepath = context.DefaultPostForm("path","")
		}
		//filepath = context.DefaultPostForm("path","")
		fmt.Println("filepath",filepath)
		if filepath == "" {
			filepath = context.DefaultQuery("path","")
		}

		fmt.Println("filepath",filepath)
		filepath = path.Join(".","docs",filepath)
		fmt.Println("filepath",filepath)
		if IsDir(filepath) {
			filepath = path.Join(filepath,"index.md")
		}
		fmt.Println("filepath",filepath)
		if IsExist(filepath){
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
	engine.Run(":"+outportStr)
}

func GetSTSToken(access_id string,access_key string, rolearn string,rolesessionrole string) (response sts.Credentials, err error){
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", access_id, access_key)
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = rolearn
	request.RoleSessionName = rolesessionrole
	resp, err :=client.AssumeRole(request)
	if err != nil {
		fmt.Println(err.Error())
		return sts.Credentials{"","","",""},nil
	}
	return (*resp).Credentials, nil
}

func Send(account,password,sign,mobile,content string) (*string,error) {
	params := make(map[string]interface{})
	params["account"] = account
	params["password"] = password
	params["phone"] = mobile
	params["msg"] = url.QueryEscape("【"+sign+"】"+content)
	params["report"] = "true"
	bytesData,err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json"
	request,err := http.NewRequest("POST",url,reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	client := http.Client{}
	resp,err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return str,nil
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func GetAllFile(pathname string, s []string) ([]string, error) {
	fromSlash := filepath.FromSlash(pathname)
	//fmt.Println(fromSlash)
	rd, err := ioutil.ReadDir(fromSlash)
	if err != nil {
		//log.LOGGER("Error").Error("read dir fail %v\n", err)
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir:= filepath.Join(fromSlash,fi.Name())
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				//log.LOGGER("Error").Error("read dir fail: %v\n", err)
				return s, err
			}
		} else {
			fullName:= filepath.Join(fromSlash,fi.Name())
			s = append(s, fullName)
		}
	}
	return s, nil
}

func GetALLFIles_walk(pathname string)([]string){
	StartTime :=time.Now();
	dst_target :=[]string{}
	err := filepath.Walk(pathname, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir(){
			return nil
		}else { //进行文件的复制
			dst_target=append(dst_target,src)

			//return s
		}
		//println(path)
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return nil
		//log.LOGGER("Error").Error("filepath.Walk() returned %v\n", err)
	}
	fmt.Println("Cost Time:",time.Since(StartTime))
	return dst_target
}
