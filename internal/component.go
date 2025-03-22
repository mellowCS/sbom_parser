package internal

type Commit struct {
	Uid       *string  // optional
	Url       *string  // optional
	Author    *GitUser // optional
	Committer *GitUser // optional
	Message   *string  // optional
}

type Pedigree struct {
	Ancestors   *[]Component // optional
	Descendants *[]Component // optional
	Variants    *[]Component // optional
	Commits     *[]Commit    // optional
	Patches     *[]Patch     // optional
	Notes       *string      // optional
}

type ComponentType string

const (
	application          ComponentType = "application"
	framework            ComponentType = "framework"
	library              ComponentType = "library"
	container            ComponentType = "container"
	platform             ComponentType = "platform"
	operatingSystem      ComponentType = "operating-system"
	device               ComponentType = "device"
	deviceDriver         ComponentType = "device-driver"
	firmware             ComponentType = "firmware"
	file                 ComponentType = "file"
	machineLearningModel ComponentType = "machine-learning-model"
	data                 ComponentType = "data"
	cryptographicAsset   ComponentType = "cryptographic-asset"
)

type Component struct {
	Type         ComponentType            // required
	MimeType     *string                  // optional
	BomRef       *string                  // optional
	Supplier     *OrganizationalEntity    // optional
	Manufacturer *OrganizationalEntity    // optional
	Authors      *[]OrganizationalContact // optional
	Author       *string                  // optional and deprecated
	Publisher    *string                  // optional
	Group        *string                  // optional
	Name         string                   // required
	Version      *string                  // optional
	Description  *string                  // optional
	Scope        *Scope                   // optional
	Hashes       *[]Hash                  // required
	// TODO: write a custom json unmarshal function to deserialize either license or spdxLicenseExpression
	Licenses           *[]any               // optional
	Copyright          *string              // optional
	Cpe                *string              // optional
	Purl               *string              // optional
	Omniborld          *[]string            // optional
	Swhid              *[]string            // optional
	Swid               *Swid                // optional
	Modified           *bool                // optional and deprecated
	Pedigree           *Pedigree            // optional
	ExternalReferences *[]ExternalReference // optional
	Components         *[]Component         // optional

}
