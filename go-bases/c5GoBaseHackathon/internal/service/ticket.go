package service

import (
	"fmt"
	"reflect"
	"strconv"
)

type Ticket struct {
	id          int
	name        string
	email       string
	destination string
	date        string
	price       int
}

func NewTicket(name string, email string, destination string, date string, price int) Ticket {
	return Ticket{name: name, email: email, destination: destination, date: date, price: price}
}

func newTicketWithId(id int, name string, email string, destination string, date string, price int) Ticket {
	return Ticket{id: id, name: name, email: email, destination: destination, date: date, price: price}
}

func NewTicketFromSlice(slice []string) Ticket {
	id, errId := strconv.Atoi(slice[0])
	if errId != nil {
		panic(errId)
	}
	price, errPrice := strconv.Atoi(slice[5])
	if errPrice != nil {
		panic(errPrice)
	}
	return newTicketWithId(id, slice[1], slice[2], slice[3], slice[4], price)
}

func (t Ticket) GetId() (id int) {
	return t.id
}

func (t Ticket) ToString() (text string) {
	rTicket := reflect.ValueOf(t)
	for i := 0; i < rTicket.NumField(); i++ {
		if i == rTicket.NumField()-1 {
			text += fmt.Sprintf("%+v", rTicket.Field(i))
		} else {
			text += fmt.Sprintf("%+v,", rTicket.Field(i))
		}
	}
	return
}

func Update(t Ticket, nt Ticket) Ticket {
	t.name = nt.name
	t.email = nt.email
	t.destination = nt.destination
	t.date = nt.date
	t.price = nt.price
	return t
}
