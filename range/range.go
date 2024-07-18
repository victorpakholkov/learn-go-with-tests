package range_

func Range(literal string, timesRepeated int) (resultingString string) {
	for range timesRepeated {
		resultingString += literal
	}
	return resultingString
}
