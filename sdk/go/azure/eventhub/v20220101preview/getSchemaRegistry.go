


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchemaRegistry(ctx *pulumi.Context, args *LookupSchemaRegistryArgs, opts ...pulumi.InvokeOption) (*LookupSchemaRegistryResult, error) {
	var rv LookupSchemaRegistryResult
	err := ctx.Invoke("azure-native:eventhub/v20220101preview:getSchemaRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaRegistryArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaGroupName   string `pulumi:"schemaGroupName"`
}


type LookupSchemaRegistryResult struct {
	CreatedAtUtc        string             `pulumi:"createdAtUtc"`
	ETag                string             `pulumi:"eTag"`
	GroupProperties     map[string]string  `pulumi:"groupProperties"`
	Id                  string             `pulumi:"id"`
	Location            string             `pulumi:"location"`
	Name                string             `pulumi:"name"`
	SchemaCompatibility *string            `pulumi:"schemaCompatibility"`
	SchemaType          *string            `pulumi:"schemaType"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
	UpdatedAtUtc        string             `pulumi:"updatedAtUtc"`
}

func LookupSchemaRegistryOutput(ctx *pulumi.Context, args LookupSchemaRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSchemaRegistryResult, error) {
			args := v.(LookupSchemaRegistryArgs)
			r, err := LookupSchemaRegistry(ctx, &args, opts...)
			var s LookupSchemaRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSchemaRegistryResultOutput)
}

type LookupSchemaRegistryOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SchemaGroupName   pulumi.StringInput `pulumi:"schemaGroupName"`
}

func (LookupSchemaRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaRegistryArgs)(nil)).Elem()
}


type LookupSchemaRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaRegistryResult)(nil)).Elem()
}

func (o LookupSchemaRegistryResultOutput) ToLookupSchemaRegistryResultOutput() LookupSchemaRegistryResultOutput {
	return o
}

func (o LookupSchemaRegistryResultOutput) ToLookupSchemaRegistryResultOutputWithContext(ctx context.Context) LookupSchemaRegistryResultOutput {
	return o
}

func (o LookupSchemaRegistryResultOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

func (o LookupSchemaRegistryResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupSchemaRegistryResultOutput) GroupProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) map[string]string { return v.GroupProperties }).(pulumi.StringMapOutput)
}

func (o LookupSchemaRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSchemaRegistryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSchemaRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSchemaRegistryResultOutput) SchemaCompatibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) *string { return v.SchemaCompatibility }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaRegistryResultOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) *string { return v.SchemaType }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSchemaRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSchemaRegistryResultOutput) UpdatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaRegistryResult) string { return v.UpdatedAtUtc }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaRegistryResultOutput{})
}
