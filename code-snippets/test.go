package logger

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New good practice to create a factory function to create Logger. Important
// to use a buffered channel so we can implement the "Drop Pattern". So we
// have a cap parameter.
func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	// The goroutine to consume ch
	// Waitgroup is to manage the below goroutine
	l.wg.Add(1)
	go func() {
		// important to note everything we are doing is just the core
		// langauge, not even the stdlib.
		for v := range l.ch {
			fmt.Fprintf(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

func (l *Logger) Stop() {
	// We can just close the channel and wait for the consuming goroutine to
	// be done.
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(s string) {
	// Non-blocking send. Because we have the default case we will not block
	// if `l.ch` is at capacity, and just drop the log message.
	select {
	case l.ch <- s:
	default:
		fmt.Println("DROP")
	}

	// ^^ Bill loves this piece of code, because how the primitives in the
	// languages are really concise but powerful for this problem.
}

// in main.go
func main() {
	// Before this would fail because something would block on device, when
	// using the stdlib logger
	var d device
	l := log.New(&d, "", 0)

	// After we use our custom logger, and the app doesn't fail when device is
	// backed up.
	var d device
	l := /*logger.*/ New(d, 10)
}

type device struct{}

func (*device) Write(p []byte) (n int, err error) {
	return 0, nil
}
