package middleware

import (
	"github.com/MCPutro/Go-MyWallet/entity/web"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func CustomMiddleware(jwtService service.JwtService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		//check the request is use auth Bearer or not
		auth := c.Get(fiber.HeaderAuthorization, "xxx")
		if !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(web.Response{
				Status:  "ERROR",
				Message: "Invalid Authorization",
				Data:    nil,
			})
		}

		//validation jwt
		validateToken, err := jwtService.ValidateToken(strings.ReplaceAll(auth, "Bearer ", ""))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
				Status:  "ERROR",
				Message: err.Error(),
				Data:    nil,
			})
		}

		//if token is valid
		if validateToken.Valid {
			claims := validateToken.Claims.(jwt.MapClaims)
			ClaimData := claims["Data"]
			ClaimUID := claims["UID"]
			ClaimId := claims["Id"]
			if ClaimData == nil && ClaimUID == nil && ClaimId == nil {
				return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
					Status:  "ERROR",
					Message: "invalid tokens",
					Data:    nil,
				})
			}

			//cek header headerUserId is same with UID in jwt? need to encrypt UID in jwt
			headerUserId := c.Get("UserId", "xxx")
			headerId := strings.Split(headerUserId, "-")
			decryption := strings.ReplaceAll(helper.Decryption(ClaimData.(string)), "#", "-")
			if decryption == headerUserId && decryption == ClaimUID.(string)+"-"+ClaimId.(string) && ClaimId.(string) == headerId[len(headerId)-1] {
				//c.Request().SetBodyRaw([]byte("{\"haha update di middleware\":1}"))
				//c.Request().SetBody([]byte("{\"haha update di middleware\":2}"))
				return c.Next()
			}

			return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
				Status:  "ERROR",
				Message: "invalid tokens",
				Data:    nil,
			})

		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
				Status:  "ERROR",
				Message: "Token invalid",
				Data:    nil,
			})
		}

		//return c.Next()

	}
}
