package exception

import (
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/simabdi/go-jwt-auth/internal/model"
	"net"
	"net/http"
)

func NewHTTPErrorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	switch e := err.(type) {
	case *HTTPError:
		response := model.Response{
			Meta: model.Meta{
				Message: http.StatusText(e.Code),
				Code:    e.Code,
				Status:  false,
				Error:   e.Error(),
			},
			Data: nil,
		}
		return ctx.Status(e.Code).JSON(response)
	case *fiber.Error:
		response := model.Response{
			Meta: model.Meta{
				Message: http.StatusText(e.Code),
				Code:    e.Code,
				Status:  false,
				Error:   e.Error(),
			},
			Data: nil,
		}
		return ctx.Status(e.Code).JSON(response)
	case *net.OpError:
		response := model.Response{
			Meta: model.Meta{
				Message: http.StatusText(fiber.StatusInternalServerError),
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Error:   e.Error() + " " + e.Net,
			},
			Data: nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	case *mysql.MySQLError:
		response := model.Response{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: http.StatusText(fiber.StatusInternalServerError),
				Status:  false,
				Error:   e.Message + " " + e.Error(),
			},
			Data: nil,
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	default:
		response := model.Response{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: http.StatusText(fiber.StatusInternalServerError),
				Status:  false,
				Error:   e.Error(),
			},
			Data: nil,
		}
		return ctx.Status(response.Meta.Code).JSON(response)
	}
}
