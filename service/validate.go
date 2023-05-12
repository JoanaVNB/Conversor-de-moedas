package service

import (
	"errors"
	"fmt"
)

func Validate(ex Exchange) error{
	valueFrom := fmt.Sprintf("%.2f", ex.ValueFrom)
	rate := fmt.Sprintf("%.2f", ex.Rate)
	m := map[int]string{0: valueFrom, 1:ex.From, 2:ex.To, 3:rate}
	err := ValidateFX(m); if err != nil{
		return err
	}
	return nil
}

func ValidateFX(m map[int]string) error {
	if m[1] == "BRL" && m[2] == "USD" {
		return nil
	}
	if m[1] == "USD" && m[2] == "BRL" {
		return nil
	}
	if m[1] == "BRL" && m[2] == "EUR" {
		return nil
	}
	if m[1] == "EUR" && m[2] == "BRL" {
		return nil
	}
	if m[1] == "BTC" && m[2] == "USD" {
		return nil
	}
	if m[1] == "BTC" && m[2] == "BRL" {
		return nil
	}
	return errors.New("invalid converted currency")
}