package main

import (
	"fmt"
	"github.com/hoanhan101/ultimate-go/go/httpc"
	"runtime"
	"ti
)

func main()  {
	for i := 0; i < 10; i++{
		dis := time.Now()
		doUpload()
		fmt.Println( "======================================")
		fmt.Println(time.Now().Sub(dis).Microseconds())
		fmt.Println("======================================")
		if i > 5{
			runtime.Goexit()
		}
	}
}

func doUpload()  {
	//新建一个http客户端
	client:=httpc.NewHttpClient()
	//新建一个请求
	req:=httpc.NewRequest(client)
	req.SetMethod("post").SetUrl("http://localhost:9096/api/v2/r0storage/upload/fileUpload/ogg")
	//设置上传的文件
	req.SetFileData("file","/Users/rice/Desktop/test.zip",true)
	//设置附加参数
	req.SetFileData("client","httpc",false)
	resp,body,err:=req.Send(true).End()
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println(resp)
		fmt.Println(body)
	}
}