package main

func SearchByIndex(people []Person, value int) (*Person, bool) {

	for _, p := range people {
		if p.ID == value {
			return &p, true
		}
	}
	return nil, false
}
