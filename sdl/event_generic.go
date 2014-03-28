// +build !windows

// On Windows, events must be handled in the thread that created
// the window, which is most probably the thread that called sdl.Init().
// Therefore. the Events channel is not provided on Windows, and you
// need to set up your own ticker and call sdl.PollEvents() from the
// main thread.

package sdl

import "time"

var events chan interface{} = make(chan interface{})

// This channel delivers SDL events. Each object received from this channel
// has one of the following types: sdl.QuitEvent, sdl.KeyboardEvent,
// sdl.MouseButtonEvent, sdl.MouseMotionEvent, sdl.ActiveEvent,
// sdl.ResizeEvent, sdl.JoyAxisEvent, sdl.JoyButtonEvent, sdl.JoyHatEvent,
// sdl.JoyBallEvent
var Events <-chan interface{} = events

// Polling interval, in milliseconds
const poll_interval_ms = 10

// Polls SDL events in periodic intervals.
// This function does not return.
func pollEvents() {
	for {
		for _, event := range PollEvents() {
			events <- event
		}
		time.Sleep(poll_interval_ms * 1e6)
	}
}

func init() {
	go pollEvents()
}
