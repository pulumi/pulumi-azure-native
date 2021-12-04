


package v20190101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdvancedThreatProtection(ctx *pulumi.Context, args *LookupAdvancedThreatProtectionArgs, opts ...pulumi.InvokeOption) (*LookupAdvancedThreatProtectionResult, error) {
	var rv LookupAdvancedThreatProtectionResult
	err := ctx.Invoke("azure-native:security/v20190101:getAdvancedThreatProtection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdvancedThreatProtectionArgs struct {
	ResourceId  string `pulumi:"resourceId"`
	SettingName string `pulumi:"settingName"`
}


type LookupAdvancedThreatProtectionResult struct {
	Id        string `pulumi:"id"`
	IsEnabled *bool  `pulumi:"isEnabled"`
	Name      string `pulumi:"name"`
	Type      string `pulumi:"type"`
}
