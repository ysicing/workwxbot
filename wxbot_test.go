package workwxbot_test

import (
	"github.com/ysicing/workwxbot"
	"log"
)

func ExampleRobot_SendMarkdown() {
	webhook := "http://wr.oa.local.com:89/api/Robot?token=xxxxxxxxxxxx"
	wxbot := workwxbot.NewRobot(webhook)
	configid := "6eb955fa-xxxx-4e36-xxxx-a6b44817181f"
	content := "### 测试\n > 测试哈哈哈哈"
	err := wxbot.SendMarkdown(configid, content, nil)
	if err != nil {
		log.Fatal(err)
	}
}
