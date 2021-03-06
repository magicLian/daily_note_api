package services

import (
	"github.com/dgrijalva/jwt-go"
	"ml_daily_record/pkg/configs"
	"ml_daily_record/pkg/db"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/util"
	"strconv"
	"time"
)

func AuthNative(an *models.AuthNative) (string, error) {
	userQuery, err := db.PGDB.FindUserByUsernameAndPassword(an.Username, an.Password)
	if err != nil {
		return "", err
	}
	if userQuery == nil {
		return "", models.UserNotFound
	}

	now := time.Now()
	if userQuery.Token != "" {
		expiredDuation := time.Second * time.Duration(configs.Cf.Jwt.ExpiredDays) * 3600 * 24
		if userQuery.TokenCreateTime.After(now.Add(-expiredDuation)) {
			return userQuery.Token, nil
		}
	}

	tokenUser := &models.TokenUser{
		Id:            userQuery.Id,
		Username:      userQuery.Username,
		Age:           userQuery.Age,
		LastLoginTime: userQuery.LastLoginTime,
	}
	token, err := CreateToken(tokenUser, configs.Cf.Jwt.Secret)
	if err != nil {
		return "", err
	}

	userQuery.Token = token
	userQuery.TokenCreateTime = &now
	userQuery.LastLoginTime = &now
	if err = db.PGDB.UpdateUserToken(userQuery); err != nil {
		return "", err
	}

	return token, nil
}

func CreateToken(tokenUser *models.TokenUser, secret string) (string, error) {
	expiredDuation := time.Second * time.Duration(configs.Cf.Jwt.ExpiredDays) * 3600 * 24
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  strconv.FormatInt(time.Now().Add(expiredDuation).Unix(), 10),
		"user": tokenUser,
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string, secret string) (map[string]interface{}, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}

func IsTokenExpired(expiredTimestamp string) (bool, error) {
	if expiredTimestamp == "" {
		return false, models.AuthTokenExpired
	}
	expiredInt, err := strconv.ParseInt(expiredTimestamp, 10, 64)
	if err != nil {
		return false, err
	}
	expiredAt := util.ParseFromInt64(expiredInt)
	expiredDuration := time.Second * time.Duration(configs.Cf.Jwt.ExpiredDays) * 3600 * 24
	if expiredAt.Before(time.Now().Add(-expiredDuration)) {
		return false, models.AuthTokenExpired
	}
	return true, nil
}
