package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)


func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"state":2000,
			"message":"success",
		})
	})
	engine.GET("/aliyun/sts", func(c *gin.Context) {
		GetSTSToken("LTAI5tKnra54xozuA3KktFur", "VyIHrtVQZxXeiuuBWUW2oG34qe87dk","acs:ram::1378573870105843:role/osser","osser")
		c.JSON(200, gin.H{
			"state":2000,
			"message":"success",
			"data":
		})
	})
	engine.Run(":7655")
}

func GetSTSToken(access_id string,access_key string, rolearn string,rolesessionrole string) {
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", access_id, access_key)
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = rolearn
	request.RoleSessionName = rolesessionrole
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

