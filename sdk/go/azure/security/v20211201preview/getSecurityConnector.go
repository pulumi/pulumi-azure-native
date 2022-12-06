


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityConnector(ctx *pulumi.Context, args *LookupSecurityConnectorArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorResult, error) {
	var rv LookupSecurityConnectorResult
	err := ctx.Invoke("azure-native:security/v20211201preview:getSecurityConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorArgs struct {
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}


type LookupSecurityConnectorResult struct {
	EnvironmentData     interface{}        `pulumi:"environmentData"`
	EnvironmentName     *string            `pulumi:"environmentName"`
	Etag                *string            `pulumi:"etag"`
	HierarchyIdentifier *string            `pulumi:"hierarchyIdentifier"`
	Id                  string             `pulumi:"id"`
	Kind                *string            `pulumi:"kind"`
	Location            *string            `pulumi:"location"`
	Name                string             `pulumi:"name"`
	Offerings           []interface{}      `pulumi:"offerings"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Tags                map[string]string  `pulumi:"tags"`
	Type                string             `pulumi:"type"`
}

func LookupSecurityConnectorOutput(ctx *pulumi.Context, args LookupSecurityConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityConnectorResult, error) {
			args := v.(LookupSecurityConnectorArgs)
			r, err := LookupSecurityConnector(ctx, &args, opts...)
			var s LookupSecurityConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityConnectorResultOutput)
}

type LookupSecurityConnectorOutputArgs struct {
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (LookupSecurityConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorArgs)(nil)).Elem()
}


type LookupSecurityConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorResult)(nil)).Elem()
}

func (o LookupSecurityConnectorResultOutput) ToLookupSecurityConnectorResultOutput() LookupSecurityConnectorResultOutput {
	return o
}

func (o LookupSecurityConnectorResultOutput) ToLookupSecurityConnectorResultOutputWithContext(ctx context.Context) LookupSecurityConnectorResultOutput {
	return o
}

func (o LookupSecurityConnectorResultOutput) EnvironmentData() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) interface{} { return v.EnvironmentData }).(pulumi.AnyOutput)
}

func (o LookupSecurityConnectorResultOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorResultOutput) HierarchyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.HierarchyIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorResultOutput) Offerings() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) []interface{} { return v.Offerings }).(pulumi.ArrayOutput)
}

func (o LookupSecurityConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSecurityConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSecurityConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityConnectorResultOutput{})
}
