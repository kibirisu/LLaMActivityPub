package domain

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json/v2"
)

type nullable interface {
	driver.Valuer
	sql.Scanner
}

func withMarshalers() json.Options {
	fn := json.MarshalFunc(func(val nullable) ([]byte, error) {
		data, err := val.Value()
		if err != nil {
			return nil, err
		}
		return json.Marshal(data)
	})
	return json.WithMarshalers(fn)
}

func withUnmarshalers() json.Options {
	fn := json.UnmarshalFunc(func(data []byte, val nullable) error {
		alias, err := val.Value()
		if err != nil {
			return err
		}
		if err = json.Unmarshal(data, &alias); err != nil {
			return err
		}
		return val.Scan(alias)
	})
	return json.WithUnmarshalers(fn)
}

func getOptions() json.Options {
	return json.JoinOptions(withMarshalers(), withUnmarshalers())
}
