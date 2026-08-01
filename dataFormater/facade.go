package dataFormater

import (
	"github.com/ViitoJooj/go-sdk/dataFormater/addresses"
	"github.com/ViitoJooj/go-sdk/dataFormater/company"
	"github.com/ViitoJooj/go-sdk/dataFormater/users"
)

// User - data format
func CPF(cpf string) string     { return users.CPF(cpf) }
func Phone(phone string) string { return users.Phone(phone) }

// Company - data format
func CNPJ(cnpj string) string { return company.CNPJ(cnpj) }

// Address - data format
func PostalCode(postal string) string { return addresses.PostalCode(postal) }
