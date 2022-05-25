// Copyright 2021, Pulumi Corporation.  All rights reserved.

//go:build coverage
// +build coverage

package test

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi"
	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

var testOutputDir = flag.String("testOutputDir", "test-results", "location to write raw test output to. Defaults to ./test-results. Creates the folder if it does not exist.")
var retainConverted = flag.Bool("retainConverted", false, "When set to true retains the converted files in 'testOutputDir'")

// These templates currently cause panics.
var excluded = map[string]bool{
	"../testdata/azure-quickstart-templates/101-attestation-provider-create/azuredeploy.json": true,
	"../testdata/azure-quickstart-templates/101-logic-app-sql-proc/azuredeploy.json":          true,
	"../testdata/azure-quickstart-templates/101-private-endpoint-sql/azuredeploy.json":        true,
	// Following templates are removed bc of non-deterministic output: https://github.com/pulumi/arm2pulumi/issues/28
	"../testdata/azure-quickstart-templates/101-asev2-appservice-sql-vpngw/azuredeploy.json": true,
	"../testdata/azure-quickstart-templates/101-functions-managed-identity/azuredeploy.json": true,
	"../testdata/azure-quickstart-templates/101-redis-cache/azuredeploy.json":                true,
	"../testdata/azure-quickstart-templates/101-sql-logical-server/azuredeploy.json":         true,
	"../testdata/azure-quickstart-templates/101-custom-rp-with-function/azuredeploy.json":    true,
}

func TestQuickstartTemplateCoverage(t *testing.T) {
	matches, err := filepath.Glob("../testdata/azure-quickstart-templates/10[01]*/azuredeploy.json")
	require.NoError(t, err)

	require.NoError(t, os.MkdirAll(*testOutputDir, 0700))

	t.Logf("Test outputs will be logged to: %s", *testOutputDir)
	renderer := arm2pulumi.NewRenderer(loadSchema(t), loadMetadata(t))
	var diagList []Diag
	for _, match := range matches {
		t.Run(match, func(t *testing.T) {
			if excluded[match] {
				t.Skip()
			}
			template := filepath.Base(filepath.Dir(match))
			body, diags, err := renderer.RenderFileIR(match, renderOptionsOverride[match]...)
			if err != nil {
				diagList = append(diagList, Diag{
					Diagnostic: arm2pulumi.Diagnostic{
						Severity:    arm2pulumi.SevFatal,
						Description: err.Error(),
					},
					Template: template,
				})
			}

			if err == nil {
				if len(diags) == 0 {
					// Success is represented by no diagnostic information apart from template name.
					diagList = append(diagList, Diag{Template: template})
				}
				for _, v := range diags {
					for _, d := range v {
						diagList = append(diagList, Diag{Template: template, Diagnostic: d})
					}
				}
				if *retainConverted {
					rendered, _, err := renderer.RenderPrograms(body, []string{"nodejs"})
					require.NoError(t, err)
					require.NoError(t, ioutil.WriteFile(filepath.Join(*testOutputDir, template+".ts"), []byte(rendered["nodejs"]), 0600))
				}
			}
		})
	}
	summarize(t, diagList)
}

type Diag struct {
	arm2pulumi.Diagnostic
	Template string `json:"template"`
}

//func TestSummarize(t *testing.T) {
//	*testOutputDir = "<location>"
//	content, err := ioutil.ReadFile(filepath.Join(*testOutputDir, "summary.json"))
//	require.NoError(t, err)
//	var diagList []Diag
//	require.NoError(t, json.Unmarshal(content, &diagList))
//	summarize(t, diagList)
//}

type Result struct {
	Number int
	Pct    float32
}

type OverallResult struct {
	NoErrors, LowSevErrors, HighSevErrors, Fatal Result
	Total                                        int
}

