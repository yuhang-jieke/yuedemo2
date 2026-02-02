package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func downloadImage(url, saveDir string) (string, error) {
	// 1. 创建保存目录（如果不存在）
	err := os.MkdirAll(saveDir, 0755)
	if err != nil {
		return "", fmt.Errorf("创建目录失败: %v", err)
	}

	// 2. 生成唯一的文件名（用时间戳+原文件名后缀）
	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(url))
	filePath := filepath.Join(saveDir, fileName)

	// 3. 发送 HTTP 请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("请求图片失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 4. 创建本地文件并写入内容
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("写入文件失败: %v", err)
	}

	return filePath, nil
}
func GetImg(address string, selector string) {

	url := launcher.New().
		Headless(false).
		Set("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0").
		MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()
	page := browser.MustPage(address)
	page.MustWaitLoad()
	err := page.Timeout(2 * time.Minute).MustElement(selector).WaitVisible()
	if err != nil {
		fmt.Println("等待元素超时")
		return
	}
	time.Sleep(5 * time.Second)
	element := page.MustElements(selector)
	fmt.Printf("获取到%d个元素\n", len(element))
	saveDir := "./downloaded_images"
	for i, e := range element {
		if e.MustVisible() {
			src, err := e.Attribute("src")
			if err == nil && src != nil {
				fmt.Printf("%d.%s\n", i+1, *src)
				localPath, err := downloadImage(*src, saveDir)
				if err != nil {
					fmt.Printf("下载图片失败: %v\n", err)
					continue
				}
				fmt.Printf("图片已保存到: %s\n", localPath)

				// 如果你需要上传本地文件到 MinIO，这里传 localPath 而不是 src
				time.Sleep(2 * time.Second)
				Upload(localPath)
			}

		}
	}
}
