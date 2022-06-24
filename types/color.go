package types

import (
	"database/sql/driver"
	"fmt"
)

type Colors string

const (
	ColorsRed   Colors = "red"
	ColorsBlue  Colors = "blue"
	ColorsGreen Colors = "green"
)

func (e *Colors) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Colors(s)
	case string:
		*e = Colors(s)
	default:
		return fmt.Errorf("unsupported scan type for Colors: %T", src)
	}
	return nil
}

func (e Colors) Valid() bool {
	switch e {
	case ColorsRed,
		ColorsBlue,
		ColorsGreen:
		return true
	}
	return false
}

type NullColors struct {
	Colors Colors
	Valid  bool
}

func (nc *NullColors) Scan(src interface{}) error {
	if src == nil {
		nc.Colors = Colors("")
		nc.Valid = false
		return nil
	}

	err := nc.Colors.Scan(src)
	if err != nil {
		nc.Valid = false
		return err
	}

	if !nc.Colors.Valid() {
		nc.Valid = false
		return fmt.Errorf("%s is not a valid color", nc.Colors)
	}

	nc.Valid = true

	return nil
}

func (nc NullColors) Value() (driver.Value, error) {
	if !nc.Valid {
		return nil, fmt.Errorf("not a valid color")
	}

	if !nc.Colors.Valid() {
		return nil, fmt.Errorf("not a valid color")
	}

	return string(nc.Colors), nil
}
