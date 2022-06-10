package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Tokens struct {
	Access string
	Refresh string
}

func GenerateNewTokens(id string, credentials []string) (*Tokens, error) {
	accessToken, err := generateNewAccessToken(id, credentials)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access: accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id string, credentials []string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	expMinutes, _ := strconv.Atoi(os.Getenv("JWT_EXP_MINUTES"))

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()
	claims["connections:list"] = false
	claims["connections:add"] = false
	claims["connections:remove"] = false
	claims["connections:update"] = false
	claims["users:list"] = false
	claims["users:add"] = false
	claims["users:remove"] = false
	claims["users:update"] = false
	claims["settings:list"] = false
	claims["settings:add"] = false
	claims["settings:update"] = false
	claims["tasks:list"] = false
	claims["tasks:add"] = false
	claims["tasks:remove"] = false
	claims["tasks:update"] = false
	claims["tasks:run"] = false
	claims["tasks:stop"] = false
	claims["tasks:restart"] = false
	claims["tasks:pause"] = false
	claims["tasks:resume"] = false
	claims["tasks:logs"] = false
	claims["tasks:status"] = false
	claims["tasks:logs:download"] = false
	claims["links:list"] = false
	claims["links:add"] = false
	claims["links:remove"] = false
	claims["links:update"] = false
	claims["roles:list"] = false
	claims["roles:add"] = false
	claims["roles:remove"] = false
	claims["roles:update"] = false
	claims["permissions:list"] = false
	claims["permissions:add"] = false
	claims["permissions:remove"] = false
	claims["permissions:update"] = false
	claims["remoteconfig:list"] = false
	claims["remoteconfig:add"] = false
	claims["remoteconfig:remove"] = false
	claims["remoteconfig:update"] = false

	for _, credential := range credentials {
		claims[credential] = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}


func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expires hours count for refresh key from .env file.
	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXP_HRS"))

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
