package test

import (
	"fmt"
	bdyp "github.com/zcxey2911/bdyp_upload_golang"
	"os"
	"testing"
)

func TestBaiduUpload(t *testing.T) {
	var bcloud = bdyp.Bcloud{}

	// 获取token
	res, err := bcloud.GetToken("9f7a8e57fe0989c6d8476c57bb1fb389", "oob", "ytshSVOuTLj0uKMBN4e75vjH4lZMGZrT", "ADfBkIM2yk1zzKeLCP7G9yO3Pij2M9nk")

	fmt.Println(res)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Printf("接口的token是: %#v\n", res.AccessToken)
	}
	// 读取文件
	f, err := os.Open("D:\\apps\\gin-vue-admin\\simple\\server\\sqlite.db")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer f.Close()

	// 上传文件
	print(bcloud.Upload(&bdyp.FileUploadReq{
		Name:  "/apps/fs/sqlite.db",
		File:  f,
		RType: nil,
	}))
}
