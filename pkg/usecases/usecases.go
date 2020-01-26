package usecases

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/repositories"
    "math"
    "time"
)

type Controller struct {
    repositories.Repository
}

// AddData adds new speed information in storage.
func (ctl Controller) AddData(ctx context.Context, data domain.SpeedInfo) error {
    err := ctl.Repository.AddRecord(ctx, data)
    return err
}

// GetInfractions returns a list of all transport that
// broke speed limit, for specified date.
func (ctl Controller) GetInfractions(ctx context.Context, date time.Time, limit float32) ([]domain.SpeedInfo, error) {
    silist, err := ctl.Repository.GetRecords(ctx, date)

    var infList []domain.SpeedInfo
    for _, si := range silist {
        if si.Speed > limit {
            infList = append(infList, si)
        }
    }

    return infList, err
}

// GetMinMax returns minimal and maximal speeds for specified date.
func (ctl Controller) GetMinMax(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {
    silist, err := ctl.Repository.GetRecords(ctx, date)

    min := domain.SpeedInfo{Speed: float32(math.MaxFloat32)}
    max := domain.SpeedInfo{Speed: float32(0)}
    for _, si := range silist {
        if si.Speed > max.Speed {
            max = si
        }
        if si.Speed < min.Speed {
            min = si
        }
    }

    return []domain.SpeedInfo{min, max}, err
}
