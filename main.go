package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
	"github.com/mswift42/gxlaunch/search"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(500, 200, "Launch")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray10))

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	searchBox := theme.CreateTextBox()
	searchBox.SetDesiredWidth(500)
	searchBox.SetMargin(math.Spacing{L: 4, T: 2, R: 4, B: 2})

	layout.AddChild(searchBox)

	adapter := gxui.CreateDefaultAdapter()

	searchBox.OnKeyDown(func(ev gxui.KeyboardEvent) {
		res := search.Search(searchBox.Text())
		adapter.SetItems(res.NameList())
	})

	droplist := theme.CreateDropDownList()
	droplist.SetAdapter(adapter)

	layout.AddChild(droplist)

	window.AddChild(layout)
	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
}
