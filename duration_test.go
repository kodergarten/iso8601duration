package duration

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFromString(t *testing.T) {
	t.Parallel()

	// test with bad format
	_, err := FromString("asdf")
	assert.Equal(t, err, ErrBadFormat)

	// test with good full string
	dur, err := FromString("P1Y2M3DT4H5M6S")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Years)
	assert.Equal(t, 2, dur.Months)
	assert.Equal(t, 3, dur.Days)
	assert.Equal(t, 4, dur.Hours)
	assert.Equal(t, 5, dur.Minutes)
	assert.Equal(t, 6, dur.Seconds)

	// test with good week string
	dur, err = FromString("P1W")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Weeks)
}

func TestString(t *testing.T) {
	t.Parallel()

	// test empty
	d := Duration{}
	assert.Equal(t, d.String(), "P")

	// test only larger-than-day
	d = Duration{Years: 1, Months: 2, Days: 3}
	assert.Equal(t, d.String(), "P1Y2M3D")

	// test only smaller-than-day
	d = Duration{Hours: 1, Minutes: 2, Seconds: 3}
	assert.Equal(t, d.String(), "PT1H2M3S")

	// test full format
	d = Duration{Years: 1, Months: 2, Days: 3, Hours: 4, Minutes: 5, Seconds: 6}
	assert.Equal(t, d.String(), "P1Y2M3DT4H5M6S")

	// test week format
	d = Duration{Weeks: 1}
	assert.Equal(t, d.String(), "P1W")
}

func TestToDuration(t *testing.T) {
	t.Parallel()

	d := Duration{Years: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24*365)

	d = Duration{Months: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24*30)

	d = Duration{Weeks: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24*7)

	d = Duration{Days: 1}
	assert.Equal(t, d.ToDuration(), time.Hour*24)

	d = Duration{Hours: 1}
	assert.Equal(t, d.ToDuration(), time.Hour)

	d = Duration{Minutes: 1}
	assert.Equal(t, d.ToDuration(), time.Minute)

	d = Duration{Seconds: 1}
	assert.Equal(t, d.ToDuration(), time.Second)
}
