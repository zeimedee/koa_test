package handlers

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/koa_test/types"
	"github.com/zeimedee/koa_test/utils"
)

func Convert(c *fiber.Ctx) error {
	req := new(types.Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	rate, err := utils.Rates(req.Currency, req.Target)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	result := rate * req.Amount
	res := types.Res{
		Currency: req.Target,
		Amount:   math.Round(result*100) / 100,
	}
	return c.Status(200).JSON(res)
}
