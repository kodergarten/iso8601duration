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
	assert.Equal(t, 1, dur.Years())
	assert.Equal(t, 2, dur.Months())
	assert.Equal(t, 3, dur.Days())
	assert.Equal(t, 4, dur.Hours())
	assert.Equal(t, 5, dur.Minutes())
	assert.Equal(t, 6, dur.Seconds())

	// test with good week string
	dur, err = FromString("P1W")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Weeks())
}

func TestString(t *testing.T) {
	t.Parallel()

	// test empty
	d := Duration{}
	assert.Equal(t, d.String(), "P")

	// test only larger-than-day
	p, _ := time.ParseDuration("10272h")
	d = Duration{p}
	assert.Equal(t, d.String(), "P1Y2M3D")

	// test only smaller-than-day
	p, _ = time.ParseDuration("1h2m3s")
	d = Duration{p}
	assert.Equal(t, d.String(), "PT1H2M3S")

	// test full format
	p, _ = time.ParseDuration("10276h5m6s")
	d = Duration{p}
	assert.Equal(t, d.String(), "P1Y2M3DT4H5M6S")

	// test week format
	p, _ = time.ParseDuration("168h")
	d = Duration{p}
	assert.Equal(t, d.String(), "P1W")
}

func TestToDuration(t *testing.T) {
	t.Parallel()
	p, _ := time.ParseDuration("8760h")
	d := Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24*365)
	p, _ = time.ParseDuration("720h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24*30)
	p, _ = time.ParseDuration("168h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24*7)
	p, _ = time.ParseDuration("24h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24)
	p, _ = time.ParseDuration("1h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour)
	p, _ = time.ParseDuration("1m")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Minute)
	p, _ = time.ParseDuration("1s")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Second)
}
