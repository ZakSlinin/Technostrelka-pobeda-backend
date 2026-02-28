package service

import (
	"context"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/repository"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionsService struct {
	repo repository.SubscriptionsRepository
}

func NewSubscriptionsService(repo repository.SubscriptionsRepository) *SubscriptionsService {
	return &SubscriptionsService{
		repo: repo,
	}
}

type ErrorMessage struct {
	Error     string                 `json:"error"`
	Message   string                 `json:"message"`
	Timestamp *timestamppb.Timestamp `json:"timestamp"`
}

func (s *SubscriptionsService) Create(ctx context.Context, userID uuid.UUID, sub *model.CreateSubscriptionRequest) (*model.Subscriptions, error) {
	subscriptionID, err := uuid.NewUUID()

	subscription := &model.Subscriptions{
		SubscriptionID:        subscriptionID,
		UserID:                userID,
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

func (s *SubscriptionsService) UpdateSubscriptionByID(ctx context.Context, userID, id uuid.UUID, req *model.UpdateSubscriptionRequest) (string, *ErrorMessage) {
	err := s.repo.UpdateSubscriptionByID(ctx, id, userID, req)

	if err != nil {
		return "", &ErrorMessage{
			Error:     "UPDATE_ERROR",
			Message:   err.Error(),
			Timestamp: timestamppb.Now(),
		}
	}

	return "Subscription updated successfully", nil
}
