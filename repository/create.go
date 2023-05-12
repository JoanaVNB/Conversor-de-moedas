package repository

import(
	"exchange/service"
)

func (r Repository) Create(ex service.Exchange) (service.Exchange, error){
	db := r.DB.Create(&ex)
	if db.Error != nil{
		return service.Exchange{}, db.Error
	}
	return ex, nil
}