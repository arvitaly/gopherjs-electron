package electron

import "github.com/gopherjs/gopherjs/js"

type App interface {
	OnReady(listener func())
}

type _App struct {
	*js.Object
}

func (a *_App) OnReady(listener func()) {
	a.Call("on", "ready", func() {
		listener()
	})
}
