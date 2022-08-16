


package dashboard

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGrafana(ctx *pulumi.Context, args *LookupGrafanaArgs, opts ...pulumi.InvokeOption) (*LookupGrafanaResult, error) {
	var rv LookupGrafanaResult
	err := ctx.Invoke("azure-native:dashboard:getGrafana", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGrafanaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupGrafanaResult struct {
	Id         string                           `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse  `pulumi:"identity"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties ManagedGrafanaPropertiesResponse `pulumi:"properties"`
	Sku        *ResourceSkuResponse             `pulumi:"sku"`
	SystemData SystemDataResponse               `pulumi:"systemData"`
	Tags       map[string]string                `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}

func LookupGrafanaOutput(ctx *pulumi.Context, args LookupGrafanaOutputArgs, opts ...pulumi.InvokeOption) LookupGrafanaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGrafanaResult, error) {
			args := v.(LookupGrafanaArgs)
			r, err := LookupGrafana(ctx, &args, opts...)
			var s LookupGrafanaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGrafanaResultOutput)
}

type LookupGrafanaOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupGrafanaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrafanaArgs)(nil)).Elem()
}


type LookupGrafanaResultOutput struct{ *pulumi.OutputState }

func (LookupGrafanaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrafanaResult)(nil)).Elem()
}

func (o LookupGrafanaResultOutput) ToLookupGrafanaResultOutput() LookupGrafanaResultOutput {
	return o
}

func (o LookupGrafanaResultOutput) ToLookupGrafanaResultOutputWithContext(ctx context.Context) LookupGrafanaResultOutput {
	return o
}

func (o LookupGrafanaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrafanaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGrafanaResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupGrafanaResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupGrafanaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrafanaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGrafanaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrafanaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGrafanaResultOutput) Properties() ManagedGrafanaPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGrafanaResult) ManagedGrafanaPropertiesResponse { return v.Properties }).(ManagedGrafanaPropertiesResponseOutput)
}

func (o LookupGrafanaResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupGrafanaResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupGrafanaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGrafanaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGrafanaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGrafanaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGrafanaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrafanaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGrafanaResultOutput{})
}
