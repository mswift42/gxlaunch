package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	font, err := driver.CreateFont(gxfont.Default, 50)
	if err != nil {
		panic(err)
	}
	window := theme.CreateWindow(300, 100, "Launch")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray90))
	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetColor(gxui.Gray10)
	label.SetText("Launch Stuff")

	window.AddChild(label)

	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
}
