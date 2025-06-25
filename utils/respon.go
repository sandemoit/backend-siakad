package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response structures
type (
	BaseResponse struct {
		Status     int         `json:"status"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
		Page       int64       `json:"page,omitempty"`
		Limit      int64       `json:"limit,omitempty"`
		Total      int64       `json:"total,omitempty"`
		TotalPages int64       `json:"total_pages,omitempty"`
		Error      string      `json:"error,omitempty"`
	}

	PaginationParams struct {
		Page       int
		Limit      int
		Total      int
		TotalPages int
	}

	PlatResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

// ResponseError returns a standardized error response
func ResponseError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(PlatResponse{
		Status:  status,
		Message: message,
	})
}

// ResponseSuccess returns a standardized success response
func ResponseSuccess(c *fiber.Ctx, status int, message string, data ...interface{}) error {
	resp := PlatResponse{
		Status:  status,
		Message: message,
	}

	// Only set Data if data is provided
	if len(data) > 0 {
		// If only one data argument, use it directly (not as array)
		resp.Data = data[0]
	}

	return c.Status(status).JSON(resp)
}

// ResponseData returns a standardized data response with optional pagination
func ResponGetData(c *fiber.Ctx, status int, data interface{}, pagination ...PaginationParams) error {
	response := BaseResponse{
		Status:  status,
		Message: "Berhasil mendapatkan data",
	}

	// Handle pagination if provided
	if len(pagination) > 0 {
		p := pagination[0]
		if p.Page > 0 {
			response.Page = int64(p.Page)
		}
		if p.Limit > 0 {
			response.Limit = int64(p.Limit)
		}
		if p.Total > 0 {
			response.Total = int64(p.Total)
		}
		if p.TotalPages > 0 {
			response.TotalPages = int64(p.TotalPages)
		}
	}

	response.Data = data

	return c.Status(status).JSON(response)
}
