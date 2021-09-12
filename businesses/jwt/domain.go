package jwt

import "context"

type Usecase interface {
	GenerateToken(ctx context.Context, userID int) (string, error)
}
