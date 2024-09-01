package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hajipy/go_todo_app/entity"
)

func User(u *entity.User) *entity.User {
	now := time.Now()
	result := &entity.User{
		ID:       entity.UserID(rand.Int()),
		Name:     "hajipy" + strconv.Itoa(rand.Int())[:5],
		Password: "password",
		Role:     "admin",
		Created:  now,
		Modified: now,
	}
	if u == nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if u.Role != "" {
		result.Role = u.Role
	}
	if !u.Created.IsZero() {
		result.Created = u.Created
	}
	if !u.Modified.IsZero() {
		result.Modified = u.Modified
	}
	return result
}
