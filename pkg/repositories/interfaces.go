package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "time"
)

type Repository interface {
    AddData(ctx context.Context, data domain.SpeedInfo) error
    GetInfractions(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error)
    GetBoundaries(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error)
}
