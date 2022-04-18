package service

import (
	"PicusFinalCase/src/models"
	"PicusFinalCase/src/pkg/config"
	"testing"
)

var (
	user = models.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin",
		Password:  "12345",
		Role:      0,
	}
	givenUser = models.User{
		Base:      models.Base{},
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
		Password:  "test",
		Role:      1,
	}

	jwtConfig = config.JWTConfig{
		SessionTime: 3600,
		SecretKey:   "test",
	}

	mockRepo = &mockRepository{users: []models.User{user, givenUser}}
	s        = NewAuthService(jwtConfig, mockRepo)
)

func TestAuthService_CreateUser(t *testing.T) {

	type args struct {
		user models.User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"WithValidUser_ShouldReturnString", args{models.User{}}, "token"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.CreateUser(tt.args.user); len(got) == 0 {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_LoginUser(t *testing.T) {

	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{"WithEmptyName_ShouldReturnEmptyString", args{"", "test"}, "", true},
		{"WithEmptyPassword_ShouldReturnEmptyString", args{"test", ""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SequenceInt() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := s.LoginUser(tt.args.name, tt.args.password); got != tt.want {
				t.Errorf("LoginUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockRepository struct {
	users []models.User
}

func (m *mockRepository) CreateUser(user models.User) *models.User {
	m.users = append(m.users, user)
	return &user
}

func (m *mockRepository) FindUser(name string) *models.User {

	if len(name) == 0 {
		return nil
	}

	for _, u := range m.users {
		if u.FirstName == name {
			return &u
		}
	}

	return nil
}
