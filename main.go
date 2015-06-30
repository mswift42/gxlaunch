package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(500, 200, "Launch")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray90))

	searchBox := theme.CreateTextBox()
	searchBox.SetDesiredWidth(400)

	window.AddChild(searchBox)
	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
}
