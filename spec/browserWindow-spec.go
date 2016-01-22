package spec

import (
	"time"

	"github.com/arvitaly/gopherjs-electron"
	jasmine "github.com/arvitaly/gopherjs-jasmine"
	"github.com/gopherjs/gopherjs/js"
)

var bwTest = func() bool {
	jasmine.Describe("BrowserWindow", func() {
		var w electron.BrowserWindow
		jasmine.It("New", func() {
			w = electron.NewBrowserWindow(&map[string]interface{}{
				"title": "title1",
			})
			jasmine.Expect(w.GetTitle() == "title1").ToBeTruthy()
		})
		jasmine.ItAsync("LoadUrl", func(done func()) {

			w = electron.NewBrowserWindow(&map[string]interface{}{
				"title": "title2",
			})
			w.LoadURL("file:///"+js.Global.Get("process").Call("cwd").String()+"/spec/page1.html", nil)
			w.GetWebContents().OnDidStopLoading(func() {
				jasmine.Expect(w.GetTitle()).ToBe("Page1")
				done()
			})

		})
		jasmine.It("OnClose and OnClosed", func() {
			var w = electron.NewBrowserWindow(nil)
			var closec = make(chan bool)
			var first = true
			w.OnClose(func(event *js.Object) {
				if first {
					event.Call("preventDefault")
				}
				first = false
				go func() {
					closec <- true
				}()

			})
			w.Close()
			<-closec
			var closed = make(chan bool)
			w.OnClosed(func() {
				go func() {
					closed <- true
				}()
			})
			w.Close()

			<-closec
			<-closed
		})
		jasmine.AfterAllAsync(func(done func()) {
			time.AfterFunc(time.Millisecond*100, func() {
				w.Close()
				done()
			})
		})
	})
	return true
}()
