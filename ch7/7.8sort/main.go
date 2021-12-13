package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks




//!+customcode
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


type sortByTitle []*Track

func (x sortByTitle) Len() int           { return len(x) }
func (x sortByTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x sortByTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type sortByYear []*Track

func (x sortByYear) Len() int           { return len(x) }
func (x sortByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x sortByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!+yearcode
type sortByLength []*Track

func (x sortByLength) Len() int           { return len(x) }
func (x sortByLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x sortByLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }


//!-customcode

type x []int

// 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，
// 第二个排序键是第二最近点击过列头的列，等等。定义一个sort.Interface的实现用在这样的表格中。
// 比较这个实现方式和重复使用sort.Stable来排序的方式。
func main() {

	customSort := customSort{t:tracks, sortKeys: []string{"Title", "Year", "Length"}}
	customSortFuncs := make(map[string]func(x, y *Track) bool)
	customSortFuncs["Title"] = byTitle
	customSortFuncs["Year"] = byYear
	customSortFuncs["Length"] = byLength
	customSort.less = customSortFuncs

	fmt.Printf("\nCustom default sort %v: \n", customSort.sortKeys)
	sort.Sort(customSort)
	printTracks(tracks)

	//fmt.Printf("\nCustom click %s:", customSort.sortKeys[2])
	//customSort.clinkAKey(2)
	//fmt.Printf("sort %v: \n", customSort.sortKeys)
	//sort.Sort(customSort)
	//printTracks(tracks)
	//
	//fmt.Printf("\nCustom click %s:", customSort.sortKeys[1])
	//customSort.clinkAKey(1)
	//fmt.Printf("sort %v: \n", customSort.sortKeys)
	//sort.Sort(customSort)
	//printTracks(tracks)

}

