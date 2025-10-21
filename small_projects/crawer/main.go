package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Get(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Printf("resp.Body.Read err= ", err)
				break
			}
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int, page chan<- int) {
	url := "https://github.com/search?q=go&type=repositories&p=" + strconv.Itoa(i)
	fmt.Printf("正在爬取第%d个网页\n", i)
	result, err := Get(url)
	if err != nil {
		fmt.Printf("http.Get err= ", err)
		return
	}
	filename := "page" + strconv.Itoa(i) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Printf("os.Create err= ", err1)
		return
	}

	f.WriteString(result)
	f.Close()
	page <- i
}

func Run(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完毕\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页数字>=1")
	fmt.Scan(&start)
	fmt.Println("请输入结束页数字：>")
	fmt.Scan(&end)
	Run(start, end)
}
