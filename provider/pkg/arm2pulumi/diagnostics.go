package arm2pulumi

import "fmt"

type Severity string

const (
	SevHigh Severity = "High"
	SevMed  Severity = "Med"
	SevLow  Severity = "Low"
)

// Diagnostic collects useful information for partial or complete failure in converting
// an item in the template to pulumi code.
type Diagnostic struct {
	SourceElement string
	SourceToken   string
	Severity      Severity
	Description   string
}

func (d Diagnostic) Error() string {
	return fmt.Sprintf("[%s] Failed converting resource %s of type %s. Desc: %s",
		d.Severity, d.SourceElement, d.SourceToken, d.Description)
}

func (t *TemplateElements) addDiagnostic(diagnostic Diagnostic) {
	diags := t.diagnostics[diagnostic.SourceElement]
	for _, d := range diags {
		if d == diagnostic {
			return
		}
	}

	diags = append(diags, diagnostic)
	t.diagnostics[diagnostic.SourceElement] = diags
}
