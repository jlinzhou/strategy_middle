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
	// Strategy_id   int    `json:strategy_id`
	// Strategy_name string `json:strategy_name`
	// Action        string `json:action`
	// Account       string `json:account`
	Contract string `bson:"contract"`
	Price float64 `bson:"price"`
	Volume int `baon:volume`
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
	ID int `bson:"id"`
	Instance_name string `bson:"instance_name"`
	Strategy_name   string    `bson:"strategy_name"`
	Strategy_args *Strategy_args `bson:"strategy_args"`
	Account string `bson:"account"`
	Create_time     time.Time `bson:"create_time"`
	Update_time     time.Time `bson:"update_time"`
	Strategy_status string    `bson:"strategy_status"`
}
