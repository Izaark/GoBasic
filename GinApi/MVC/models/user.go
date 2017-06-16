package models

type Users struct {
	Id    int `json:"id"`
	FirstName  string `json:"firstname" binding:"required" form:"firstName" gorethink:"firstName"`
	LastName string `json:"lastname" binding:"required" form:"lastName" gorethink:"lastName"`
	Age int `json:"age" binding:"required" form:"age" gorethink:"age"`
}
