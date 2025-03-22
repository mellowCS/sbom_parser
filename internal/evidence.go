package internal

type Field string

const (
	fieldGroup Field = "group"
	name       Field = "name"
	version    Field = "version"
	purl       Field = "purl"
	cpe        Field = "cpe"
	omniborld  Field = "omniborld"
	swhid      Field = "swhid"
	swid       Field = "swid"
	hash       Field = "hash"
)

type Technique string

const (
	sourceCodeAnalysis   Technique = "source-code-analysis"
	binaryAnalysis       Technique = "binary-analysis"
	manifestAnalysis     Technique = "manifest-analysis"
	astFingerprint       Technique = "ast-fingerprint"
	hashComparison       Technique = "hash-comparison"
	instrumentation      Technique = "instrumentation"
	dynamicAnalysis      Technique = "dynamic-analysis"
	filename             Technique = "filename"
	techniqueAttestation Technique = "attestation"
	techniqueOther       Technique = "other"
)

type Method struct {
	Technique  Technique // required
	Confidence float64   // required
	Value      *string   // optional
}

type Identity struct {
	Field          Field
	Confidence     *float64  // optional
	ConcludedValue *string   // optional
	Methods        *[]Method // optional
	// TODO: Validate that it is either BOM ref or BOM-Link Element
	Tools *[]string // optional
}

type Occurrence struct {
	BomRef   *string // optional
	Location string  // required
	Line     int
}

type Evidence struct {
	Identity *[]Identity // optional

}
