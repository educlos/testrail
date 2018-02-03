package testrail

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimespanUnmarshal(t *testing.T) {
	var testData = []struct{ json, stringDuration string }{
		{`null`, "0s"},
		{`"15s"`, "15s"},
		{`"12m"`, "12m"},
		{`"11h"`, "11h"},
		{`"4h 5m 6s"`, "4h5m6s"},
		{`"1d"`, "8h"},
		{`"1w"`, "40h"},
		{`"1d 2h"`, "8h2h"},
		{`"1w 2d 3h"`, "40h16h3h"},
	}

	for _, data := range testData {
		var results []Result
		js := []byte(fmt.Sprintf(`[{"elapsed":%v}]`, data.json))
		if err := json.Unmarshal(js, &results); err != nil {
			t.Fatal(err)
		}

		r := results[0]
		expected, err := time.ParseDuration(data.stringDuration)
		if err != nil {
			t.Fatal(err)
		}

		if r.Elapsed.Duration != expected {
			t.Fatalf("Wrong duration: %v", r.Elapsed.Duration)
		}
	}
}

func TestTimespanMarshal(t *testing.T) {
	var testData = []struct{ json, stringDuration string }{
		{`null`, "0s"},
		{`"0h 0m 15s"`, "15s"},
		{`"0h 12m 0s"`, "12m"},
		{`"11h 0m 0s"`, "11h"},
		{`"4h 5m 6s"`, "4h5m6s"},
		{`"8h 0m 0s"`, "8h"},
		{`"40h 0m 0s"`, "40h"},
		{`"10h 0m 0s"`, "8h2h"},
		{`"59h 0m 0s"`, "40h16h3h"},
	}

	for _, td := range testData {
		duration, err := time.ParseDuration(td.stringDuration)
		if err != nil {
			t.Fatal(err)
		}

		result := Result{
			Elapsed: timespan{
				Duration: duration,
			},
		}
		data, err := json.Marshal(result.Elapsed)
		if err != nil {
			t.Fatal(err)
		}

		if string(data) != td.json {
			t.Fatalf("Wrong data: %v", string(data))
		}
	}
}

func TestTimespanFromDurationValidDuration(t *testing.T) {
	start := time.Now()
	time.Sleep(5 * time.Millisecond)
	d := time.Since(start)
	ts := TimespanFromDuration(d)
	assert.NotNil(t, ts)
	assert.Equal(t, d, ts.Duration)
}

func TestTimespanFromDurationInvalidDuration(t *testing.T) {
	d, _ := time.ParseDuration("0s")
	ts := TimespanFromDuration(d)
	assert.Nil(t, ts)
}
