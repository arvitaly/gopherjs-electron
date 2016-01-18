package spec

import "github.com/arvitaly/gopherjs-electron"

import "github.com/arvitaly/gopherjs-jasmine"

var _ = jasmine.Run(func() {
	jasmine.Describe("App", func() {
		var app = electron.GetApp()
		jasmine.ItAsync("OnReady", func(done func()) {
			app.OnReady(func() {
				done()
			})
		})
		jasmine.It("GetAppPath", func() {
			jasmine.Expect(app.GetAppPath() != "").ToBeTruthy()
		})
	})
})
