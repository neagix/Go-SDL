package sdl

// Polls all SDL events that are currently available.
func PollEvents() []interface{} {
	events := make([]interface{}, 0)
	event := &Event{}
	for event.Poll() {
		var cookedEvent interface{}
		switch event.Type {
		case QUIT:
			cookedEvent = *(*QuitEvent)(cast(event))

		case KEYDOWN, KEYUP:
			cookedEvent = *(*KeyboardEvent)(cast(event))

		case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
			cookedEvent = *(*MouseButtonEvent)(cast(event))

		case MOUSEMOTION:
			cookedEvent = *(*MouseMotionEvent)(cast(event))

		case JOYAXISMOTION:
			cookedEvent = *(*JoyAxisEvent)(cast(event))

		case JOYBUTTONDOWN, JOYBUTTONUP:
			cookedEvent = *(*JoyButtonEvent)(cast(event))

		case JOYHATMOTION:
			cookedEvent = *(*JoyHatEvent)(cast(event))

		case JOYBALLMOTION:
			cookedEvent = *(*JoyBallEvent)(cast(event))

		case ACTIVEEVENT:
			cookedEvent = *(*ActiveEvent)(cast(event))

		case VIDEORESIZE:
			cookedEvent = *(*ResizeEvent)(cast(event))
		}
		if cookedEvent != nil {
			events = append(events, cookedEvent)
		}
		event = &Event{}
	}
	return events
}
