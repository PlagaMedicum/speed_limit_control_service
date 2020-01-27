package storage

import (
    "context"
    "github.com/golang/protobuf/ptypes"
    "os"
    "testing"
)

func newCancelledContext() context.Context {
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    return ctx
}

func TestStorage_InsertRead(t *testing.T) {
    ts := ptypes.TimestampNow()
    msg := SpeedInfo{
        Time:   ts,
        Number: "123 PP-7",
        Speed:  33.5,
    }

    s := Storage{DataPath: "./"}
    fname := s.genFileName(ts.Seconds)

    os.Remove(fname)

    ctx := context.Background()
    err := s.InsertMessage(ctx, &msg)
    if err != nil {
        t.Errorf("Error inserting message: %v", err)
    }
    defer os.Remove(fname)

    res, err := s.ReadMessages(ctx, ts)
    if err != nil {
        t.Errorf("Error reading message: %v", err)
    }

    rmsg := res[0]
    if rmsg.Time.Seconds != msg.Time.Seconds || rmsg.Speed != msg.Speed || rmsg.Number != msg.Number {
        t.Errorf("inserted and readed messages are not equal.\nInserted: %v\nReaded: %v", rmsg, msg)
    }
}

func TestStorage_InsertMessage(t *testing.T) {
    ctx := context.Background()
    s := Storage{DataPath: "./"}
    ts := ptypes.TimestampNow()

    tclist := []struct {
        Name string
        Msg  SpeedInfo
        Ctx  context.Context
        Err  error
    }{
        {Name: "Ok", Msg: SpeedInfo{
            Time:   ts,
            Number: "1234 P",
            Speed:  43,
        }, Ctx: ctx},
        {Name: "Cancelled context",
            Ctx: newCancelledContext(), Err: context.Canceled},
    }

    for _, tc := range tclist {
        t.Run(tc.Name, func(t *testing.T) {
            var fname string
            if tc.Msg.Time != nil {
                fname = s.genFileName(tc.Msg.Time.Seconds)
                os.Remove(fname)
            }

            err := s.InsertMessage(tc.Ctx, &tc.Msg)
            if err != tc.Err {
                t.Errorf("Wrong error value!\nExpected: %v\nGot: %v", tc.Err, err)
            }

            defer os.Remove(fname)
        })
    }
}
