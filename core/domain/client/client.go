package client

import "time"

type Client struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"nome"`
	BirthDate     time.Time `json:"data_de_nascimento"`
	Occupation    string    `json:"profissão"`
	Gender        string    `json:"sexo"`
	Age           uint      `json:"idade"`
	MaritalStatus string    `json:"estado_civil"`
}
