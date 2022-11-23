


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:azurestackhci/v20220501:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	ArcSettingName    string `pulumi:"arcSettingName"`
	ClusterName       string `pulumi:"clusterName"`
	ExtensionName     string `pulumi:"extensionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	AggregateState          string                          `pulumi:"aggregateState"`
	AutoUpgradeMinorVersion *bool                           `pulumi:"autoUpgradeMinorVersion"`
	CreatedAt               *string                         `pulumi:"createdAt"`
	CreatedBy               *string                         `pulumi:"createdBy"`
	CreatedByType           *string                         `pulumi:"createdByType"`
	ForceUpdateTag          *string                         `pulumi:"forceUpdateTag"`
	Id                      string                          `pulumi:"id"`
	LastModifiedAt          *string                         `pulumi:"lastModifiedAt"`
	LastModifiedBy          *string                         `pulumi:"lastModifiedBy"`
	LastModifiedByType      *string                         `pulumi:"lastModifiedByType"`
	Name                    string                          `pulumi:"name"`
	PerNodeExtensionDetails []PerNodeExtensionStateResponse `pulumi:"perNodeExtensionDetails"`
	ProtectedSettings       interface{}                     `pulumi:"protectedSettings"`
	ProvisioningState       string                          `pulumi:"provisioningState"`
	Publisher               *string                         `pulumi:"publisher"`
	Settings                interface{}                     `pulumi:"settings"`
	Type                    string                          `pulumi:"type"`
	TypeHandlerVersion      *string                         `pulumi:"typeHandlerVersion"`
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtensionResult, error) {
			args := v.(LookupExtensionArgs)
			r, err := LookupExtension(ctx, &args, opts...)
			var s LookupExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	ArcSettingName    pulumi.StringInput `pulumi:"arcSettingName"`
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ExtensionName     pulumi.StringInput `pulumi:"extensionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}


type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.AggregateState }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupExtensionResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) PerNodeExtensionDetails() PerNodeExtensionStateResponseArrayOutput {
	return o.ApplyT(func(v LookupExtensionResult) []PerNodeExtensionStateResponse { return v.PerNodeExtensionDetails }).(PerNodeExtensionStateResponseArrayOutput)
}

func (o LookupExtensionResultOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupExtensionResult) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o LookupExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupExtensionResult) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o LookupExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
