package user

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_userDelivery_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		beforeTest     func(registerUserUseCase *MockRegisterUserUseCase)
		wantStatusCode int
		wantHeader     http.Header
		wantBody       string
	}{
		{
			name: "success registering new user",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://0.0.0.0/users",
					strings.NewReader(`
					{
						"first_name": "John",
						"last_name": "Doe",
						"email": "john.doe@example.com",
						"gender": "male"
					  }
					`),
				),
			},
			beforeTest: func(registerUserUseCase *MockRegisterUserUseCase) {
				registerUserUseCase.EXPECT().
					Execute(
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
			wantStatusCode: http.StatusCreated,
			wantHeader:     http.Header{"Content-Type": {"application/json"}},
			wantBody:       `{"id":1,"email":"john.doe@example.com","first_name":"John","last_name":"Doe","gender":"male"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegisterUserUC := NewMockRegisterUserUseCase(ctrl)

			u := &userDelivery{
				registerUserUseCase: mockRegisterUserUC,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockRegisterUserUC)
			}

			u.Register(tt.args.w, tt.args.r)

			rec := tt.args.w.(*httptest.ResponseRecorder)
			res := rec.Result()

			if !reflect.DeepEqual(res.StatusCode, tt.wantStatusCode) {
				t.Errorf("userDelivery.Register() = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}

			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("userDelivery.Register() = %v, want %v", res.Header, tt.wantHeader)
			}

			bodyBuffer := new(bytes.Buffer)
			bodyBuffer.ReadFrom(res.Body)
			body := strings.TrimSpace(bodyBuffer.String())

			if !reflect.DeepEqual(body, tt.wantBody) {
				t.Errorf("userDelivery.Register() = %s, want %s", bodyBuffer.String(), tt.wantBody)
			}
		})
	}
}
