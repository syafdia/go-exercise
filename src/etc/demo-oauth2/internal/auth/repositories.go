package auth

import "context"

type Repo struct {
	User *UserRepo
}

func NewRepo() *Repo {
	return &Repo{
		User: NewUserRepo(),
	}
}

type UserRepo struct {
	stores []User
}

func NewUserRepo() *UserRepo {
	u := &UserRepo{}
	u.seed()

	return u
}

func (u *UserRepo) FindByUsernameAndPassword(_ctx context.Context, username string, password string) (User, error) {
	for _, user := range u.stores {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return User{}, ErrNotFound
}

func (u *UserRepo) seed() {
	u.stores = []User{
		{ID: "usr-0001", Username: "john123", Password: "123"},
		{ID: "usr-0002", Username: "doe456", Password: "456"},
	}
}
