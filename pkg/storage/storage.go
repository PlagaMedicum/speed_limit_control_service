package storage

import (
    "encoding/binary"
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/pkg/errors"
    "os"
)

const dataPath = "data/"

type Storage struct {

}

func (s Storage) Insert(data *SpeedInfo) error {
    b, err := proto.Marshal(data)
    if err != nil {
        return errors.Wrap(err, "Error marshalling proto message")
    }

    fname := fmt.Sprintf("%srecord%d.pb", dataPath, data.Date.Seconds)
    f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        return errors.Wrapf(err, "Error opening file %s", fname)
    }

    err = binary.Write(f, binary.LittleEndian, int64(len(b)))
    if err != nil {
        return errors.Wrap(err, "Error encoding length of message")
    }
    _, err = f.Write(b)
    if err != nil {
        return errors.Wrap(err, "Error writing data to file")
    }

    err = f.Close()
    return errors.Wrap(err, "Error closing file")
}
