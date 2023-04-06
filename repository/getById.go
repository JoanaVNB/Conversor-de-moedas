package repository

import "exchange/domain"

func (r Repository) Get(id int, e domain.Exchange) (domain.Exchange, error){
	db := r.DB.First(&e, "id = ?", id)
	if db.Error != nil{
		return domain.Exchange{}, db.Error
	}
	return e, nil
}

