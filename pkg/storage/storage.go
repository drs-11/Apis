package storage

import "Apis/internal/utils"

type DataBase map[string]*User

func InitDB() *DataBase {
	return &DataBase{}
}

func (db *DataBase) CreateUser(username string, passwdh string) {
	user := &User{username, passwdh, map[string]*Table{}}
	(*db)[username] = user
}

func (db *DataBase) Authenticate(payload *PayLoad) bool {
	user, found := (*db)[payload.Username]
	if !found {
		return false
	}
	return utils.CheckPasswordHash(payload.Password, user.PasswordHash)
}

func (db *DataBase) CreateTable(payload *PayLoad) {
	table := &Table{payload.Tablename, map[string]*Item{}}
	user := (*db)[payload.Username]
	user.Tables[payload.Tablename] = table
}

func (db *DataBase) AddItems(payload *PayLoad) {
	items := payload.Items
	user := (*db)[payload.Username]
	table := user.Tables[payload.Tablename]
	for key, item := range items {
		table.Fields[key] = item
	}
}

func (db *DataBase) GetItems(payload *PayLoad) []*Item {
	keys := payload.Keys
	user := (*db)[payload.Username]
	table := user.Tables[payload.Tablename]
	var items []*Item
	for _, k := range keys {
		item, found := table.Fields[k]
		if found {
			items = append(items, item)
		}
	}
	return items
}

func (db *DataBase) DeleteItems(payload *PayLoad) {
	keys := payload.Keys
	user := (*db)[payload.Username]
	table := user.Tables[payload.Tablename]
	for _, k := range keys {
		_, found := table.Fields[k]
		if found {
			delete(table.Fields, k)
		}
	}
}

type User struct {
	Username     string
	PasswordHash string
	Tables       map[string]*Table
}

type Table struct {
	Name   string
	Fields map[string]*Item
}

type Item struct {
	Key string `json:"key"`
	Num int    `json:"num"`
	Bin []byte `json:"bin"`
	Str string `json:"str"`
}

// Allowed value types in Item
const (
	num = "num"
	str = "str"
	bin = "bin"
)

type PayLoad struct {
	Username  string           `json:"username"`
	Password  string           `json:"passwd"`
	Tablename string           `json:"tablename"`
	Items     map[string]*Item `json:"values"`
	Keys      []string         `json:"keys"`
}
