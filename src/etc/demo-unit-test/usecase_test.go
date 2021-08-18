package main

import (
	"reflect"
	"testing"

	"github.com/syafdia/demo-unit-test/internal/user"
)

func TestNewRegisterUserUseCase(t *testing.T) {
	type args struct {
		userRepo user.UserRepo
	}
	tests := []struct {
		name string
		args args
		want user.RegisterUserUseCase
	}{
		{
			name: "success creating RegisterUserUseCase instance",
			args: args{},
			want: user.NewRegisterUserUseCase(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := user.NewRegisterUserUseCase(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
