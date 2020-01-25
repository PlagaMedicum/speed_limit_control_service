package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/storage"
    "github.com/golang/protobuf/ptypes"
    "github.com/pkg/errors"
    "time"
)

type Controller struct {
    storage.Storage
}

func (ctl Controller) AddData(ctx context.Context, data domain.SpeedInfo) error {
    date, err := ptypes.TimestampProto(data.Date)
    if err != nil {
        return errors.Wrapf(err, "Error converting date from %T to *tspb.Timestamp", data.Date)
    }

    si := storage.SpeedInfo{
        Date:                 date,
        Number:               data.Number,
        Speed:                data.Speed,
    }
    err = ctl.Storage.Insert(&si)
    if err != nil {
        return errors.Wrap(err, "Error inserting data into storage")
    }

    return nil
}

func (ctl Controller) GetInfractions(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}

func (ctl Controller) GetBoundaries(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {

    return nil, nil
}
