


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHcxEnterpriseSite(ctx *pulumi.Context, args *LookupHcxEnterpriseSiteArgs, opts ...pulumi.InvokeOption) (*LookupHcxEnterpriseSiteResult, error) {
	var rv LookupHcxEnterpriseSiteResult
	err := ctx.Invoke("azure-native:avs/v20210101preview:getHcxEnterpriseSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHcxEnterpriseSiteArgs struct {
	HcxEnterpriseSiteName string `pulumi:"hcxEnterpriseSiteName"`
	PrivateCloudName      string `pulumi:"privateCloudName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupHcxEnterpriseSiteResult struct {
	ActivationKey string `pulumi:"activationKey"`
	Id            string `pulumi:"id"`
	Name          string `pulumi:"name"`
	Status        string `pulumi:"status"`
	Type          string `pulumi:"type"`
}

func LookupHcxEnterpriseSiteOutput(ctx *pulumi.Context, args LookupHcxEnterpriseSiteOutputArgs, opts ...pulumi.InvokeOption) LookupHcxEnterpriseSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHcxEnterpriseSiteResult, error) {
			args := v.(LookupHcxEnterpriseSiteArgs)
			r, err := LookupHcxEnterpriseSite(ctx, &args, opts...)
			var s LookupHcxEnterpriseSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHcxEnterpriseSiteResultOutput)
}

type LookupHcxEnterpriseSiteOutputArgs struct {
	HcxEnterpriseSiteName pulumi.StringInput `pulumi:"hcxEnterpriseSiteName"`
	PrivateCloudName      pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHcxEnterpriseSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHcxEnterpriseSiteArgs)(nil)).Elem()
}


type LookupHcxEnterpriseSiteResultOutput struct{ *pulumi.OutputState }

func (LookupHcxEnterpriseSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHcxEnterpriseSiteResult)(nil)).Elem()
}

func (o LookupHcxEnterpriseSiteResultOutput) ToLookupHcxEnterpriseSiteResultOutput() LookupHcxEnterpriseSiteResultOutput {
	return o
}

func (o LookupHcxEnterpriseSiteResultOutput) ToLookupHcxEnterpriseSiteResultOutputWithContext(ctx context.Context) LookupHcxEnterpriseSiteResultOutput {
	return o
}

func (o LookupHcxEnterpriseSiteResultOutput) ActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHcxEnterpriseSiteResult) string { return v.ActivationKey }).(pulumi.StringOutput)
}

func (o LookupHcxEnterpriseSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHcxEnterpriseSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHcxEnterpriseSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHcxEnterpriseSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHcxEnterpriseSiteResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHcxEnterpriseSiteResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupHcxEnterpriseSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHcxEnterpriseSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHcxEnterpriseSiteResultOutput{})
}
