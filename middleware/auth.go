package middleware

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	log.Println("Secret : ", secret)
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(err)
		},
	})
}

func GetCredential(ctx *fiber.Ctx) (err error) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}()
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println("CREDENTIALS: ", claims)
	/*
		TODO Contoh untuk mapping credential ke user model

		credential := userModel.User{
				Branch:            claims["branch"].(string),
				Nama:              claims["nama"].(string),
				Organisasi_Unit:   claims["org_unit"].(string),
				Personal_Number:   claims["personal_number"].(string),
				Jabatan:           claims["jabatan"].(string),
				JGPG:              claims["jgpg"].(string),
				Deskripsi_Jabatan: claims["deskripsi_jabatan"].(string),
				Role:              claims["role"].(string),
			}
			ctx.Locals("credential", credential)

	*/

	return ctx.Next()
}
