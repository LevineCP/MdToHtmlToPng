package common

import (
	"bytes"
	"fmt"
	"github.com/qiniu/x/log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func MdToHtmlToPngS(fileContent string) string {
	nameTime := NumToStr(time.Now().UnixMilli()) 
	
	//fileContent = "# **马斯克的首款多模态大模型**\n## **模型介绍**\n### **模型名**\n- Grok-1.5V \n### **模型特性**\n- 能理解文本，处理文档、图表、截图和照片中的内容 \n- 可进行多学科推理，理解物理世界 \n- 在RealWorldQA基准测试中，表现优于同类产品 \n### **模型功能**\n- 能将流程图的白板草图转换为Python代码 \n- 能计算卡路里 \n- 能根据孩子的绘画生成睡前故事 \n- 能解释流行语，将表格转化为CSV文件格式 \n\n## **模型发展**\n### **开源情况**\n- Grok-1已经开源 \n### **模型更新**\n- Grok-1.5V即将发布"

	fileContent = strings.Replace(fileContent, "### ", "### <!-- markmap: fold --> ", -1)
	// 将Markdown内容写入文件
	mdName := nameTime + ".md"
	err := os.WriteFile(mdName, []byte(fileContent), 0644)
	if err != nil {
		log.Error("Error writing to file:", err)
		return ""
	}

	fileName := nameTime + ".html"
	fileArg := "-o" + fileName
	// 创建一个*exec.Cmd对象，表示要执行的外部命令
	cmd := exec.Command("npx", "markmap-cli", "--no-open", "--no-toolbar", fileArg, mdName) // 假设我们要执行的是Unix下的"ls -l"命令

	// 创建一个缓冲区，用于存储命令的标准输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令，如果命令执行出错，则返回错误
	err = cmd.Run()
	if err != nil {
		log.Error("命令执行出错:", err)
		return ""
	}

	for {
		if _, err := os.Stat(fileName); err == nil {
			break
		}
	}

	// 换成自己的 上传文件
	/*fileUrl, _ := s3client.NewSo().UploadFile("xmind/html/"+fileName, fileName)
	if fileUrl == `` {
		log.Error("Error uploading file to S3")
		return ""
	}
	log.Info("fileUrl:%s", fileUrl)

	// 运行Chrome，加载临时HTML文件并截图
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(300, 200, chromedp.EmulateScale(1.5)),
		chromedp.Navigate(fileUrl),
		chromedp.Sleep(1000*time.Millisecond),
		chromedp.FullScreenshot(&buf, 100),
	); err != nil {
		log.Error("Chrome截图失败:", err)
		return ""
	}

	imageName := nameTime + ".png"
	if err := os.WriteFile(imageName, buf, 0o644); err != nil {
		log.Error("写入PNG文件失败:", err)
		return ""
	}

	// 换成自己的 上传文件到S3
	imageUrl, _ := s3client.NewSo().UploadFile("xmind/html/"+imageName, imageName)
	if imageUrl == `` {
		log.Error("Error uploading image file to S3")
		return ""
	}
	log.Info("imageUrl:%s", imageUrl)

 	// 删除生成的文件
	for _, v := range []string{mdName, fileName, imageName} {

		err = os.Remove(v)
		if err != nil {
			// 如果遇到其他类型的错误，则打印错误信息并退出
			log.Error("Failed to delete file:", v, err)
		}
	}*/
	return fileName
}

func NumToStr(num interface{}) string {
	switch num.(type) {
	case string:
		return num.(string)
	case int:
		return strconv.Itoa(num.(int))
	case int64:
		return strconv.FormatInt(int64(num.(int64)), 10)
	case float64:
		return strconv.FormatFloat(float64(num.(float64)), 'f', -1, 64)
	}
	return fmt.Sprintf("%d", num)
}
