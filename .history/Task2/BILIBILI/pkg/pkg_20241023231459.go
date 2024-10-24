package pkg

import "regexp"

// 定义数据结构以解析JSON
type ReplyData struct {
	Data struct {
		Replies []struct {
			Rpid string `json:"rpid"`
		} `json:"replies"`
		Page struct {
			Num int `json:"num"`
		} `json:"page"`
	} `json:"data"`
}

// 正则表达式全局变量
var (
	Writer   = regexp.MustCompile(`target=_blank class="lm_a" style="float:left;">【((.*?))】<\/a>`)
	Title    = regexp.MustCompile(`target=_blank title="((.*?))" style="">`)
	Text     = regexp.MustCompile(`<a href="((.*?))" target=_blank title=`)
	Time     = regexp.MustCompile(`<span class="fr">((.*?))</span>`)
	Maintext = regexp.MustCompile(`<META Name="description" Content=((.*?))/>`)
)
