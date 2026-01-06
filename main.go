package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/hexlet/hello-hexlet/greeting"
)

func main() {
	fmt.Println(greeting.Hello())

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

	color.RGB(255, 128, 0).Println("foreground orange")
	color.RGB(230, 42, 42).Println("foreground red")

	color.BgRGB(255, 128, 0).Println("background orange")
	color.BgRGB(230, 42, 42).Println("background red")

	color.RGB(255, 128, 0).AddBgRGB(0, 0, 0).Println("orange with black background")

	color.BgRGB(255, 128, 0).AddRGB(255, 255, 255).Println("orange background with white foreground")

}
