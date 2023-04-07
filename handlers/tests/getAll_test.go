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

func TestGetAllUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	wantStatusCode := 200
	wantBody := `{"valorConvertido":25,"simboloMoeda":"$"}`

	mockRepository := handlers.NewMockGetAllHandler(controller)

	mockRepository.
  EXPECT().
	Execute(gomock.Any()).
  Return([]domain.Exchange{{
		ID:1, From:"BRL", To:"USD", SymbolFrom: "R$", SymbolTo: "$", ValueFrom: 100, Rate: 4, ValueConverted: 25, Converted: "R$100->$25"},
	}, nil)

	handler := handlers.NewGetAllHandler(mockRepository)

		r := gin.Default()
		r.GET("/exchanges", handler.GetAll)

		req, err := http.NewRequest(http.MethodGet,
			"/exchanges", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		gotBody := w.Body.String()


		assert.EqualValues(t, wantStatusCode, w.Code)
		assert.EqualValues(t, wantBody, gotBody)
}
