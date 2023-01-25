


package v20221102preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrustedAccessRoleBinding(ctx *pulumi.Context, args *LookupTrustedAccessRoleBindingArgs, opts ...pulumi.InvokeOption) (*LookupTrustedAccessRoleBindingResult, error) {
	var rv LookupTrustedAccessRoleBindingResult
	err := ctx.Invoke("azure-native:containerservice/v20221102preview:getTrustedAccessRoleBinding", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrustedAccessRoleBindingArgs struct {
	ResourceGroupName            string `pulumi:"resourceGroupName"`
	ResourceName                 string `pulumi:"resourceName"`
	TrustedAccessRoleBindingName string `pulumi:"trustedAccessRoleBindingName"`
}


type LookupTrustedAccessRoleBindingResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Roles             []string           `pulumi:"roles"`
	SourceResourceId  string             `pulumi:"sourceResourceId"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupTrustedAccessRoleBindingOutput(ctx *pulumi.Context, args LookupTrustedAccessRoleBindingOutputArgs, opts ...pulumi.InvokeOption) LookupTrustedAccessRoleBindingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrustedAccessRoleBindingResult, error) {
			args := v.(LookupTrustedAccessRoleBindingArgs)
			r, err := LookupTrustedAccessRoleBinding(ctx, &args, opts...)
			var s LookupTrustedAccessRoleBindingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrustedAccessRoleBindingResultOutput)
}

type LookupTrustedAccessRoleBindingOutputArgs struct {
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                 pulumi.StringInput `pulumi:"resourceName"`
	TrustedAccessRoleBindingName pulumi.StringInput `pulumi:"trustedAccessRoleBindingName"`
}

func (LookupTrustedAccessRoleBindingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustedAccessRoleBindingArgs)(nil)).Elem()
}


type LookupTrustedAccessRoleBindingResultOutput struct{ *pulumi.OutputState }

func (LookupTrustedAccessRoleBindingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustedAccessRoleBindingResult)(nil)).Elem()
}

func (o LookupTrustedAccessRoleBindingResultOutput) ToLookupTrustedAccessRoleBindingResultOutput() LookupTrustedAccessRoleBindingResultOutput {
	return o
}

func (o LookupTrustedAccessRoleBindingResultOutput) ToLookupTrustedAccessRoleBindingResultOutputWithContext(ctx context.Context) LookupTrustedAccessRoleBindingResultOutput {
	return o
}

func (o LookupTrustedAccessRoleBindingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTrustedAccessRoleBindingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTrustedAccessRoleBindingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTrustedAccessRoleBindingResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o LookupTrustedAccessRoleBindingResultOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.SourceResourceId }).(pulumi.StringOutput)
}

func (o LookupTrustedAccessRoleBindingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTrustedAccessRoleBindingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedAccessRoleBindingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrustedAccessRoleBindingResultOutput{})
}
