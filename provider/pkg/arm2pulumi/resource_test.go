package arm2pulumi

import (
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/arm2pulumi/internal/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddResource(t *testing.T) {
	pkgSpec = test.LoadSchema(t)
	metadata = test.LoadMetadata(t)

	for _, testCase := range []struct{
		name string
		input map[string]interface{}
		expectedResourceName string
		expectedResourceToken string
		expectedResourceBody string
		err error
	} {
		{
			name: "DeepNesting",
			expectedResourceName: "rancherSecurityGroup",
				expectedResourceToken: "azure-nextgen:network/v20200501:NetworkSecurityGroup",
				expectedResourceBody: "",
			input: map[string]interface{}{
				"type": "Microsoft.Network/networkSecurityGroups",
				"apiVersion": "2020-05-01",
				"name": "rancher-security-group",
				"location": "westus2",
				"properties": map[string]interface{}{
					"securityRules": []map[string]interface{}{
						{
							"name": "SSH",
							"properties": map[string]interface{}{
								"description":                "SSH",
								"protocol":                   "*",
								"sourcePortRange":            "*",
								"destinationPortRange":       "22",
								"sourceAddressPrefix":        "*",
								"destinationAddressPrefix":   "*",
								"access":                     "Allow",
								"priority":                   100,
								"direction":                  "Inbound",
							},
						},
						{
							"name": "Docker",
							"properties": map[string]interface{}{
								"description":                "Docker",
								"protocol":                   "Tcp",
								"sourcePortRange":            "*",
								"destinationPortRange":       "2375",
								"sourceAddressPrefix":        "*",
								"destinationAddressPrefix":   "*",
								"access":                     "Allow",
								"priority":                   300,
								"direction":                  "Inbound",
							},
						},
					},
				},
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			templateElements := NewTemplateElements()
			err := templateElements.AddResource(testCase.input)
			if testCase.err != nil {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Len(t, templateElements.elements, 1)
			require.Contains(t, templateElements.elements, testCase.expectedResourceName)
			require.Equal(t, testCase.input, templateElements.elements[testCase.expectedResourceName].TemplateElement.Args())

			body, err := templateElements.RenderPCL(metadata, pkgSpec)
			require.NoError(t, err)
			require.Equal(t, testCase.expectedResourceBody, fmt.Sprintf("%v", body))
		})
	}
}

func Test_toResourceToken(t *testing.T) {
	type args struct {
		resourceType string
		apiVersion   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Preview",
			args: args{
				resourceType: "Microsoft.DocumentDb/databaseAccounts",
				apiVersion: "2020-06-01-preview",
			},
			want: "azure-nextgen:documentdb/v20200601preview:DatabaseAccount",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toResourceToken(tt.args.resourceType, tt.args.apiVersion)
			assert.Equal(t, tt.want, got)
		})
	}
}
