package middlewares

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/spf13/cast"
)

// middleware for context management
func ContextManagement(appTimeout time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, cancel := context.WithTimeout(context.Background(), appTimeout)
			defer cancel()
			ctx = context.WithValue(ctx, "request_id", cast.ToString(xid.NewWithTime(time.Now()).String()))

			//jwt
			if c.Get("user") != nil {
				user := c.Get("user").(*jwt.Token)
				claims := user.Claims.(*JwtCustomClaims)
				ctx = context.WithValue(ctx, "user_id", cast.ToInt(claims.UserId))
			}

			c.Set("ctx", ctx)
			return next(c)
		}
	}
}
