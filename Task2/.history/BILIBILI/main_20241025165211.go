package main

import (
	"bilibili/client"
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 200; i++ {
		go func() {
			fmt.Println(i)
			data, err := client.Fetch("https://api.bilibili.com/x/v2/reply?&type=1&oid=420981979&pn=" + strconv.Itoa(i) + "&sort=1")
			if err != nil {
				fmt.Println(err)
			}
			ids, pn := client.GetSecondId(data)
			go func() {
				for _, id := range ids {
					subData, err := client.Fetch("https://api.bilibili.com/x/v2/reply/reply?oid=420981979&pn=1&ps=10&root=" + id + "&type=1")
					if err != nil {
						fmt.Println(err)
					}
					path := "第 " + strconv.Itoa(pn) + "页" + id + ".txt"
					client.CheckAndWriteToFile(path, subData)
				}
			}()
		}()
	}
}
