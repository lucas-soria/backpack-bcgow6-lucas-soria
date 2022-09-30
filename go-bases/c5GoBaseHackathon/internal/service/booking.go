package service

import "fmt"

const (
	TicketNotFound = "ticket not found"
	CantUpdate     = "problems updating ticket"
	CantDelete     = "problems deleting ticket"
)

type booking struct {
	tickets []Ticket
	lastId  int
}

func NewBookings(tickets []Ticket) Booking {
	lastId := max(tickets)
	return &booking{tickets: tickets, lastId: lastId}
}

func (b *booking) Create(t Ticket) (Ticket, error) {
	b.lastId += 1
	t.id = b.lastId
	b.tickets = append(b.tickets, t)
	return t, nil
}

func (b *booking) Read(id int) (t Ticket, err error) {
	for _, ticket := range b.tickets {
		if ticket.id == id {
			t = ticket
			return
		}
	}
	err = fmt.Errorf("problems reading ticket:\n\t%s: ticket with id %d", TicketNotFound, id)
	return
}

func (b *booking) Update(id int, t Ticket) (ticket Ticket, err error) {
	update, err := b.Read(id)
	if err != nil {
		err = fmt.Errorf("%s:\n\t%w", CantUpdate, err)
		return
	}
	index, _ := b.findIndex(update.id)
	b.tickets[index] = Update(update, t)
	ticket = b.tickets[index]
	return
}

func (b *booking) Delete(id int) (deletedId int, err error) {
	index, err := b.findIndex(id)
	if err != nil {
		err = fmt.Errorf("%s:\n\t%w", CantDelete, err)
		return
	}
	copy(b.tickets[index:], b.tickets[index+1:])
	b.tickets[len(b.tickets)-1] = Ticket{}
	b.tickets = b.tickets[:len(b.tickets)-1]
	deletedId = id
	return
}

func (b *booking) findIndex(id int) (index int, err error) {
	for i, ticket := range b.tickets {
		if ticket.id == id {
			index = i
			return
		}
	}
	err = fmt.Errorf("%s: ticket with id %d", TicketNotFound, id)
	return
}

func max(tickets []Ticket) (max int) {
	max = tickets[len(tickets)-1].id
	for _, ticket := range tickets {
		if ticket.id > max {
			max = ticket.id
		}
	}
	return
}
