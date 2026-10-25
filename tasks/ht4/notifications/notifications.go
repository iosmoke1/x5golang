package notifications

import (
	"errors"
	"fmt"

	ht3 "x5golangrepo/tasks/ht3/notifications"
)

var ErrPanic = errors.New("паника при отправке уведомления")

type Channeler interface {
	ChannelName() string
}

type PanicError struct {
	Channel string
	Value   any
	Cause   error
}

func (e *PanicError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("канал %s: паника: %v", e.Channel, e.Value)
}

func (e *PanicError) Unwrap() error {
	if e == nil || e.Cause == nil {
		return ErrPanic
	}
	return e.Cause
}

func (e *PanicError) Is(target error) bool {
	if target == ErrPanic {
		return true
	}

	other, ok := target.(*PanicError)
	return ok && other != nil && (other.Channel == "" || e.Channel == other.Channel)
}

func SafeSend(notifier ht3.Notifier, message string) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = &PanicError{
				Channel: channelName(notifier),
				Value:   value,
				Cause:   ErrPanic,
			}
		}
	}()

	return notifier.Send(message)
}

func DispatchAtBoundary(dispatch func() error) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = &PanicError{
				Channel: "dispatcher",
				Value:   value,
				Cause:   ErrPanic,
			}
		}
	}()

	return dispatch()
}

func RunAtBoundary(run func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()

	run()
	return nil
}

func channelName(notifier ht3.Notifier) string {
	if channeler, ok := notifier.(Channeler); ok {
		return channeler.ChannelName()
	}

	switch notifier.(type) {
	case ht3.EmailNotifier:
		return "email"
	case ht3.PushNotifier:
		return "push"
	default:
		return "unknown"
	}
}
