package simple_http

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
	"github.com/Allen9012/AllenGame/sysservice/httpservice"
	"net/http"
)

// HttpService是origin引擎中系统实现的http服务，http接口中常用的GET,POST以及url路由处理。
// 注意，要在main.go中加入import _ "orginserver/simple_service"，并且在config/cluster/cluster.json中的ServiceList加入服务。

func init() {
	//必须先启动HttpService
	node.Setup(&httpservice.HttpService{})
	node.Setup(&ServiceHttp{})
}

type ServiceHttp struct {
	service.Service
}

func (slf *ServiceHttp) OnInit() error {
	//获取系统httpservice服务
	httpervice := node.GetService("HttpService").(*httpservice.HttpService)

	//新建并设置路由对象
	httpRouter := httpservice.NewHttpHttpRouter()
	httpervice.SetHttpRouter(httpRouter, slf.GetEventHandler())

	//GET方法，请求url:http://127.0.0.1:9402/get/query?nickname=boyce
	//并header中新增key为uid,value为1000的头,则用postman测试返回结果为：
	//head uid:1000, nickname:boyce
	httpRouter.GET("/get/query", slf.HttpGet)

	//POST方法 请求url:http://127.0.0.1:9402/post/query
	//返回结果为：{"msg":"hello world"}
	httpRouter.POST("/post/query", slf.HttpPost)

	//GET方式获取目录下的资源，http://127.0.0.1:port/img/head/a.jpg
	httpRouter.SetServeFile(httpservice.METHOD_GET, "/img/head/", "d:/img")

	fmt.Printf("【http服务启动】\n")

	return nil
}

func (slf *ServiceHttp) HttpGet(session *httpservice.HttpSession) {
	//从头中获取key为uid对应的值
	uid := session.GetHeader("uid")
	//从url参数中获取key为nickname对应的值
	nickname, _ := session.Query("nickname")
	//向body部分写入数据
	session.Write([]byte(fmt.Sprintf("head uid:%s, nickname:%s", uid, nickname)))
	//写入http状态
	session.WriteStatusCode(http.StatusOK)
	//完成返回
	session.Done()
}

type HttpRespone struct {
	Msg string `json:"msg"`
}

func (slf *ServiceHttp) HttpPost(session *httpservice.HttpSession) {
	//也可以采用直接返回数据对象方式，如下：
	session.WriteJsonDone(http.StatusOK, &HttpRespone{Msg: "hello world"})
}
