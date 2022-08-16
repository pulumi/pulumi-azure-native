


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentName        string  `pulumi:"environmentName"`
	EnvironmentSettingName string  `pulumi:"environmentSettingName"`
	Expand                 *string `pulumi:"expand"`
	LabAccountName         string  `pulumi:"labAccountName"`
	LabName                string  `pulumi:"labName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentResult struct {
	ClaimedByUserName        string                        `pulumi:"claimedByUserName"`
	ClaimedByUserObjectId    string                        `pulumi:"claimedByUserObjectId"`
	ClaimedByUserPrincipalId string                        `pulumi:"claimedByUserPrincipalId"`
	Id                       string                        `pulumi:"id"`
	IsClaimed                bool                          `pulumi:"isClaimed"`
	LastKnownPowerState      string                        `pulumi:"lastKnownPowerState"`
	LatestOperationResult    LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location                 *string                       `pulumi:"location"`
	Name                     string                        `pulumi:"name"`
	NetworkInterface         NetworkInterfaceResponse      `pulumi:"networkInterface"`
	PasswordLastReset        string                        `pulumi:"passwordLastReset"`
	ProvisioningState        *string                       `pulumi:"provisioningState"`
	ResourceSets             *ResourceSetResponse          `pulumi:"resourceSets"`
	Tags                     map[string]string             `pulumi:"tags"`
	TotalUsage               string                        `pulumi:"totalUsage"`
	Type                     string                        `pulumi:"type"`
	UniqueIdentifier         *string                       `pulumi:"uniqueIdentifier"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	EnvironmentName        pulumi.StringInput    `pulumi:"environmentName"`
	EnvironmentSettingName pulumi.StringInput    `pulumi:"environmentSettingName"`
	Expand                 pulumi.StringPtrInput `pulumi:"expand"`
	LabAccountName         pulumi.StringInput    `pulumi:"labAccountName"`
	LabName                pulumi.StringInput    `pulumi:"labName"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}


type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ClaimedByUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ClaimedByUserName }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) ClaimedByUserObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ClaimedByUserObjectId }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) ClaimedByUserPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ClaimedByUserPrincipalId }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) IsClaimed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) bool { return v.IsClaimed }).(pulumi.BoolOutput)
}

func (o LookupEnvironmentResultOutput) LastKnownPowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.LastKnownPowerState }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LookupEnvironmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) NetworkInterface() NetworkInterfaceResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) NetworkInterfaceResponse { return v.NetworkInterface }).(NetworkInterfaceResponseOutput)
}

func (o LookupEnvironmentResultOutput) PasswordLastReset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.PasswordLastReset }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) ResourceSets() ResourceSetResponsePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *ResourceSetResponse { return v.ResourceSets }).(ResourceSetResponsePtrOutput)
}

func (o LookupEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnvironmentResultOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.TotalUsage }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupEnvironmentResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
