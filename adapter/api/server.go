package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paulovitor-dock/clean-architecture-go/entity"
	"github.com/paulovitor-dock/clean-architecture-go/usecase/process_transaction"
)

type WebServer struct {
	Repository entity.TransactionRepository
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.POST("/transactions", w.process)
	e.Logger.Fatal(e.Start(":8085"))
}

func (w WebServer) process(c echo.Context) error {
	transactionDto := &process_transaction.TransactionDtoInput{}
	c.Bind(transactionDto)
	usecase := process_transaction.NewProcessTransaction(w.Repository)
	output, _ := usecase.Execute(*transactionDto)
	return c.JSON(http.StatusOK, output)
}
