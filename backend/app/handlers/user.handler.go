package handlers

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*models.User, error) {
	db := database.DB.Db
	var user models.User
	if err := db.Where(&models.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(u string) (*models.User, error) {
	db := database.DB.Db
	var user models.User
	if err := db.Where(&models.User{Username: u}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.APIResponse("error", "error on parsing user", err, nil))
	}

	userModel, err := new(models.User), *new(error)
	userModel, err = getUserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on geting email", err, nil))
	} else if userModel != nil {
		return c.Status(fiber.StatusConflict).JSON(utils.APIResponse("error", "email is already taken", errors.New("duplication Error"), nil))
	}
	userModel, err = getUserByUsername(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on geting username", err, nil))
	} else if userModel != nil {
		return c.Status(fiber.StatusConflict).JSON(utils.APIResponse("error", "username is already taken", errors.New("duplication Error"), nil))
	}

	hash_pass, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on password hashing", err, nil))
	}

	access_token, refresh_token, err := utils.GenerateAllTokens(user.Email, user.FirstName, user.LastName, user.ID.String())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on creating auth token", err, nil))
	}

	user.ID = uuid.New()
	user.Password = hash_pass
	user.AccessToken = access_token
	user.RefreshToken = refresh_token

	if err := database.DB.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on inserting user", err, nil))
	}

	userRes := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	data := fiber.Map{
		"user": userRes,
	}

	c.Cookie(&fiber.Cookie{
		Name:  "refresh_token",
		Value: *refresh_token,
		// HTTPOnly: true,
		// Secure:   true,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: *access_token,
		// HTTPOnly: true,
		// Secure:   true,
	})

	return c.Status(fiber.StatusCreated).JSON(utils.APIResponse("success", "user created", nil, data))
}

func Login(c *fiber.Ctx) error {
	var userCredential LoginRequest
	if err := c.BodyParser(&userCredential); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.APIResponse("error", "error on parsing credential", err, nil))
	}

	var userRes UserResponse
	userModel, err := new(models.User), *new(error)

	if valid(userCredential.Identity) {
		userModel, err = getUserByEmail(userCredential.Identity)
	} else {
		userModel, err = getUserByUsername(userCredential.Identity)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "invalid username or password", errors.New("unauthorised user"), nil))
	} else if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "invalid username or password", errors.New("unauthorised user"), nil))
	} else if !CheckPasswordHash(userCredential.Password, userModel.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "invalid username or password", errors.New("unauthorised user"), nil))
	} else {
		userRes = UserResponse{
			ID:        userModel.ID,
			FirstName: userModel.FirstName,
			LastName:  userModel.LastName,
			Username:  userModel.Username,
			Email:     userModel.Email,
			Phone:     userModel.Phone,
		}
	}

	access_token, refresh_token, err := utils.GenerateAllTokens(userRes.Email, userRes.FirstName, userRes.LastName, (userModel.ID).String())
	fmt.Println(userModel.ID.String())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on creating auth token", err, nil))
	}

	data := fiber.Map{
		"user": userRes,
	}

	c.Cookie(&fiber.Cookie{
		Name:  "refresh_token",
		Value: *refresh_token,
		// HTTPOnly: true,
		// Secure:   true,
		Expires: time.Now().Local().Add(time.Hour * time.Duration(168)),
	})
	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: *access_token,
		// HTTPOnly: true,
		// Secure:   true,
		Expires: time.Now().Local().Add(time.Hour * time.Duration(24)),
	})

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "user created", nil, data))
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: time.Now(),
		// HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: time.Now(),
		// HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "log out successfully", nil, nil))
}
