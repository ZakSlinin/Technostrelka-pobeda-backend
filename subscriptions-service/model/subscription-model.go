package model

import (
	"github.com/google/uuid"
	"time"
)

type Subscriptions struct {
	SubscriptionID        uuid.UUID `gorm:"primaryKey;column:subscription_id" json:"subscription_id"`
	UserID                uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Name                  string    `gorm:"column:name" json:"name"`
	Cost                  float64   `gorm:"column:cost" json:"cost"`
	NextBilling           time.Time `gorm:"column:next_billing" json:"next_billing"`
	Status                bool      `gorm:"column:status" json:"status"`
	SubscriptionAvatarUrl string    `gorm:"column:subscription_avatar_url" json:"subscription_avatar_url"`
	Category              string    `gorm:"column:category" json:"category"`
	UrlService            string    `gorm:"column:url_service" json:"url_service"`
	UseInThisMonth        bool      `gorm:"column:use_in_this_month" json:"use_in_this_month"`
	CancellationLink      string    `gorm:"column:cancellation_link" json:"cancellation_link"`
}

type UpdateSubscriptionRequest struct {
	Name               *string    `json:"name"`
	Cost               *float64   `json:"cost"`
	NextBilling        *time.Time `json:"next_billing"`
	SubscriptionAvatar *string    `json:"subscription_avatar"`
	Category           *string    `json:"category"`
	UrlService         *string    `json:"url_service"`
	UseInThisMonth     *bool      `json:"use_in_this_month"`
	CancellationLink   *string    `json:"cancellation_link"`
}

type CreateSubscriptionRequest struct {
	Name                   string    `gorm:"column:name" json:"name"`
	Cost                   float64   `gorm:"column:cost" json:"cost"`
	NextBilling            time.Time `gorm:"column:next_billing" json:"next_billing"`
	Status                 bool      `gorm:"column:status" json:"status"`
	SubscriptionsAvatarUrl string    `gorm:"column:subscription_avatar_url" json:"subscription_avatar_url"`
	Category               string    `gorm:"column:category" json:"category"`
	UrlService             string    `gorm:"column:url_service" json:"url_service"`
	UseInThisMonth         bool      `gorm:"column:use_in_this_month" json:"use_in_this_month"`
	CancellationLink       string    `gorm:"column:cancellation_link" json:"cancellation_link"`
}
