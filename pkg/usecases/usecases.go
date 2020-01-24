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
func (con Controller) AddData(ctx context.Context, data domain.SpeedInfo) error {

    return nil
}

// GetInfractions returns a list of all transport that
// broke speed limit, for specified date.
func (con Controller) GetInfractions(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}

// GetMinMax returns minimal and maximal speeds for specified date.
func (con Controller) GetMinMax(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}