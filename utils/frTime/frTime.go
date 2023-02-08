package frTime

import "time"

var (
	isFreeze   bool       = false
	frozenTime *time.Time = nil
)

func Mock(now time.Time) {
	if isFreeze {
		return
	}

	isFreeze = true
	frozenTime = &now
}

func Now() time.Time {
	if isFreeze {
		return *frozenTime
	}

	return time.Now()
}

func ResetMock() {
	if !isFreeze {
		return
	}

	isFreeze = false
	frozenTime = nil
}
