package uuid

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type UUID strfmt.UUID4

func Parse(id string) (UUID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", err
	}

	return UUID(id), nil
}

func New() UUID {
	return UUID(uuid.New().String())
}

func Nil() UUID {
	return ""
}

func (t UUID) String() string {
	return strfmt.UUID4(t).String()
}

func (t UUID) Validate() error {
	_, err := Parse(t.String())
	return err
}
