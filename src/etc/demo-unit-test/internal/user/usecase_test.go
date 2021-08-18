package user

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_registerUserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewMockUserRepo(ctrl)
	fmt.Println(c)

	type args struct {
		ctx   context.Context
		input CreateUserInput
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(userRepo *MockUserRepo)
		want       User
		wantErr    bool
	}{
		{
			name: "success creating new user",
			args: args{
				ctx: context.TODO(),
				input: CreateUserInput{
					Email:     "john.doe@example.com",
					FirstName: "John",
					LastName:  "Doe",
					Gender:    GenderMale,
				},
			},
			beforeTest: func(userRepo *MockUserRepo) {
				userRepo.EXPECT().
					Create(
						context.TODO(),
						CreateUserInput{
							Email:     "john.doe@example.com",
							FirstName: "John",
							LastName:  "Doe",
							Gender:    GenderMale,
						},
					).
					Return(
						User{
							ID:        1,
							Email:     "john.doe@example.com",
							FirstName: "John",
							LastName:  "Doe",
							Gender:    GenderMale,
						},
						nil,
					)
			},
			want: User{
				ID:        1,
				Email:     "john.doe@example.com",
				FirstName: "John",
				LastName:  "Doe",
				Gender:    GenderMale,
			},
		},
		// Next test cases.
		// ...
		// ...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := NewMockUserRepo(ctrl)

			w := &registerUserUseCase{
				userRepo: mockUserRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockUserRepo)
			}

			got, err := w.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerUserUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerUserUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRegisterUserUseCase(t *testing.T) {
	type args struct {
		userRepo UserRepo
	}
	tests := []struct {
		name string
		args args
		want RegisterUserUseCase
	}{
		{
			name: "success creating RegisterUserUseCase instance",
			args: args{},
			want: NewRegisterUserUseCase(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegisterUserUseCase(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
