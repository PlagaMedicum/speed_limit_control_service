package usecases

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/repositories"
    "time"
)

type Controller struct {
    repositories.Repository
}

// AddData adds new speed information in storage.
func (ctl Controller) AddData(ctx context.Context, data domain.SpeedInfo) error {
    err := ctl.Repository.AddData(ctx, data)
    return err
}

// GetInfractions returns a list of all transport that
// broke speed limit, for specified date.
func (ctl Controller) GetInfractions(ctx context.Context, date time.Time, limit float32) ([]domain.SpeedInfo, error) {
    silist, err := ctl.Repository.GetInfractions(ctx, date)
    return silist, err
}

// GetMinMax returns minimal and maximal speeds for specified date.
func (ctl Controller) GetMinMax(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {
    silist, err := ctl.Repository.GetBoundaries(ctx, date)
    return silist, err
}
