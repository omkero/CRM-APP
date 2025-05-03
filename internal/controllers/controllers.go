package controllers

import (
	services "crm_system/internal/services"
)

type Controllers struct {
	Service *services.Services
}

func NewController(services *services.Services) *Controllers {
	return &Controllers{
		Service: services,
	}
}
