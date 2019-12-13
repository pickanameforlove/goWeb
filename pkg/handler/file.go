package handler

import (
	"fmt"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	fmt.Println(123654)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}

	xlsx, err1 := excelize.OpenReader(file)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	rows := xlsx.GetRows("Sheet1")
	for irow, row := range rows {
		if irow > 0 {
			for _, cell := range row {
				fmt.Print(cell)
			}
			fmt.Println()
		}
	}
	//content, err := ioutil.ReadAll(file)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
	//	return
	//}
	//
	fmt.Println(header.Filename)
	//fmt.Println(string(content))
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}
