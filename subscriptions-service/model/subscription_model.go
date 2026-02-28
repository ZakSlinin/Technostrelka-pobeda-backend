package model

import (
	"time"
)

type SubscriptionsModel struct {
	SubscriptionID        string    `json:"subscription_id"`
	UserID                string    `json:"user_id"`
	Name                  string    `json:"name"`
	Cost                  float64   `json:"cost"`
	NextBilling           time.Time `json:"next_billing"`
	Status                string    `json:"status"`
	SubscriptionAvatarUrl string    `json:"subscription_avatar_url"`
	Category              string    `json:"category"`
	UrlService            string    `json:"url_service"`
	UseInThisMonth        bool      `json:"use_in_this_month"`
}
