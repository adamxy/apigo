package main

import (
    "io"
    "log"
    "net/http"
    //"github.com/tidwall/gjson"
)
const json = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}`

func helloHandler(w http.ResponseWriter, r *http.Request) {
    //result := gjson.Get(json, "friends")
    io.WriteString(w, json)
    //result.ForEach(func(key, value gjson.Result) bool {
    //    println(value.String())
    //    return true // keep iterating
    //})
}

func main() {
    http.HandleFunc("/test", helloHandler)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }

}
