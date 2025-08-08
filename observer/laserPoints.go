package main

import "fmt"

type Observer struct {
	Name string
}

type LaserPointer struct {
	observers []Observer
	position  string
}

func (lp *LaserPointer) RegisterObserver(o Observer) {
	fmt.Printf("Регистрируем слушателя %s\n", o)
	lp.observers = append(lp.observers, o)
}

func (lp *LaserPointer) RemoveObserver(o Observer) {
	fmt.Printf("Слушатель %s - устал и больше не хочет слушать\n", o)
	for i, observer := range lp.observers {
		if observer == o {
			lp.observers = append(lp.observers[:i], lp.observers[i+1:]...)
			break
		}
	}
}

func (lp *LaserPointer) NotifyObservers() {
	for _, observer := range lp.observers {
		observer.Update(lp.position)
	}
}

func (lp *LaserPointer) Move(newPosition string) {
	fmt.Printf("Перемещение лазера %s\n", newPosition)
	lp.position = newPosition
	lp.NotifyObservers()
}

func (o *Observer) Update(position string) {
	fmt.Printf("Оповестил %s - %s\n", o, position)
}
