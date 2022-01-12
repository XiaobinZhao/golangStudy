package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
)
// 练习 8.5： 使用一个已有的CPU绑定的顺序程序，比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，
// 并将他们的主循环改为并发形式，使用channel来进行通信。在多核计算机上这个程序得到了多少速度上的改进？
// 使用多少个goroutine是最合适的呢？
//
// 程序运行中发现，当分辨率是1024*1024时，imgchan使用有缓存的chan,此时chan的缓存size会影响到内存的占用情况。我的机器是12G内存，
// 程序没有运行时，内存占用率为50%左右，剩余大概有6G左右内存。
// case1: imgchan := make(chan point1, 8)     报错：out of memory
// case2: imgchan := make(chan point1, 1024)     运行成功，但是内存飙升到99%
// case2: imgchan := make(chan point1, 1024*1024)     瞬间运行成功,没有看到内存占用率有明显大幅度波动
// 计划使用pprof工具，把本程序改为server版本，来看下内存占用情况，进行分析。 但是发现pprof抓大的也是一次请求完成时的内存结果，
// 需要手速很快的时候抓到内存情况。但是依然没有分析出来结果，仍需努力

type point2 struct {
	x,y int
	color color.Color
}
const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func myHandler(w http.ResponseWriter, r *http.Request) {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup

	imgchan := make(chan point2, 102400)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			wg.Add(1)
			go func(_px, _py int, _z complex128) {
				p := point2{_px, _py, mandelbrot2(_z)}
				imgchan <- p
				//img.Set(px, py, mandelbrot(z))
				defer wg.Done()
			}(px, py, z)
		}
	}
	go func() {
		wg.Wait()
		close(imgchan)
	}()

	for point := range imgchan {
		img.Set(point.x, point.y, point.color)
	}
	// why out of memory? width, height =1024
	// 因为每个goroutine都有一个单独的栈，栈初始大小为2KB,会按需扩展或者缩减。那么如果imgchan使用无缓存chan,
	// 就意味着，内存最少需要1024*1024*2KiB=2GiB的内存

	f, _ := os.Create("./xx.png")     //创建文件
	defer f.Close()
	png.Encode(f, img) // NOTE: ignoring errors
	fmt.Fprintf(w, "Success")
}


func main() {
	http.HandleFunc("/", myHandler)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		fmt.Printf("ListenAndServe: %s", err)
	}
}


func mandelbrot2(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
