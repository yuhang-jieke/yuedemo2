package pkg

import (
	"fmt"

	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/yuhang-jieke/yuedemo2/srv/user-server/basic/config"
	"github.com/yuhang-jieke/yuedemo2/srv/user-server/model"
)

func GetTitle(address string, selector string) {
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
	for i, e := range element {
		if e.MustVisible() {

			text := e.MustText()
			title := model.TitleContent{
				Content: text,
			}
			if err = title.CreateTitle(config.DB); err != nil {
				fmt.Println("添加失败")
				return
			}
			fmt.Printf("%d.%s\n", i+1, text)
		}
	}
}
