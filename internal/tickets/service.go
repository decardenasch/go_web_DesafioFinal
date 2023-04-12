package tickets

import(
	"context"
)
//Creamos una interfaz para el servicio, y este tiene los metodos que va a ejecutar, con lo que nos va a retornar
type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}
//Creamos la estructura del servicio, y le pasamos el repositorio, esta es como 
//la capa de enlace entre el repositorio y el controlador
type service struct {
	repo Repository
}
//Creamos el metodo
// Esste metodo sirve para instanciar el servicio
func NewService(repo Repository) Service {
	return service{repo}
}
//Se crea el metodo para obtener el total de tickets, y se le pasa el contexto y el destino
func (s service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}
//Se crea el metodo para obtener el promedio de tickets, y se le pasa el contexto y el destino
func (s service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	//Se traen todos los tickets. Se divide la cantidad de tickets por el destino, entre la cantidad de tickets totales
	totalTickets, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return  float64(len(tickets))/float64(len(totalTickets)), nil
}