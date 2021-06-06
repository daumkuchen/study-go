package main

import (
  "github.com/gomodule/redigo/redis"
)

type Data struct {
  Key string
  Value string
}

func Connection() redis.Conn {
  c, err := redis.Dial("tcp", "127.0.0.1:6379")
  if err != nil {
    panic(err)
  }
  return c
}

func Set(c redis.Conn, key string, value int) string {
  res, err := redis.String(c.Do("SET", key, value))
  if err != nil {
    panic(err)
  }
  return res
}

func Get(c redis.Conn, key string) int {
  res, err := redis.Int(c.Do("GET", key))
  if err != nil {
    panic(err)
  }
  return res
}

func Increment(c redis.Conn, key string) int {
  res, err := redis.Int(c.Do("INCR", key))
  if err != nil {
    panic(err)
  }
  return res
}

func Decrement(c redis.Conn, key string) int {
  res, err := redis.Int(c.Do("DECR", key))
  if err != nil {
    panic(err)
  }
  return res
}

func main() {

  // 接続
  c := Connection()
  defer c.Close()

  // データの登録
  //res_set := Set(c, "count", 1)
  //fmt.Println(res_set)

  // データの取得
  //res_get := Get(c, "count")
  //fmt.Println(res_get)

  // インクリメント
  //res_inc := Increment(c, "count")
  //fmt.Println(res_inc)

  // デクリメント
  //res_dec := Decrement(c, "count")
  //fmt.Println(res_dec)

}