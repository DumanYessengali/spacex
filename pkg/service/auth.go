package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"garyshker"
	"garyshker/pkg/repository"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	salt = "ztxciubnimdwefojrsih"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthService) CreateUser(user *garyshker.User) (uint64, error) {
	if len(user.Name) < 2 {
		return 0, errors.New("INITIALIZING_NAME_ERROR")
	}

	if len(user.Username) < 2 {
		return 0, errors.New("INITIALIZING_USERNAME_ERROR")
	}

	if len(user.Password) < 4 {
		return 0, errors.New("INITIALIZING_PASSWORD_ERROR")
	}
	user.Password = generatePasswordHash(user.Password)
	return a.repos.CreateUser(user)
}
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	if email != "" {
		if err := checkmail.ValidateFormat(email); err != nil {
			return false
		}
	}
	return true
}
func (a *AuthService) GetUser(usernameOrEmail, password string) (*garyshker.User, error) {
	return a.repos.GetUser(usernameOrEmail, generatePasswordHash(password), ValidateEmail(usernameOrEmail))
}

func (a *AuthService) FetchAuth(authD *garyshker.AuthDetails) (*garyshker.Auth, error) {
	return a.repos.FetchAuth(authD)
}

func (a *AuthService) DeleteAuth(authD *garyshker.AuthDetails) error {
	return a.repos.DeleteAuth(authD)
}

func (a *AuthService) CreateAuth(userId uint64) (*garyshker.Auth, error) {
	return a.repos.CreateAuth(userId)
}

func CreateToken(authD garyshker.AuthDetails) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = authD.AuthUuid
	claims["user_id"] = authD.UserId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func (a *AuthService) TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (a *AuthService) ExtractTokenAuth(r *http.Request) (*garyshker.AuthDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &garyshker.AuthDetails{
			AuthUuid: authUuid,
			UserId:   userId,
		}, nil
	}
	return nil, err
}

func (a *AuthService) SignIn(authD garyshker.AuthDetails) (string, error) {
	token, err := CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}
