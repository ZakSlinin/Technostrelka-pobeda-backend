package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"gorm.io/gorm"
)

type SubscriptionsRepository interface {
	Create(ctx context.Context, sub *model.Subscriptions) (*model.Subscriptions, error)
	UpdateSubscriptionByID(id string, req *model.UpdateSubscriptionRequest)
	GetAllByUserID(userID string) (*model.Subscriptions, error)
	Delete(ctx context.Context, subscriptionID string) error
}

type PostgresSubscriptionsRepository struct {
	db *gorm.DB
}

func NewPostgresSubscriptionsRepository(db *gorm.DB) *PostgresSubscriptionsRepository {
	return &PostgresSubscriptionsRepository{db: db}
}

func (r *PostgresSubscriptionsRepository) Create(ctx context.Context, sub *model.Subscriptions) (*model.Subscriptions, error) {
	err := r.db.WithContext(ctx).Create(sub).Error

	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (r *PostgresSubscriptionsRepository) UpdateSubscriptionByID(ctx context.Context, id string, req *model.UpdateSubscriptionRequest) error {
	result := r.db.WithContext(ctx).
		Where("subscription_id = ?", id).
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
		return nil, fmt.Errorf("failed to fetch subscriptions: %w", result.Error)
	}

	return subs, nil
}

func (r *PostgresSubscriptionsRepository) Delete(ctx context.Context, subID string) error {
	result := r.db.WithContext(ctx).
		Where("subscription_id = ?", subID).
		Delete(&model.Subscriptions{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("subscription not found")
	}

	return nil
}
