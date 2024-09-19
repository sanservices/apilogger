package apilogger

import "testing"

func TestLogCategories(t *testing.T) {
	logCatCodes := make(map[string]string)

	for _, cat := range allLogCats {
		code, ok := logCatCodes[cat.Code]
		if ok {
			t.Errorf("Duplicate log category code - [%s]", code)
		}
		logCatCodes[cat.Code] = cat.Code
	}
}
