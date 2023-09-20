package contact

import "github.com/ehcp10/consulta/core/domain/client"

type ContactType struct {
	ID          int
	Description string
}

type Contact struct {
	ID          int
	Client      client.Client
	ContactType ContactType
	Value       string
}

type EmergenceContact struct {
	ID          int
	Client      client.Client
	ContactType ContactType
	Value       string
}
