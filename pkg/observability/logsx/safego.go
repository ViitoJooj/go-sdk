package logsx

// SafeGo runs fn in a goroutine that won't take the process down with it.
// A panic inside fn is recovered and logged instead of crashing the server.
// Use it for fire-and-forget work (welcome emails, alerts) where losing the
// task on panic is acceptable but losing the whole process is not.
//
// ponytail: in-process fire-and-forget; if these tasks must survive a restart
// or scale across instances, replace the call sites with a durable queue
// (outbox table, Redis/NATS/RabbitMQ), not more goroutines.
func SafeGo(name string, fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				Errors().Error("background task panicked", "task", name, "panic", r)
			}
		}()
		fn()
	}()
}
