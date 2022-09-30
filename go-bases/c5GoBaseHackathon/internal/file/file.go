package file

import (
	"fmt"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/service"
	"os"
	"strings"
)

type file struct {
	path string
}

func NewFile(path string) file {
	return file{path: path}
}

func (f file) Read() (bookings service.Booking, err error) {
	text := theReadingPart(f.path)
	dataMatrix := splitText(text)
	bookings = service.NewBookings(getTicketsFromMatrix(dataMatrix))
	return
}

func (f *file) Write(ticket service.Ticket) (err error) {
	// strings.Join(reg[:], ",")
	text := theReadingPart(f.path)
	err = os.WriteFile(
		f.path,
		[]byte(text+"\n"+ticket.ToString()),
		777,
	)
	return
}

func theReadingPart(path string) string {
	textB, errRead := os.ReadFile(path)
	if errRead != nil {
		panic(fmt.Errorf("error while reading file:\n\t%w\n", errRead))
	}
	return string(textB)
}

func splitText(raw string) (dataMatrix [][]string) {
	rows := strings.Split(raw, "\n")
	for _, row := range rows {
		dataMatrix = append(dataMatrix, strings.SplitN(row, ",", 6))
	}
	return dataMatrix
}

func getTicketsFromMatrix(dataMatrix [][]string) (tickets []service.Ticket) {
	for _, row := range dataMatrix {
		tickets = append(tickets, service.NewTicketFromSlice(row))
	}
	return
}
