package internal

type Encoding string

const base64 Encoding = "base64"

type Text struct {
	contentType *string   // optional
	encoding    *Encoding // optional
	content     string    // required
}
