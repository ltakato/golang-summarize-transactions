package core

import "encoding/json"

type Currency int

func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(c) / 100)
}

func (c *Currency) UnmarshalJSON(data []byte) error {
	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Currency(value * 100)
	return nil
}
