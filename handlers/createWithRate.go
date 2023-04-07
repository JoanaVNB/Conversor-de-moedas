package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"exchange/domain"
	"log"
	"strconv"
	"exchange/presenter"
	"github.com/gin-gonic/gin"
)

type CreateRHandler interface {
	Execute(float64, string, string) (domain.Exchange, error)
}

type createRHandler struct{
	repository CreateRHandler
}

func NewCreateRHandler (repository CreateRHandler) *createRHandler{
	return &createRHandler{repository: repository}
}

func (cr createRHandler) Create(c *gin.Context) {
	amount, _ := strconv.ParseFloat(c.Param("amount"), 64)
	from := c.Param("from")
	to := c.Param("to")

	ex, err := cr.repository.Execute(amount, from, to)
	if err != nil{
		log.Fatal("error to create")
	}

	presenter := *presenter.PresenterFX(ex)
	c.JSON(201, presenter)
}