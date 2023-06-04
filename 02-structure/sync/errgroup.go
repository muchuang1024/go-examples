package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.baidu.com/",
		"https://www.sina.com.cn/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	err := g.Wait()
	if err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println("fetched error:", err.Error())
	}
}
