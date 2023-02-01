package utils

func FindAndAppend(array []any, value any) []any {

	if _, ok := FindExistInArray(array, value); ok {
		return array
	}

	array = append(array, value)
	return array
}

func FindAndDelete(array []any, value any) []any {
	i, ok := FindExistInArray(array, value)
	if !ok {
		return array
	}

	array[i] = array[len(array)-1]
	array[len(array)-1] = nil
	array = array[:len(array)-1]

	return array
}

func FindExistInArray(array []any, value any) (int, bool) {
	for i, v := range array {
		if v == value {
			return i, true
		}
	}

	return -1, false
}

func FindAndAppendInt(array []int, value int) []int {

	if _, ok := FindExistInArrayInt(array, value); ok {
		return array
	}

	array = append(array, value)
	return array
}

func FindAndDeleteInt(array []int, value int) []int {
	i, ok := FindExistInArrayInt(array, value)
	if !ok {
		return array
	}

	array[i] = array[len(array)-1]
	array[len(array)-1] = 0
	array = array[:len(array)-1]

	return array
}

func FindExistInArrayInt(array []int, value int) (int, bool) {
	for i, v := range array {
		if v == value {
			return i, true
		}
	}

	return -1, false
}
