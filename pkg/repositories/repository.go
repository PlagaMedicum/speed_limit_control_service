package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "time"
)

type Controller struct {

}

func (ctl Controller) AddData(ctx context.Context, data domain.SpeedInfo) error {

    return nil
}

func (ctl Controller) GetInfractions(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}

func (ctl Controller) GetBoundaries(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}
