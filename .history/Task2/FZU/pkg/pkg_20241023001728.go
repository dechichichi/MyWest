package pkg

import "regexp"

// 正则表达式全局变量
var (
	writer = regexp.MustCompile(`target=_blank class="lm_a" style="float:left;">【((.*?))】<\/a>`)
	title  = regexp.MustCompile(`target=_blank title="((.*?))" style="">`)
	text   = regexp.MustCompile(`<a href="((.*?))" target=_blank title=`)
	time   = regexp.MustCompile(`<span class="fr">((.*?))</span>`)
)
