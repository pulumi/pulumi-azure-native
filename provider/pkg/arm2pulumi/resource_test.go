package arm2pulumi

import (
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_toResourceToken(t *testing.T) {
	type args struct {
		resourceType string
		apiVersion   string
		alias        string
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
				apiVersion:   "2020-06-01-preview",
			},
			want: "azure-nextgen:documentdb/v20200601preview:DatabaseAccount",
		},
		{
			name: "SubResource",
			args: args{
				resourceType: "Microsoft.Web/sites/sourcecontrols",
				apiVersion:   "2019-08-01",
				alias:        "azure-nextgen:web/v20190801:SiteSourceControl",
			},
			want: "azure-nextgen:web/v20150801:SiteSourceControl",
		},
		{
			name: "SecurityRules",
			args: args{
				resourceType: "Microsoft.Network/networkSecurityGroups/securityRules",
				apiVersion:   "2020-05-01",
			},
			want: "azure-nextgen:network/v20200501:SecurityRule",
		},
		{
			name: "VirtualNetworkLink",
			args: args{
				resourceType: "Microsoft.Network/privateDnsZones/virtualNetworkLinks",
				apiVersion:   "2018-09-01",
			},
			want: "azure-nextgen:network/v20180901:VirtualNetworkLink",
		},
	}

	pkgSpec := schema.PackageSpec{
		Resources: map[string]schema.ResourceSpec{},
	}
	metadata := provider.AzureAPIMetadata{
		Resources: map[string]provider.AzureAPIResource{},
	}

	for i, t := range tests {
		res := schema.ResourceSpec{
			Aliases: []schema.AliasSpec{},
		}
		if t.args.alias != "" {
			res.Aliases = append(res.Aliases, schema.AliasSpec{Type: &tests[i].args.alias})
		}
		pkgSpec.Resources[t.want] = res
		metadata.Resources[t.want] = provider.AzureAPIResource{}
	}

	ctx := pclRenderContext{pkgSpec: &pkgSpec, metadata: &metadata}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := toResourceToken(&ctx, tt.args.resourceType, tt.args.apiVersion)
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}
