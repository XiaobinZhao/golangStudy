package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type xkcdDetail struct {
	Number int `json:"num"`
	Title string
	Year string
	Month string
	Day string
	//Transcript string
}

func main() {
	/*
	练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，
	一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。
	 下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
	*/
	param := os.Args[1]
	xkcdDetails := fetchDetails()
	var searchResult []xkcdDetail
	for _, detail := range xkcdDetails {
		if strings.Contains(detail.Title, param) {
			searchResult = append(searchResult, detail)
		}
	}
	fmt.Printf("search result: %+v", searchResult)

}

func fetchDetails()[]xkcdDetail  {
	var details []xkcdDetail
	for num:=1;num<10;num++ {
		detail, err := fetch(fmt.Sprintf("https://xkcd.com/%d/info.0.json", num))
		if err != nil {
			fmt.Println(err)
		}
		var xkcdDetail xkcdDetail
		detailB := []byte(detail)
		json.Unmarshal(detailB, &xkcdDetail)
		details = append(details, xkcdDetail)
	}

	fmt.Printf("-------------got total: %d \n", len(details))
	return details
}

func fetch(url string)(string, error) {
	/**
	 * @Description 根据指定url获取response
	 * @Param url string，路径url
	 * @return "", nil
	 **/
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("获取url: %s 失败。\n", url)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("读取url: %s response body失败。\n", url)
	}
	fmt.Printf("++++++++++++++get %s success \n", url)
	return string(respBody), nil
}