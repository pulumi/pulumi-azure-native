


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutoscaleSetting(ctx *pulumi.Context, args *LookupAutoscaleSettingArgs, opts ...pulumi.InvokeOption) (*LookupAutoscaleSettingResult, error) {
	var rv LookupAutoscaleSettingResult
	err := ctx.Invoke("azure-native:insights/v20210501preview:getAutoscaleSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAutoscaleSettingArgs struct {
	AutoscaleSettingName string `pulumi:"autoscaleSettingName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupAutoscaleSettingResult struct {
	Enabled                   *bool                              `pulumi:"enabled"`
	Id                        string                             `pulumi:"id"`
	Location                  string                             `pulumi:"location"`
	Name                      string                             `pulumi:"name"`
	Notifications             []AutoscaleNotificationResponse    `pulumi:"notifications"`
	PredictiveAutoscalePolicy *PredictiveAutoscalePolicyResponse `pulumi:"predictiveAutoscalePolicy"`
	Profiles                  []AutoscaleProfileResponse         `pulumi:"profiles"`
	SystemData                SystemDataResponse                 `pulumi:"systemData"`
	Tags                      map[string]string                  `pulumi:"tags"`
	TargetResourceLocation    *string                            `pulumi:"targetResourceLocation"`
	TargetResourceUri         *string                            `pulumi:"targetResourceUri"`
	Type                      string                             `pulumi:"type"`
}


func (val *LookupAutoscaleSettingResult) Defaults() *LookupAutoscaleSettingResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := true
		tmp.Enabled = &enabled_
	}
	return &tmp
}
