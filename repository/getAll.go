package repository

import "exchange/domain"

func (r Repository) GetAll (e []domain.Exchange) ([]domain.Exchange, error){
	db := r.DB.Find(&e)
	if db.Error != nil{
		return []domain.Exchange{}, db.Error
	}
	return e, nil
}