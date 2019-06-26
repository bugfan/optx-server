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
		Name: "admin",
	}
	has, _ := x.Where("name=?", user.Name).Exist(&user)
	if !has {
		user.Password = utils.EncryptedPassword("123456")
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
		Name: "admin",
	}
	has, _ := x.Where("name=?", admin.Name).Exist(&admin)
	if !has {
		admin.Password = utils.EncryptedPassword("123456")
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

func ResetAdmin(user, pwd string) error {
	u := &Admin{
		Name: user,
	}
	has, _ := x.Where("name=?", u.Name).Exist(u)
	if has {
		u.JWT = ""
		u.Password = utils.EncryptedPassword(pwd)
		_, err := x.Where("name =?", u.Name).AllCols().Update(u)
		return err
	} else {
		u.Password = utils.EncryptedPassword(pwd)
		u.JWT = ""
		_, err := x.Insert(u)
		return err
	}
}
