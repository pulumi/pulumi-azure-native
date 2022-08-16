


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityConnectorApplication(ctx *pulumi.Context, args *LookupSecurityConnectorApplicationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorApplicationResult, error) {
	var rv LookupSecurityConnectorApplicationResult
	err := ctx.Invoke("azure-native:security/v20220701preview:getSecurityConnectorApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorApplicationArgs struct {
	ApplicationId         string `pulumi:"applicationId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}


type LookupSecurityConnectorApplicationResult struct {
	Description        *string `pulumi:"description"`
	DisplayName        *string `pulumi:"displayName"`
	Id                 string  `pulumi:"id"`
	Name               string  `pulumi:"name"`
	SourceResourceType string  `pulumi:"sourceResourceType"`
	Type               string  `pulumi:"type"`
}

func LookupSecurityConnectorApplicationOutput(ctx *pulumi.Context, args LookupSecurityConnectorApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityConnectorApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityConnectorApplicationResult, error) {
			args := v.(LookupSecurityConnectorApplicationArgs)
			r, err := LookupSecurityConnectorApplication(ctx, &args, opts...)
			var s LookupSecurityConnectorApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityConnectorApplicationResultOutput)
}

type LookupSecurityConnectorApplicationOutputArgs struct {
	ApplicationId         pulumi.StringInput `pulumi:"applicationId"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (LookupSecurityConnectorApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorApplicationArgs)(nil)).Elem()
}


type LookupSecurityConnectorApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityConnectorApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorApplicationResult)(nil)).Elem()
}

func (o LookupSecurityConnectorApplicationResultOutput) ToLookupSecurityConnectorApplicationResultOutput() LookupSecurityConnectorApplicationResultOutput {
	return o
}

func (o LookupSecurityConnectorApplicationResultOutput) ToLookupSecurityConnectorApplicationResultOutputWithContext(ctx context.Context) LookupSecurityConnectorApplicationResultOutput {
	return o
}

func (o LookupSecurityConnectorApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorApplicationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorApplicationResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityConnectorApplicationResultOutput{})
}
