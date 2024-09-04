// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{name, age, address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || r.Address == nil {
		return false
	}

	if r.Address["street"] == "" {
		return false
	}

	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Address = nil
	r.Name = ""
	r.Age = 0
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, v := range residents {
		if v.Address != nil && v.Age != 0 && v.Name != "" {
			count++
		}
	}

	return count
}
