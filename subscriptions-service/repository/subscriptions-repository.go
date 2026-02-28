package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionsRepository interface {
	Create(ctx context.Context, sub *model.Subscriptions) (*model.Subscriptions, error)
	UpdateSubscriptionByID(ctx context.Context, id, userID uuid.UUID, req *model.UpdateSubscriptionRequest) error
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.Subscriptions, error)
	Delete(ctx context.Context, id, userID uuid.UUID) error
}

type PostgresSubscriptionsRepository struct {
	db *gorm.DB
}

func NewPostgresSubscriptionsRepository(db *gorm.DB) *PostgresSubscriptionsRepository {
	return &PostgresSubscriptionsRepository{db: db}
}

func (r *PostgresSubscriptionsRepository) Create(ctx context.Context, sub *model.Subscriptions) (*model.Subscriptions, error) {
	err := r.db.WithContext(ctx).Model(&model.Subscriptions{}).Create(sub).Error

	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (r *PostgresSubscriptionsRepository) UpdateSubscriptionByID(ctx context.Context, id, userID uuid.UUID, req *model.UpdateSubscriptionRequest) error {
	result := r.db.WithContext(ctx).Model(&model.Subscriptions{}).
		Where("subscription_id = ? AND user_id = ?", id, userID).
		Updates(req)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("subscription not found")
	}

	return nil
}

func (r *PostgresSubscriptionsRepository) GetAllByUserID(ctx context.Context, userID string) ([]model.Subscriptions, error) {
	var subs []model.Subscriptions

	result := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&subs)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("subscriptions not found")
		}
		return nil, fmt.Errorf("failed to fetch subscriptions: %w", result.Error)
	}

	return subs, nil
}

func (r *PostgresSubscriptionsRepository) Delete(ctx context.Context, id, userID uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("subscription_id = ? and user_id = ?", id, userID).
		Delete(&model.Subscriptions{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("subscription not found")
	}

	return nil
}
