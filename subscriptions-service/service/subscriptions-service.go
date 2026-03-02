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

func NewSubscriptionsService(repo *repository.PostgresSubscriptionsRepository) *SubscriptionsService {
	return &SubscriptionsService{
		repo: repo,
	}
}

type ErrorMessage struct {
	Error     string                 `json:"error"`
	Message   string                 `json:"message"`
	Timestamp *timestamppb.Timestamp `json:"timestamp"`
}

func NewValidationError(message string) *ErrorMessage {
	return &ErrorMessage{
		Error:     "VALIDATION_ERROR",
		Message:   message,
		Timestamp: timestamppb.Now(),
	}
}

func NewSubscriptionNotFound() *ErrorMessage {
	return &ErrorMessage{
		Error:     "SUBSCRIPTION_NOT_FOUND",
		Message:   "Subscription not found",
		Timestamp: timestamppb.Now(),
	}
}

func (s *SubscriptionsService) Create(ctx context.Context, userID uuid.UUID, sub *model.CreateSubscriptionRequest) (*model.Subscriptions, error) {
	subscriptionID, err := uuid.NewUUID()

	subscription := &model.Subscriptions{
		SubscriptionID:        subscriptionID,
		UserID:                userID,
		Name:                  sub.Name,
		NextBilling:           sub.NextBilling.Time,
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

func (s *SubscriptionsService) UpdateSubscriptionByID(ctx context.Context, userID, id uuid.UUID, req *model.UpdateSubscriptionRequest) *ErrorMessage {
	if req.Cost != nil && *req.Cost < 0 {
		return NewValidationError("Cost cannot be negative")
	}

	err := s.repo.UpdateSubscriptionByID(ctx, id, userID, req)
	if err != nil {
		if err.Error() == "subscription not found" {
			return NewSubscriptionNotFound()
		}

		return &ErrorMessage{
			Error:     "INTERNAL_ERROR",
			Message:   err.Error(),
			Timestamp: timestamppb.Now(),
		}
	}

	return nil
}

func (s *SubscriptionsService) GetAllUserByID(ctx context.Context, userID uuid.UUID) ([]model.Subscriptions, *ErrorMessage) {
	subs, err := s.repo.GetAllByUserID(ctx, userID)

	if err != nil {
		if err.Error() == "subscription not found" {
			return nil, NewSubscriptionNotFound()
		}

		return nil, &ErrorMessage{
			Error:     "INTERNAL_ERROR",
			Message:   err.Error(),
			Timestamp: timestamppb.Now(),
		}
	}

	return subs, nil
}

func (s *SubscriptionsService) DeleteSubscriptionByID(ctx context.Context, id, userID uuid.UUID) *ErrorMessage {
	err := s.repo.Delete(ctx, id, userID)
	if err != nil {
		if err.Error() == "subscription not found" {
			return NewSubscriptionNotFound()
		}

		return &ErrorMessage{
			Error:     "INTERNAL_ERROR",
			Message:   err.Error(),
			Timestamp: timestamppb.Now(),
		}
	}

	return nil
}
