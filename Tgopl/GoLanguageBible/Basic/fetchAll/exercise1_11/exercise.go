package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second, // 设置10秒超时
}

func main() {
	start := time.Now()
	ch := make(chan string) // 创建无缓冲通道

	// 启动 goroutines 执行并发请求
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// 读取所有的 goroutine 发送回来的结果
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	// 如果 URL 没有前缀，加上 http:// 前缀
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// 使用带超时控制的 http.Client 发送请求
	resp, err := client.Get(url)
	if err != nil {
		// 错误情况，向通道发送错误信息
		ch <- fmt.Sprintf("error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体，丢弃数据
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	// 将请求结果发送到通道
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
