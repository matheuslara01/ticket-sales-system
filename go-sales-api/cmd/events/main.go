package events

import (
	"database/sql"
	"net/http"

	"github.com/matheuslara01/ticket-sales-system/go-sales-api/internal/events/infra/repository"
	"github.com/matheuslara01/ticket-sales-system/go-sales-api/internal/events/infra/service"
	"github.com/matheuslara01/ticket-sales-system/go-sales-api/internal/events/usecase"

	httpHandler "github.com/matheuslara01/ticket-sales-system/go-sales-api/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	createEventUseCase := usecase.NewCreateEventUseCase(eventRepo)
	createSpotsUseCase := usecase.NewCreateSpotsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		createEventUseCase,
		buyTicketsUseCase,
		createSpotsUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /events", eventsHandler.CreateEvent)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)
	r.HandleFunc("POST /events/{eventID}/spots", eventsHandler.CreateSpots)

	http.ListenAndServe(":8080", r)
}
