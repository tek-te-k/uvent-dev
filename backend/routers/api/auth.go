package api

import (
	"net/http"
	"time"
	"uvent/database"
	"uvent/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
}

type SignupForm struct {
	Name                 string `json:"name" validate:"required"`
	Birth                string `json:"birth"`
	Address              string `json:"address"`
	Gender               string `json:"gender" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8"`
}

func Signup(c echo.Context) error {
	form := new(SignupForm)
	err := c.Bind(&form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if form.Password != form.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, "password and password_confirmation must be same")
	}
	if len(form.Password) < 8 {
		return c.JSON(http.StatusBadRequest, "password must be at least 8 characters")
	}
	// if form.Gender != "male" && form.Gender != "female" && form.Gender != "other" {
	// 	return c.JSON(http.StatusBadRequest, "gender must be 'male' or 'female' or 'other'")
	// }
	password, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 8)
	dateFormat := "2006-01-02"
	birth, err := time.Parse(dateFormat, form.Birth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "birth must be in format '2006-01-02'")
	}
	u, _ := database.DB.Model(&models.User{}).Where("email = ?", form.Email).Rows()
	if u.Next() {
		return c.JSON(http.StatusBadRequest, "email already exists")
	}
	user := models.User{
		Email:    form.Email,
		Password: password,
		Name:     form.Name,
		Birth:    birth,
		Address:  form.Address,
		Gender:   form.Gender,
	}
	database.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}

type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func Login(c echo.Context) error {
	form := new(LoginForm)
	err := c.Bind(&form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var user models.User
	res := database.DB.Where("email = ?", form.Email).First(&user)
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, "email or password is incorrect")
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(form.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "email or password is incorrect")
	}
	// JWT Claimsの発行
	claims := jwt.StandardClaims{
		Issuer:    user.Email,                            // ユーザIDをstringに変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // JWTトークンの有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	cookie := new(http.Cookie)
	cookie.Name = "email"
	cookie.Value = user.Email
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func GetUserInfo(c echo.Context) error {
	cookie, err := c.Cookie("email")
	if err != nil {
		c.JSON(http.StatusBadRequest, "bed req")
	}
	// User IDを取得
	email := cookie.Value
	var user models.User
	database.DB.Where("email = ?", email).First(&user)

	return c.JSON(http.StatusOK, user)
}
