package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)
// 练习 8.5： 使用一个已有的CPU绑定的顺序程序，比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，
// 并将他们的主循环改为并发形式，使用channel来进行通信。在多核计算机上这个程序得到了多少速度上的改进？
// 使用多少个goroutine是最合适的呢？


type point1 struct {
	x,y int
	color color.Color
}


func main() {

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup

	imgchan := make(chan point1, 1024*1024)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			wg.Add(1)

			go func(_px, _py int, _z complex128) {
				p := point1{_px, _py, mandelbrot1(_z)}
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
	// why out of memory? width, height =1024, imgchan := make(chan point1)
	// 因为每个goroutine都有一个单独的栈，栈初始大小为2KB,会按需扩展或者缩减。那么如果imgchan使用无缓存chan,
	// 就意味着，内存最少需要1024*1024*2KiB=2GiB的内存。但是运行下来发现，占用内存比2G高得多，达到6G多。
	// 解决方法：1. 扩大chan缓存，比如1024*1024  2. 约束同时进行的goroutine个数


	f, _ := os.Create("./xx.png")     //创建文件
	defer f.Close()
	png.Encode(f, img) // NOTE: ignoring errors
	fmt.Println("Success!")

}

func mandelbrot1(z complex128) color.Color {
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
