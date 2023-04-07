package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"exchange/domain"
	"exchange/presenter"

	"github.com/gin-gonic/gin"
)

type GetByFromHandler interface {
	Execute(string, []domain.Exchange) ([]domain.Exchange, error)
}

type getByFromHandler struct{
	repository GetByFromHandler
}

func NewGetByFromHandler (repository GetByFromHandler) *getByFromHandler{
	return &getByFromHandler{repository: repository}
}

func (ga getByFromHandler) GetByFrom(c *gin.Context){
	var e []domain.Exchange

	from := c.Param("from")
	exchanges, err := ga.repository.Execute(from, e)
	if err != nil {
		c.JSON(400, "not found")
		return
	}

	for _, eachExchange := range exchanges{
		exchange := eachExchange
		presenterExchange := *presenter.PresenterFX(exchange)
		c.JSON(200, presenterExchange)
	}
}