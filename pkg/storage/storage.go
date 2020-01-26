package storage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
)

const dataPath = "data/"
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
}

// genFileName generates file name for provided date in timestamp.
// All files have timestamp, represented by days of UTC time since
// Unix epoch 1970-01-01T00:00:00Z.
func genFileName(seconds int64) string {
	days := seconds / secondsInDay // division with no remainder.
	return fmt.Sprintf("%srecord%d.pb", dataPath, days)
}

// Insert creates a new binary record for the data in
// the data storage.
func (s Storage) Insert(data *SpeedInfo) error {
	b, err := proto.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "Error marshalling proto message")
	}

	fname := genFileName(data.Time.Seconds)
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
		return errors.Wrap(err, "Error writing data to file")
	}

	err = f.Close()
	return errors.Wrap(err, "Error closing file")
}

// Read retrieves all data from the record file in the storage
// for specified date in the ts timestamp.
func (s Storage) Read(ts *timestamp.Timestamp) ([]SpeedInfo, error) {
	fname := genFileName(ts.Seconds)
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, errors.Wrapf(err, "Error reading file %s", fname)
	}

	var res []SpeedInfo
	for len(b) > 0 {
		if len(b) < recLengthSize {
			return nil, errors.Errorf("Odd %d bytes occurred instead of record length mark.", len(b))
		}

		var rlen recLength
		err := binary.Read(bytes.NewReader(b[:recLengthSize]), byteOrder, &rlen)
		if err != nil {
			return nil, errors.Wrap(err, "Error decoding record's length")
		}

		b = b[recLengthSize:]
		var si SpeedInfo
		err = proto.Unmarshal(b[:rlen], &si)
		if err != nil {
			return nil, errors.Wrap(err, "Error reading record's data")
		}

		b = b[rlen:]
		res = append(res, si)
	}

	return res, nil
}
