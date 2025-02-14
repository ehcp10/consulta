package usecase

import (
	"context"
	"github.com/ehcp10/consulta/internal/domain"
	"github.com/ehcp10/consulta/internal/repository"
)

type ClientUseCase interface {
	CreateNewClient(ctx context.Context, client *domain.Client) error
	GetClientByID(ctx context.Context, clientID int) (*domain.Client, error)
	EditClientByID(ctx context.Context, clientID int, client *domain.Client) error
	DeleteClientByID(ctx context.Context, clientID int) error
}

type clientUseCase struct {
	clientRepo repository.ClientRepository
}

func NewClientUseCase(cr repository.ClientRepository) ClientUseCase {
	return &clientUseCase{clientRepo: cr}
}

func (uc *clientUseCase) CreateNewClient(ctx context.Context, client *domain.Client) error {
	err := uc.clientRepo.Create(ctx, client)

	if err != nil {
		return err
	}

	return nil
}

func (uc *clientUseCase) GetClientByID(ctx context.Context, clientID int) (*domain.Client, error) {
	client, err := uc.clientRepo.FindByID(ctx, clientID)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (uc *clientUseCase) EditClientByID(ctx context.Context, clientID int, client *domain.Client) error {
	storedClient, err := uc.clientRepo.FindByID(ctx, clientID)

	if err != nil {
		//TODO add log trace
		return err
	}

	storedClient.BirthDate = client.BirthDate
	storedClient.FirstName = client.FirstName
	storedClient.LastName = client.LastName
	storedClient.Gender = client.Gender
	storedClient.CPF = client.CPF

	err = uc.clientRepo.Update(ctx, storedClient)

	if err != nil {
		//TODO add log trace
		return err
	}

	return nil
}

func (uc *clientUseCase) DeleteClientByID(ctx context.Context, clientID int) error {
	return nil
}
