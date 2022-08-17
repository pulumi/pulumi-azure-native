


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLabAccount(ctx *pulumi.Context, args *LookupLabAccountArgs, opts ...pulumi.InvokeOption) (*LookupLabAccountResult, error) {
	var rv LookupLabAccountResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getLabAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabAccountArgs struct {
	Expand            *string `pulumi:"expand"`
	LabAccountName    string  `pulumi:"labAccountName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLabAccountResult struct {
	EnabledRegionSelection *bool                               `pulumi:"enabledRegionSelection"`
	Id                     string                              `pulumi:"id"`
	LatestOperationResult  LatestOperationResultResponse       `pulumi:"latestOperationResult"`
	Location               *string                             `pulumi:"location"`
	Name                   string                              `pulumi:"name"`
	ProvisioningState      *string                             `pulumi:"provisioningState"`
	SizeConfiguration      SizeConfigurationPropertiesResponse `pulumi:"sizeConfiguration"`
	Tags                   map[string]string                   `pulumi:"tags"`
	Type                   string                              `pulumi:"type"`
	UniqueIdentifier       *string                             `pulumi:"uniqueIdentifier"`
}

func LookupLabAccountOutput(ctx *pulumi.Context, args LookupLabAccountOutputArgs, opts ...pulumi.InvokeOption) LookupLabAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabAccountResult, error) {
			args := v.(LookupLabAccountArgs)
			r, err := LookupLabAccount(ctx, &args, opts...)
			var s LookupLabAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabAccountResultOutput)
}

type LookupLabAccountOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabAccountName    pulumi.StringInput    `pulumi:"labAccountName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupLabAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabAccountArgs)(nil)).Elem()
}


type LookupLabAccountResultOutput struct{ *pulumi.OutputState }

func (LookupLabAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabAccountResult)(nil)).Elem()
}

func (o LookupLabAccountResultOutput) ToLookupLabAccountResultOutput() LookupLabAccountResultOutput {
	return o
}

func (o LookupLabAccountResultOutput) ToLookupLabAccountResultOutputWithContext(ctx context.Context) LookupLabAccountResultOutput {
	return o
}

func (o LookupLabAccountResultOutput) EnabledRegionSelection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLabAccountResult) *bool { return v.EnabledRegionSelection }).(pulumi.BoolPtrOutput)
}

func (o LookupLabAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabAccountResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupLabAccountResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LookupLabAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLabAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabAccountResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabAccountResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupLabAccountResultOutput) SizeConfiguration() SizeConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupLabAccountResult) SizeConfigurationPropertiesResponse { return v.SizeConfiguration }).(SizeConfigurationPropertiesResponseOutput)
}

func (o LookupLabAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLabAccountResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabAccountResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabAccountResultOutput{})
}
