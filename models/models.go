package models

import (
	"time"
)

type Student struct {
	//Id_   bson.ObjectId `bson:"_id"`
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Email string `bson:"email"`
	Sex   string `bson:"sex"`
}
type Auth_args struct {
	User string `json:user`
	Pwd  string `json:Pwd`
}
type Auth struct {
	Id   int        `json:id`
	Op   string     `json:op`
	Args *Auth_args `json:args`
}

type Strategy_args struct {
	Strategy_id   int    `json:strategy_id`
	Strategy_name string `json:strategy_name`
	Action        string `json:action`
	Account       string `json:account`
}
type Strategy struct {
	Id   int            `json:id`
	Op   string         `json:op`
	Args *Strategy_args `json:args`
}

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

type User struct {
	ID       uint
	Username string
	Password string
}

type Onorder struct {
	ID               uint
	Investorid       string
	Userid           string
	Ordersysid       string
	Instrumentid     string
	Direction        string
	Offsetflag       string
	Limitprice       float64
	Volume           uint
	Orderstatus      string
	Volumetraded     uint
	Volumeremain     uint
	Volumecancled    uint
	Inserttime       string
	Localdate        time.Time
	Exchangeid       string
	Userorderlocalid string
}
type Strategy_info struct {
	Strategy_name   string    `json:"strategy_name"`
	Create_time     time.Time `json:"create_time"`
	Update_time     time.Time `json:"update_time"`
	Strategy_status string    `json:"strategy_status"`
}
