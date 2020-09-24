package arm2pulumi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
				apiVersion:   "2020-06-01-preview",
			},
			want: "azure-nextgen:documentdb/v20200601preview:DatabaseAccount",
		},
		{
			name: "SubResource",
			args: args{
				resourceType: "Microsoft.Web/sites/sourcecontrols",
				apiVersion:   "2019-08-01",
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := toResourceToken(tt.args.resourceType, tt.args.apiVersion)
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}
