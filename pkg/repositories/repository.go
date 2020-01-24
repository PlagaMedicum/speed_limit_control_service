package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "time"
)

type Controller struct {

}

func (con Controller) AddData(ctx context.Context, data domain.SpeedInfo) error {

    return nil
}

func (con Controller) GetInfractions(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}

func (con Controller) GetBoundaries(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}
