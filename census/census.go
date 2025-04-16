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
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Address["street"] == "" || r.Name == "" {
		return false
	} else {
		return true
	}
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	var name string
	var age int 
	var address map[string]string
	r.Name = name
	r.Age = age
	r.Address = address
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	counter := 0
	for _,x := range residents {
		if x.HasRequiredInfo(){
			counter++
		}
	}
	return counter
}
