package set

// Make new set
func Make(a []string) []string {
	set := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if itemExists(set, a[i]) {
			continue
		}
		set = append(set, a[i])
	}
	return set
}

// Intersection ...
func Intersection(a []string, b []string) []string {
	set := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if contains(b, a[i]) {
			set = append(set, a[i])
		}
	}
	return set
}

//Difference ...
func Difference(a []string, b []string) []string {
	set := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if !contains(b, a[i]) {
			set = append(set, a[i])
		}
	}
	return set
}

func contains(a []string, e string) bool {

	for i := 0; i < len(a); i++ {
		if a[i] == e {
			return true
		}
	}
	return false
}

func itemExists(array []string, item string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}
