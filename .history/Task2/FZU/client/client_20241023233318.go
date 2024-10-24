package client

import (
	"fmt"
	"fzu/pkg"
	"io"
	"net/http"
	"os"
	"sync"
)

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result := ""
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		result += string(buf[:n])
	}
	return result, nil
}

func HttpGetSingle(i int, page chan<- int, f *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=%d&wbtreeid=1460", i)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Printf("爬取失败: %s\n", err.Error())
		return
	}

	var ans string
	ans1 := pkg.Writer.FindAllStringSubmatch(result, -1)
	ans2 := pkg.Title.FindAllStringSubmatch(result, -1)
	ans3 := pkg.Text.FindAllStringSubmatch(result, -1)
	ans4 := pkg.Time.FindAllStringSubmatch(result, -1)
	for a := 0; a < len(ans1) && a < len(ans2) && a < len(ans3) && a < len(ans4); a++ {
		ans += fmt.Sprintf("%s  %s  %s  %s\n", ans1[a][1], ans2[a][1], ans3[a][1], ans4[a][1])
	}

	_, err = f.WriteString(ans)
	if err != nil {
		fmt.Printf("写入文件失败: %s\n", err.Error())
		return
	}

	page <- i
}

func SHttpGetSingle(i int, page chan<- int, f *os.File) {
	url := fmt.Sprintf("https://info22.fzu.edu.cn/lm_list.jsp?totalpage=1015&PAGENUM=%d&wbtreeid=1460", i)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Printf("爬取失败: %s\n", err.Error())
		return
	}

	var ans string
	ans1 := pkg.Writer.FindAllStringSubmatch(result, -1)
	ans2 := pkg.Title.FindAllStringSubmatch(result, -1)
	ans3 := pkg.Text.FindAllStringSubmatch(result, -1)
	ans4 := pkg.Time.FindAllStringSubmatch(result, -1)
	for a := 0; a < len(ans1) && a < len(ans2) && a < len(ans3) && a < len(ans4); a++ {
		ans += fmt.Sprintf("%s  %s  %s  %s\n", ans1[a][1], ans2[a][1], ans3[a][1], ans4[a][1])
	}

	_, err = f.WriteString(ans)
	if err != nil {
		fmt.Printf("写入文件失败: %s\n", err.Error())
		return
	}

	page <- i
}
func Work(start int, end int) {
	page := make(chan int, end-start+1)
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Printf("创建文件失败: %s\n", err.Error())
		return
	}
	defer f.Close()

	var wg sync.WaitGroup
	for i := start; i <= end; i++ {
		wg.Add(1)
		go HttpGetSingle(i, page, f, &wg)
	}

	go func() {
		wg.Wait()
		close(page)
	}()

	for i := start; i <= end; i++ {
		fmt.Printf("第%d页爬取成功\n", <-page)
	}
}

func SlowWork(start int, end int) {
	page := make(chan int, end-start+1)
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Printf("创建文件失败: %s\n", err.Error())
		return
	}
	defer f.Close()

	for i := start; i <= end; i++ {
		SHttpGetSingle(i, page, f)
		fmt.Printf("第%d页爬取成功\n", <-page)
	}
}
