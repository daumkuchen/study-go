package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  ID int
  Name string
}

func createDb() (db *sql.DB, err error) {

  USER := "root"
  PASS := "pass"
  ENDPOINT := "127.0.0.1"
  PORT := "3306"
  DB_NAME := "sample_db"

  db, err = sql.Open("mysql", USER + ":" + PASS + "@tcp(" + ENDPOINT + ":" + PORT + ")/" + DB_NAME)
  return db, err
}

func selectUsersAll(db *sql.DB) {
  rows, err := db.Query("SELECT * FROM users")
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

  for rows.Next() {
    var user User
    err := rows.Scan(&user.ID, &user.Name)
    if err != nil {
      panic(err.Error())
    }
    fmt.Println(user.ID, user.Name)
  }
}

func selectUserIndex(db *sql.DB, index int) {
  var user User
  q := db.QueryRow("SELECT * FROM users WHERE id = ?", index)
  err := q.Scan(&user.ID, &user.Name)
  switch {
    case err == sql.ErrNoRows:
      fmt.Println("レコードが存在しません")
    case err != nil:
      panic(err.Error())
    default:
      fmt.Println(user.ID, user.Name)
  }
}

func insertUser(db *sql.DB, name string)  {
  stmtInsert, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
  if err != nil {
    panic(err.Error())
  }
  defer stmtInsert.Close()

  result, err := stmtInsert.Exec(name)
  if err != nil {
    panic(err.Error())
  }

  lastInsertID, err := result.LastInsertId()
  if err != nil {
    panic(err.Error())
  }
  fmt.Println(lastInsertID)
}

func updateUser(db *sql.DB, name string, index int)  {
  stmtUpdate, err := db.Prepare("UPDATE users SET name=? WHERE id=?")
  if err != nil {
    panic(err.Error())
  }
  defer stmtUpdate.Close()

  result, err := stmtUpdate.Exec(name, index)
  if err != nil {
    panic(err.Error())
  }

  rowsAffect, err := result.RowsAffected()
  if err != nil {
    panic(err.Error())
  }
  fmt.Println(rowsAffect)
}

func deleteUser(db *sql.DB, index int)  {
  stmtDelete, err := db.Prepare("DELETE FROM users WHERE id=?")
  if err != nil {
    panic(err.Error())
  }
  defer stmtDelete.Close()

  result, err := stmtDelete.Exec(index)
  if err != nil {
    panic(err.Error())
  }

  rowsAffect, err := result.RowsAffected()
  if err != nil {
    panic(err.Error())
  }
  fmt.Println(rowsAffect)
}

func main() {

  // DBの生成
  db, err := createDb()
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  // 追加
  //insertUser(db, "Dave");

  // 更新
  //updateUser(db, "Alex", 1)

  // 削除
  //deleteUser(db, 6)

  // 出力
  selectUsersAll(db);

  // 出力（順番指定）
  //selectUserIndex(db, 1);

}