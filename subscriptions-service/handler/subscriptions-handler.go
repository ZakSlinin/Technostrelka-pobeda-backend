package handler

import (
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
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
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := g.FormFile("subscription_avatar")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		newFilename := uuid.New().String() + ext
		dst := "/app/uploads/" + newFilename

		if err := g.SaveUploadedFile(file, dst); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "failed to save uploaded file"})
			return
		}

		req.SubscriptionAvatarUrl = "/uploads/" + newFilename
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

	var req model.UpdateSubscriptionRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := g.FormFile("subscription_avatar")

	if err == nil {
		newFileName := uuid.New().String() + filepath.Ext(file.Filename)
		dst := "/app/uploads/" + newFileName
		if err := g.SaveUploadedFile(file, dst); err == nil {
			path := "/uploads/" + newFileName
			req.SubscriptionAvatarUrl = &path
		}
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
	userID, err := uuid.Parse(g.GetHeader("X-User-Id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	result, errMsg := h.subscriptionsService.GetAllUserByID(g.Request.Context(), userID)

	if errMsg != nil {
		g.JSON(http.StatusInternalServerError, errMsg)
		return
	}

	g.JSON(http.StatusOK, result)
}

func (h *SubscriptionsHandler) DeleteSubscriptionByID(g *gin.Context) {
	userID, errU := uuid.Parse(g.GetHeader("X-User-Id"))
	subscriptionID, errS := uuid.Parse(g.Param("subscription_id"))

	if errU != nil || errS != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	errMsg := h.subscriptionsService.DeleteSubscriptionByID(g.Request.Context(), subscriptionID, userID)
	if errMsg != nil {
		status := http.StatusInternalServerError
		if errMsg.Error == "SUBSCRIPTION_NOT_FOUND" {
			status = http.StatusNotFound
		}
		g.JSON(status, errMsg)
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "subscription deleted successfully"})
}

func (h *SubscriptionsHandler) GetSubscriptionByID(g *gin.Context) {
	userID, err := uuid.Parse(g.GetHeader("X-User-Id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	subID, err := uuid.Parse(g.Param("subscription_id"))
 if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription id format"})
		return
	}

	result, errMsg := h.subscriptionsService.GetSubscriptionById(g.Request.Context(), subID, userID)

	if errMsg != nil {
		g.JSON(http.StatusInternalServerError, errMsg)
		return
	}

	g.JSON(http.StatusOK, result)
}