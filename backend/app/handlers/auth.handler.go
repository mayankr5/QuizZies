package handlers

import (
	"errors"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mayankr5/quizzies/app/models"
	"github.com/mayankr5/quizzies/database"
	"github.com/mayankr5/quizzies/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Identity string `json:"Identity"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
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
		return c.Status(fiber.StatusConflict).JSON(utils.APIResponse("error", "email is already taken", errors.New("Duplicating Error"), nil))
	}
	userModel, err = getUserByUsername(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on geting username", err, nil))
	} else if userModel != nil {
		return c.Status(fiber.StatusConflict).JSON(utils.APIResponse("error", "username is already taken", errors.New("Duplicating Error"), nil))
	}

	hash_pass, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on password hashing", err, nil))
	}

	user.ID = uuid.New()
	user.Password = hash_pass

	if err := database.DB.Db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on inserting user", err, nil))
	}

	userRes := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}

	token, err := "nil", nil

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on creating auth token", err, nil))
	}

	data := fiber.Map{
		"user":  userRes,
		"token": token,
	}
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
	} else {
		userRes = UserResponse{
			ID:       userModel.ID,
			Name:     userModel.Name,
			Username: userModel.Username,
			Email:    userModel.Email,
		}
	}

	if !CheckPasswordHash(userCredential.Password, userModel.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.APIResponse("error", "invalid username or password", errors.New("unauthorised user"), nil))
	}

	token, err := "utils.GenerateToken(utils.User(userRes))", nil

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.APIResponse("error", "error on creating auth token", err, nil))
	}

	data := fiber.Map{
		"user":  userRes,
		"token": token,
	}

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "user created", nil, data))
}

func Logout(c *fiber.Ctx) error {
	auth_token := c.Locals("auth_token").(models.AuthToken)
	database.DB.Db.Delete(&auth_token)

	return c.Status(fiber.StatusOK).JSON(utils.APIResponse("success", "log out successfully", nil, nil))
}
