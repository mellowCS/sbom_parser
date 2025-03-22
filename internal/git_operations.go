package internal

type GitUser struct {
	Timestamp *string // optional
	Name      *string // optional
	Email     *string // optional
}

type PatchType string

const (
	unofficial PatchType = "unofficial"
	monkey     PatchType = "monkey"
	backport   PatchType = "backport"
	cherryPick PatchType = "cherry-pick"
)

type Diff struct {
	Text *Text   // optional
	Url  *string // optional
}

type ResolveType string

const (
	defect      ResolveType = "defect"
	enhancement ResolveType = "enhancement"
	security    ResolveType = "security"
)

type Source struct {
	Name *string // optional
	Url  *string // optional
}

type Resolves struct {
	Type        ResolveType // required
	Id          *string     // optional
	Name        *string     // optional
	Description *string     // optional
	Source      *Source     // optional
	References  *[]string   // optional
}

type Patch struct {
	Type     PatchType   // required
	Diff     *Diff       // optional
	Resolves *[]Resolves // optional
}
