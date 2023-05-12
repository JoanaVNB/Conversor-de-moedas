package domain

import (
	"time"
)

type Exchange struct{
	ID int	
	From string 
	To string 
	SymbolTo string 
	SymbolFrom string 
	ValueFrom float64 
	Rate float64 
	ValueConverted	float64 
	Converted	string 
	CreatedAt time.Time  
	UpdatedAt	 time.Time
}

