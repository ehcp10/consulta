package payment

import "github.com/ehcp10/consulta/core/domain/client"

type ClientCost struct {
	ID          int
	Client      client.Client
	SessionCost float64
}
