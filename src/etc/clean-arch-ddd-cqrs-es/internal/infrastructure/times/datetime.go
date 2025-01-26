package times

import "time"

type TimeProvider interface {
	Now() time.Time
}
