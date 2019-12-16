package models

import (
	"fmt"
	"optx-server/settings"
	"time"

	"github.com/bugfan/to"
)

func init() {
	Register(new(Options))
}

type Options struct {
	ID        int64
	Kind      int64 `xorm:"default(0)"`
	Question  string
	Answer    int64
	Desc      string
	Options   []string
	Created   time.Time
	Updated   time.Time
	DeletedAt time.Time `xorm:"deleted"`
}

type Temp struct {
	Question string
	Options  []string
	Desc     string
	Answer   int64
}

// 小程序需要的结构
type WordList struct {
	Question string  `json:"question"`
	Items    []*Item `json:"items"`
	Detail   string  `json:"detail"`
	Pic      string  `json:"pic"`
}
type Item struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

func getItems(ans []string, answer int64) (items []*Item) {
	items = make([]*Item, 0, 0)
	for i, t := range ans {
		value := "0"
		if to.Int64(i+1) == answer {
			value = "1"
		}
		items = append(items, &Item{
			Value:   value,
			Name:    t,
			Checked: false,
		})
	}
	return
}
func GetCountOptions(count int) (wl []*WordList, err error) {
	options := make([]*Options, 0, 0)
	r := "RANDOM()"
	fmt.Println("t:")
	if settings.Get("db_obj") != "sqlite3" {
		r = "RAND()"
	}
	err = GetEngine().OrderBy(r).Limit(count, 0).Find(&options)
	if err == nil {
		wl = make([]*WordList, 0, 0)
		for _, v := range options {
			items := getItems(v.Options, v.Answer)
			obj := &WordList{
				Question: v.Question,
				Detail:   v.Desc,
				Items:    items,
				Pic:      "",
			}
			wl = append(wl, obj)
		}
	}
	return
}

func GetOptionsPages(offset, limit int) (wl []*WordList, err error) {
	options := make([]*Options, 0, 0)
	err = GetEngine().Limit(limit, offset).Find(&options)
	if err == nil {
		wl = make([]*WordList, 0, 0)
		for _, v := range options {
			items := getItems(v.Options, v.Answer)
			obj := &WordList{
				Question: v.Question,
				Detail:   v.Desc,
				Items:    items,
				Pic:      "",
			}
			wl = append(wl, obj)
		}
	}
	return
}
