package repository

import "exchange/service"

func (r Repository) GetByFrom(from string, e []service.Exchange) ([]service.Exchange, error){
	db := r.DB.Where("`from` = ?", from).Find(&e)
	if db.Error != nil { 
		return []service.Exchange{{
					From:           "not found",
		}}, db.Error
	}
	return e, nil
}
