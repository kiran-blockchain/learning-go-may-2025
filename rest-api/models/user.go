package models

type User struct {
    ID   int    `json:"id"` //to convert data from db to user 
    Name string `json:"name"`
}
