package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

type User struct {
  Id int `json:"id"`
  Name string `json:"Name"`
}

const path = "./user.json"

func Save()  {
  user := new(User)
  user.Id = 1
  user.Name = "Alex"

  data, _ := json.Marshal(user)

  file, err := os.Create(path)
  if(err != nil){
    panic(err.Error())
  }
  defer file.Close()

  file.Write(([]byte)(string(data)))

  fmt.Println("save json success!")
}

func Load()  {
  data, err := ioutil.ReadFile(path)
  if(err != nil){
    panic(err.Error())
  }

  var user User
  json.Unmarshal(data, &user)

  fmt.Println("id:", user.Id)
  fmt.Println("name:", user.Name)
}

func main() {
  Save()
  defer Load()
}