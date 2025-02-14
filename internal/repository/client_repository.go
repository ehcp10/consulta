package repository

import (
	"context"
	"github.com/ehcp10/consulta/internal/domain"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(ctx context.Context, client *domain.Client) error
	FindByID(ctx context.Context, clientID int) (*domain.Client, error)
	Update(ctx context.Context, client *domain.Client) error
	Delete(ctx context.Context, clientID string) error
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) Create(ctx context.Context, client *domain.Client) error {
	return r.db.WithContext(ctx).Create(client).Error
}

func (r *clientRepository) FindByID(ctx context.Context, clientID int) (*domain.Client, error) {
	var c domain.Client
	err := r.db.WithContext(ctx).
		Preload("Contacts").
		Preload("EmergencyContacts").
		Preload("Appointments").
		First(&c, clientID).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *clientRepository) Update(ctx context.Context, client *domain.Client) error {
	return r.db.WithContext(ctx).Save(client).Error
}

func (r *clientRepository) Delete(ctx context.Context, clientID string) error {
	return r.db.WithContext(ctx).Delete(&domain.Client{}, "client_id = ?", clientID).Error
}
