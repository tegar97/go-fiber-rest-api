package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest-api/database"
	"go-fiber-rest-api/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "tegarADad123415125512"

func Register(c *fiber.Ctx) error {
	var input map[string]string

	c.BodyParser(&input)

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input["password"]), 12)

	users := models.Users{
		Name:     input["name"],
		Email:    input["email"],
		Password: password,
	}
	database.DB.Create(&users)

	return c.JSON(users)
}

func Login(c *fiber.Ctx) error {
	var input map[string]string
	c.BodyParser(&input)
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	var user models.Users

	database.DB.Where("email = ?", input["email"]).First(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  false,
			"message": "Users Not Found",
		})

	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(input["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  false,
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  false,
			"message": "jwt error",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":   true,
		"password": "login success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})

	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.Users

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
