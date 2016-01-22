package electron

import "github.com/gopherjs/gopherjs/js"

type BrowserWindow interface {
	GetWebContents() WebContents
	LoadURL(url string, opts *map[string]interface{})
	GetTitle() string
	Close()
	Destroy()

	//Emitted when the window is going to be closed. It's emitted before the beforeunload and unload event of the DOM. Calling event.preventDefault() will cancel the close.
	OnClose(listener func(event *js.Object))
	//Emitted when the window is closed. After you have received this event you should remove the reference to the window and avoid using it anymore.
	OnClosed(listener func())
}
type _BrowserWindow struct {
	*js.Object
	WebContents WebContents
}

func (w *_BrowserWindow) GetWebContents() WebContents {
	return w.WebContents
}
func (w *_BrowserWindow) LoadURL(url string, opts *map[string]interface{}) {
	w.Call("loadURL", url, opts)
}

func (w *_BrowserWindow) GetTitle() string {
	return w.Call("getTitle").String()
}
func (w *_BrowserWindow) Destroy() {
	w.Call("destroy")
}
func (w *_BrowserWindow) Close() {
	w.Call("close")
}
func (w *_BrowserWindow) OnClose(listener func(event *js.Object)) {
	w.Call("on", "close", func(event *js.Object) {
		listener(event)
	})
}
func (w *_BrowserWindow) OnClosed(listener func()) {
	w.Call("on", "closed", func() {
		listener()
	})
}
func (w *_BrowserWindow) HookWindowMessage(message int, callback func(result ...interface{})) {

}
