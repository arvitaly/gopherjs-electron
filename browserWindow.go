package electron

import "github.com/gopherjs/gopherjs/js"

type BrowserWindow interface {
	GetWebContents() WebContents
	LoadURL(url string, opts *map[string]interface{})
	GetTitle() string
	Close()
	Destroy()
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

func (w *_BrowserWindow) HookWindowMessage(message int, callback func(result ...interface{})) {

}
