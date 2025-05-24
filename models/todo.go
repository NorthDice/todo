package models

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UsersLists struct {
	Id     int
	UserId int
	ListId int
}

type ListsItems struct {
	Id     int
	ListId int
	ItemId int
}
