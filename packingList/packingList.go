package packingList

import (
	"errors"
	"sort"
)

type PacksUsed struct {
	remainder int
	packsUsed int
}

func fullPacksDeliverableCount(requestedCount int, packSize int) (PacksUsed, error) {
	if packSize < 1 {
		// 🤔 Not sure about returning an empty PacksUsed below 🤔
		return PacksUsed{} , errors.New("invalid pack size, supplied size was less than 1")
	}
	remainder := requestedCount % packSize
	packsUsed := (requestedCount - remainder) / packSize
	return PacksUsed{remainder: remainder, packsUsed: packsUsed}, nil
}

func GetPackingList(packSizes []int, requestedCount int) (map[int]int, error) {
	// Minimal Validation
	if len(packSizes) == 0 {
		return nil, errors.New("must supply at least 1 packSize")
	}
	if requestedCount < 1 {
		return nil, errors.New("must request at least 1 pack")
	}

	packingList := make(map[int]int)

	// This sort operation mutates the underlying variable 🙈
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	remainingItemsToFulfill := requestedCount
	for _, packSize := range packSizes {
		packsUsedForAPackSize, err := fullPacksDeliverableCount(remainingItemsToFulfill, packSize)
		if err != nil {
			return nil, errors.New("unable to calculate packs required")
		}
		packingList[packSize] = packsUsedForAPackSize.packsUsed
		remainingItemsToFulfill = packsUsedForAPackSize.remainder
	}

	if remainingItemsToFulfill > 0 {
		// There is going to be over over fulfilment
		// ToDo: It would be nice to track the over fulfilment in orders
		minimumPackSizeIndex := len(packSizes) - 1
		packingList[packSizes[minimumPackSizeIndex]] = 1
	}

	return packingList, nil
}