package packingList

import (
	"reflect"
	"testing"
)

// I suspect in a real world situation you could structure these with a set of inputs and expected results
// and have it run through the lot in a Java cucumber sort of way 🤔 Here is a basic set of tests that allowed
// me to believe that the current implementation matches requirements.

func TestShouldCalculatePackList(t *testing.T) {
	t.Run("should CalculatePackList", func(t *testing.T) {
		// Given
		packSizes := []int{20, 10, 5, 3, 2}
		var requestedCount = 44

		// When
		packingListForRequest, err := GetPackingList(packSizes, requestedCount)
		if err != nil {
			t.Fail()
		}

		// Then
		expectedPackingList := make(map[int]int)
		expectedPackingList[20] = 2
		expectedPackingList[10] = 0
		expectedPackingList[5] = 0
		expectedPackingList[3] = 1
		expectedPackingList[2] = 1

		isExpectedResult := reflect.DeepEqual(packingListForRequest, expectedPackingList)
		if !isExpectedResult {
			t.Fail()
		}
	},
	)
}

func TestShouldCalculatePackListWithSinglePackSize(t *testing.T) {
	// Given
	packSizes := []int{1}
	var requestedCount = 9

	// When
	packingListForRequest, err := GetPackingList(packSizes, requestedCount)
	if err != nil {
		t.Fail()
	}

	// Then
	expectedPackingList := make(map[int]int)
	expectedPackingList[1] = 9

	isExpectedResult := reflect.DeepEqual(packingListForRequest, expectedPackingList)
	if !isExpectedResult {
		t.Fail()
	}
}

func TestShouldReturnErrorWithInvalidRequestedCount(t *testing.T) {
	// Given
	packSizes := []int{10, 5, 3}
	var requestedCount = -3

	// When
	_, err := GetPackingList(packSizes, requestedCount)

	// Then
	if err == nil {
		t.Fail()
	}

}

func TestShouldReturnErrorWithInvalidPackSizes(t *testing.T) {
	// Given
	packSizes := []int{0}
	var requestedCount = 2

	// When
	_, err := GetPackingList(packSizes, requestedCount)

	// Then
	if err == nil {
		t.Fail()
	}

}