


package v20200808preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupControllerDetails(ctx *pulumi.Context, args *LookupControllerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupControllerDetailsResult, error) {
	var rv LookupControllerDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork/v20200808preview:getControllerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControllerDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupControllerDetailsResult struct {
	DncAppId          string            `pulumi:"dncAppId"`
	DncEndpoint       string            `pulumi:"dncEndpoint"`
	DncTenantId       string            `pulumi:"dncTenantId"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ResourceGuid      string            `pulumi:"resourceGuid"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupControllerDetailsOutput(ctx *pulumi.Context, args LookupControllerDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupControllerDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupControllerDetailsResult, error) {
			args := v.(LookupControllerDetailsArgs)
			r, err := LookupControllerDetails(ctx, &args, opts...)
			var s LookupControllerDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupControllerDetailsResultOutput)
}

type LookupControllerDetailsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupControllerDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerDetailsArgs)(nil)).Elem()
}


type LookupControllerDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupControllerDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerDetailsResult)(nil)).Elem()
}

func (o LookupControllerDetailsResultOutput) ToLookupControllerDetailsResultOutput() LookupControllerDetailsResultOutput {
	return o
}

func (o LookupControllerDetailsResultOutput) ToLookupControllerDetailsResultOutputWithContext(ctx context.Context) LookupControllerDetailsResultOutput {
	return o
}

func (o LookupControllerDetailsResultOutput) DncAppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.DncAppId }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) DncEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.DncEndpoint }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) DncTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.DncTenantId }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupControllerDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupControllerDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupControllerDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControllerDetailsResultOutput{})
}
