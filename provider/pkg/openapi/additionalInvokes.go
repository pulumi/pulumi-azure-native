package openapi

import (
	"strings"

	"github.com/go-openapi/spec"
)

// We only create getFoo or listFoo functions in combination with a Foo resource. We don't turn all
// GET endpoints to functions. The assumption is that most of those aren't very useful. If we do want
// to include specific ones, we add them here. See, e.g., #2419.
func shouldIncludeGETInvoke(path string, op *spec.Operation) bool {
	return op != nil &&
		!op.Deprecated &&

		// #2419
		strings.HasSuffix(path, "/providers/Microsoft.Insights/diagnosticSettingsCategories") &&
		op.OperationProps.ID == "DiagnosticSettingsCategory_List"
}
