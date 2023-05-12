package repository

import "exchange/service"

func (r Repository) GetAll (e []service.Exchange) ([]service.Exchange, error){
	db := r.DB.Find(&e)
	if db.Error != nil{
		return []service.Exchange{}, db.Error
	}
	return e, nil
}