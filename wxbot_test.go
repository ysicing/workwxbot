package workwxbot_test

import (
	"github.com/ysicing/workwxbot"
	"log"
	"testing"
)

func ExampleRobot_SendMarkdown() {
	webhook := "http://wr.oa.local.com:89/api/Robot?token=xxxxxxxxxxxx"
	wxbot := workwxbot.NewRobot(webhook)
	configid := "6eb955fa-xxxx-4e36-xxxx-a6b44817181f"
	content := "### 测试\n > 测试哈哈哈哈"
	err := wxbot.SendMarkdown("text", configid, content, "all")
	if err != nil {
		log.Fatal(err)
	}
}

func TestRobot_SendMarkdown(t *testing.T) {
	webhook := "http://wr.oa.local.com:89/api/Robot?token=xxxxxxxxxxxxxx"
	wxbot := workwxbot.NewRobot(webhook)
	configid := "6eb955fa-xxxx-4e36-xxxx-a6b44817181f"
	content := "测试 测试哈哈哈哈"
	t.Log(wxbot.SendMarkdown("text", configid, content, "@all"))
}
