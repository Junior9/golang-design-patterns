package main

import "fmt"

type Topico interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	Observers []Observer
	name      string
	avaible   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvariable() {
	fmt.Printf("Item %s is avariable\n", i.name)
	i.avaible = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, ob := range i.Observers {
		ob.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.Observers = append(i.Observers, observer)
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s avaible from client %s\n", value, eC.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3000")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	seconObserver := &EmailClient{
		id: "34ac",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(seconObserver)
	nvidiaItem.updateAvariable()
}
