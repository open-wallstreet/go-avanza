package client

import (
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EncoderTestSuite struct {
	suite.Suite
}

func TestEncoderTestSuite(t *testing.T) {
	suite.Run(t, new(EncoderTestSuite))
}

func (suite *EncoderTestSuite) SetupTest() {
}

func (suite *EncoderTestSuite) TestEncode_Params() {
	testPath := "/v1/{float}/{str}"

	type Params struct {
		Float float64 `validate:"required" path:"float"`
		Str   string  `validate:"required" path:"str"`

		FloatQuery *float64 `query:"float"`
		StrQuery   *string  `query:"str"`
	}

	num := 2.1234
	str := "testing"
	params := Params{
		Float:      2.1234,
		Str:        str,
		FloatQuery: &num,
		StrQuery:   &str,
	}

	expected := "/v1/2.1234/testing?float=2.1234&str=testing"
	actual, err := NewEncoder().EncodeParams(testPath, params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *EncoderTestSuite) TestEncode_Time() {
	testPath := "/v1/{time}"

	type Params struct {
		Time  models.Time  `validate:"required" path:"time"`
		TimeQ *models.Time `query:"time"`
	}

	ptime := models.Time(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	params := Params{
		Time:  ptime,
		TimeQ: &ptime,
	}

	expected := "/v1/2020-01-01T00:00:00.000Z?time=2020-01-01T00%3A00%3A00.000Z"
	actual, err := NewEncoder().EncodeParams(testPath, params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *EncoderTestSuite) TestEncode_Date() {
	testPath := "/v1/{date}"

	type Params struct {
		Time  models.Date  `validate:"required" path:"date"`
		TimeQ *models.Date `query:"date"`
	}

	ptime := models.Date(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	params := Params{
		Time:  ptime,
		TimeQ: &ptime,
	}

	expected := "/v1/2020-01-01?date=2020-01-01"
	actual, err := NewEncoder().EncodeParams(testPath, params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *EncoderTestSuite) TestEncode_Milliseconds() {
	testPath := "/v1/{millis}"

	type Params struct {
		Millis  models.Millis  `validate:"required" path:"millis"`
		MillisQ *models.Millis `query:"millis"`
	}

	pMillis := models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := Params{
		Millis:  pMillis,
		MillisQ: &pMillis,
	}

	expected := "/v1/1626912000000?millis=1626912000000"
	actual, err := NewEncoder().EncodeParams(testPath, params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *EncoderTestSuite) TestEncode_Nanoseconds() {
	testPath := "/v1/{nanos}"

	type Params struct {
		Nanos  models.Nanos  `validate:"required" path:"nanos"`
		NanosQ *models.Nanos `query:"nanos"`
	}

	pNanos := models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := Params{
		Nanos:  pNanos,
		NanosQ: &pNanos,
	}

	expected := "/v1/1626912000000000000?nanos=1626912000000000000"
	actual, err := NewEncoder().EncodeParams(testPath, params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *EncoderTestSuite) TestEncodeValidationError() {
	_, err := NewEncoder().EncodeParams("/v1/test", nil)
	assert.NotNil(suite.T(), err)
}
