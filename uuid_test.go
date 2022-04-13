package uuid

import "testing"

func TestParse(t *testing.T) {
	type testData struct {
		uuid        string
		mustBeError bool
	}

	tests := []testData{
		{"ca0408b7-de11-4c8d-bc5a-3290c5d4e0db", false},
		{"063f9003-6d5f-4967-9b21-17856a23556f", false},
		{"063f9003-6d5f-4967-9b21-17856a23556", true},
		{"test", true},
		{"", true},
	}

	for _, test := range tests {
		_, err := Parse(test.uuid)
		if test.mustBeError && err == nil {
			t.Error(
				"For", test.uuid,
				"expected error, but got nil",
			)
		}

		if !test.mustBeError && err != nil {
			t.Error(
				"For", test.uuid,
				"expected error nil, but got error", err,
			)
		}
	}
}

func TestString(t *testing.T) {
	type testData struct {
		uuid       UUID
		uuidString string
	}

	tests := []testData{
		{UUID("0d53a49c-6a2d-44a6-85fe-7a8a11347212"), "0d53a49c-6a2d-44a6-85fe-7a8a11347212"},
		{UUID("e20ffb8d-6adf-4e14-83e5-e79d4eb1e4d6"), "e20ffb8d-6adf-4e14-83e5-e79d4eb1e4d6"},
		{UUID("8f488b76-f8c8-4f5a-ad1e-042a835bf27c"), "8f488b76-f8c8-4f5a-ad1e-042a835bf27c"},
		{UUID(""), ""},
	}

	for _, test := range tests {
		if test.uuid.String() != test.uuidString {
			t.Error(
				"Fail on uuid stringify:",
				"uuid", test.uuid,
				"uuid string", test.uuidString,
			)
		}
	}
}

func TestValidate(t *testing.T) {
	type testData struct {
		uuid        UUID
		mustBeError bool
	}

	tests := []testData{
		{UUID("ca0408b7-de11-4c8d-bc5a-3290c5d4e0db"), false},
		{UUID("063f9003-6d5f-4967-9b21-17856a23556f"), false},
		{UUID("063f9003-6d5f-4967-9b21-17856a23556"), true},
		{UUID("test"), true},
		{UUID(""), true},
	}

	for _, test := range tests {
		err := test.uuid.Validate()
		if test.mustBeError && err == nil {
			t.Error(
				"For", test.uuid,
				"expected error, but got nil",
			)
		}

		if !test.mustBeError && err != nil {
			t.Error(
				"For", test.uuid,
				"expected error nil, but got error", err,
			)
		}
	}
}
