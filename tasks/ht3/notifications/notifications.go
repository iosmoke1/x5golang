package notifications

import "fmt"

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	Channel string
}

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func NewEmailNotifier(address string) EmailNotifier {
	return EmailNotifier{
		BaseNotifier: BaseNotifier{Channel: "email"},
		Address:      address,
	}
}

func (n EmailNotifier) Send(string) error {
	return nil
}

type PushNotifier struct {
	BaseNotifier
	DeviceID string
}

func NewPushNotifier(deviceID string) PushNotifier {
	return PushNotifier{
		BaseNotifier: BaseNotifier{Channel: "push"},
		DeviceID:     deviceID,
	}
}

func (n PushNotifier) Send(string) error {
	return nil
}

type Dispatcher struct {
	notifiers []Notifier
}

func NewDispatcher(notifiers []Notifier) Dispatcher {
	return Dispatcher{notifiers: notifiers}
}

func (d Dispatcher) Dispatch(message string) error {
	for _, notifier := range d.notifiers {
		if notifier == nil {
			return fmt.Errorf("отправка nil notifier")
		}
		if err := notifier.Send(message); err != nil {
			return fmt.Errorf("отправить уведомление: %w", err)
		}
	}
	return nil
}
