package cli

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type DataType int

// We doon't want too offer too many options, because that goes out of scope,
// but providing the tools that will enable validation for 98% of the input is
// desirable if its not too much. Since it will enable developers too have input
// validation and consistent input across all applications, even simple scripts
// with little cost. Like MAC could be included, but its not nearly as common as
// IP or Filename as a datatype. We could do a automated survey of GNU coreutils to
// try to determine the most common types in the future with greater certainty.
const (
	Bool DataType = iota
	Int
	Float
	String
	Directory
	Filename
	Filenames // Via globbing or comma separated values
	URL
	Port
	IPv4
	IPv6
)

// TODO: This should migrate into its own package (the generic equivialence
// helpers) the idea would be to provide a collection of helpers to make all of
// our code cleaner and more expressive in the same way rails extends the
// epxressiveness of default ruby
// These may seem pointless but they will also simplify validation of values and
// provide helpers for developers to simplify validation
const (
	Space = " "
	Blank = ""
	Zero  = 0
)

var trueValues = []string{"t", "true", "y", "yes", "1"}
var falseValues = []string{"f", "false", "n", "no", "0"}

//
// Transform
///////////////////////////////////////////////////////////////////////////////

//
// Validate
///////////////////////////////////////////////////////////////////////////////

//
// Encode
///////////////////////////////////////////////////////////////////////////////
func toString(value interface{}) string { return value.(string) }

// String Subtypes
func toDirectory(value interface{}) string { return value.(string) }
func toFilename(value interface{}) string  { return value.(string) }
func toFilenames(value interface{}) string { return value.(string) }
func toURL(value interface{}) string       { return value.(string) }
func toIPv4(value interface{}) string      { return value.(string) }
func toIPv6(value interface{}) string      { return value.(string) }

func toInt(value interface{}) int {
	intValue, err := strconv.Atoi(value.(string))
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

// Int Subtypes
func toPort(value interface{}) int { return toInt(value) }

func toFloat(value interface{}) float64 {
	floatParts := strings.Split(value.(string), ".")
	if len(floatParts) > 1 {
		floatValue, err := strconv.ParseFloat(value.(string), len(floatParts[1]))
		if err != nil {
			return float64(0.00)
		} else {
			return floatValue
		}
	}
	return float64(0)
}

func toBool(value interface{}) bool {
	for _, trueValue := range trueValues {
		if value.(string) == trueValue {
			return true
		}
	}
	return false
}

//
// Public Methods
///////////////////////////////////////////////////////////////////////////////
// TODO: Basic validation should move here, and basic conversion (which is in
// flags since this is where it orginated, but it became clear it would also be
// important for parameters) [BUT maybe parameters don't have a type? Just hand
// off the string, just seems if we are dealing with flag type we should go
// ahead and extend this to parameters, but it may be wise to just ignore both]
func Valid(flagType DataType, value interface{}) (bool, error) {
	switch flagType {
	case Bool:
		boolValues := append(trueValues, falseValues...)
		for _, boolValue := range boolValues {
			if boolValue == value {
				return true, nil
			}
		}
		return false, errors.New("[error] could not parse valid boolean value")
	//case Int:
	//case String:
	//case Directory:
	case Filename:
		_, err := os.Stat(value.(string))
		return (err == nil), nil
	//case Filenames:
	//case URL:
	//case IPv4:
	//case IPv6:
	//case Port:
	default:
		return false, errors.New("[error] failed to parse data type")
	}
}

// Not expressive because it makes "Not is zero"
// func Not(value bool) bool { return !value }
func IsZero(value int) bool            { return value == Zero }
func IsBlank(str string) bool          { return IsZero(len(str)) }
func IsEmpty(value []interface{}) bool { return IsZero(len(value)) }
func IsNil(value interface{}) bool {
	switch value.(type) {
	case int:
		return value.(int) == Zero
	case string:
		return value.(string) == Blank
	case error:
		return value.(error) == nil
	default:
		return value == nil
	}
}
func IsGreaterThan(gt, value int) bool         { return (gt > value) }
func IsGreaterOrEqualThan(gte, value int) bool { return (gte >= value) }
func IsLessThan(lt, value int) bool            { return (lt < value) }
func IsLessOrEqualThan(lte, value int) bool    { return (lte <= value) }
func IsBetween(start, end, value int) bool     { return start < value && value < end }
