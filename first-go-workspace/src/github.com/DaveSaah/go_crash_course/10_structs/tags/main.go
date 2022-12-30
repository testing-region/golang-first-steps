package main

import (
    "fmt"
    "reflect"
)


type food struct {
    ingredients []string `required min:"3"`
    origin string
}


func main() {
    t := reflect.TypeOf(food{})
    field, _ := t.FieldByName("ingredients")
    tag := field.Tag
    fmt.Println(tag)
}
