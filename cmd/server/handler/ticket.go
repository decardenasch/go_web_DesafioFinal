package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

//Se crea la capa de conexion con el servicio, como una estrutura que tiene el servicio
type Service struct {
	service tickets.Service
}
//Se instancia el handler anterior (la estructura con el servicio)
func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

//Empezamos a crear los handlers
//Se crea el handler para obtener todos los tickets por destino, 
//Esta apunta a la conexion con el serviciio, utiliza el nombre del handler
//se utiliza el framework gin.HandlerFunc y retorna el contexto.
//Estas tienen una estructura
// 1. Obtener la peticion y validarla
	// Aqui se pasa la estructura que creamos con los tags para el request, y se utiliza
	//el ShouldBindJSON(&request) para validarla, las validaciones se deben discriminar (los errores)
	// si hay error se envia un BadRequest, sino entonces
// 2. Obtener los parametros de la peticion
	//Si se tienen parametros por url, se obtienen con el c.Param("parametro"), si se
	//tienen parametros por query se obtienen con el c.Query("parametro")
// 3. Llamar al servicio
	//Se llama al servicio con el contexto y los parametros, el servicio en este caso obtiene
	//el ticket por destino.
// 4. Retornar la respuesta
	// Se crea tambien estructuras de respuesta estas pueden ir en pkg
	//En este caso se retorna un JSON con el status 200 y los tickets con ese destino
// 5. Manejar los errores

//De la ruta se saca el parametro dest que corresponde al destino
func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
