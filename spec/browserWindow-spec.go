package spec

import (
	"time"

	"github.com/arvitaly/gopherjs-electron"
	jasmine "github.com/arvitaly/gopherjs-jasmine"
	"github.com/gopherjs/gopherjs/js"
)

var bwTest = func() bool {
	jasmine.Describe("BrowserWindow", func() {
		jasmine.It("New", func() {
			var w = electron.NewBrowserWindow(&map[string]interface{}{
				"title": "title1",
			})
			jasmine.Expect(w.GetTitle() == "title1").ToBeTruthy()
		})
		jasmine.ItAsync("LoadUrl", func(done func()) {

			var w = electron.NewBrowserWindow(&map[string]interface{}{
				"title": "title2",
			})
			w.LoadURL("file:///"+js.Global.Get("process").Call("cwd").String()+"/spec/page1.html", nil)
			w.GetWebContents().OnDidStopLoading(func() {
				jasmine.Expect(w.GetTitle()).ToBe("Page1")
				done()
			})

		})
		jasmine.AfterEachAsync(func(done func()) {
			time.AfterFunc(time.Millisecond*10, done)
		})
	})
	return true
}()
