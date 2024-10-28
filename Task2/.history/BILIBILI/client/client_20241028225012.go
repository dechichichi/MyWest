package client

import (
	"bilibili/pkg"
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Fetch(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func GetSecondId(data string) ([]string, int) {
	var Data pkg.Data
	err := json.Unmarshal([]byte(data), &Data)
	if err != nil {
		return nil, 0
	}
	ids := make([]string, len(Data.Replies))
	for i, reply := range Data.Replies {
		ids[i] = reply.Rpid
	}
	return ids, Data.Page.Num
}

func CheckAndWriteToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}
	return writer.Flush()
}
