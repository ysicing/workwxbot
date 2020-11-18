package workwxbot_test

import (
	"github.com/ysicing/workwxbot"
	"testing"
)

func TestClient_Send(t *testing.T) {
	corpid := "wwxxxx"
	agentid := int64(1000004)
	secretkey := "xxxxxxx"
	client := workwxbot.Client{CropID: corpid, AgentID: agentid, AgentSecret: secretkey}
	msg := workwxbot.Message{
		ToUser:   "ysicing",
		MsgType:  "markdown",
		Markdown: workwxbot.Content{Content: "### 测试"},
	}
	t.Log(client.Send(msg))
}
