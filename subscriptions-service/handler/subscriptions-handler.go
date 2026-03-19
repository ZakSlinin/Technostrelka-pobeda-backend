package handler

import (
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/model"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
		openedFile, err := file.Open()
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
			return
		}
		defer openedFile.Close()

		fileBytes, err := io.ReadAll(openedFile)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
			return
		}

		ext := filepath.Ext(file.Filename)
		newFilename := uuid.New().String() + ext
		dst := "/app/uploads/" + newFilename

		err = os.WriteFile(dst, fileBytes, 0644)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file"})
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

	if v, ok := g.GetPostForm("name"); ok {
		req.Name = &v
	}

	if v, ok := g.GetPostForm("category"); ok {
		req.Category = &v
	}

	if v, ok := g.GetPostForm("url_service"); ok {
		req.UrlService = &v
	}

	if v, ok := g.GetPostForm("next_billing"); ok {
		req.NextBilling = &v
	}

	if v, ok := g.GetPostForm("cancellation_link"); ok {
		req.CancellationLink = &v
	}

	if v, ok := g.GetPostForm("cost"); ok {
		cost, err := strconv.ParseFloat(v, 64)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid cost"})
			return
		}
		req.Cost = &cost
	}

	if v, ok := g.GetPostForm("status"); ok {
		status, err := strconv.ParseBool(v)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
			return
		}
		req.Status = &status
	}

	if v, ok := g.GetPostForm("use_in_this_month"); ok {
		use, err := strconv.ParseBool(v)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid use_in_this_month"})
			return
		}
		req.UseInThisMonth = &use
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
