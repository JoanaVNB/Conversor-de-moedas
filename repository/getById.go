package repository

import "exchange/service"

func (r Repository) Get(id int, e service.Exchange) (service.Exchange, error){
	db := r.DB.First(&e, "id = ?", id)
	if db.Error != nil{
		return service.Exchange{}, db.Error
	}
	return e, nil
}

