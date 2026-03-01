package handler

import (
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type SubscriptionsHandler struct {
	subscriptionsService *service.SubscriptionsService
}

func NewSubscriptionsHandler(subscriptionsService *service.SubscriptionsService) *SubscriptionsHandler {
	return &SubscriptionsHandler{subscriptionsService: subscriptionsService}
}

func (h *SubscriptionsHandler) Create(g *gin.Context) {
	userIDStr := g.GetHeader("X-User-Id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	var req model.CreateSubscriptionRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.subscriptionsService.Create(g.Request.Context(), userID, &req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create subscription"})
		return
	}

	g.JSON(http.StatusCreated, result)
}

func (h *SubscriptionsHandler) UpdateSubscriptionByID(g *gin.Context) {
	userIDStr := g.GetHeader("X-User-Id")
	subscriptionIDStr := g.Param("X-Subscription-Id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	subscriptionID, err := uuid.Parse(subscriptionIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription id format"})
		return
	}

	var req model.UpdateSubscriptionRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errMsg := h.subscriptionsService.UpdateSubscriptionByID(g.Request.Context(), userID, subscriptionID, &req)

	if errMsg != nil {
		status := http.StatusInternalServerError
		if errMsg.Error == "SUBSCRIPTION_NOT_FOUND" {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}

		g.JSON(status, errMsg)
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "subscription updated successfully"})

}

func (h *SubscriptionsHandler) GetAllUserByID(g *gin.Context) {
	userIDStr := g.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	result, errMsg := h.subscriptionsService.GetAllUserByID(g.Request.Context(), userID)

	if errMsg != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	g.JSON(http.StatusOK, result)
}

func (h *SubscriptionsHandler) DeleteSubscriptionByID(g *gin.Context) {
	userIDStr := g.Param("user_id")
	subscriptionIDStr := g.Param("subscription_id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	subscriptionID, err := uuid.Parse(subscriptionIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription id format"})
		return
	}

	errMsg := h.subscriptionsService.DeleteSubscriptionByID(g.Request.Context(), userID, subscriptionID)
	if errMsg != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "subscription deleted successfully"})
}
