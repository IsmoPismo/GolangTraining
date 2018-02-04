package main

import "fmt"

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
	cars := []car{prius, tacoma, bmw528}

	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	planes := []plane{boeing747, boeing757, boeing767}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	boats := []boat{sanger, nautique, malibu}

	for key, value := range cars {
		fmt.Println(key, " - ", value)
	}

	for key, value := range planes {
		fmt.Println(key, " - ", value)
	}

	for key, value := range boats {
		fmt.Println(key, " - ", value)
	}
}
