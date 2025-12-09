package generator

import "fmt"

// Address represent an address
type Address struct {
	Address    string `json:"address,omitempty" validate:"required"`
	Address2   string `json:"address_2,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
}

// ToString output address as string
// Line break are added for new lines
func (a *Address) ToString() string {
	addrString := a.Address

	if len(a.Address2) > 0 {
		addrString += "\n"
		addrString += a.Address2
	}

	if len(a.State) > 0 {
		addrString += fmt.Sprintf("\n%s, %s %s", a.City, a.State, a.PostalCode)
	} else {
		if len(a.PostalCode) > 0 {
			addrString += "\n"
			addrString += a.PostalCode
		} else {
			addrString += "\n"
		}

		if len(a.City) > 0 {
			addrString += " "
			addrString += a.City
		}
	}

	if len(a.Country) > 0 {
		addrString += "\n"
		addrString += a.Country
	}

	return addrString
}

func (a *Address) Lines() int {
	lines := 0
	if len(a.Address) > 0 {
		lines++
	}

	if len(a.Address2) > 0 {
		lines++
	}

	if len(a.State) > 0 {
		lines++
	} else {
		if len(a.PostalCode) > 0 {
			lines++
		} else {
			lines++
		}

		if len(a.City) > 0 {
			lines++
		}
	}

	if len(a.Country) > 0 {
		lines++
	}

	return lines
}
