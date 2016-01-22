package spec

import (
	"time"

	"github.com/arvitaly/gopherjs-electron"
)
import jasmine "github.com/arvitaly/gopherjs-jasmine"
import "github.com/gopherjs/gopherjs/js"

var wsTests = func() bool {
	jasmine.Describe("WebContents", func() {
		jasmine.ItAsync("ExecuteJavascript", func(done func()) {
			var w = electron.NewBrowserWindow(nil)
			w.LoadURL("file:///"+js.Global.Get("process").Call("cwd").String()+"/spec/page1.html", nil)
			w.GetWebContents().OnWillNavigate(func(event *js.Object, url string) {
				jasmine.Expect(url).ToBe("file:///url2")
				w.Close()
				done()
			})
			w.GetWebContents().ExecuteJavaScript("location.href='file:///url2'", nil)
		})
		jasmine.It("UserAgent", func() {
			var w = electron.NewBrowserWindow(nil)
			w.GetWebContents().SetUserAgent("user1")
			jasmine.Expect(w.GetWebContents().GetUserAgent()).ToBe("user1")
			w.Close()
		})
		jasmine.It("OnCrashed", func() {
			var w = electron.NewBrowserWindow(&map[string]interface{}{
				"webSecurity": false,
			})
			var crashed = make(chan bool)
			w.GetWebContents().OnCrashed(func() {
				go func() {
					crashed <- true
				}()
			})
			w.LoadURL("chrome://crash/", nil)
			<-crashed
			w.Close()
		})
		jasmine.AfterAllAsync(func(done func()) {
			time.AfterFunc(time.Millisecond*10, done)
		})
	})
	return true
}()
