package client

import "time"

type Client struct {
	ID            int           `json:"id,omitempty"`
	Name          string        `json:"nome"`
	BirthDate     time.Time     `json:"data_de_nascimento"`
	Occupation    Occupation    `json:"profissão"`
	Gender        Gender        `json:"sexo"`
	Age           uint          `json:"idade"`
	MaritalStatus MaritalStatus `json:"estado_civil"`
}

type Occupation struct {
	ID         int
	Occupation string
}

type Gender struct {
	ID     int
	Gender string
}

type MaritalStatus struct {
	ID            int
	MaritalStatus string
}
