package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 200; i++ {
		fmt.Println(i)
		data, err := Fetch("https://api.bilibili.com/x/v2/reply?&type=1&oid=420981979&pn=" + strconv.Itoa(i) + "&sort=1")
		if err != nil {
			fmt.Println(err)
			continue
		}
		ids, pn := GetSecondId(data)
		for _, id := range ids {
			subData, err := Fetch("https://api.bilibili.com/x/v2/reply/reply?oid=420981979&pn=1&ps=10&root=" + id + "&type=1")
			if err != nil {
				fmt.Println(err)
				continue
			}
			path := "第 " + strconv.Itoa(pn) + "页" + id + ".txt"
			CheckAndWriteToFile(path, subData)
		}
	}
}
