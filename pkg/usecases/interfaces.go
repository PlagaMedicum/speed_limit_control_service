package usecases

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "time"
)

type Usecases interface {
    AddData(ctx context.Context, data domain.SpeedInfo) error
    GetInfractions(ctx context.Context, date time.Time, limit float32) ([]domain.SpeedInfo, error)
    GetMinMax(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error)
}
