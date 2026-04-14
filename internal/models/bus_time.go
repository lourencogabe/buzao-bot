package models

struct BusTime {
	ID int
	NumberLine int //chave externa
	stop string
	day string
	time string
}