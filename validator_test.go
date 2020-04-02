package validator_test

import (
	"errors"
	"github.com/smartwalle/validator"
	"testing"
	"time"
)

var ErrEmptyName = errors.New("请输入名字")
var ErrEmptyAge = errors.New("请输入年龄")
var ErrInvalidAge = errors.New("你也太长命了吧")
var ErrEmptyNumber = errors.New("应该有学号")
var ErrEmptyTime = errors.New("应该有时间")

type Human struct {
	Name string
	Age  int
}

func (this *Human) NameValidator(n string) error {
	if n == "" {
		return ErrEmptyName
	}
	return nil
}

func (this Human) AgeValidator(a int) error {
	if a <= 0 {
		return ErrEmptyAge
	}
	if a > 100 {
		return ErrInvalidAge
	}
	return nil
}

type Student struct {
	Human  *Human
	Number int
	Time   *time.Time
}

func (this Student) NumberValidator(n int) error {
	if n <= 0 {
		return ErrEmptyNumber
	}
	return nil
}

func (this Student) TimeValidator(p *time.Time) error {
	if p == nil {
		return ErrEmptyTime
	}
	return nil
}

func TestCheck(t *testing.T) {
	var tables = []struct {
		val interface{}
		r   error
	}{{val: &Human{}, r: ErrEmptyName}}

	for _, item := range tables {
		if err := validator.Check(item.val); err != item.r {
			t.Fatal("应该得到:", item.r, "实际得到:", err)
		}
	}
}

func BenchmarkCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = &Student{Human: &Human{}}
		validator.Check(s)
	}
}
