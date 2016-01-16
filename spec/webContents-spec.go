package spec

import (
	"time"

	"github.com/arvitaly/gopherjs-electron"
)
import jasmine "github.com/arvitaly/gopherjs-jasmine"
import "github.com/gopherjs/gopherjs/js"

var wsTests = func() bool {
	jasmine.Describe("WebContents", func() {
		var w electron.BrowserWindow
		jasmine.ItAsync("ExecuteJavascript", func(done func()) {
			w = electron.NewBrowserWindow(nil)
			w.LoadURL("file:///"+js.Global.Get("process").Call("cwd").String()+"/spec/page1.html", nil)
			w.GetWebContents().OnWillNavigate(func(event *js.Object, url string) {
				jasmine.Expect(url).ToBe("file:///url2")
				done()
			})
			w.GetWebContents().ExecuteJavaScript("location.href='file:///url2'", nil)
		})
		jasmine.It("UserAgent", func() {
			w = electron.NewBrowserWindow(nil)
			w.GetWebContents().SetUserAgent("user1")
			jasmine.Expect(w.GetWebContents().GetUserAgent()).ToBe("user1")
		})
		jasmine.AfterAllAsync(func(done func()) {
			w.Close()
			time.AfterFunc(time.Millisecond*10, done)
		})
	})
	return true
}()
