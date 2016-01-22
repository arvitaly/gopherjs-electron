package electron

import "github.com/gopherjs/gopherjs/js"

type App interface {
	OnWillQuit(listener func(event *js.Object))
	OnReady(listener func())
	GetAppPath() string
}

type _App struct {
	*js.Object
}

func (a *_App) OnReady(listener func()) {
	a.Call("on", "ready", func() {
		listener()
	})
}
func (a *_App) OnWillQuit(listener func(event *js.Object)) {
	a.Call("on", "will-quit", func(event *js.Object) {
		listener(event)
	})
}

func (a *_App) GetAppPath() string {
	return a.Call("getAppPath").String()
}
