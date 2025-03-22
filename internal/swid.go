package internal

type Swid struct {
	TagId      string  // required
	Name       string  // required
	Version    *string // optional, default: 0.0
	TagVersion *string // optional, default 0
	Patch      *bool   // optional, default: false
	Text       *Text   // optional
	Url        *string // optional
}
