package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "time"
)

type Repository interface {
    AddRecord(ctx context.Context, data domain.SpeedInfo) error
    GetRecords(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error)
}
