package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)
// 练习 8.5： 使用一个已有的CPU绑定的顺序程序，比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，
// 并将他们的主循环改为并发形式，使用channel来进行通信。在多核计算机上这个程序得到了多少速度上的改进？
// 使用多少个goroutine是最合适的呢？


type point struct {
	x,y int
	color color.Color
}


func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 10, 10
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	imgchan := make(chan point)
	count := 0
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			count ++
			go func(_px, _py int, _z complex128) {
				p := point{_px, _py, mandelbrot(_z)}
				imgchan <- p
				fmt.Println(count)
			}(px, py, z)
			fmt.Printf("after count :%d \n", count)  // 为啥不是一个count一个after count ?
			point := <- imgchan
			img.Set(point.x, point.y, point.color)
		}
	}

	//for point := range imgchan {
	//	img.Set(point.x, point.y, point.color)
	//}

	f, _ := os.Create("./xx.png")     //创建文件
	defer f.Close()
	png.Encode(f, img) // NOTE: ignoring errors
	fmt.Println("Success!")
}

func mandelbrot(z complex128) color.Color {
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
