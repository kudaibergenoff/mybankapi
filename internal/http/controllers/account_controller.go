package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"github.com/kudaibergenoff/mybankapi/internal/services"
)

type AccountController struct {
	service *services.AccountService
}

func NewAccountController(service *services.AccountService) *AccountController {
	return &AccountController{service: service}
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account with the input payload
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body models.AccountCreateRequest true "Account"
// @Success 201 {object} models.Account
// @Failure 400 {object} map[string]string "cannot parse JSON"
// @Failure 500 {object} map[string]string "failed to create account"
// @Router /accounts [post]
func (ctrl *AccountController) CreateAccount(c *fiber.Ctx) error {
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := ctrl.service.CreateAccount(account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create account"})
	}

	return c.Status(fiber.StatusCreated).JSON(account)
}

// UpdateAccount godoc
// @Summary Update an existing account
// @Description Update an existing account with the input payload
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Param account body models.AccountCreateRequest true "Account"
// @Success 200 {object} models.Account
// @Failure 400 {object} map[string]string "cannot parse JSON"
// @Failure 500 {object} map[string]string "failed to update account"
// @Router /accounts/{id} [put]
func (ctrl *AccountController) UpdateAccount(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	account.ID = id

	if err := ctrl.service.UpdateAccount(account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update account"})
	}

	return c.JSON(account)
}

// DeleteAccount godoc
// @Summary Delete an account
// @Description Delete an account by ID
// @Tags accounts
// @Param id path string true "Account ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]string "failed to delete account"
// @Router /accounts/{id} [delete]
func (ctrl *AccountController) DeleteAccount(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	if err := ctrl.service.DeleteAccount(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete account"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetAccountByID godoc
// @Summary Get an account by ID
// @Description Get details of an account by its ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account
// @Failure 500 {object} map[string]string "account not found"
// @Router /accounts/{id} [get]
func (ctrl *AccountController) GetAccountByID(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	account, err := ctrl.service.GetAccountByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "account not found"})
	}

	return c.JSON(account)
}

// FreezeAccount godoc
// @Summary Freeze an account
// @Description Freeze an account by ID
// @Tags accounts
// @Param id path string true "Account ID"
// @Success 200 {object} map[string]string "account frozen"
// @Failure 500 {object} map[string]string "failed to freeze account"
// @Router /accounts/{id}/freeze [post]
func (ctrl *AccountController) FreezeAccount(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))

	if err := ctrl.service.FreezeAccount(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to freeze account"})
	}

	return c.JSON(fiber.Map{"message": "account frozen"})
}

// UnfreezeAccount godoc
// @Summary Unfreeze an account
// @Description Unfreeze an account by ID
// @Tags accounts
// @Param id path string true "Account ID"
// @Success 200 {object} map[string]string "account unfrozen"
// @Failure 500 {object} map[string]string "failed to unfreeze account"
// @Router /accounts/{id}/unfreeze [post]
func (ctrl *AccountController) UnfreezeAccount(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))

	if err := ctrl.service.UnfreezeAccount(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to unfreeze account"})
	}

	return c.JSON(fiber.Map{"message": "account unfrozen"})
}
