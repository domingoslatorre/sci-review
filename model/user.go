package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
	"time"
)

type User struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	Role      UserRole  `db:"role" json:"role"`
	Active    bool      `db:"active" json:"active"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func NewUser(name string, email string, password string) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Id:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  string(passwordHash),
		Role:      UserReviewer,
		Active:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u User) IsAdmin() bool {
	return u.Role == UserAdmin
}

func (u *User) SetNewPassword(password string) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	u.UpdatedAt = time.Now()
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", u.Id.String()),
		slog.String("name", u.Name),
		slog.String("email", u.Email),
		slog.String("role", string(u.Role)),
		slog.Bool("active", u.Active),
		slog.Time("created_at", u.CreatedAt),
		slog.Time("updated_at", u.UpdatedAt),
	)
}
