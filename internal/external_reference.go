package internal

type ExternalReferenceType string

const (
	vcs                     ExternalReferenceType = "vcs"
	issueTracker            ExternalReferenceType = "issue-tracker"
	website                 ExternalReferenceType = "website"
	advisories              ExternalReferenceType = "advisories"
	bom                     ExternalReferenceType = "bom"
	mailingList             ExternalReferenceType = "mailing-list"
	social                  ExternalReferenceType = "social"
	chat                    ExternalReferenceType = "chat"
	documentation           ExternalReferenceType = "documentation"
	support                 ExternalReferenceType = "support"
	sourceDistribution      ExternalReferenceType = "source-distribution"
	distribution            ExternalReferenceType = "distribution"
	distributionIntake      ExternalReferenceType = "distribution-intake"
	license                 ExternalReferenceType = "license"
	buildMeta               ExternalReferenceType = "build-meta"
	buildSystem             ExternalReferenceType = "build-system"
	releaseNotes            ExternalReferenceType = "release-notes"
	securityContact         ExternalReferenceType = "security-contact"
	modelCard               ExternalReferenceType = "model-card"
	log                     ExternalReferenceType = "log"
	configuration           ExternalReferenceType = "configuration"
	evidence                ExternalReferenceType = "evidence"
	formulation             ExternalReferenceType = "formulation"
	attestation             ExternalReferenceType = "attestation"
	threatModel             ExternalReferenceType = "threat-model"
	adversaryModel          ExternalReferenceType = "adversary-model"
	riskAssessment          ExternalReferenceType = "risk-assessment"
	vulnerabilityAssertion  ExternalReferenceType = "vulnerability-assertion"
	exploitabilityStatement ExternalReferenceType = "exploitability-statement"
	pentestReport           ExternalReferenceType = "pentest-report"
	staticAnalysisReport    ExternalReferenceType = "static-analysis-report"
	dynamicAnalysisReport   ExternalReferenceType = "dynamic-analysis-report"
	runtimeAnalysisReport   ExternalReferenceType = "runtime-analysis-report"
	componentAnalysisReport ExternalReferenceType = "component-analysis-report"
	maturityReport          ExternalReferenceType = "maturity-report"
	certificationReport     ExternalReferenceType = "certification-report"
	codifiedInfrastructure  ExternalReferenceType = "codified-infrastructure"
	qualityMetrics          ExternalReferenceType = "quality-metrics"
	poam                    ExternalReferenceType = "poam"
	electronicSignature     ExternalReferenceType = "electronic-signature"
	digitalSignature        ExternalReferenceType = "digital-signature"
	rfc9116                 ExternalReferenceType = "rfc-9116"
	other                   ExternalReferenceType = "other"
)

type ExternalReference struct {
	// TODO: Validate between URL, BOM-Link Document and BOM-Link Element
	Url     string                // required
	Comment *string               // optional
	Type    ExternalReferenceType // required
	Hashes  *[]Hash               // optional
}
