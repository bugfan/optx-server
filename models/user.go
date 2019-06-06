package models

import (
	"log"
	"optx-server/utils"
	"time"
)

func init() {
	Register(new(User))
	Register(new(Admin))
}

func initDefaultUser() {
	user := User{
		Name:     "admin",
		Password: "123456",
	}
	has, _ := x.Where("name=?", user.Name).Exist(&user)
	if !has {
		x.Insert(&user)
	}
}

type User struct {
	ID       int64
	Name     string
	Password string
	JWT      string
	Created  time.Time
	Updated  time.Time
}

func (s *User) Exist() bool {
	s.Password = utils.EncryptedPassword(s.Password)
	has, _ := x.Where("name = ? and password = ?", s.Name, s.Password).Exist(s)
	return has
}

func initDefaultAdminUser() {
	admin := Admin{
		Name:     "admin",
		Password: "123456",
	}
	has, _ := x.Where("name=?", admin.Name).Exist(&admin)
	if !has {
		x.Insert(&admin)
	}
}

type Admin struct {
	ID       int64
	Name     string
	Password string
	JWT      string
	Created  time.Time
	Updated  time.Time
}

func (s *Admin) Exist() bool {
	s.Password = utils.EncryptedPassword(s.Password)
	has, _ := x.Where("name = ? and password = ?", s.Name, s.Password).Exist(s)
	log.Println()
	return has
}
