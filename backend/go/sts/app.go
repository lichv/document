package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"unsafe"
)


func main() {
	var outport int
	flag.IntVar(&outport,"o",7654,"输出端口号")
	flag.Parse()

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"state":2000,
			"message":"success",
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
