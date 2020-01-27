package storage

import (
    "bytes"
    "context"
    "encoding/binary"
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/ptypes/timestamp"
    "github.com/pkg/errors"
    "io/ioutil"
    "os"
)

const secondsInDay int64 = 24 * 60 * 60

// recLength is a type for representing length of a record in
// the storage
type recLength int64

// size in bytes of recLength
const recLengthSize = 8

var byteOrder = binary.LittleEndian

// Storage handles all operations of writing and reading records
// from the data storage.
type Storage struct {
    DataPath string
}

// genFileName generates file name for provided date in timestamp.
// All files have timestamp, represented by days of UTC time since
// Unix epoch 1970-01-01T00:00:00Z.
func (s Storage) genFileName(seconds int64) string {
    days := seconds / secondsInDay // division with no remainder.
    return fmt.Sprintf("%srecords%d.pb", s.DataPath, days)
}

// InsertMessage creates a new binary record for the message, in
// the data storage.
func (s Storage) InsertMessage(ctx context.Context, msg *SpeedInfo) error {
    select {
    case <-ctx.Done():
        return context.Canceled
    default:
        break
    }

    b, err := proto.Marshal(msg)
    if err != nil {
        return errors.Wrap(err, "Error marshalling proto message")
    }

    fname := s.genFileName(msg.Time.Seconds)
    f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
    if err != nil {
        return errors.Wrapf(err, "Error opening file %s", fname)
    }

    err = binary.Write(f, byteOrder, recLength(len(b)))
    if err != nil {
        return errors.Wrap(err, "Error encoding and writing length of record in file")
    }
    _, err = f.Write(b)
    if err != nil {
        return errors.Wrap(err, "Error writing message's data to file")
    }

    err = f.Close()
    return errors.Wrap(err, "Error closing file")
}

// readMessage reads only one message
func readMessage(b []byte) (recLength, SpeedInfo, error) {
    if len(b) < recLengthSize {
        return 0, SpeedInfo{}, errors.Errorf("Odd %d bytes occurred instead of record length mark.", len(b))
    }

    var rlen recLength
    err := binary.Read(bytes.NewReader(b[:recLengthSize]), byteOrder, &rlen)
    if err != nil {
        return 0, SpeedInfo{}, errors.Wrap(err, "Error decoding record's length")
    }

    b = b[recLengthSize:]
    var si SpeedInfo
    err = proto.Unmarshal(b[:rlen], &si)
    if err != nil {
        return 0, SpeedInfo{}, errors.Wrap(err, "Error unmarshalling message")
    }

    return recLengthSize + rlen, si, nil
}

// ReadMessages retrieves all data from the record file in the storage
// for specified date in the ts timestamp, and returns it as a slice of
// unmarshalled messages.
func (s Storage) ReadMessages(ctx context.Context, ts *timestamp.Timestamp) ([]SpeedInfo, error) {
    fname := s.genFileName(ts.Seconds)
    b, err := ioutil.ReadFile(fname)
    if err != nil {
        return nil, errors.Wrapf(err, "Error reading file %s", fname)
    }

    var res []SpeedInfo
    for len(b) > 0 {
        select {
        case <-ctx.Done():
            return nil, context.Canceled
        default:
            break
        }

        rlen, si, err := readMessage(b)
        if err != nil {
            return nil, errors.Wrap(err, "Error reading message")
        }
        b = b[rlen:]

        res = append(res, si)
    }

    return res, nil
}
