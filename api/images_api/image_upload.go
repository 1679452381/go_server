/**
 * @author X
 * @date 2023/4/6
 * @description
 **/
package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_server/global"
	"go_server/models/res"
	"os"
	"path/filepath"
	"strings"
)

type ImageUploadResponse struct {
	FilName   string
	IsSuccess bool
	Msg       string
}

var WhiteImageList = map[string]bool{
	"jpg":  true,
	"png":  true,
	"jpeg": true,
	"ico":  true,
	"tiff": true,
	"svg":  true,
}

func (ImageApi) ImageUploadView(c *gin.Context) {
	//获取一张图片
	//fileHeader, err := c.FormFile("image")
	//if err != nil {
	//	global.Log.Error(err.Error())
	//	return
	//}

	//获取多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		global.Log.Error(err.Error())
		return
	}
	imageList, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("文件不存在", c)
	}
	//判断路径是否存在
	fileInfo, err := os.Stat(filepath.Dir(global.Config.Local.Path))
	if os.IsNotExist(err) {
		//不存在创建
		err = os.MkdirAll(filepath.Dir(global.Config.Local.Path), os.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}
	fmt.Println(fileInfo)

	var resList []ImageUploadResponse
	for _, image := range imageList {

		//白名单
		fileName := image.Filename
		nameList := strings.Split(fileName, ".")
		//获取图片后缀
		suffix := strings.ToLower(nameList[len(nameList)-1])
		if !WhiteImageList[suffix] {
			resList = append(resList, ImageUploadResponse{
				fileName,
				false,
				"不支持上传此类型文件",
			})
			continue
		}
		//限制文件大小
		fileSize := float64(image.Size) / float64(1024*1024)
		if fileSize > float64(global.Config.Local.FileSize) {
			resList = append(resList, ImageUploadResponse{
				fileName,
				false,
				fmt.Sprintf("图片大小为:%.2fMB,超过设定大小:%dMB", fileSize, global.Config.Local.FileSize),
			})
			continue
		}
		//保存文件
		filePath := filepath.Join("uploads", image.Filename)
		err := c.SaveUploadedFile(image, filePath)
		if err != nil {
			resList = append(resList, ImageUploadResponse{
				fileName,
				false,
				err.Error(),
			})
			continue
		}
		resList = append(resList, ImageUploadResponse{
			filePath,
			true,
			fmt.Sprintf("上传成功"),
		})
	}
	res.OkWithData(resList, c)
}
