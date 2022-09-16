package entity

import "encoding/json"

type Category struct {
	tableName struct{} `pg:"categories"`

	Id    int    `pg:"id,pk" json:"id"`
	Name  string `pg:"name" json:"name"`
	Image string `pg:"image" json:"image"`
}

type Liquid struct {
	tableName struct{} `pg:"liquids"`

	Id           int    `pg:"id,pk" json:"id"`
	Name         string `pg:"name" json:"name"`
	Image        string `pg:"image" json:"image"`
	Availability bool   `pg:"avaialability" json:"availability"`
	BriefDesc    string `pg:"brief_desc" json:"brief_desc"`
	Desc         string `pg:"desc" json:"desc"`

	CategoryId int      `pg:"category_id" json:"-"`
	Category   Category `pg:"rel:has-one" json:"-"`

	Strengths []Strength `pg:"many2many:strengths" json:"strenghts"`
	Volumes   []Volume   `pg:"many2many:volumes" json:"volumes"`
	Doppings  []Dopping  `pg:"many2many:doppings" json:"doppings"`
}

type Strength struct {
	tableName struct{} `pg:"strengths"`

	LiquidId int    `pg:"liquid_id"`
	Strength string `pg:"strength"`
}

func (s *Strength) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Strength)
}

type Volume struct {
	tableName struct{} `pg:"volumes"`

	LiquidId int `pg:"liquid_id"`
	Volume   int `pg:"volume"`
}

func (v *Volume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Volume)
}

type Dopping struct {
	tableName struct{} `pg:"doppings"`

	LiquidId int    `pg:"liquid_id"`
	Dopping  string `pg:"dopping"`
}

func (d *Dopping) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Dopping)
}

type Comment struct {
	tableName struct{} `pg:"comments"`

	LiquidId int    `pg:"liquid_id" json:"-"`
	UserId   int    `pg:"user_id" json:"user_id"`
	Text     string `pg:"text" json:"text"`
	Rate     int    `pg:"rate" json:"rate"`
}

type Price struct {
	tableName struct{} `pg:"prices"`

	Roubles int `pg:"roubles" json:"roubles"`
	Cents   int `pg:"cents" json:"cents"`
}

type User struct {
	tableName struct{} `pg:"users"`

	Id int `pg:"id,pk" json:"-"`

	Name   string `pg:"name" json:"name"`
	Phone  string `pg:"phone" json:"phone"`
	VkLink string `pg:"vk_link" json:"vk_link,omitempty"`
	TgLink string `pg:"tg_link" json:"tg_link,omitempty"`
}

type AuthedUser struct {
	tableName struct{} `pg:"auth_data"`

	UserId int  `pg:"user_id"`
	User   User `pg:"rel:has-one"`

	Token    string `pg:"token"`
	Login    string `pg:"login"`
	Password string `pg:"password"`
}
