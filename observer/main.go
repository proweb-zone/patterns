package main

import "time"

func main() {

	laser := &LaserPointer{}

	// создаем слушателей
	cat1 := NewCat("vasya")
	cat2 := NewCat("seryi")
	cat3 := NewCat("rigiy")

	// регистриуем их
	laser.RegisterObserver(cat1)
	laser.RegisterObserver(cat2)
	laser.RegisterObserver(cat3)

	positions := []string{"лево", "право", "вверх", "низ"}

	for _, pos := range positions {
		laser.Move(pos)
		time.Sleep(1 * time.Second)
	}

	laser.RemoveObserver(cat2)
	laser.Move("Центр")
}

func NewCat(cat string) Observer {
	return Observer{Name: cat}
}
