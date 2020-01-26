package domain

import (
    "time"
)

// SpeedInfo represents the data, used in server requests.
type SpeedInfo struct {
    Time   time.Time // time of the message
    Number string    // number of the transport
    Speed  float32   // movement speed of the transport
}
