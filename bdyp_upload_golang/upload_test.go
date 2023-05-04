package bdyp_upload_golang

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestBaiduUpload(t *testing.T) {
	var bcloud = Bcloud{}
	bcloud.SetAppKey("ytshSVOuTLj0uKMBN4e75vjH4lZMGZrT")
	bcloud.SetSecret("ADfBkIM2yk1zzKeLCP7G9yO3Pij2M9nk")
	bcloud.SetAccessToken("121.e8e380d8edd3e7c75252afb331033708.YmwnHN8gLXwceO2oRKfcjEEtohoiiS1b2YclXe5.ywMBCQ")
	bcloud.SetRefreshToken("122.ffa358a244253409c7aaf168df0e87d2.YsKTg2aZH7d62jzK9Nw3UfLvM1IZF-KeWF7jP6n.WMbFrg")

	// 读取文件
	f, err := os.ReadFile("D:\\apps\\gin-vue-admin\\simple\\server\\sqlite.db")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	// 上传文件
	err = bcloud.Upload(&FileUploadReq{
		Name:  "/apps/fs/aa/sqlite.db",
		File:  bytes.NewBuffer(f),
		RType: nil,
	})
	if err != nil {
		fmt.Println("Upload err", err)
	} else {
		fmt.Println("上传成功")
	}

}
