package crawler

func Generate() []string {
	return []string{}
}

func checkMembership(set []string, element string) bool {
	var check bool
	for _, e := range set {
		if e == element {
			check = true
			break
		}
	}
	return check
}

func AddElement(set *[]string, element string) {
	if checkMembership(*set, element) == false {
		*set = append(*set, element)
	}
}
