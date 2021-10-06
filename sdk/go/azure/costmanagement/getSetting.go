


package costmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSetting(ctx *pulumi.Context, args *LookupSettingArgs, opts ...pulumi.InvokeOption) (*LookupSettingResult, error) {
	var rv LookupSettingResult
	err := ctx.Invoke("azure-native:costmanagement:getSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSettingArgs struct {
	SettingName string `pulumi:"settingName"`
}


type LookupSettingResult struct {
	Cache   []SettingsPropertiesResponseCache `pulumi:"cache"`
	Id      string                            `pulumi:"id"`
	Kind    string                            `pulumi:"kind"`
	Name    string                            `pulumi:"name"`
	Scope   string                            `pulumi:"scope"`
	StartOn *string                           `pulumi:"startOn"`
	Type    string                            `pulumi:"type"`
}
