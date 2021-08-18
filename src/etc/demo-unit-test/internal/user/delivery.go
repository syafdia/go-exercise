package user

import (
	"encoding/json"
	"net/http"
)

type userDelivery struct {
	registerUserUseCase RegisterUserUseCase
}

func (u *userDelivery) Register(w http.ResponseWriter, r *http.Request) {
	var createUserInput CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&createUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := u.registerUserUseCase.Execute(r.Context(), createUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (u *userDelivery) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// TODO
}
