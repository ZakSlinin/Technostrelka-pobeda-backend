package service

import (
	"context"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/repository"
)

type SubscriptionsService struct {
	repo repository.SubscriptionsRepository
}

func NewSubscriptionsService(repo repository.SubscriptionsRepository) *SubscriptionsService {
	return &SubscriptionsService{
		repo: repo,
	}
}

func (s *SubscriptionsService) Create(ctx context.Context, sub *model.CreateSubscriptionRequest) (*model.Subscriptions, error) {
	subscription := &model.Subscriptions{
		SubscriptionID:        sub.SubscriptionID,
		UserID:                sub.UserID,
		Name:                  sub.Name,
		NextBilling:           sub.NextBilling,
		Status:                sub.Status,
		SubscriptionAvatarUrl: sub.SubscriptionsAvatarUrl,
		Category:              sub.Category,
		UrlService:            sub.UrlService,
		UseInThisMonth:        sub.UseInThisMonth,
		CancellationLink:      sub.CancellationLink,
	}

	createdSubscription, err := s.repo.Create(ctx, subscription)
	if err != nil {
		return nil, err
	}

	return createdSubscription, nil
}
