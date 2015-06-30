package main

import (
	"fmt"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"github.com/mswift42/gxlaunch/search"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(300, 100, "Launch")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray90))

	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
	fmt.Println(search.Search("living"))
}
