package repository

import(
	"exchange/domain"
)

func (r Repository) Create(ex domain.Exchange) (domain.Exchange, error){
	db := r.DB.Create(&ex)
	if db.Error != nil{
		return domain.Exchange{}, db.Error
	}
	return ex, nil
}