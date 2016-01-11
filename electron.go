package electron

import "github.com/gopherjs/gopherjs/js"

var jsElectron = js.Global.Get("require").Invoke("electron")

func GetApp() App {
	return &_App{jsElectron.Get("app")}
}

//It creates a new BrowserWindow with native properties as set by the options
//type BrowserWindowOptions struct {
//	Width                  int         //Window's width in pixels. Default is 800
//	Height                 int         //Window's height in pixels. Default is 600
//	X                      int         //Window's left offset from screen. Default is to center the window.
//	Y                      int         //Window's top offset from screen. Default is to center the window.
//	UseContentSize         bool        //The width and height would be used as web page's size, which means the actual window's size will include window frame's size and be slightly larger. Default is false.
//	Center                 bool        //Show window in the center of the screen.
//	MinWidth               int         //Window's minimum width. Default is 0.
//	MinHeight              int         //Window's minimum height. Default is 0.
//	MaxWidth               int         //Window's maximum width. Default is no limit.
//	MaxHeight              int         //Window's maximum height. Default is no limit.
//	Resizable              bool        //Whether window is resizable. Default is true.
//	AlwaysOnTop            bool        //Whether the window should always stay on top of other windows. Default is false.
//	Fullscreen             bool        //Whether the window should show in fullscreen. When set to false the fullscreen button will be hidden or disabled on OS X. Default is false.
//	SkipTaskbar            bool        //Whether to show the window in taskbar. Default is false.
//	Kiosk                  bool        //The kiosk mode. Default is false.
//	Title                  string      //Default window title. Default is "Electron".
//	Icon                   interface{} //The window icon, when omitted on Windows the executable's icon would be used as window icon.
//	Show                   bool        //Whether window should be shown when created. Default is true.
//	Frame                  bool        //Specify false to create a Frameless Window. Default is true
//	AcceptFirstMouse       bool
//	DisableAutoHideCursor  bool
//	AutoHideMenuBar        bool
//	EnableLargerThanScreen bool
//	BackgroundColor        string `js:"backgroundColor"`
//	DarkTheme              bool
//	Transparent            bool
//	Type                   string
//	TitleBarStyle          string
//	WebPreferences         interface{}
//	NodeIntegration        bool
//}
func NewBrowserWindow(opts *map[string]interface{}) BrowserWindow {
	var bw = jsElectron.Get("BrowserWindow").New(opts)
	return &_BrowserWindow{Object: bw, WebContents: &_WebContents{Object: bw.Get("webContents")}}
}
