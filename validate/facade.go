package validate

import (
	"database/sql"

	"github.com/ViitoJooj/go-sdk/validate/addresses"
	"github.com/ViitoJooj/go-sdk/validate/company"
	"github.com/ViitoJooj/go-sdk/validate/finance"
	"github.com/ViitoJooj/go-sdk/validate/ids"
	"github.com/ViitoJooj/go-sdk/validate/networks"
	"github.com/ViitoJooj/go-sdk/validate/users"
	"github.com/google/uuid"
)

func Email(email string) error         { return users.Email(email) }
func Password(password string) error   { return users.Password(password) }
func Username(username string) error   { return users.Username(username) }
func FullName(fullName string) error   { return users.FullName(fullName) }
func FirstName(firstName string) error { return users.FirstName(firstName) }
func LastName(lastName string) error   { return users.LastName(lastName) }
func Phone(phone string) error         { return users.Phone(phone) }
func CPF(cpf string) error             { return users.CPF(cpf) }

func Country(country string) error         { return addresses.Country(country) }
func PostalCode(postalCode string) error   { return addresses.PostalCode(postalCode) }
func CEP(cep string) error                 { return addresses.CEP(cep) }
func CEPv2(cep string) error               { return addresses.CEPv2(cep) }
func DDD(ddd string) error                 { return addresses.DDD(ddd) }
func Label(label string) error             { return addresses.Label(label) }
func Street(street string) error           { return addresses.Street(street) }
func HouseNumber(houseNumber string) error { return addresses.HouseNumber(houseNumber) }
func Complement(complement string) error   { return addresses.Complement(complement) }
func District(district string) error       { return addresses.District(district) }
func City(city string) error               { return addresses.City(city) }
func StateRegion(stateRegion string) error { return addresses.StateRegion(stateRegion) }

func CNPJ(cnpj string) error                   { return company.CNPJ(cnpj) }
func CorporateName(corporateName string) error { return company.CorporateName(corporateName) }
func TradeName(tradeName string) error         { return company.TradeName(tradeName) }
func IM(im string) error                       { return company.IM(im) }
func IE(ie string) error                       { return company.IE(ie) }
func CNAE(cnae string) error                   { return company.CNAE(cnae) }

func URL(url string) error           { return networks.URL(url) }
func Domain(domain string) error     { return networks.Domain(domain) }
func Ipv4(ip string) error           { return networks.Ipv4(ip) }
func Ipv6(ip string) error           { return networks.Ipv6(ip) }
func Hostname(hostname string) error { return networks.Hostname(hostname) }

func PixKey(key string) error { return finance.PixKey(key) }
func IBAN(iban string) error  { return finance.IBAN(iban) }
func Swift(code string) error { return finance.Swift(code) }

func IntID(id int, table string, conn *sql.DB) error      { return ids.IntID(id, table, conn) }
func StrID(id string, table string, conn *sql.DB) error   { return ids.StrID(id, table, conn) }
func UUID(id uuid.UUID, table string, conn *sql.DB) error { return ids.UUID(id, table, conn) }
func UUIDv7(id string, table string, conn *sql.DB) error  { return ids.UUIDv7(id, table, conn) }
