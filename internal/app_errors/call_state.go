package app_errors

import (
	"time"
)

type CallNotFound struct{}

func (e *CallNotFound) Error() string {
	return "Звонок не найден"
}

type BeginAfterEnd struct {
	Begin time.Time
	End   time.Time
}

func (e *BeginAfterEnd) Error() string {
	return "Недопустимый диапазон дат: время начала находится после времени окончания"
}
