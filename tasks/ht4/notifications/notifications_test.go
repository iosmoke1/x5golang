package notifications

import (
	"errors"
	"testing"
)

type panicNotifier struct {
	channel string
	value   any
}

func (n panicNotifier) Send(string) error {
	panic(n.value)
}

func (n panicNotifier) ChannelName() string {
	return n.channel
}

func TestSafeSendRecoversAtNotifierLevel(t *testing.T) {
	err := SafeSend(panicNotifier{channel: "email", value: "driver failed"}, "Заказ готов")

	var panicErr *PanicError
	if !errors.As(err, &panicErr) {
		t.Fatalf("ошибка = %v, ожидалась PanicError", err)
	}
	if panicErr.Channel != "email" {
		t.Fatalf("канал ошибки = %q, ожидался email", panicErr.Channel)
	}
	if !errors.Is(err, ErrPanic) {
		t.Fatal("PanicError не сохраняет ErrPanic через Unwrap")
	}
	if !errors.Is(err, &PanicError{Channel: "email"}) {
		t.Fatal("PanicError не сопоставляется по каналу через Is")
	}
}

func TestDispatchAtBoundaryRecoversDispatcherPanic(t *testing.T) {
	err := DispatchAtBoundary(func() error {
		panic("dispatcher failed")
	})

	var panicErr *PanicError
	if !errors.As(err, &panicErr) {
		t.Fatalf("ошибка = %v, ожидалась PanicError", err)
	}
	if panicErr.Channel != "dispatcher" {
		t.Fatalf("канал ошибки = %q, ожидался dispatcher", panicErr.Channel)
	}
}

func TestRunAtBoundaryRecoversWithoutReturningError(t *testing.T) {
	var recovered any
	recovered = RunAtBoundary(func() {
		panic("top-level failure")
	})

	if recovered != "top-level failure" {
		t.Fatalf("перехваченное значение = %v, ожидалось %q", recovered, "top-level failure")
	}
}

func TestSafeSendPreservesRegularErrors(t *testing.T) {
	wantErr := errors.New("канал недоступен")
	notifier := errorNotifier{err: wantErr}

	if err := SafeSend(notifier, "Заказ готов"); !errors.Is(err, wantErr) {
		t.Fatalf("ошибка = %v, ожидалась исходная ошибка", err)
	}
}

type errorNotifier struct {
	err error
}

func (n errorNotifier) Send(string) error {
	return n.err
}
