package domain

import (
    "time"
)

// SpeedInfo represents the data, used in server requests.
type SpeedInfo struct {
    Date   time.Time // date and time of the message
    Number string     // number of the transport
    Speed  float32   // movement speed of the transport
}
