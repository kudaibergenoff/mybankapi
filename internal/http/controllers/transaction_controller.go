package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/services"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// DebitAccount godoc
// @Summary Debit an account
// @Description Debit an account by a given amount
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Param amount path float64 true "Amount"
// @Success 200 {object} map[string]string "account debited"
// @Failure 400 {object} map[string]string "cannot parse JSON"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /accounts/{id}/debit [post]
func (ctrl *TransactionController) DebitAccount(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := ctrl.service.DebitAccount(id, req.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "account debited"})
}

// CreditAccount godoc
// @Summary Credit an account
// @Description Credit an account by a given amount
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Param amount path float64 true "Amount"
// @Success 200 {object} map[string]string "account credited"
// @Failure 400 {object} map[string]string "cannot parse JSON"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /accounts/{id}/credit [post]
func (ctrl *TransactionController) CreditAccount(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := ctrl.service.CreditAccount(id, req.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "account credited"})
}

// TransferFunds godoc
// @Summary Transfer funds between accounts
// @Description Transfer funds from one account to another
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param from_account_id path string true "From Account ID"
// @Param to_account_id path string true "To Account ID"
// @Param amount path float64 true "Amount"
// @Success 200 {object} map[string]string "transfer successful"
// @Failure 400 {object} map[string]string "cannot parse JSON"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /transfer [post]
func (ctrl *TransactionController) TransferFunds(c *fiber.Ctx) error {
	var req struct {
		FromAccountID uuid.UUID `json:"from_account_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		Amount        float64   `json:"amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := ctrl.service.TransferFunds(req.FromAccountID, req.ToAccountID, req.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "transfer successful"})
}