func summarize(t *testing.T, diagList []Diag) {
	jsonOutputLocation := filepath.Join(*testOutputDir, "summary.json")
	marshalled, err := json.MarshalIndent(diagList, "", "\t")
	require.NoError(t, err)
	require.NoError(t, ioutil.WriteFile(jsonOutputLocation, marshalled, 0600))

	db, err := sql.Open("sqlite", filepath.Join(*testOutputDir, "summary.db"))
	require.NoError(t, err)
	_, err = db.Exec(`
DROP TABLE IF EXISTS errors;
CREATE TABLE errors(
    template TEXT NOT NULL,
    severity TEXT,
    source_token TEXT,
    description TEXT,
    source_element TEXT,
    PRIMARY KEY (template, severity, source_token, description, source_element)
);`)
	require.NoError(t, err)

	nullable := func(val string) sql.NullString {
		if val == "" {
			return sql.NullString{}
		}
		return sql.NullString{
			String: val,
			Valid:  true,
		}
	}
	for _, d := range diagList {
		_, err = db.Exec(`INSERT INTO errors(
                   template,
                   severity,
                   source_token,
                   description,
                   source_element
        	) values(?, ?, ?, ?, ?)`,
			d.Template,
			nullable(string(d.Severity)),
			nullable(d.SourceToken),
			nullable(d.Description),
			nullable(d.SourceElement))
		require.NoError(t, err)
	}

	var numTemplates, fatalErrors, highSevErrors, medSevErrors, success int
	row := db.QueryRow(`SELECT COUNT(DISTINCT template) FROM errors`)
	require.NoError(t, row.Scan(&numTemplates))

	row = db.QueryRow(`SELECT COUNT(DISTINCT template) FROM errors WHERE severity='Fatal'`)
	require.NoError(t, row.Scan(&fatalErrors))

	row = db.QueryRow(`SELECT COUNT(DISTINCT template) FROM errors WHERE severity='High'`)
	require.NoError(t, row.Scan(&highSevErrors))

	row = db.QueryRow(`SELECT COUNT(DISTINCT template) FROM errors WHERE severity='Med'`)
	require.NoError(t, row.Scan(&medSevErrors))

	row = db.QueryRow(`SELECT COUNT(DISTINCT template) FROM errors WHERE severity is NULL`)
	require.NoError(t, row.Scan(&success))

	// Stores the overall results in a JSON object to compare in future tests.
	data := OverallResult{
		NoErrors:      Result{success, float32(success) / float32(numTemplates) * 100.0},
		LowSevErrors:  Result{medSevErrors, float32(medSevErrors) / float32(numTemplates) * 100.0},
		HighSevErrors: Result{highSevErrors, float32(highSevErrors) / float32(numTemplates) * 100.0},
		Fatal:         Result{fatalErrors, float32(fatalErrors) / float32(numTemplates) * 100.0},
		Total:         numTemplates,
	}

	file, err := json.MarshalIndent(data, "", "\t")
	require.NoError(t, err)

	// Stores JSON result in "results.json" file in current directory.
	err = ioutil.WriteFile("results.json", file, 0600)
	require.NoError(t, err)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "Overall Summary of Conversions")
	table.SetHeader([]string{"Result", "Number", "Perc."})
	table.Append([]string{"No Errors", fmt.Sprintf("%d", success), fmt.Sprintf("%.2f%%", float32(success)/float32(numTemplates)*100.0)})
	table.Append([]string{"Low Sev. Errors", fmt.Sprintf("%d", medSevErrors), fmt.Sprintf("%.2f%%", float32(medSevErrors)/float32(numTemplates)*100.0)})
	table.Append([]string{"High Sev. Errors", fmt.Sprintf("%d", highSevErrors), fmt.Sprintf("%.2f%%", float32(highSevErrors)/float32(numTemplates)*100.0)})
	table.Append([]string{"Fatal", fmt.Sprintf("%d", fatalErrors), fmt.Sprintf("%.2f%%", float32(fatalErrors)/float32(numTemplates)*100.0)})
	table.Append([]string{"Total", fmt.Sprintf("%d", numTemplates), ""})
	table.Render()
	fmt.Print("\n\n")

	table = tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "Top Reasons For Fatal Errors")
	table.SetRowLine(true)
	table.SetHeader([]string{"# of Times", "Reason"})
	var desc string
	var cnt int
	rows, err := db.Query(`SELECT description, COUNT(*) AS cnt FROM errors WHERE severity='Fatal' GROUP BY description ORDER BY cnt DESC LIMIT 5`)
	require.NoError(t, err)
	for rows.Next() {
		rows.Scan(&desc, &cnt)
		table.Append([]string{fmt.Sprintf("%d", cnt), desc})
	}
	table.Render()
	fmt.Print("\n\n")

	table = tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "Top Reasons For High Sev Errors")
	table.SetRowLine(true)
	table.SetHeader([]string{"# of Times", "Reason"})
	rows, err = db.Query(`SELECT description, COUNT(*) AS cnt FROM errors WHERE severity='High' GROUP BY description ORDER BY cnt DESC LIMIT 5`)
	require.NoError(t, err)
	for rows.Next() {
		rows.Scan(&desc, &cnt)
		table.Append([]string{fmt.Sprintf("%d", cnt), desc})
	}
	table.Render()
}
