package electron

type EventEmitter interface {
	On(Event string, listener func(args ...interface{}))
}
