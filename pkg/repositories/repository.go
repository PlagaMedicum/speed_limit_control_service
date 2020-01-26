package repositories

import (
	"context"
	"github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
	"github.com/PlagaMedicum/speed_limit_control_service/pkg/storage"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"time"
)

// Controller ...
type Controller struct {
	storage.Storage
}

// AddRecord ...
func (ctl Controller) AddRecord(ctx context.Context, data domain.SpeedInfo) error {
	ts, err := ptypes.TimestampProto(data.Time)
	if err != nil {
		return errors.Wrapf(err, "Error converting time from %T to *timestamp.Timestamp", data.Time)
	}

	si := storage.SpeedInfo{
		Time:   ts,
		Number: data.Number,
		Speed:  data.Speed,
	}
	err = ctl.Storage.Insert(&si)
	return errors.Wrap(err, "Error inserting data into storage")
}

// GetRecords ...
func (ctl Controller) GetRecords(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {
	ts, err := ptypes.TimestampProto(date)
	if err != nil {
		return nil, errors.Wrapf(err, "Error converting time from %T to *timestamp.Timestamp", date)
	}

	data, err := ctl.Storage.Read(ts)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading records from the storage")
	}

	var siList []domain.SpeedInfo
	for _, rec := range data {
		t, err := ptypes.Timestamp(rec.Time)
		if err != nil {
			return nil, errors.Wrap(err, "Error decoding time from record's timestamp")
		}

		siList = append(siList, domain.SpeedInfo{
			Time:   t,
			Number: rec.Number,
			Speed:  rec.Speed,
		})
	}
	return siList, nil
}
