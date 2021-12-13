package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

//!+main
type Track struct {
	Id  int
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{1, "Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{2, "Go", "Aly", "Moby", 1992, length("3m37s")},
	{3, "Go", "Moby", "Moby", 1992, length("2m37s")},
	{4, "Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m24s")},
	{5, "Ready 2 Go", "Martin Solveig", "Smash", 2012, length("4m24s")},
}

var tracksTemplate = template.Must(template.New("trackslist").Parse(`
<h1>{{len .}} tracks</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th Onclick="submitForm('Title')">Title
	<form action="" name="Title" method="post">
		<input type="hidden" name="orderBy" value="0" />
	</form>
  </th>
  <th>Artist</th>
  <th>Album</th>
  <th Onclick="submitForm('Year')">Year
	<form action="" name="Year" method="post">
		<input type="hidden" name="orderBy" value="1" />
	</form>
  </th>
  <th Onclick="submitForm('Length')">Length
	<form action="" name="Length" method="post">
		<input type="hidden" name="orderBy" value="2" />
	</form>
  </th>
</tr>
{{range .}}
<tr>
  <td>{{.Id}}</td>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitForm(formName) {
    document[formName].submit();
}
</script>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}


type customSort struct {
	t    []*Track
	less map[string]func(x, y *Track) bool  // 字段对应的排序算法，map的key对应字段名字，value就是排序方法
	sortKeys []string  // 排序的字段，按照顺序，sortKeys[0]就是第一排序字段，依次后推
}

func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool {
	for _,k := range x.sortKeys {
		if x.less[k](x.t[i], x.t[j]) {
			return true  // 大于
		} else if !x.less[k](x.t[j], x.t[i]) {  // i和j换一下位置，来判断等于
			continue  // 等于时，看下一个条件
		} else {
			return false  // 小于
		}
	}
	return false
}

func byTitle(i, j *Track) bool {
	return i.Title < j.Title
}
func byYear(i, j *Track) bool {
	return i.Year < j.Year
}
func byLength(i, j *Track) bool {
	return i.Length < j.Length
}

/**
 * @Description 点击了任何一个字段,更新排序优先级： 把该字段放在排序第一位，其他的往后顺移
 * @Param index：字段所在的序号
 * @return nil
 **/
func (x *customSort) clinkAKey(index int) {
	click := x.sortKeys[index]
	for i:=index -1 ; i>=0; i-- {
		x.sortKeys[i+1] = x.sortKeys[i]
	}
	x.sortKeys[0] = click
}


// 练习 7.9： 使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。将这个解决方案用在前一个练习中，
// 让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	http.HandleFunc("/", handler)
	log.Println("start server localhost:8000 ...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("accept get /")
	customSort := customSort{t:tracks, sortKeys: []string{"Title", "Year", "Length"}}
	customSortFuncs := make(map[string]func(x, y *Track) bool)
	customSortFuncs["Title"] = byTitle
	customSortFuncs["Year"] = byYear
	customSortFuncs["Length"] = byLength
	customSort.less = customSortFuncs

	if err := r.ParseForm(); err != nil {
		fmt.Printf("ParseForm: %v\n", err)
	}

	index := 0

	for k,v := range r.Form {
		if k == "orderBy" {
			index,_ = strconv.Atoi(v[0])
			customSort.clinkAKey(index)
		}
	}

	fmt.Printf("\nclick %s sort %v: \n", customSort.sortKeys[index], customSort.sortKeys)
	sort.Sort(customSort)

	if err := tracksTemplate.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}
