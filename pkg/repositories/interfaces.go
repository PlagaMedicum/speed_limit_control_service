package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "time"
)

type Repository interface {
    Insert(ctx context.Context, data domain.SpeedInfo) error
    SelectByDate(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error)
}
