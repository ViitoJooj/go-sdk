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

// User - data validate
func (f *Config) Email(email string) error         { return f.EmailRules.Validate(email) }
func (f *Config) Password(password string) error   { return f.PasswordRules.Validate(password) }
func (f *Config) Username(username string) error   { return users.Username(username) }
func (f *Config) FullName(fullName string) error   { return users.FullName(fullName) }
func (f *Config) FirstName(firstName string) error { return users.FirstName(firstName) }
func (f *Config) LastName(lastName string) error   { return users.LastName(lastName) }
func (f *Config) Phone(phone string) error         { return users.Phone(phone) }
func (f *Config) CPF(cpf string) error             { return users.CPF(cpf) }

// Address - data validate
func (f *Config) Country(country string) error         { return addresses.Country(country) }
func (f *Config) PostalCode(postalCode string) error   { return addresses.PostalCode(postalCode) }
func (f *Config) CEP(cep string) error                 { return addresses.CEP(cep) }
func (f *Config) CEPv2(cep string) error               { return addresses.CEPv2(cep) }
func (f *Config) DDD(ddd string) error                 { return addresses.DDD(ddd) }
func (f *Config) Label(label string) error             { return addresses.Label(label) }
func (f *Config) Street(street string) error           { return addresses.Street(street) }
func (f *Config) HouseNumber(houseNumber string) error { return addresses.HouseNumber(houseNumber) }
func (f *Config) Complement(complement string) error   { return addresses.Complement(complement) }
func (f *Config) District(district string) error       { return addresses.District(district) }
func (f *Config) City(city string) error               { return addresses.City(city) }
func (f *Config) StateRegion(stateRegion string) error { return addresses.StateRegion(stateRegion) }

// Company - data validate
func (f *Config) CNPJ(cnpj string) error { return company.CNPJ(cnpj) }
func (f *Config) CorporateName(corporateName string) error {
	return company.CorporateName(corporateName)
}
func (f *Config) TradeName(tradeName string) error { return company.TradeName(tradeName) }
func (f *Config) IM(im string) error               { return company.IM(im) }
func (f *Config) IE(ie string) error               { return company.IE(ie) }
func (f *Config) CNAE(cnae string) error           { return company.CNAE(cnae) }

// Network - data validate
func (f *Config) URL(url string) error           { return networks.URL(url) }
func (f *Config) Domain(domain string) error     { return networks.Domain(domain) }
func (f *Config) Ipv4(ip string) error           { return networks.Ipv4(ip) }
func (f *Config) Ipv6(ip string) error           { return networks.Ipv6(ip) }
func (f *Config) Hostname(hostname string) error { return networks.Hostname(hostname) }

// Finance - data validate
func (f *Config) PixKey(key string) error { return finance.PixKey(key) }
func (f *Config) IBAN(iban string) error  { return finance.IBAN(iban) }
func (f *Config) Swift(code string) error { return finance.Swift(code) }

// IDs - data validate
func (f *Config) IntID(id int, table string, conn *sql.DB) error { return ids.IntID(id, table, conn) }
func (f *Config) StrID(id string, table string, conn *sql.DB) error {
	return ids.StrID(id, table, conn)
}
func (f *Config) UUID(id uuid.UUID, table string, conn *sql.DB) error {
	return ids.UUID(id, table, conn)
}
func (f *Config) UUIDv7(id string, table string, conn *sql.DB) error {
	return ids.UUIDv7(id, table, conn)
}
