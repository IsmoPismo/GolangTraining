package main

import "fmt"

type vehicles interface {
}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{vehicle: vehicle{Seats: 2, MaxSpeed: 270, Color: "Yellow"}, Wheels: 4, Doors: 2}
	tacoma := car{vehicle: vehicle{Seats: 4, MaxSpeed: 230, Color: "Royalblue"}, Wheels: 4, Doors: 4}
	bmw528 := car{vehicle: vehicle{Seats: 5, MaxSpeed: 250, Color: "Darkblue"}, Wheels: 4, Doors: 4}
	boeing747 := plane{vehicle{150, 600, "White"}, false}
	boeing757 := plane{vehicle{130, 550, "White"}, false}
	boeing767 := plane{vehicle{160, 700, "White"}, true}
	sanger := boat{vehicle{5, 40, "Gray"}, 7}
	nautique := boat{vehicle{15, 29, "Black"}, 12}
	malibu := boat{vehicle{25, 33, "Indigo"}, 16}
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
