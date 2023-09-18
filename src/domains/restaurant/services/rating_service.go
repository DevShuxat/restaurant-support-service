package services

import (
	"context"

	"github.com/DevShuxat/restaurant-suppor-service/src/domains/restaurant/models"
)

type RatingService interface {
	CreateRating(ctx context.Context,ratingID, entityID string,restaurantInfo *models.RestaurantInfo,eaterInfo *models.EaterInfo,rating int,commment string) (*models.Rating, error)
	UpdateRating(ctx context.Context, ratingID string, newRating int, newComment string) (*models.Rating, error)
	ListRatingByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error)
}