package main

import (
  "fmt"

  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

type Users struct {
  Id   int `json:id`
  Name string `json:name`
}

func Connect() *gorm.DB {
  dsn := "root:pass@tcp(127.0.0.1:3306)/sample_db"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }

  return db
}

func main() {
  db := Connect()

  user := Users{}
  db.First(&user, "id=?", 1)
  fmt.Println(user)

  users := []Users{}
  db.Find(&users, "id=?", 2)
  fmt.Println(users)
}