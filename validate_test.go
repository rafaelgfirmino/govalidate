package govalidate

import (
	"testing"
)
type User struct {
	Email string `validate:"required,email"`
	User2
}
type User2 struct {
	Email2 string `validate:"email,required"`
	User3
}
type User3 struct {
	Email3 string `validate:"email,required"`
}

func BenchmarkValidateStruct(b *testing.B) {
	user := User{
		"rafaelgfirmino",
		User2{
			Email2: "rafaelgfrimino.com.br",
		},
	}
	for i := 0; i < b.N; i++ {
		Struct(user)
	}
}