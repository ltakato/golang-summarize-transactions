package core

import (
	"encoding/json"
	"testing"
)

func TestCurrency_MarshalJSON(t *testing.T) {
	t.Run("should return passed currency as float with decimal", func(t *testing.T) {
		c := Currency(3095)
		value, err := json.Marshal(c)
		expected := "30.95"

		if err != nil {
			t.Errorf("failed to marshall currency value")
		}

		if string(value) != expected {
			t.Errorf("expected value to be %s, but got %s", expected, string(value))
		}
	})
}

func TestCurrency_UnmarshalJSON(t *testing.T) {
	t.Run("should return passed currency as float with decimal", func(t *testing.T) {
		expected := Currency(3095)
		value, err := json.Marshal(expected)
		var result Currency
		err = json.Unmarshal(value, &result)

		if err != nil {
			t.Errorf("failed to unmarshall currency value")
		}

		if result != expected {
			t.Errorf("expected value to be %d, but got %d", expected, result)
		}
	})
}
