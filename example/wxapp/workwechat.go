package main

import (
	"github.com/ysicing/workwxbot"
)

func main() {
	corpid := "xxxx"
	agentid := int64(1000000)
	secretkey := "xxxx"
	client := workwxbot.Client{CropID: corpid, AgentID: agentid, AgentSecret: secretkey}
	md := workwxbot.Message{
		ToUser:   "xxxx",
		MsgType:  "markdown",
		Markdown: workwxbot.Content{Content: "### 测试"},
	}
	_ = client.Send(md)
	text := workwxbot.Message{
		ToUser:  "xxxx",
		MsgType: "text",
		Text:    workwxbot.Content{Content: "文本测试"},
	}
	_ = client.Send(text)
	textcard := workwxbot.Message{
		ToUser:  "xxx",
		MsgType: "textcard",
		Textcard: workwxbot.TextCard{
			Title:       "hahha",
			Description: "xxxx",
			Url:         "https://jira.baidu.com",
			Btntxt:      "更多",
		},
	}
	_ = client.Send(textcard)
	news := workwxbot.Message{
		ToUser:  "xxxx",
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
	_ = client.Send(news)
}
