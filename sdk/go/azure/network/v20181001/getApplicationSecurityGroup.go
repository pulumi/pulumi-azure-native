


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationSecurityGroup(ctx *pulumi.Context, args *LookupApplicationSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSecurityGroupResult, error) {
	var rv LookupApplicationSecurityGroupResult
	err := ctx.Invoke("azure-native:network/v20181001:getApplicationSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationSecurityGroupArgs struct {
	ApplicationSecurityGroupName string `pulumi:"applicationSecurityGroupName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupApplicationSecurityGroupResult struct {
	Etag              string            `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ResourceGuid      string            `pulumi:"resourceGuid"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupApplicationSecurityGroupOutput(ctx *pulumi.Context, args LookupApplicationSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationSecurityGroupResult, error) {
			args := v.(LookupApplicationSecurityGroupArgs)
			r, err := LookupApplicationSecurityGroup(ctx, &args, opts...)
			var s LookupApplicationSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationSecurityGroupResultOutput)
}

type LookupApplicationSecurityGroupOutputArgs struct {
	ApplicationSecurityGroupName pulumi.StringInput `pulumi:"applicationSecurityGroupName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSecurityGroupArgs)(nil)).Elem()
}


type LookupApplicationSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSecurityGroupResult)(nil)).Elem()
}

func (o LookupApplicationSecurityGroupResultOutput) ToLookupApplicationSecurityGroupResultOutput() LookupApplicationSecurityGroupResultOutput {
	return o
}

func (o LookupApplicationSecurityGroupResultOutput) ToLookupApplicationSecurityGroupResultOutputWithContext(ctx context.Context) LookupApplicationSecurityGroupResultOutput {
	return o
}

func (o LookupApplicationSecurityGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationSecurityGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSecurityGroupResultOutput{})
}
