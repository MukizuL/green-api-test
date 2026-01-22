package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"greep-api-test/internal/dto"
	"io"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

func (c *Controller) GetSettings(ctx *gin.Context) {
	var req dto.GetSettingsRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	url := fmt.Sprintf(
		"https://1103.api.green-api.com/waInstance%s/getSettings/%s",
		req.IdInstance,
		req.ApiTokenInstance,
	)

	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Data(resp.StatusCode, "application/json", body)
}

func (c *Controller) GetStateInstance(ctx *gin.Context) {
	var req dto.GetStateInstanceRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	url := fmt.Sprintf(
		"https://1103.api.green-api.com/waInstance%s/getStateInstance/%s",
		req.IdInstance,
		req.ApiTokenInstance,
	)

	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Data(resp.StatusCode, "application/json", body)
}

func (c *Controller) SendMessage(ctx *gin.Context) {
	var req dto.SendMessageRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	url := fmt.Sprintf(
		"https://1103.api.green-api.com/waInstance%s/sendMessage/%s",
		req.IdInstance,
		req.ApiTokenInstance,
	)

	message := struct {
		ChatID  string `json:"chatId"`
		Message string `json:"message"`
	}{
		ChatID:  req.ChatID + "@c.us",
		Message: req.Message,
	}

	data, err := json.Marshal(&message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Data(resp.StatusCode, "application/json", body)
}

func (c *Controller) SendFileByUrl(ctx *gin.Context) {
	var req dto.SendFileByUrlRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	url := fmt.Sprintf(
		"https://1103.api.green-api.com/waInstance%s/sendFileByUrl/%s",
		req.IdInstance,
		req.ApiTokenInstance,
	)

	message := struct {
		ChatID   string `json:"chatId"`
		UrlFile  string `json:"urlFile"`
		Filename string `json:"fileName"`
	}{
		ChatID:   req.ChatID + "@c.us",
		UrlFile:  req.FileUrl,
		Filename: path.Base(req.FileUrl),
	}

	data, err := json.Marshal(&message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Data(resp.StatusCode, "application/json", body)
}
