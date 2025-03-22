package internal

type Acknowledgement string

const (
	// Declared licenses represent the initial intentions of authors regarding the licensing terms of their code.
	declared Acknowledgement = "declared"
	// Concluded licenses are verified and confirmed.
	concluded Acknowledgement = "concluded"
)

type LicenseType string

const (
	academic        LicenseType = "academic"
	appliance       LicenseType = "appliance"
	clientAccess    LicenseType = "client-access"
	concurrentUser  LicenseType = "concurrent-user"
	corePoints      LicenseType = "core-points"
	customMetric    LicenseType = "custom-metric"
	licenseDevice   LicenseType = "device"
	evaluation      LicenseType = "evaluation"
	namedUser       LicenseType = "named-user"
	nodeLocked      LicenseType = "node-locked"
	oem             LicenseType = "oem"
	perpetual       LicenseType = "perpetual"
	processorPoints LicenseType = "processor-points"
	subscription    LicenseType = "subscription"
	user            LicenseType = "user"
	licenseOther    LicenseType = "other"
)

type Licensing struct {
	AltIds *[]string
	// TODO: implement validation to check if it's either OrganizationalEntity or OrganizationalContact
	Licensor *any
	// TODO: implement validation to check if it's either OrganizationalEntity or OrganizationalContact
	Licensee *any
	// TODO: implement validation to check if it's either OrganizationalEntity or OrganizationalContact
	Purchaser     *any
	PurchaseOrder *string
	LicenseTypes  *[]LicenseType
	LastRenewal   *string
	Expiration    *string
}

type Property struct {
	Name  string  // required
	Value *string // optional
}

type GenericLicense interface {
	License | SPDXLicenseExpression
}

// TODO: implement validation function so that Licenses have either an SPDX ID or a name
type License struct {
	// TODO: get all SPDX IDs
	Id              *string          // optional
	Name            *string          // optional
	BomRef          *string          // optional
	Acknowledgement *Acknowledgement // optional
	Text            *Text            // optional
	Url             *string          // optional
	LicenseType     *Licensing       // optional
	Properties      *[]Property      // optional
}

type SPDXLicenseExpression struct {
	Expression      string           // required
	Acknowledgement *Acknowledgement // optional
	BomRef          *string          // optional
}
