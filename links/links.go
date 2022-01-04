// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
//!+Extract

// Package links provides a link-extraction function.
package links

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)
/**
 * @Description 保存页面到文件
 * @Param s string: 页面内容
 * @Param urlStr string: url全路径
 * @Param domainUrl string: 网站的域名
 * @return
 **/
func saveAsFile(s, urlStr string) {
	// 创建根文件夹，名字为域名
	urlObj, _ := url.Parse(urlStr)
	folderPath := urlObj.Host
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, os.ModePerm)
		os.Chmod(folderPath, os.ModePerm)
	}

	// 得到合适的文件名字
	fileName := "./" + folderPath
	urlStr,_ = url.PathUnescape(urlStr)
	// 此处没有创建目录，以url path为文件名，以_隔开
	if len(urlObj.Path) >0 && urlObj.Path != "/"{
		if strings.HasSuffix(urlStr, "/") {  // `/`结尾的页面，去除末尾的`/`
			urlObj.Path = urlObj.Path[:len(urlObj.Path)-1]
		}
		pathSplitStr := strings.Split(urlObj.Path, "/")
		if len(pathSplitStr) > 2 {
			dirName := "./" + folderPath + strings.Join(pathSplitStr[:len(pathSplitStr)-1], "/")
			if _, err := os.Stat(dirName); os.IsNotExist(err) {
				// 必须分成两步：先创建文件夹、再修改权限
				os.MkdirAll(dirName, os.ModePerm)
				os.Chmod(dirName, os.ModePerm)
			}
		}
		fileName = fileName + urlObj.Path
	} else {
		fileName = fileName + "/index"
	}

	// 写文件
	file, err := os.OpenFile(fileName + ".html", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(s)
}

func Extract(url, domainUrl string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	resp.Body.Close()
	doc, err := html.Parse(strings.NewReader(bodyStr))
	saveAsFile(bodyStr, url)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if link.Fragment != "" || !strings.HasPrefix(link.String(), domainUrl) {  // 不处理非本站的和页内链接的url
					break
				}
				if err != nil {
					break // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
