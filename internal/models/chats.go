package models

type PeopleNotConnect struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	/// CAN ADD SOMETHING ELSE
}

type UserChats struct {
	Event string     `json:"event"`
	Data  *DataChats `json:"data"`
}

type DataChats struct {
	User2 int    `json:"user_id"`
	Info  string `json:"info"`
}

type PeopleNotConnectDTO struct {
	User   *User
	People *[]PeopleNotConnect
	/// CAN ADD SOMETHING ELSE
}
