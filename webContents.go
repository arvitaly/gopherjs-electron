package electron

import "github.com/gopherjs/gopherjs/js"

type WebContents interface {
	SetUserAgent(userAgent string)
	ExecuteJavaScript(code string, UserGesture *bool)
	GetUserAgent() string
	/*Events*/
	//Corresponds to the points in time when the spinner of the tab stopped spinning.
	OnDidStopLoading(listener func())
	OnDidFailLoad(listener func(event *js.Object, errorCode int, errorDescription string, validatedURL string))
	OnWillNavigate(listener func(event *js.Object, url string))
}

type _WebContents struct {
	*js.Object
	EventEmitter
}

func (w *_WebContents) GetUserAgent() string {
	return w.Call("getUserAgent").String()
}

func (w *_WebContents) SetUserAgent(userAgent string) {
	w.Call("setUserAgent", userAgent)
}

//	HttpReferrer string //A HTTP Referrer url.
//	UserAgent    string //A user agent originating the request.
//	ExtraHeaders string //Extra headers separated by "\n"
func (w *_WebContents) LoadURL(url string, opts *map[string]interface{}) {
	w.Call("loadURL", url, opts)
}
func (w *_WebContents) ExecuteJavaScript(code string, userGesture *bool) {
	w.Call("executeJavaScript", code, userGesture)
}
func (w *_WebContents) OnDidFailLoad(listener func(event *js.Object, errorCode int, errorDescription string, validatedURL string)) {
	w.Call("on", "did-fail-load", func(event *js.Object, errorCode *js.Object, errorDescription *js.Object, validatedURL *js.Object) {
		listener(event, errorCode.Int(), errorDescription.String(), validatedURL.String())
	})
}
func (w *_WebContents) OnWillNavigate(listener func(event *js.Object, url string)) {
	w.Call("on", "will-navigate", func(event *js.Object, url *js.Object) {
		listener(event, url.String())
	})
}

func (w *_WebContents) OnDidStopLoading(listener func()) {
	w.Call("on", "did-stop-loading", func() {
		go func() {
			listener()
		}()
	})
}
