package data

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json/v2"
)

type Type interface {
	driver.Valuer
	sql.Scanner
}

func marshalValue(val Type) ([]byte, error) {
	data, err := val.Value()
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

func unmarshalValue(data []byte, val Type) error {
	alias, err := val.Value()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &alias); err != nil {
		return err
	}
	return val.Scan(alias)
}

func WithMarshalers() json.Options {
	return json.WithMarshalers(json.MarshalFunc(marshalValue))
}

func WithUnmarshalers() json.Options {
	return json.WithUnmarshalers(json.UnmarshalFunc(unmarshalValue))
}
