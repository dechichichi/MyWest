package pkg

import "regexp"

// 正则表达式全局变量
var (
	Writer = regexp.MustCompile(`target=_blank class="lm_a" style="float:left;">【((.*?))】<\/a>`)
	Title  = regexp.MustCompile(`target=_blank title="((.*?))" style="">`)
	Text   = regexp.MustCompile(`<a href="((.*?))" target=_blank title=`)
	Time   = regexp.MustCompile(`<span class="fr">((.*?))</span>`)
)
