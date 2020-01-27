package repositories

import (
    "context"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/domain"
    "github.com/PlagaMedicum/speed_limit_control_service/pkg/storage"
    "github.com/golang/protobuf/ptypes"
    "github.com/pkg/errors"
    "sync"
    "time"
)

// Controller ...
type Controller struct {
    storage.Storage
}

// Insert ...
func (ctl Controller) Insert(ctx context.Context, data domain.SpeedInfo) error {
    ts, err := ptypes.TimestampProto(data.Time)
    if err != nil {
        return errors.Wrapf(err, "Error converting time from %T to *timestamp.Timestamp", data.Time)
    }

    msg := storage.SpeedInfo{
        Time:   ts,
        Number: data.Number,
        Speed:  data.Speed,
    }
    err = ctl.Storage.InsertMessage(ctx, &msg)
    return errors.Wrap(err, "Error inserting message into the storage")
}

// SelectByDate selects all recorded info from the data storage
// for the specified date.
func (ctl Controller) SelectByDate(ctx context.Context, date time.Time) ([]domain.SpeedInfo, error) {
    ts, err := ptypes.TimestampProto(date)
    if err != nil {
        return nil, errors.Wrapf(err, "Error converting time from %T to *timestamp.Timestamp", date)
    }

    msglist, err := ctl.Storage.ReadMessages(ctx, ts)
    if err != nil {
        return nil, errors.Wrap(err, "Error reading records from the storage")
    }

    var siList []domain.SpeedInfo
    mux := sync.Mutex{}
    wg := sync.WaitGroup{}
    errc := make(chan error)
    for _, msg := range msglist {
        wg.Add(1)
        go func(msg storage.SpeedInfo) {
            t, err := ptypes.Timestamp(msg.Time)
            if err != nil {
                errc <- errors.Wrap(err, "Error decoding time from record's timestamp")
            }

            mux.Lock()
            siList = append(siList, domain.SpeedInfo{
                Time:   t,
                Number: msg.Number,
                Speed:  msg.Speed,
            })
            mux.Unlock()

            wg.Done()
        }(msg)
    }

    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(done)
    }()

    select {
    case <-ctx.Done():
        return nil, context.Canceled
    case err = <-errc:
        return nil, err
    case <-done:
        break
    }

    return siList, nil
}
