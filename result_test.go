package testrail

import (
	"encoding/json"
	"testing"
	"time"
)

const resultsJSON = `[{
    "id": 3179,
    "test_id": 91193,
    "status_id": 5,
    "created_by": 11,
    "created_on": 1475001797,
    "elapsed": "37m"
}]`

func TestFixedGetResults(t *testing.T) {
	var results []Result
	if err := json.Unmarshal([]byte(resultsJSON), &results); err != nil {
		t.Fatal(err)
	}

	r := results[0]

	createdOn, err := time.Parse(time.RFC3339, "2016-09-27T18:43:17Z")
	if err != nil {
		t.Fatal(err)
	}

	if !r.CreatedOn.Equal(createdOn) {
		t.Fatalf("Got: %v, Expected: %v", r.CreatedOn, createdOn)
	}
	if r.Elapsed.Duration != 37*time.Minute {
		t.Fatalf("Wrong duration: %v", r.Elapsed.Duration)
	}
	if r.ID != 3179 {
		t.Fatalf("Wrong ID: %v", r.ID)
	}
	if r.TestID != 91193 {
		t.Fatalf("Wrong testID: %v", r.TestID)
	}
	if r.StatusID != 5 {
		t.Fatalf("Wrong status: %v", r.StatusID)
	}
	if r.CreatedBy != 11 {
		t.Fatalf("Wrong createdBy: %v", r.CreatedBy)
	}
}
