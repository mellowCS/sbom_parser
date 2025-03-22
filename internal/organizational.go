package internal

type OrganizationalContact struct {
	BomRef *string // optional
	Name   *string // optional
	Email  *string // optional
	Phone  *string // optional
}

type Address struct {
	BomRef              *string // optional
	Country             *string // optional
	Region              *string // optional
	Locality            *string // optional
	PostOfficeBoxNumber *string // optional
	PostalCode          *string // optional
	StreetAddress       *string // optional
}

type OrganizationalEntity struct {
	BomRef  *string                  // optional
	Name    *string                  // optional
	Address *Address                 // optional
	Url     *[]string                // optional
	Contact *[]OrganizationalContact // optional
}
