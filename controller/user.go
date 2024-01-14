package controller

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct{}

func (u User) PageHome(ctx *gin.Context) {
	var users []struct {
		models.User
		FileName string `gorm:"column:file_name"`
	}

	db.Conn.Table("users").Select("users.*, file_storage_systems.file_name").
		Joins("INNER JOIN file_storage_systems ON users.id = file_storage_systems.ref_id").
		Where("file_storage_systems.ref_table = ?", "users").
		Where("users.deleted_at IS NULL").
		Scan(&users)

	ctx.HTML(http.StatusOK, "home-user.html", gin.H{"Users": users})
}

func (u User) PageCreate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create-user.html", gin.H{})
}

func (u User) Create(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	file, _ := ctx.FormFile("profile")

	var user models.User
	if db.Conn.Where("username = ? OR password = ?", username, password).First(&user).RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "username or password is exiting.."})
		return
	}

	user = models.User{
		Username: username,
		Password: password,
		Email:    username + "@gmail.com",
		Profile:  "profile",
	}

	if err := db.Conn.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "create user error.. " + err.Error()})
		return
	}

	fileExtension := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + fileExtension
	fileSystem := models.FileStorageSystem{
		FileName:      fileName,
		FileSize:      file.Size,
		FileExtension: fileExtension,
		RefID:         user.ID,
		RefTable:      "users",
		RefField:      "profile",
	}

	ctx.SaveUploadedFile(file, "./assets/uploads/"+fileName)
	db.Conn.Create(&fileSystem)

	ctx.JSON(http.StatusOK, gin.H{"data": user.ID, "msg": "create user success.. "})
}

func (u User) PageUpdate(ctx *gin.Context) {
	var user models.User
	userId := ctx.Param("id")

	if db.Conn.Where("id = ?", userId).First(&user).RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "user not found.."})
		return
	}

	var userUpdate struct {
		models.User
		FileID   uint   `gorm:"column:file_storage_id"`
		FileName string `gorm:"column:file_name"`
	}

	db.Conn.Model(&models.User{}).
		Joins("INNER JOIN file_storage_systems AS f ON users.id = f.ref_id").
		Where("users.id = ?", userId).
		Select("users.*, f.id as file_storage_id, f.file_name").
		First(&userUpdate)

	ctx.HTML(http.StatusOK, "update-user.html", gin.H{"User": userUpdate})
}

func (u User) Update(ctx *gin.Context) {
	var user models.User

	userId := ctx.Param("id")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fileId := ctx.PostForm("fileId")
	file, _ := ctx.FormFile("profile")

	if db.Conn.Where("username = ? OR password = ?", username, password).First(&user).RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "username or password is exiting.."})
		return
	}

	updates := map[string]interface{}{
		"Username":      username,
		"Password":      password,
		"FirstName":     "",
		"LastName":      "",
		"Email":         username + "@gmail.com",
		"Address":       "",
		"PhoneNumber":   "",
		"Profile":       "profile",
		"RememberToken": "",
		"CreatedAt":     time.Now(),
		"UpdatedAt":     time.Now(),
	}

	if err := db.Conn.Model(&models.User{}).Where("id = ?", userId).Updates(updates).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "update user error: " + err.Error()})
		return
	}

	fileExtension := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + fileExtension

	updateFile := map[string]interface{}{
		"FileName":      fileName,
		"FileSize":      file.Size,
		"FileType":      "",
		"FileExtension": fileExtension,
		"Content":       "",
		"RefID":         userId,
		"RefTable":      "users",
		"RefField":      "profile",
		"UpdatedAt":     time.Now(),
	}

	if err := db.Conn.Model(&models.FileStorageSystem{}).Where("id = ?", fileId).Updates(updateFile).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "update file error: " + err.Error()})
		return
	}

	ctx.SaveUploadedFile(file, "./assets/uploads/"+fileName)
	ctx.JSON(http.StatusOK, gin.H{"data": userId, "msg": "update user success.. "})
}

func (u User) Delete(ctx *gin.Context) {
	var user models.User
	var fileSystem models.FileStorageSystem

	userId := ctx.Param("id")

	if err := db.Conn.Delete(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "delete user error: " + err.Error()})
		return
	}

	db.Conn.Where("ref_id = ? AND ref_table = ?", userId, "users").First(&fileSystem)

	if err := db.Conn.Delete(&fileSystem, fileSystem.ID).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "msg": "delete file error: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userId, "fileName": fileSystem.FileName, "msg": "delete user success.. "})
}
