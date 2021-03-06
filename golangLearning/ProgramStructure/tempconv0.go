package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZerosC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CRoF(c Celsius) Fahrenheit{ return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func init() {
	fmt.Println("DDDDDDDDDDDDDDD")
}