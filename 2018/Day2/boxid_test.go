package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BoxIdSuite struct {
	suite.Suite
}

func TestBoxIdSuite(t *testing.T) {
	suite.Run(t, new(BoxIdSuite))
}

func (s *BoxIdSuite) TestBoxId() {
	cases := []struct {
		Name        string
		InputData   []string
		ExpectedOut string
		ExpectedErr bool
	}{
		{
			Name:        "standard case",
			InputData:   loadData(s),
			ExpectedOut: "fgij",
		},
	}

	for _, c := range cases {
		f := boxListReader()

		assert.Equal(s.T(), c.ExpectedOut, f, c.Name)

	}

}
