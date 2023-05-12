package repository

import "exchange/service"

func (r Repository) Delete (id int) (error){
	var e service.Exchange
	db := r.DB.First(&e, "id = ?", id).Delete(&e)
	if db.Error != nil{
		return db.Error
	}
	return nil
}