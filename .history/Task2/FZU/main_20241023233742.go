package main

import (
	"fmt"
	"fzu/client"
	"time"
)

func main() {
	start := 1
	end := 100
	t1 := time.Now()
	client.SlowWork(start, end)
	t2 := time.Now()
	fmt.Println("Time taken by SlowWork:", t2.Sub(t1)) //48.5185141s
	client.Work(start, end)
	t3 := time.Now()
	fmt.Println("Time taken by Work:", t3.Sub(t2)) //268.6945ms
}
