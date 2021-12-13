package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

/*
//!+flagvalue
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

//!+celsiusFlag
// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct { Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}


var temp = CelsiusFlag("temp", 20.0, "the temperature")


// 练习 7.6： 对tempFlag加入支持开尔文温度。
// 摄氏度（℃）是由瑞典人摄尔修斯提出并命名的，它是以水银为测温物质，并将水的冰点定为0度，沸点定为100度，
// 将这两个固定温度之间分为100等分，每一等分就是1摄氏度。
//
// 华氏度（℉）是以水银温度计发明人华伦海特姓名所命名的，它是以冰水与氯化铵作为测温物质，将水的冰点定为32度，沸点定为212度，
// 将这两个固定温度之间分为180等分，每一等分就是1华氏度。
// 华氏度换算成摄氏度的公式是：（华氏度-32）/1.8，这样的温度标准使用并不方便，所以目前也只有巴哈马，帕劳等少数几个国家还在使用华氏度。

// 开尔文（K) 是国际温度单位，它是将水的沸点定为373.15K，冰点定为273.15K，并作为计算起点的温度，也就是273.15K与常使用的0摄氏度相等，
// 所以与我们常使用的摄氏度就是相差273.15度，如5摄氏度换算成卡尔文温度就是273.15+5。


// 练习 7.7： 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。
// 答： fmt.Println(*temp)输出时，会调用*temp的Stringer()接口的具体实现，Stringer()接口声明了String()方法。
//     *temp是Celsius的一个实例，也就是调用Celsius的String()方法，该方法会在输出后加上°C

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	flag.Parse()
	fmt.Println(*temp)
}
