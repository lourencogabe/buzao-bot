package models

type BusTime struct {
	ID int
	NumberLine int //chave externa
	stop string
	day string
	time string
}
