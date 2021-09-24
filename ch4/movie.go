package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const apikey = "837a1b8b"
const apiUrl = "http://www.omdbapi.com/"

type Movie struct {
	imdbID string
	Title string
	Poster string
}


func fetchMovie(url string)(*Movie, error) {
	/**
	 * @Description 根据指定url获取response
	 * @Param url string，路径url
	 * @return "", nil
	 **/
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("获取url: %s 失败。\n", url)
	}
	var movie Movie
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return nil, fmt.Errorf("读取url: %s response body失败。\n", url)
	}
	fmt.Printf("++++++++++++++get %s success: %s \n", url, movie)
	return &movie, nil
}

func downloadPoster(posterUrl string, fileName string) {
	resp, err := http.Get(posterUrl)
	if err != nil {
		fmt.Printf("下载失败： %s", err)
		os.Exit(1)
	}
	var suffixes = "jpg"
	if names := strings.Split(posterUrl, ".");len(names) == 2 {
		suffixes = names[1]
	}
	file, _ := os.Create(fmt.Sprintf("%s.%s", fileName, suffixes))
	defer file.Close()
	respB, _:= ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	_, err = file.Write(respB)
	if err != nil {
		fmt.Printf("写入新的文件失败： %s", err)
		os.Exit(1)
	}

}

func main() {
	detail, err := fetchMovie(fmt.Sprintf("%s?t=%s&apikey=%s", apiUrl, url.QueryEscape("Sherlock Holmes"), apikey))
	if err != nil {
		fmt.Printf("%s", err)
	}

	downloadPoster(detail.Poster, detail.Title)
}
