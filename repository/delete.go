package repository

import "exchange/domain"

func (r Repository) Delete (id int) (error){
	var e domain.Exchange
	db := r.DB.First(&e, "id = ?", id).Delete(&e)
	if db.Error != nil{
		return db.Error
	}
	return nil
}