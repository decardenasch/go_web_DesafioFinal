package tickets

import (
	"context"
	"fmt"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)
//Tiene el servicio de tickets, y este tiene el metodo de obtener todos los tickets, y el metodo de obtener los tickets por destino
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}
//Se crea la estructura del repositorio, y se le pasa el parametro de la base de datos, es donde tenemos
//nuestro modelo
type repository struct {
	db []domain.Ticket
}
//Esta funcion recibe el slice de la estructura y retorna el repositorio,
//pero es un puntero a la estructura
func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}
//Se realizan las implementaciones de los metodos de la interfaz
//Se observa si el slice que se llena de la estructura de tickets en el domain esta vacio
//Se observa mirando su longitud, si esta vacio, se retorna un error, sino se retorna el slice
func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

//Se recorre el slice para encontrar los tickets que coincidan con el destino
//Se va llenando un slice llamado ticketsDest si se encuentra coincidencia con el destino ingresado
func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}
