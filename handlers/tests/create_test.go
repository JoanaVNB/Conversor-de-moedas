package handlers_test

import (
	"exchange/domain"
	"exchange/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	exchangeMock := domain.Exchange{
		From:           "BRL",
		To:             "USD",
		SymbolFrom:       "R$",
		SymbolTo:     	"$",
		ValueFrom:      100,
		Rate:           4,
		ValueConverted: 25,
		Converted:      "R$100->$25",
	}

	wantStatusCode := 201
	wantBody := `{"valorConvertido":25,"simboloMoeda":"$"}`

	mockRepository := handlers.NewMockCreateHandler(controller)

		mockRepository.
			EXPECT().
			Execute(100.00, "BRL", "USD", 4.00).
			Return(exchangeMock, nil)

		handler := handlers.NewCreateHandler(mockRepository)

		r := gin.Default()
		r.GET("/exchanges/:amount/:from/:to/:rate", handler.Create)

		req, err := http.NewRequest(http.MethodGet,
			"/exchanges/100/BRL/USD/4", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		gotBody := w.Body.String()

		assert.EqualValues(t, wantStatusCode, w.Code)
		assert.EqualValues(t, wantBody, gotBody)
}