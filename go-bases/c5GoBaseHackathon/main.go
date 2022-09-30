package main

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/file"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/service"
)

const (
	TicketsFile = "./tickets.csv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic! :%w\n", err)
		}
	}()
	// Create file struct
	f := file.NewFile(TicketsFile)
	// Get all bookings -> Test Read File
	bookings, err := f.Read()
	if err != nil {
		panic("Ay caray")
	}
	// Test of booking struct and Create
	newTicket, _ := bookings.Create(
		service.NewTicket(
			"Lucas Soria",
			"lucas.soria@mercadolibre.com",
			"spain",
			"10:30",
			220_000,
		))
	fmt.Println("Ticket created:", newTicket.ToString())
	// Test Read booking
	if readTicket, err := bookings.Read(newTicket.GetId()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket read from Booking:", readTicket.ToString())
	}
	// Update booking
	upTicket := service.NewTicket(
		"Lucas Soria",
		"lucas.soria@mercadolibre.com",
		"France",
		"10:30",
		220_000,
	)
	if updatedTicket, err := bookings.Update(newTicket.GetId(), upTicket); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket updated:", updatedTicket.ToString())
	}
	// Check booking is updated
	if readTicket, err := bookings.Read(newTicket.GetId()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket updated from booking: ", readTicket.ToString())
	}
	// Test delete
	if id, err := bookings.Delete(newTicket.GetId()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ticket deleted with id:", id)
	}
	// test writing file
	if err := f.Write(newTicket); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println("Write test (outdated ticket):", newTicket)
	}
}
