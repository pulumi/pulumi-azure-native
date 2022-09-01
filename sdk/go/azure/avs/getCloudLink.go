


package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudLink(ctx *pulumi.Context, args *LookupCloudLinkArgs, opts ...pulumi.InvokeOption) (*LookupCloudLinkResult, error) {
	var rv LookupCloudLinkResult
	err := ctx.Invoke("azure-native:avs:getCloudLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudLinkArgs struct {
	CloudLinkName     string `pulumi:"cloudLinkName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCloudLinkResult struct {
	Id          string  `pulumi:"id"`
	LinkedCloud *string `pulumi:"linkedCloud"`
	Name        string  `pulumi:"name"`
	Status      string  `pulumi:"status"`
	Type        string  `pulumi:"type"`
}

func LookupCloudLinkOutput(ctx *pulumi.Context, args LookupCloudLinkOutputArgs, opts ...pulumi.InvokeOption) LookupCloudLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudLinkResult, error) {
			args := v.(LookupCloudLinkArgs)
			r, err := LookupCloudLink(ctx, &args, opts...)
			var s LookupCloudLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudLinkResultOutput)
}

type LookupCloudLinkOutputArgs struct {
	CloudLinkName     pulumi.StringInput `pulumi:"cloudLinkName"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCloudLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudLinkArgs)(nil)).Elem()
}


type LookupCloudLinkResultOutput struct{ *pulumi.OutputState }

func (LookupCloudLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudLinkResult)(nil)).Elem()
}

func (o LookupCloudLinkResultOutput) ToLookupCloudLinkResultOutput() LookupCloudLinkResultOutput {
	return o
}

func (o LookupCloudLinkResultOutput) ToLookupCloudLinkResultOutputWithContext(ctx context.Context) LookupCloudLinkResultOutput {
	return o
}

func (o LookupCloudLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudLinkResultOutput) LinkedCloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudLinkResult) *string { return v.LinkedCloud }).(pulumi.StringPtrOutput)
}

func (o LookupCloudLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudLinkResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudLinkResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCloudLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudLinkResultOutput{})
}
