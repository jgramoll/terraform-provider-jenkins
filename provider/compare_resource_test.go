package provider

import (
	"fmt"
	"strconv"
)

func testCompareResourceInt(className string, fieldName string, expectedValue string, actualValue int) error {
	expectedInt, err := strconv.Atoi(expectedValue)
	if err != nil {
		return err
	}
	if expectedInt != actualValue {
		return fmt.Errorf("%v %v should be %v, was %v", className, fieldName, expectedValue, actualValue)
	}
	return nil
}

func testCompareResourceBool(className string, fieldName string, expectedValue string, actualValue bool) error {
	expectedBool, err := strconv.ParseBool(expectedValue)
	if err != nil {
		return err
	}
	if expectedBool != actualValue {
		return fmt.Errorf("%v %v should be %v, was %v", className, fieldName, expectedValue, actualValue)
	}
	return nil
}
