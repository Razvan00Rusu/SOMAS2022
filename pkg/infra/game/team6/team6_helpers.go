package team6

type MapDefault map[string]uint

func (m MapDefault) GetValOrDefault(key string, val uint) (result uint) {
	if v, ok := m[key]; ok {
		return v
	} else {
		m[key] = val
		return val
	}
}

func SubUpToMinimum(a uint, b uint, min uint) uint {
	if a-b < min {
		return min
	} else {
		return a - b
	}
}

func AddUpToMaximum(a uint, b uint, max uint) uint {
	if a+b > max {
		return max
	} else {
		return a + b
	}
}
