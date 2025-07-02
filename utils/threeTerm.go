package utils

func ThreeTerm(condition bool, True interface{}, False interface{}) interface{} {
	if condition {
		return True
	} else {
		return False
	}
}

func ThreeTermString(condition bool, True string, False string) string {
	if condition {
		return True
	} else {
		return False
	}
}

func ThreeTermFloat64(condition bool, True float64, False float64) float64 {
	if condition {
		return True
	} else {
		return False
	}
}

func ThreeTermFloat32(condition bool, True float32, False float32) float32 {
	if condition {
		return True
	} else {
		return False
	}
}

func ThreeTermInt64(condition bool, True int64, False int64) int64 {
	if condition {
		return True
	} else {
		return False
	}
}

func ThreeTermInt32(condition bool, True int32, False int32) int32 {
	if condition {
		return True
	} else {
		return False
	}
}

func ThreeTermInt(condition bool, True int, False int) int {
	if condition {
		return True
	} else {
		return False
	}
}
