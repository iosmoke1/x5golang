package notifications

import (
	"errors"
	"testing"
)

type recordingNotifier struct {
	messages []string
	err      error
}

func (n *recordingNotifier) Send(message string) error {
	n.messages = append(n.messages, message)
	return n.err
}

func TestDispatchSendsToAllNotifiers(t *testing.T) {
	email := &recordingNotifier{}
	push := &recordingNotifier{}
	dispatcher := NewDispatcher([]Notifier{email, push})

	if err := dispatcher.Dispatch("Заказ готов"); err != nil {
		t.Fatalf("Dispatch вернул неожиданную ошибку: %v", err)
	}

	for name, notifier := range map[string]*recordingNotifier{
		"email": email,
		"push":  push,
	} {
		if len(notifier.messages) != 1 {
			t.Errorf("%s получил %d сообщений, ожидалось 1", name, len(notifier.messages))
		}
		if len(notifier.messages) == 1 && notifier.messages[0] != "Заказ готов" {
			t.Errorf("%s получил сообщение %q", name, notifier.messages[0])
		}
	}
}

func TestDispatchStopsAfterError(t *testing.T) {
	wantErr := errors.New("канал недоступен")
	failed := &recordingNotifier{err: wantErr}
	notCalled := &recordingNotifier{}
	dispatcher := NewDispatcher([]Notifier{failed, notCalled})

	err := dispatcher.Dispatch("Заказ готов")
	if !errors.Is(err, wantErr) {
		t.Fatalf("ошибка Dispatch = %v, ожидалась ошибка канала", err)
	}
	if len(notCalled.messages) != 0 {
		t.Fatal("dispatcher вызвал notifier после ошибки")
	}
}

func TestConcreteNotifiersEmbedChannel(t *testing.T) {
	email := NewEmailNotifier("student@example.com")
	if email.Channel != "email" {
		t.Fatalf("канал email = %q, ожидалось %q", email.Channel, "email")
	}

	push := NewPushNotifier("device-42")
	if push.Channel != "push" {
		t.Fatalf("канал push = %q, ожидалось %q", push.Channel, "push")
	}
}
