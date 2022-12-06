


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRosettaNetProcessConfiguration(ctx *pulumi.Context, args *LookupRosettaNetProcessConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupRosettaNetProcessConfigurationResult, error) {
	var rv LookupRosettaNetProcessConfigurationResult
	err := ctx.Invoke("azure-native:logic/v20160601:getRosettaNetProcessConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRosettaNetProcessConfigurationArgs struct {
	IntegrationAccountName             string `pulumi:"integrationAccountName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
	RosettaNetProcessConfigurationName string `pulumi:"rosettaNetProcessConfigurationName"`
}


type LookupRosettaNetProcessConfigurationResult struct {
	ActivitySettings      RosettaNetPipActivitySettingsResponse `pulumi:"activitySettings"`
	ChangedTime           string                                `pulumi:"changedTime"`
	CreatedTime           string                                `pulumi:"createdTime"`
	Description           *string                               `pulumi:"description"`
	Id                    string                                `pulumi:"id"`
	InitiatorRoleSettings RosettaNetPipRoleSettingsResponse     `pulumi:"initiatorRoleSettings"`
	Location              *string                               `pulumi:"location"`
	Metadata              map[string]string                     `pulumi:"metadata"`
	Name                  string                                `pulumi:"name"`
	ProcessCode           string                                `pulumi:"processCode"`
	ProcessName           string                                `pulumi:"processName"`
	ProcessVersion        string                                `pulumi:"processVersion"`
	ResponderRoleSettings RosettaNetPipRoleSettingsResponse     `pulumi:"responderRoleSettings"`
	Tags                  map[string]string                     `pulumi:"tags"`
	Type                  string                                `pulumi:"type"`
}

func LookupRosettaNetProcessConfigurationOutput(ctx *pulumi.Context, args LookupRosettaNetProcessConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupRosettaNetProcessConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRosettaNetProcessConfigurationResult, error) {
			args := v.(LookupRosettaNetProcessConfigurationArgs)
			r, err := LookupRosettaNetProcessConfiguration(ctx, &args, opts...)
			var s LookupRosettaNetProcessConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRosettaNetProcessConfigurationResultOutput)
}

type LookupRosettaNetProcessConfigurationOutputArgs struct {
	IntegrationAccountName             pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName                  pulumi.StringInput `pulumi:"resourceGroupName"`
	RosettaNetProcessConfigurationName pulumi.StringInput `pulumi:"rosettaNetProcessConfigurationName"`
}

func (LookupRosettaNetProcessConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRosettaNetProcessConfigurationArgs)(nil)).Elem()
}


type LookupRosettaNetProcessConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupRosettaNetProcessConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRosettaNetProcessConfigurationResult)(nil)).Elem()
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ToLookupRosettaNetProcessConfigurationResultOutput() LookupRosettaNetProcessConfigurationResultOutput {
	return o
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ToLookupRosettaNetProcessConfigurationResultOutputWithContext(ctx context.Context) LookupRosettaNetProcessConfigurationResultOutput {
	return o
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ActivitySettings() RosettaNetPipActivitySettingsResponseOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) RosettaNetPipActivitySettingsResponse {
		return v.ActivitySettings
	}).(RosettaNetPipActivitySettingsResponseOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) InitiatorRoleSettings() RosettaNetPipRoleSettingsResponseOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) RosettaNetPipRoleSettingsResponse {
		return v.InitiatorRoleSettings
	}).(RosettaNetPipRoleSettingsResponseOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ProcessCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.ProcessCode }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ProcessName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.ProcessName }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ProcessVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.ProcessVersion }).(pulumi.StringOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) ResponderRoleSettings() RosettaNetPipRoleSettingsResponseOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) RosettaNetPipRoleSettingsResponse {
		return v.ResponderRoleSettings
	}).(RosettaNetPipRoleSettingsResponseOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRosettaNetProcessConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRosettaNetProcessConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRosettaNetProcessConfigurationResultOutput{})
}
