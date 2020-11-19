package wxbot

import (
	"github.com/ysicing/workwxbot"
)

//func SendBot() {
//	webhook := "http://wr.oa.baidu.com/api/Robot?token=12345678"
//	wxbot := workwxbot.NewRobot(webhook)
//	configid := "12b9a5fa-xxxx-xxxx-8d7f-a6b4f8c7184f"
//	content := "测试 测试哈哈哈哈"
//	msg := workwxbot.BotMessage{
//		MsgType:       "markdown",
//		ProgramType:   "OA",
//		IsSendNow:     true,
//		ConfigID:      configid,
//		Content:       content,
//		MentionedList: "@all",
//	}
//	wxbot.Send(msg)
//}

func main() {
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=12b9a5fa-xxxx-xxxx-8d7f-a6b4f8c7184f"
	wxbot := workwxbot.NewRobot(webhook)
	text := workwxbot.WxBotMessage{
		MsgType: "text",
		BotText: workwxbot.BotText{
			Content:       "测试",
			MentionedList: []string{"@all"},
		},
	}
	wxbot.Send(text)
	markdown := workwxbot.WxBotMessage{
		MsgType:  "markdown",
		MarkDown: workwxbot.BotMarkDown{Content: "### 哈哈哈"}}
	wxbot.Send(markdown)
	news := workwxbot.WxBotMessage{
		MsgType: "news",
		News: workwxbot.News{
			Articles: []workwxbot.Article{
				{
					Title:       "中秋节礼品领取",
					Description: "今年中秋节公司有豪礼相送",
					Url:         "https://jira.baidu.com",
					Picurl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
				},
				{
					Title:       "国庆节礼品领取",
					Description: "今年国庆节公司有豪礼相送",
					Url:         "https://wiki.baidu.com",
					Picurl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
				},
			},
		},
	}
	wxbot.Send(news)
}
