package repository

import "exchange/domain"

func (r Repository) GetByFrom(from string, e []domain.Exchange) ([]domain.Exchange, error){
	db := r.DB.Where("`from` = ?", from).Find(&e)
	if db.Error != nil { 
		return []domain.Exchange{{
					From:           "not found",
		}}, db.Error
	}
	return e, nil
}
