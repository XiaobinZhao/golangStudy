package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"studygolang/githubAPIIssues"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	//练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
	//issueTimeCategory := map[string][]*githubAPIIssues.Issue {
	//	"lessMonth": nil,
	//	"lessYear": nil,
	//	"moreYear": nil,
	//}
	//result, err := githubAPIIssues.SearchIssues(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",  // 9.9表示：9+0.9，其中0.9表示最多显示9个字符，9表示最小宽度，不满的使用空格补充。默认右对齐
	//		item.Number, item.User.Login, item.Title)
	//	if time.Now().Sub(item.CreatedAt).Seconds() < 30*24*60*60 {  // 一个月的秒
	//		issueTimeCategory["lessMonth"] = append(issueTimeCategory["lessMonth"], item)
	//	} else if time.Now().Sub(item.CreatedAt).Seconds() < 365*24*60*60 {
	//		issueTimeCategory["lessYear"] = append(issueTimeCategory["lessYear"], item)
	//	} else {
	//		issueTimeCategory["moreYear"] = append(issueTimeCategory["moreYear"], item)
	//	}
	//}

	//练习 4.11： 编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf(`githus issue contr:
                       - bye 
                       - SearchIssues 
                       - CreateIssue {title} {body}
                       - UpdateIssue {issueNumber} {title} {body}
                       - CloseIssue {issueNumber}
                      `)
	fmt.Println("")
	fmt.Printf(">")
	for input.Scan() {
		key := input.Text()
		if strings.HasPrefix(key, "SearchIssues") {
			keys := strings.Split(key, " ")[1:]
			result, err := githubAPIIssues.SearchIssues(keys)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d issues:\n", result.TotalCount)
			for _, item := range result.Items {
				fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
			}
		} else if strings.HasPrefix(key, "CreateIssue") {
			keys := strings.Split(key, " ")[1:]
			requestBody := githubAPIIssues.CreateIssueRequest{
				Title: keys[0],
				Body:  keys[1],
			}
			result, err := githubAPIIssues.CreateIssue(requestBody)
			if err != nil {
				log.Fatal(err)
			}
			responseResult , _ := json.MarshalIndent(result, "", " ")
			fmt.Printf("%s \n", responseResult)
		} else if strings.HasPrefix(key, "UpdateIssue") {
			keys := strings.Split(key, " ")[1:]
			requestBody := githubAPIIssues.CreateIssueRequest{
				Title: keys[1],
				Body:  keys[2],
			}
			issueNumber,_ := strconv.ParseInt(keys[0], 10,64)
			result, err := githubAPIIssues.UpdateIssue(issueNumber, requestBody)
			if err != nil {
				log.Fatal(err)
			}
			responseResult , _ := json.MarshalIndent(result, "", " ")
			fmt.Printf("%s \n", responseResult)
		} else if strings.HasPrefix(key, "CloseIssue") {
			keys := strings.Split(key, " ")[1:]
			issueNumber,_ := strconv.ParseInt(keys[0], 10,64)
			result, err := githubAPIIssues.CloseIssue(issueNumber)
			if err != nil {
				log.Fatal(err)
			}
			responseResult , _ := json.MarshalIndent(result, "", " ")
			fmt.Printf("%s \n", responseResult)
		} else if key == "bye" {
			fmt.Printf("exit")
			os.Exit(0)
		} else {
			fmt.Printf("invalidate param,continue \n")
		}
		fmt.Printf(">")
	}

}