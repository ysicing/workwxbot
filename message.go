package workwxbot

const (
	msgTypeMarkdown = "markdown"
	programType     = "OA"
	isSendNow       = true
)

type markdownMessage struct {
	MsgType       string   `json:"msgtype"`
	ProgramType   string   `json:"program"`
	IsSendNow     bool     `json:"issendimmediately"`
	ConfigID      string   `json:"configid"`
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list"`
}
