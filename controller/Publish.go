package controller

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"net/http"
	"os"
	"simpledouyin/role"
	"simpledouyin/service"
	"strings"
	"time"
)

func Publish(c *gin.Context) {

	//video实例化
	var video = role.Video{}

	// 获得数据
	data, _ := c.FormFile("data")
	token := c.PostForm("token")
	title := c.PostForm("title")

	if _, exist := usersLoginInfo[token]; exist {

		// dst 是想要放到的位置
		dst := "/public/" + data.Filename
		err := c.SaveUploadedFile(data, "."+dst)
		if err != nil {
			return
		}

		strArray := strings.Split(data.Filename, ".")
		ImageName := strArray[0]

		imgPath, _ := GetSnapshot("."+dst, ImageName, 1)

		//修改数据库中的

		userID := usersLoginInfo[token]

		var author role.Author
		service.Db.Where("id = ?", userID).Find(&author)

		// 增加作者的作品数量
		author.WorkCount++

		// 更新数据库中的作品数量
		service.Db.Model(&role.Author{}).Where("name = ?", author.Name).Update("work_count", author.WorkCount)

		video.Title = title
		video.Author = author
		video.AuthorID = author.ID
		video.CoverUrl = "http://192.168.1.4:8080" + imgPath
		video.PlayUrl = "http://192.168.1.4:8080" + dst
		video.CreateTime = int(time.Now().Unix())

		service.Db.Create(&video)

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "上传成功",
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "没有这个用户",
		})

	}

}

func GetSnapshot(videoPath, imageName string, frameNum int) (ImagePath string, err error) {
	snapshotPath := "/public/picture/" + imageName
	buf := bytes.NewBuffer(nil)
	err = ffmpeg_go.Input(videoPath).Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败1：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败2：", err)
		return "", err
	}

	err = imaging.Save(img, "."+snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败3：", err)
		return "", err
	}

	imgPath := snapshotPath + ".png"

	return imgPath, nil
}
