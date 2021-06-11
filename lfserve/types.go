package lfserve

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Update struct {
	UpdateId int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Command struct {
	Command string `json:"command"`
	Desc    string `json:"description"`
}
