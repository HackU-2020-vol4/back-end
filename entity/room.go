package entity

type Room struct {
	ID           uint   `json:"id"`
	RoomName     string `json:"roomname"`
	RoomPassword string `json:"roompassword"`
}
