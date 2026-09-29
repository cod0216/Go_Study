package main

import "fmt"

type Room uint8

const(
	MasterRoom Room = 1 << iota
	BathRoom
	Kitchen
	LivingRoom
)

func SetLight(rooms, room Room) Room {
	return rooms | room
}

func ResetLight(rooms, room Room) Room {
	return rooms &^ room
}
func IsTurnon(rooms, room Room) bool {
	return rooms == room
}

func TurnonLights(rooms Room){
	if IsTurnon(rooms, MasterRoom){
		fmt.Println("Turn on MasterRoom Light")
	}
	if IsTurnon(rooms, BathRoom){
		fmt.Println("Turn on BathRoomLight")
	}
	if IsTurnon(rooms, Kitchen){
		fmt.Println("Turn on Kitchen Light")
	}
	if IsTurnon(rooms, LivingRoom) {
		fmt.Println("Turn on LivingRoom Light")
	}
}

func main() {
	var rooms Room = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)
	TurnonLights(rooms)
}