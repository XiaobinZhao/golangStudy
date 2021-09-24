package main

import (
	"html/template"
	"log"
	"net/http"
	"studygolang/githubAPI"
)

//练习 4.14：创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Milestone</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  {{if .Milestone}}
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
  {{end}}
</tr>
{{end}}
</table>
`))


func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	http.HandleFunc("/", handler)
	log.Println("start server localhost:8000 ...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("accept get /")
	result, err := githubAPI.SearchIssues([]string{"repo:golang/go","is:open","json decoder"})
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
