package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.edit"] = &GuiEdit{}
}

type GuiEdit struct{}

func (t *GuiEdit) Initialize(app *app.App) {

	// Edit 1
	ed1 := gui.NewEdit(100, "")
	ed1.SetPosition(10, 10)
	app.GuiPanel().Add(ed1)

	// Edit 2
	ed2 := gui.NewEdit(200, "edit 2")
	ed2.SetPosition(10, 40)
	ed2.SetFontSize(20)
	ed2.MaxLength = 12
	ed2.SetText("max 12 chars")
	ed2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		app.Log().Info("Edit 2 OnChange:%s", ed2.Text())
	})
	app.GuiPanel().Add(ed2)

	// Edit 3
	ed3 := gui.NewEdit(250, "style changed")
	ed3.SetPosition(10, 80)
	ed3.SetFontSize(20)
	// Copy and change the default style
	styles := gui.StyleDefault().Edit
	styles.Over.BgColor = math32.Color4Name("red")
	styles.Over.FgColor = math32.Color4Name("green")
	ed3.SetStyles(&styles)
	ed3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		app.Log().Info("Edit 3 OnChange:%s", ed3.Text())
	})
	app.GuiPanel().Add(ed3)

	// Edit 4
	ed4 := gui.NewEdit(300, "edit 3")
	ed4.SetPosition(10, 120)
	ed4.SetFontSize(20)
	ed4.SetText("initial text")
	ed4.SetEnabled(false)
	ed4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		app.Log().Info("Edit 3 OnChange:%s", ed4.Text())
	})
	app.GuiPanel().Add(ed4)
	// Edit 4 - Checkbox control
	cb1 := gui.NewCheckBox("Enabled")
	cb1.SetPosition(350, 120)
	cb1.SetValue(false)
	cb1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		app.Log().Info("checkbox value:%v", cb1.Value())
		ed4.SetEnabled(cb1.Value())
	})
	app.GuiPanel().Add(cb1)
}

func (t *GuiEdit) Render(app *app.App) {
}
