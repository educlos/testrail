package testrail

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type IntList []int

func (il IntList) MarshalJSON() ([]byte, error) {
	return json.Marshal(intsList([]int(il)))
}

func intsList(ints []int) string {
	var str []string
	for _, i := range ints {
		str = append(str, fmt.Sprintf("%d", i))
	}
	return strings.Join(str, ",")
}

// loadOptionalFilters takes those RequestFilters and transforms them
// into a Query string.  You need to pass a *list* of filters here, with
// optionally the first one being used.
func loadOptionalFilters(vals url.Values, filters interface{}) {
	cnt, err := json.Marshal(filters)
	if err != nil {
		panic("should only be passing types that we know marshal over here !")
	}

	var kv []map[string]interface{}
	err = json.Unmarshal(cnt, &kv)
	if err != nil {
		panic("should only pass a list of objects that marshal to map[string]string in here !")
	}

	if len(kv) == 0 {
		return
	}

	for k, v := range kv[0] {
		switch val := v.(type) {
		case string:
			vals.Set(k, val)
		case float64:
			vals.Set(k, fmt.Sprintf("%v", val))
		default:
			panic(fmt.Sprintf("type %T not supported in loadOptionalFilters", v))
		}
	}
}

func boolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
