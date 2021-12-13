


package v20150408

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Capability struct {
	Name *string `pulumi:"name"`
}





type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(context.Context) CapabilityOutput
}

type CapabilityArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (i CapabilityArgs) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i CapabilityArgs) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}





type CapabilityArrayInput interface {
	pulumi.Input

	ToCapabilityArrayOutput() CapabilityArrayOutput
	ToCapabilityArrayOutputWithContext(context.Context) CapabilityArrayOutput
}

type CapabilityArray []CapabilityInput

func (CapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (i CapabilityArray) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return i.ToCapabilityArrayOutputWithContext(context.Background())
}

func (i CapabilityArray) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityArrayOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CapabilityArrayOutput struct{ *pulumi.OutputState }

func (CapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) Index(i pulumi.IntInput) CapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Capability {
		return vs[0].([]Capability)[vs[1].(int)]
	}).(CapabilityOutput)
}

type CapabilityResponse struct {
	Name *string `pulumi:"name"`
}

type CapabilityResponseOutput struct{ *pulumi.OutputState }

func (CapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutput() CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutputWithContext(ctx context.Context) CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (CapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutputWithContext(ctx context.Context) CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) Index(i pulumi.IntInput) CapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapabilityResponse {
		return vs[0].([]CapabilityResponse)[vs[1].(int)]
	}).(CapabilityResponseOutput)
}

type CassandraKeyspaceResource struct {
	Id string `pulumi:"id"`
}





type CassandraKeyspaceResourceInput interface {
	pulumi.Input

	ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput
	ToCassandraKeyspaceResourceOutputWithContext(context.Context) CassandraKeyspaceResourceOutput
}

type CassandraKeyspaceResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (CassandraKeyspaceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceResource)(nil)).Elem()
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput {
	return i.ToCassandraKeyspaceResourceOutputWithContext(context.Background())
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourceOutputWithContext(ctx context.Context) CassandraKeyspaceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceResourceOutput)
}

type CassandraKeyspaceResourceOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceResource)(nil)).Elem()
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput {
	return o
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourceOutputWithContext(ctx context.Context) CassandraKeyspaceResourceOutput {
	return o
}

func (o CassandraKeyspaceResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceResource) string { return v.Id }).(pulumi.StringOutput)
}

type CassandraPartitionKey struct {
	Name *string `pulumi:"name"`
}





type CassandraPartitionKeyInput interface {
	pulumi.Input

	ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput
	ToCassandraPartitionKeyOutputWithContext(context.Context) CassandraPartitionKeyOutput
}

type CassandraPartitionKeyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CassandraPartitionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKey)(nil)).Elem()
}

func (i CassandraPartitionKeyArgs) ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput {
	return i.ToCassandraPartitionKeyOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyArgs) ToCassandraPartitionKeyOutputWithContext(ctx context.Context) CassandraPartitionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyOutput)
}





type CassandraPartitionKeyArrayInput interface {
	pulumi.Input

	ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput
	ToCassandraPartitionKeyArrayOutputWithContext(context.Context) CassandraPartitionKeyArrayOutput
}

type CassandraPartitionKeyArray []CassandraPartitionKeyInput

func (CassandraPartitionKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKey)(nil)).Elem()
}

func (i CassandraPartitionKeyArray) ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput {
	return i.ToCassandraPartitionKeyArrayOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyArray) ToCassandraPartitionKeyArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyArrayOutput)
}

type CassandraPartitionKeyOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKey)(nil)).Elem()
}

func (o CassandraPartitionKeyOutput) ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput {
	return o
}

func (o CassandraPartitionKeyOutput) ToCassandraPartitionKeyOutputWithContext(ctx context.Context) CassandraPartitionKeyOutput {
	return o
}

func (o CassandraPartitionKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CassandraPartitionKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CassandraPartitionKeyArrayOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKey)(nil)).Elem()
}

func (o CassandraPartitionKeyArrayOutput) ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput {
	return o
}

func (o CassandraPartitionKeyArrayOutput) ToCassandraPartitionKeyArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyArrayOutput {
	return o
}

func (o CassandraPartitionKeyArrayOutput) Index(i pulumi.IntInput) CassandraPartitionKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraPartitionKey {
		return vs[0].([]CassandraPartitionKey)[vs[1].(int)]
	}).(CassandraPartitionKeyOutput)
}

type CassandraPartitionKeyResponse struct {
	Name *string `pulumi:"name"`
}

type CassandraPartitionKeyResponseOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKeyResponse)(nil)).Elem()
}

func (o CassandraPartitionKeyResponseOutput) ToCassandraPartitionKeyResponseOutput() CassandraPartitionKeyResponseOutput {
	return o
}

func (o CassandraPartitionKeyResponseOutput) ToCassandraPartitionKeyResponseOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseOutput {
	return o
}

func (o CassandraPartitionKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CassandraPartitionKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CassandraPartitionKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKeyResponse)(nil)).Elem()
}

func (o CassandraPartitionKeyResponseArrayOutput) ToCassandraPartitionKeyResponseArrayOutput() CassandraPartitionKeyResponseArrayOutput {
	return o
}

func (o CassandraPartitionKeyResponseArrayOutput) ToCassandraPartitionKeyResponseArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseArrayOutput {
	return o
}

func (o CassandraPartitionKeyResponseArrayOutput) Index(i pulumi.IntInput) CassandraPartitionKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraPartitionKeyResponse {
		return vs[0].([]CassandraPartitionKeyResponse)[vs[1].(int)]
	}).(CassandraPartitionKeyResponseOutput)
}

type CassandraSchema struct {
	ClusterKeys   []ClusterKey            `pulumi:"clusterKeys"`
	Columns       []Column                `pulumi:"columns"`
	PartitionKeys []CassandraPartitionKey `pulumi:"partitionKeys"`
}





type CassandraSchemaInput interface {
	pulumi.Input

	ToCassandraSchemaOutput() CassandraSchemaOutput
	ToCassandraSchemaOutputWithContext(context.Context) CassandraSchemaOutput
}

type CassandraSchemaArgs struct {
	ClusterKeys   ClusterKeyArrayInput            `pulumi:"clusterKeys"`
	Columns       ColumnArrayInput                `pulumi:"columns"`
	PartitionKeys CassandraPartitionKeyArrayInput `pulumi:"partitionKeys"`
}

func (CassandraSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchema)(nil)).Elem()
}

func (i CassandraSchemaArgs) ToCassandraSchemaOutput() CassandraSchemaOutput {
	return i.ToCassandraSchemaOutputWithContext(context.Background())
}

func (i CassandraSchemaArgs) ToCassandraSchemaOutputWithContext(ctx context.Context) CassandraSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaOutput)
}

func (i CassandraSchemaArgs) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return i.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (i CassandraSchemaArgs) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaOutput).ToCassandraSchemaPtrOutputWithContext(ctx)
}









type CassandraSchemaPtrInput interface {
	pulumi.Input

	ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput
	ToCassandraSchemaPtrOutputWithContext(context.Context) CassandraSchemaPtrOutput
}

type cassandraSchemaPtrType CassandraSchemaArgs

func CassandraSchemaPtr(v *CassandraSchemaArgs) CassandraSchemaPtrInput {
	return (*cassandraSchemaPtrType)(v)
}

func (*cassandraSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchema)(nil)).Elem()
}

func (i *cassandraSchemaPtrType) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return i.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (i *cassandraSchemaPtrType) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaPtrOutput)
}

type CassandraSchemaOutput struct{ *pulumi.OutputState }

func (CassandraSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchema)(nil)).Elem()
}

func (o CassandraSchemaOutput) ToCassandraSchemaOutput() CassandraSchemaOutput {
	return o
}

func (o CassandraSchemaOutput) ToCassandraSchemaOutputWithContext(ctx context.Context) CassandraSchemaOutput {
	return o
}

func (o CassandraSchemaOutput) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return o.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (o CassandraSchemaOutput) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraSchema) *CassandraSchema {
		return &v
	}).(CassandraSchemaPtrOutput)
}

func (o CassandraSchemaOutput) ClusterKeys() ClusterKeyArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []ClusterKey { return v.ClusterKeys }).(ClusterKeyArrayOutput)
}

func (o CassandraSchemaOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []Column { return v.Columns }).(ColumnArrayOutput)
}

func (o CassandraSchemaOutput) PartitionKeys() CassandraPartitionKeyArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []CassandraPartitionKey { return v.PartitionKeys }).(CassandraPartitionKeyArrayOutput)
}

type CassandraSchemaPtrOutput struct{ *pulumi.OutputState }

func (CassandraSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchema)(nil)).Elem()
}

func (o CassandraSchemaPtrOutput) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return o
}

func (o CassandraSchemaPtrOutput) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return o
}

func (o CassandraSchemaPtrOutput) Elem() CassandraSchemaOutput {
	return o.ApplyT(func(v *CassandraSchema) CassandraSchema {
		if v != nil {
			return *v
		}
		var ret CassandraSchema
		return ret
	}).(CassandraSchemaOutput)
}

func (o CassandraSchemaPtrOutput) ClusterKeys() ClusterKeyArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []ClusterKey {
		if v == nil {
			return nil
		}
		return v.ClusterKeys
	}).(ClusterKeyArrayOutput)
}

func (o CassandraSchemaPtrOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []Column {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnArrayOutput)
}

func (o CassandraSchemaPtrOutput) PartitionKeys() CassandraPartitionKeyArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []CassandraPartitionKey {
		if v == nil {
			return nil
		}
		return v.PartitionKeys
	}).(CassandraPartitionKeyArrayOutput)
}

type CassandraSchemaResponse struct {
	ClusterKeys   []ClusterKeyResponse            `pulumi:"clusterKeys"`
	Columns       []ColumnResponse                `pulumi:"columns"`
	PartitionKeys []CassandraPartitionKeyResponse `pulumi:"partitionKeys"`
}

type CassandraSchemaResponseOutput struct{ *pulumi.OutputState }

func (CassandraSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchemaResponse)(nil)).Elem()
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponseOutput() CassandraSchemaResponseOutput {
	return o
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponseOutputWithContext(ctx context.Context) CassandraSchemaResponseOutput {
	return o
}

func (o CassandraSchemaResponseOutput) ClusterKeys() ClusterKeyResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []ClusterKeyResponse { return v.ClusterKeys }).(ClusterKeyResponseArrayOutput)
}

func (o CassandraSchemaResponseOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []ColumnResponse { return v.Columns }).(ColumnResponseArrayOutput)
}

func (o CassandraSchemaResponseOutput) PartitionKeys() CassandraPartitionKeyResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []CassandraPartitionKeyResponse { return v.PartitionKeys }).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraSchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (CassandraSchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchemaResponse)(nil)).Elem()
}

func (o CassandraSchemaResponsePtrOutput) ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput {
	return o
}

func (o CassandraSchemaResponsePtrOutput) ToCassandraSchemaResponsePtrOutputWithContext(ctx context.Context) CassandraSchemaResponsePtrOutput {
	return o
}

func (o CassandraSchemaResponsePtrOutput) Elem() CassandraSchemaResponseOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) CassandraSchemaResponse {
		if v != nil {
			return *v
		}
		var ret CassandraSchemaResponse
		return ret
	}).(CassandraSchemaResponseOutput)
}

func (o CassandraSchemaResponsePtrOutput) ClusterKeys() ClusterKeyResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []ClusterKeyResponse {
		if v == nil {
			return nil
		}
		return v.ClusterKeys
	}).(ClusterKeyResponseArrayOutput)
}

func (o CassandraSchemaResponsePtrOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []ColumnResponse {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnResponseArrayOutput)
}

func (o CassandraSchemaResponsePtrOutput) PartitionKeys() CassandraPartitionKeyResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []CassandraPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKeys
	}).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraTableResource struct {
	DefaultTtl *int             `pulumi:"defaultTtl"`
	Id         string           `pulumi:"id"`
	Schema     *CassandraSchema `pulumi:"schema"`
}





type CassandraTableResourceInput interface {
	pulumi.Input

	ToCassandraTableResourceOutput() CassandraTableResourceOutput
	ToCassandraTableResourceOutputWithContext(context.Context) CassandraTableResourceOutput
}

type CassandraTableResourceArgs struct {
	DefaultTtl pulumi.IntPtrInput      `pulumi:"defaultTtl"`
	Id         pulumi.StringInput      `pulumi:"id"`
	Schema     CassandraSchemaPtrInput `pulumi:"schema"`
}

func (CassandraTableResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableResource)(nil)).Elem()
}

func (i CassandraTableResourceArgs) ToCassandraTableResourceOutput() CassandraTableResourceOutput {
	return i.ToCassandraTableResourceOutputWithContext(context.Background())
}

func (i CassandraTableResourceArgs) ToCassandraTableResourceOutputWithContext(ctx context.Context) CassandraTableResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableResourceOutput)
}

type CassandraTableResourceOutput struct{ *pulumi.OutputState }

func (CassandraTableResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableResource)(nil)).Elem()
}

func (o CassandraTableResourceOutput) ToCassandraTableResourceOutput() CassandraTableResourceOutput {
	return o
}

func (o CassandraTableResourceOutput) ToCassandraTableResourceOutputWithContext(ctx context.Context) CassandraTableResourceOutput {
	return o
}

func (o CassandraTableResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraTableResourceOutput) Schema() CassandraSchemaPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *CassandraSchema { return v.Schema }).(CassandraSchemaPtrOutput)
}

type ClusterKey struct {
	Name    *string `pulumi:"name"`
	OrderBy *string `pulumi:"orderBy"`
}





type ClusterKeyInput interface {
	pulumi.Input

	ToClusterKeyOutput() ClusterKeyOutput
	ToClusterKeyOutputWithContext(context.Context) ClusterKeyOutput
}

type ClusterKeyArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
}

func (ClusterKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKey)(nil)).Elem()
}

func (i ClusterKeyArgs) ToClusterKeyOutput() ClusterKeyOutput {
	return i.ToClusterKeyOutputWithContext(context.Background())
}

func (i ClusterKeyArgs) ToClusterKeyOutputWithContext(ctx context.Context) ClusterKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyOutput)
}





type ClusterKeyArrayInput interface {
	pulumi.Input

	ToClusterKeyArrayOutput() ClusterKeyArrayOutput
	ToClusterKeyArrayOutputWithContext(context.Context) ClusterKeyArrayOutput
}

type ClusterKeyArray []ClusterKeyInput

func (ClusterKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKey)(nil)).Elem()
}

func (i ClusterKeyArray) ToClusterKeyArrayOutput() ClusterKeyArrayOutput {
	return i.ToClusterKeyArrayOutputWithContext(context.Background())
}

func (i ClusterKeyArray) ToClusterKeyArrayOutputWithContext(ctx context.Context) ClusterKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyArrayOutput)
}

type ClusterKeyOutput struct{ *pulumi.OutputState }

func (ClusterKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKey)(nil)).Elem()
}

func (o ClusterKeyOutput) ToClusterKeyOutput() ClusterKeyOutput {
	return o
}

func (o ClusterKeyOutput) ToClusterKeyOutputWithContext(ctx context.Context) ClusterKeyOutput {
	return o
}

func (o ClusterKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClusterKeyOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKey) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

type ClusterKeyArrayOutput struct{ *pulumi.OutputState }

func (ClusterKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKey)(nil)).Elem()
}

func (o ClusterKeyArrayOutput) ToClusterKeyArrayOutput() ClusterKeyArrayOutput {
	return o
}

func (o ClusterKeyArrayOutput) ToClusterKeyArrayOutputWithContext(ctx context.Context) ClusterKeyArrayOutput {
	return o
}

func (o ClusterKeyArrayOutput) Index(i pulumi.IntInput) ClusterKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterKey {
		return vs[0].([]ClusterKey)[vs[1].(int)]
	}).(ClusterKeyOutput)
}

type ClusterKeyResponse struct {
	Name    *string `pulumi:"name"`
	OrderBy *string `pulumi:"orderBy"`
}

type ClusterKeyResponseOutput struct{ *pulumi.OutputState }

func (ClusterKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKeyResponse)(nil)).Elem()
}

func (o ClusterKeyResponseOutput) ToClusterKeyResponseOutput() ClusterKeyResponseOutput {
	return o
}

func (o ClusterKeyResponseOutput) ToClusterKeyResponseOutputWithContext(ctx context.Context) ClusterKeyResponseOutput {
	return o
}

func (o ClusterKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClusterKeyResponseOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKeyResponse) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

type ClusterKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKeyResponse)(nil)).Elem()
}

func (o ClusterKeyResponseArrayOutput) ToClusterKeyResponseArrayOutput() ClusterKeyResponseArrayOutput {
	return o
}

func (o ClusterKeyResponseArrayOutput) ToClusterKeyResponseArrayOutputWithContext(ctx context.Context) ClusterKeyResponseArrayOutput {
	return o
}

func (o ClusterKeyResponseArrayOutput) Index(i pulumi.IntInput) ClusterKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterKeyResponse {
		return vs[0].([]ClusterKeyResponse)[vs[1].(int)]
	}).(ClusterKeyResponseOutput)
}

type Column struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ColumnInput interface {
	pulumi.Input

	ToColumnOutput() ColumnOutput
	ToColumnOutputWithContext(context.Context) ColumnOutput
}

type ColumnArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (i ColumnArgs) ToColumnOutput() ColumnOutput {
	return i.ToColumnOutputWithContext(context.Background())
}

func (i ColumnArgs) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnOutput)
}





type ColumnArrayInput interface {
	pulumi.Input

	ToColumnArrayOutput() ColumnArrayOutput
	ToColumnArrayOutputWithContext(context.Context) ColumnArrayOutput
}

type ColumnArray []ColumnInput

func (ColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (i ColumnArray) ToColumnArrayOutput() ColumnArrayOutput {
	return i.ToColumnArrayOutputWithContext(context.Background())
}

func (i ColumnArray) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnArrayOutput)
}

type ColumnOutput struct{ *pulumi.OutputState }

func (ColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (o ColumnOutput) ToColumnOutput() ColumnOutput {
	return o
}

func (o ColumnOutput) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return o
}

func (o ColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnArrayOutput struct{ *pulumi.OutputState }

func (ColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (o ColumnArrayOutput) ToColumnArrayOutput() ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) Index(i pulumi.IntInput) ColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Column {
		return vs[0].([]Column)[vs[1].(int)]
	}).(ColumnOutput)
}

type ColumnResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ColumnResponseOutput struct{ *pulumi.OutputState }

func (ColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseOutput) ToColumnResponseOutput() ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) ToColumnResponseOutputWithContext(ctx context.Context) ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (ColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutput() ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutputWithContext(ctx context.Context) ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) Index(i pulumi.IntInput) ColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ColumnResponse {
		return vs[0].([]ColumnResponse)[vs[1].(int)]
	}).(ColumnResponseOutput)
}

type ConflictResolutionPolicy struct {
	ConflictResolutionPath      *string `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure *string `pulumi:"conflictResolutionProcedure"`
	Mode                        *string `pulumi:"mode"`
}


func (val *ConflictResolutionPolicy) Defaults() *ConflictResolutionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "LastWriterWins"
		tmp.Mode = &mode_
	}
	return &tmp
}





type ConflictResolutionPolicyInput interface {
	pulumi.Input

	ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput
	ToConflictResolutionPolicyOutputWithContext(context.Context) ConflictResolutionPolicyOutput
}

type ConflictResolutionPolicyArgs struct {
	ConflictResolutionPath      pulumi.StringPtrInput `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure pulumi.StringPtrInput `pulumi:"conflictResolutionProcedure"`
	Mode                        pulumi.StringPtrInput `pulumi:"mode"`
}

func (ConflictResolutionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicy)(nil)).Elem()
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput {
	return i.ToConflictResolutionPolicyOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyOutputWithContext(ctx context.Context) ConflictResolutionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyOutput)
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return i.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyOutput).ToConflictResolutionPolicyPtrOutputWithContext(ctx)
}









type ConflictResolutionPolicyPtrInput interface {
	pulumi.Input

	ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput
	ToConflictResolutionPolicyPtrOutputWithContext(context.Context) ConflictResolutionPolicyPtrOutput
}

type conflictResolutionPolicyPtrType ConflictResolutionPolicyArgs

func ConflictResolutionPolicyPtr(v *ConflictResolutionPolicyArgs) ConflictResolutionPolicyPtrInput {
	return (*conflictResolutionPolicyPtrType)(v)
}

func (*conflictResolutionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicy)(nil)).Elem()
}

func (i *conflictResolutionPolicyPtrType) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return i.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (i *conflictResolutionPolicyPtrType) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyPtrOutput)
}

type ConflictResolutionPolicyOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicy)(nil)).Elem()
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput {
	return o
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyOutputWithContext(ctx context.Context) ConflictResolutionPolicyOutput {
	return o
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return o.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConflictResolutionPolicy) *ConflictResolutionPolicy {
		return &v
	}).(ConflictResolutionPolicyPtrOutput)
}

func (o ConflictResolutionPolicyOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.ConflictResolutionPath }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.ConflictResolutionProcedure }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyPtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicy)(nil)).Elem()
}

func (o ConflictResolutionPolicyPtrOutput) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return o
}

func (o ConflictResolutionPolicyPtrOutput) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return o
}

func (o ConflictResolutionPolicyPtrOutput) Elem() ConflictResolutionPolicyOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) ConflictResolutionPolicy {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionPolicy
		return ret
	}).(ConflictResolutionPolicyOutput)
}

func (o ConflictResolutionPolicyPtrOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPath
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyPtrOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionProcedure
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyResponse struct {
	ConflictResolutionPath      *string `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure *string `pulumi:"conflictResolutionProcedure"`
	Mode                        *string `pulumi:"mode"`
}


func (val *ConflictResolutionPolicyResponse) Defaults() *ConflictResolutionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "LastWriterWins"
		tmp.Mode = &mode_
	}
	return &tmp
}

type ConflictResolutionPolicyResponseOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponseOutput() ConflictResolutionPolicyResponseOutput {
	return o
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponseOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponseOutput {
	return o
}

func (o ConflictResolutionPolicyResponseOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.ConflictResolutionPath }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.ConflictResolutionProcedure }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (o ConflictResolutionPolicyResponsePtrOutput) ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput {
	return o
}

func (o ConflictResolutionPolicyResponsePtrOutput) ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponsePtrOutput {
	return o
}

func (o ConflictResolutionPolicyResponsePtrOutput) Elem() ConflictResolutionPolicyResponseOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) ConflictResolutionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionPolicyResponse
		return ret
	}).(ConflictResolutionPolicyResponseOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPath
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionProcedure
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ConsistencyPolicy struct {
	DefaultConsistencyLevel DefaultConsistencyLevel `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    *int                    `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      *float64                `pulumi:"maxStalenessPrefix"`
}





type ConsistencyPolicyInput interface {
	pulumi.Input

	ToConsistencyPolicyOutput() ConsistencyPolicyOutput
	ToConsistencyPolicyOutputWithContext(context.Context) ConsistencyPolicyOutput
}

type ConsistencyPolicyArgs struct {
	DefaultConsistencyLevel DefaultConsistencyLevelInput `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    pulumi.IntPtrInput           `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      pulumi.Float64PtrInput       `pulumi:"maxStalenessPrefix"`
}

func (ConsistencyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicy)(nil)).Elem()
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyOutput() ConsistencyPolicyOutput {
	return i.ToConsistencyPolicyOutputWithContext(context.Background())
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyOutputWithContext(ctx context.Context) ConsistencyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyOutput)
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return i.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyOutput).ToConsistencyPolicyPtrOutputWithContext(ctx)
}









type ConsistencyPolicyPtrInput interface {
	pulumi.Input

	ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput
	ToConsistencyPolicyPtrOutputWithContext(context.Context) ConsistencyPolicyPtrOutput
}

type consistencyPolicyPtrType ConsistencyPolicyArgs

func ConsistencyPolicyPtr(v *ConsistencyPolicyArgs) ConsistencyPolicyPtrInput {
	return (*consistencyPolicyPtrType)(v)
}

func (*consistencyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicy)(nil)).Elem()
}

func (i *consistencyPolicyPtrType) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return i.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (i *consistencyPolicyPtrType) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyPtrOutput)
}

type ConsistencyPolicyOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicy)(nil)).Elem()
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyOutput() ConsistencyPolicyOutput {
	return o
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyOutputWithContext(ctx context.Context) ConsistencyPolicyOutput {
	return o
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return o.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsistencyPolicy) *ConsistencyPolicy {
		return &v
	}).(ConsistencyPolicyPtrOutput)
}

func (o ConsistencyPolicyOutput) DefaultConsistencyLevel() DefaultConsistencyLevelOutput {
	return o.ApplyT(func(v ConsistencyPolicy) DefaultConsistencyLevel { return v.DefaultConsistencyLevel }).(DefaultConsistencyLevelOutput)
}

func (o ConsistencyPolicyOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConsistencyPolicy) *int { return v.MaxIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConsistencyPolicy) *float64 { return v.MaxStalenessPrefix }).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyPtrOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicy)(nil)).Elem()
}

func (o ConsistencyPolicyPtrOutput) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return o
}

func (o ConsistencyPolicyPtrOutput) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return o
}

func (o ConsistencyPolicyPtrOutput) Elem() ConsistencyPolicyOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) ConsistencyPolicy {
		if v != nil {
			return *v
		}
		var ret ConsistencyPolicy
		return ret
	}).(ConsistencyPolicyOutput)
}

func (o ConsistencyPolicyPtrOutput) DefaultConsistencyLevel() DefaultConsistencyLevelPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *DefaultConsistencyLevel {
		if v == nil {
			return nil
		}
		return &v.DefaultConsistencyLevel
	}).(DefaultConsistencyLevelPtrOutput)
}

func (o ConsistencyPolicyPtrOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyPtrOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxStalenessPrefix
	}).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyResponse struct {
	DefaultConsistencyLevel string   `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    *int     `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      *float64 `pulumi:"maxStalenessPrefix"`
}

type ConsistencyPolicyResponseOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicyResponse)(nil)).Elem()
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponseOutput() ConsistencyPolicyResponseOutput {
	return o
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponseOutputWithContext(ctx context.Context) ConsistencyPolicyResponseOutput {
	return o
}

func (o ConsistencyPolicyResponseOutput) DefaultConsistencyLevel() pulumi.StringOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) string { return v.DefaultConsistencyLevel }).(pulumi.StringOutput)
}

func (o ConsistencyPolicyResponseOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) *int { return v.MaxIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyResponseOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) *float64 { return v.MaxStalenessPrefix }).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicyResponse)(nil)).Elem()
}

func (o ConsistencyPolicyResponsePtrOutput) ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput {
	return o
}

func (o ConsistencyPolicyResponsePtrOutput) ToConsistencyPolicyResponsePtrOutputWithContext(ctx context.Context) ConsistencyPolicyResponsePtrOutput {
	return o
}

func (o ConsistencyPolicyResponsePtrOutput) Elem() ConsistencyPolicyResponseOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) ConsistencyPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ConsistencyPolicyResponse
		return ret
	}).(ConsistencyPolicyResponseOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) DefaultConsistencyLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultConsistencyLevel
	}).(pulumi.StringPtrOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxStalenessPrefix
	}).(pulumi.Float64PtrOutput)
}

type ContainerPartitionKey struct {
	Kind  *string  `pulumi:"kind"`
	Paths []string `pulumi:"paths"`
}


func (val *ContainerPartitionKey) Defaults() *ContainerPartitionKey {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type ContainerPartitionKeyInput interface {
	pulumi.Input

	ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput
	ToContainerPartitionKeyOutputWithContext(context.Context) ContainerPartitionKeyOutput
}

type ContainerPartitionKeyArgs struct {
	Kind  pulumi.StringPtrInput   `pulumi:"kind"`
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (ContainerPartitionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKey)(nil)).Elem()
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput {
	return i.ToContainerPartitionKeyOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyOutputWithContext(ctx context.Context) ContainerPartitionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyOutput)
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return i.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyOutput).ToContainerPartitionKeyPtrOutputWithContext(ctx)
}









type ContainerPartitionKeyPtrInput interface {
	pulumi.Input

	ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput
	ToContainerPartitionKeyPtrOutputWithContext(context.Context) ContainerPartitionKeyPtrOutput
}

type containerPartitionKeyPtrType ContainerPartitionKeyArgs

func ContainerPartitionKeyPtr(v *ContainerPartitionKeyArgs) ContainerPartitionKeyPtrInput {
	return (*containerPartitionKeyPtrType)(v)
}

func (*containerPartitionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKey)(nil)).Elem()
}

func (i *containerPartitionKeyPtrType) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return i.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (i *containerPartitionKeyPtrType) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyPtrOutput)
}

type ContainerPartitionKeyOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKey)(nil)).Elem()
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput {
	return o
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyOutputWithContext(ctx context.Context) ContainerPartitionKeyOutput {
	return o
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return o.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerPartitionKey) *ContainerPartitionKey {
		return &v
	}).(ContainerPartitionKeyPtrOutput)
}

func (o ContainerPartitionKeyOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKey) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerPartitionKey) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type ContainerPartitionKeyPtrOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKey)(nil)).Elem()
}

func (o ContainerPartitionKeyPtrOutput) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return o
}

func (o ContainerPartitionKeyPtrOutput) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return o
}

func (o ContainerPartitionKeyPtrOutput) Elem() ContainerPartitionKeyOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) ContainerPartitionKey {
		if v != nil {
			return *v
		}
		var ret ContainerPartitionKey
		return ret
	}).(ContainerPartitionKeyOutput)
}

func (o ContainerPartitionKeyPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyPtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

type ContainerPartitionKeyResponse struct {
	Kind  *string  `pulumi:"kind"`
	Paths []string `pulumi:"paths"`
}


func (val *ContainerPartitionKeyResponse) Defaults() *ContainerPartitionKeyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}

type ContainerPartitionKeyResponseOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKeyResponse)(nil)).Elem()
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponseOutput() ContainerPartitionKeyResponseOutput {
	return o
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponseOutputWithContext(ctx context.Context) ContainerPartitionKeyResponseOutput {
	return o
}

func (o ContainerPartitionKeyResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type ContainerPartitionKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKeyResponse)(nil)).Elem()
}

func (o ContainerPartitionKeyResponsePtrOutput) ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput {
	return o
}

func (o ContainerPartitionKeyResponsePtrOutput) ToContainerPartitionKeyResponsePtrOutputWithContext(ctx context.Context) ContainerPartitionKeyResponsePtrOutput {
	return o
}

func (o ContainerPartitionKeyResponsePtrOutput) Elem() ContainerPartitionKeyResponseOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) ContainerPartitionKeyResponse {
		if v != nil {
			return *v
		}
		var ret ContainerPartitionKeyResponse
		return ret
	}).(ContainerPartitionKeyResponseOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

type DatabaseAccountConnectionStringResponse struct {
	ConnectionString string `pulumi:"connectionString"`
	Description      string `pulumi:"description"`
}

type ExcludedPath struct {
	Path *string `pulumi:"path"`
}





type ExcludedPathInput interface {
	pulumi.Input

	ToExcludedPathOutput() ExcludedPathOutput
	ToExcludedPathOutputWithContext(context.Context) ExcludedPathOutput
}

type ExcludedPathArgs struct {
	Path pulumi.StringPtrInput `pulumi:"path"`
}

func (ExcludedPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPath)(nil)).Elem()
}

func (i ExcludedPathArgs) ToExcludedPathOutput() ExcludedPathOutput {
	return i.ToExcludedPathOutputWithContext(context.Background())
}

func (i ExcludedPathArgs) ToExcludedPathOutputWithContext(ctx context.Context) ExcludedPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathOutput)
}





type ExcludedPathArrayInput interface {
	pulumi.Input

	ToExcludedPathArrayOutput() ExcludedPathArrayOutput
	ToExcludedPathArrayOutputWithContext(context.Context) ExcludedPathArrayOutput
}

type ExcludedPathArray []ExcludedPathInput

func (ExcludedPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPath)(nil)).Elem()
}

func (i ExcludedPathArray) ToExcludedPathArrayOutput() ExcludedPathArrayOutput {
	return i.ToExcludedPathArrayOutputWithContext(context.Background())
}

func (i ExcludedPathArray) ToExcludedPathArrayOutputWithContext(ctx context.Context) ExcludedPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathArrayOutput)
}

type ExcludedPathOutput struct{ *pulumi.OutputState }

func (ExcludedPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPath)(nil)).Elem()
}

func (o ExcludedPathOutput) ToExcludedPathOutput() ExcludedPathOutput {
	return o
}

func (o ExcludedPathOutput) ToExcludedPathOutputWithContext(ctx context.Context) ExcludedPathOutput {
	return o
}

func (o ExcludedPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ExcludedPathArrayOutput struct{ *pulumi.OutputState }

func (ExcludedPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPath)(nil)).Elem()
}

func (o ExcludedPathArrayOutput) ToExcludedPathArrayOutput() ExcludedPathArrayOutput {
	return o
}

func (o ExcludedPathArrayOutput) ToExcludedPathArrayOutputWithContext(ctx context.Context) ExcludedPathArrayOutput {
	return o
}

func (o ExcludedPathArrayOutput) Index(i pulumi.IntInput) ExcludedPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExcludedPath {
		return vs[0].([]ExcludedPath)[vs[1].(int)]
	}).(ExcludedPathOutput)
}

type ExcludedPathResponse struct {
	Path *string `pulumi:"path"`
}

type ExcludedPathResponseOutput struct{ *pulumi.OutputState }

func (ExcludedPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPathResponse)(nil)).Elem()
}

func (o ExcludedPathResponseOutput) ToExcludedPathResponseOutput() ExcludedPathResponseOutput {
	return o
}

func (o ExcludedPathResponseOutput) ToExcludedPathResponseOutputWithContext(ctx context.Context) ExcludedPathResponseOutput {
	return o
}

func (o ExcludedPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ExcludedPathResponseArrayOutput struct{ *pulumi.OutputState }

func (ExcludedPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPathResponse)(nil)).Elem()
}

func (o ExcludedPathResponseArrayOutput) ToExcludedPathResponseArrayOutput() ExcludedPathResponseArrayOutput {
	return o
}

func (o ExcludedPathResponseArrayOutput) ToExcludedPathResponseArrayOutputWithContext(ctx context.Context) ExcludedPathResponseArrayOutput {
	return o
}

func (o ExcludedPathResponseArrayOutput) Index(i pulumi.IntInput) ExcludedPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExcludedPathResponse {
		return vs[0].([]ExcludedPathResponse)[vs[1].(int)]
	}).(ExcludedPathResponseOutput)
}

type FailoverPolicyResponse struct {
	FailoverPriority *int    `pulumi:"failoverPriority"`
	Id               string  `pulumi:"id"`
	LocationName     *string `pulumi:"locationName"`
}

type FailoverPolicyResponseOutput struct{ *pulumi.OutputState }

func (FailoverPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverPolicyResponse)(nil)).Elem()
}

func (o FailoverPolicyResponseOutput) ToFailoverPolicyResponseOutput() FailoverPolicyResponseOutput {
	return o
}

func (o FailoverPolicyResponseOutput) ToFailoverPolicyResponseOutputWithContext(ctx context.Context) FailoverPolicyResponseOutput {
	return o
}

func (o FailoverPolicyResponseOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o FailoverPolicyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FailoverPolicyResponseOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

type FailoverPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (FailoverPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverPolicyResponse)(nil)).Elem()
}

func (o FailoverPolicyResponseArrayOutput) ToFailoverPolicyResponseArrayOutput() FailoverPolicyResponseArrayOutput {
	return o
}

func (o FailoverPolicyResponseArrayOutput) ToFailoverPolicyResponseArrayOutputWithContext(ctx context.Context) FailoverPolicyResponseArrayOutput {
	return o
}

func (o FailoverPolicyResponseArrayOutput) Index(i pulumi.IntInput) FailoverPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverPolicyResponse {
		return vs[0].([]FailoverPolicyResponse)[vs[1].(int)]
	}).(FailoverPolicyResponseOutput)
}

type GremlinDatabaseResource struct {
	Id string `pulumi:"id"`
}





type GremlinDatabaseResourceInput interface {
	pulumi.Input

	ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput
	ToGremlinDatabaseResourceOutputWithContext(context.Context) GremlinDatabaseResourceOutput
}

type GremlinDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (GremlinDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseResource)(nil)).Elem()
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput {
	return i.ToGremlinDatabaseResourceOutputWithContext(context.Background())
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourceOutputWithContext(ctx context.Context) GremlinDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseResourceOutput)
}

type GremlinDatabaseResourceOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseResource)(nil)).Elem()
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput {
	return o
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourceOutputWithContext(ctx context.Context) GremlinDatabaseResourceOutput {
	return o
}

func (o GremlinDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type GremlinGraphResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                      `pulumi:"defaultTtl"`
	Id                       string                    `pulumi:"id"`
	IndexingPolicy           *IndexingPolicy           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKey    `pulumi:"partitionKey"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `pulumi:"uniqueKeyPolicy"`
}


func (val *GremlinGraphResource) Defaults() *GremlinGraphResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type GremlinGraphResourceInput interface {
	pulumi.Input

	ToGremlinGraphResourceOutput() GremlinGraphResourceOutput
	ToGremlinGraphResourceOutputWithContext(context.Context) GremlinGraphResourceOutput
}

type GremlinGraphResourceArgs struct {
	ConflictResolutionPolicy ConflictResolutionPolicyPtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput               `pulumi:"defaultTtl"`
	Id                       pulumi.StringInput               `pulumi:"id"`
	IndexingPolicy           IndexingPolicyPtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyPtrInput    `pulumi:"partitionKey"`
	UniqueKeyPolicy          UniqueKeyPolicyPtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (GremlinGraphResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphResource)(nil)).Elem()
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourceOutput() GremlinGraphResourceOutput {
	return i.ToGremlinGraphResourceOutputWithContext(context.Background())
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourceOutputWithContext(ctx context.Context) GremlinGraphResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphResourceOutput)
}

type GremlinGraphResourceOutput struct{ *pulumi.OutputState }

func (GremlinGraphResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphResource)(nil)).Elem()
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourceOutput() GremlinGraphResourceOutput {
	return o
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourceOutputWithContext(ctx context.Context) GremlinGraphResourceOutput {
	return o
}

func (o GremlinGraphResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *ConflictResolutionPolicy { return v.ConflictResolutionPolicy }).(ConflictResolutionPolicyPtrOutput)
}

func (o GremlinGraphResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o GremlinGraphResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinGraphResourceOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *IndexingPolicy { return v.IndexingPolicy }).(IndexingPolicyPtrOutput)
}

func (o GremlinGraphResourceOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *ContainerPartitionKey { return v.PartitionKey }).(ContainerPartitionKeyPtrOutput)
}

func (o GremlinGraphResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *UniqueKeyPolicy { return v.UniqueKeyPolicy }).(UniqueKeyPolicyPtrOutput)
}

type IncludedPath struct {
	Indexes []Indexes `pulumi:"indexes"`
	Path    *string   `pulumi:"path"`
}





type IncludedPathInput interface {
	pulumi.Input

	ToIncludedPathOutput() IncludedPathOutput
	ToIncludedPathOutputWithContext(context.Context) IncludedPathOutput
}

type IncludedPathArgs struct {
	Indexes IndexesArrayInput     `pulumi:"indexes"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (IncludedPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPath)(nil)).Elem()
}

func (i IncludedPathArgs) ToIncludedPathOutput() IncludedPathOutput {
	return i.ToIncludedPathOutputWithContext(context.Background())
}

func (i IncludedPathArgs) ToIncludedPathOutputWithContext(ctx context.Context) IncludedPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathOutput)
}





type IncludedPathArrayInput interface {
	pulumi.Input

	ToIncludedPathArrayOutput() IncludedPathArrayOutput
	ToIncludedPathArrayOutputWithContext(context.Context) IncludedPathArrayOutput
}

type IncludedPathArray []IncludedPathInput

func (IncludedPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPath)(nil)).Elem()
}

func (i IncludedPathArray) ToIncludedPathArrayOutput() IncludedPathArrayOutput {
	return i.ToIncludedPathArrayOutputWithContext(context.Background())
}

func (i IncludedPathArray) ToIncludedPathArrayOutputWithContext(ctx context.Context) IncludedPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathArrayOutput)
}

type IncludedPathOutput struct{ *pulumi.OutputState }

func (IncludedPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPath)(nil)).Elem()
}

func (o IncludedPathOutput) ToIncludedPathOutput() IncludedPathOutput {
	return o
}

func (o IncludedPathOutput) ToIncludedPathOutputWithContext(ctx context.Context) IncludedPathOutput {
	return o
}

func (o IncludedPathOutput) Indexes() IndexesArrayOutput {
	return o.ApplyT(func(v IncludedPath) []Indexes { return v.Indexes }).(IndexesArrayOutput)
}

func (o IncludedPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncludedPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type IncludedPathArrayOutput struct{ *pulumi.OutputState }

func (IncludedPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPath)(nil)).Elem()
}

func (o IncludedPathArrayOutput) ToIncludedPathArrayOutput() IncludedPathArrayOutput {
	return o
}

func (o IncludedPathArrayOutput) ToIncludedPathArrayOutputWithContext(ctx context.Context) IncludedPathArrayOutput {
	return o
}

func (o IncludedPathArrayOutput) Index(i pulumi.IntInput) IncludedPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncludedPath {
		return vs[0].([]IncludedPath)[vs[1].(int)]
	}).(IncludedPathOutput)
}

type IncludedPathResponse struct {
	Indexes []IndexesResponse `pulumi:"indexes"`
	Path    *string           `pulumi:"path"`
}

type IncludedPathResponseOutput struct{ *pulumi.OutputState }

func (IncludedPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPathResponse)(nil)).Elem()
}

func (o IncludedPathResponseOutput) ToIncludedPathResponseOutput() IncludedPathResponseOutput {
	return o
}

func (o IncludedPathResponseOutput) ToIncludedPathResponseOutputWithContext(ctx context.Context) IncludedPathResponseOutput {
	return o
}

func (o IncludedPathResponseOutput) Indexes() IndexesResponseArrayOutput {
	return o.ApplyT(func(v IncludedPathResponse) []IndexesResponse { return v.Indexes }).(IndexesResponseArrayOutput)
}

func (o IncludedPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncludedPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type IncludedPathResponseArrayOutput struct{ *pulumi.OutputState }

func (IncludedPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPathResponse)(nil)).Elem()
}

func (o IncludedPathResponseArrayOutput) ToIncludedPathResponseArrayOutput() IncludedPathResponseArrayOutput {
	return o
}

func (o IncludedPathResponseArrayOutput) ToIncludedPathResponseArrayOutputWithContext(ctx context.Context) IncludedPathResponseArrayOutput {
	return o
}

func (o IncludedPathResponseArrayOutput) Index(i pulumi.IntInput) IncludedPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncludedPathResponse {
		return vs[0].([]IncludedPathResponse)[vs[1].(int)]
	}).(IncludedPathResponseOutput)
}

type Indexes struct {
	DataType  *string `pulumi:"dataType"`
	Kind      *string `pulumi:"kind"`
	Precision *int    `pulumi:"precision"`
}


func (val *Indexes) Defaults() *Indexes {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataType) {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type IndexesInput interface {
	pulumi.Input

	ToIndexesOutput() IndexesOutput
	ToIndexesOutputWithContext(context.Context) IndexesOutput
}

type IndexesArgs struct {
	DataType  pulumi.StringPtrInput `pulumi:"dataType"`
	Kind      pulumi.StringPtrInput `pulumi:"kind"`
	Precision pulumi.IntPtrInput    `pulumi:"precision"`
}

func (IndexesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Indexes)(nil)).Elem()
}

func (i IndexesArgs) ToIndexesOutput() IndexesOutput {
	return i.ToIndexesOutputWithContext(context.Background())
}

func (i IndexesArgs) ToIndexesOutputWithContext(ctx context.Context) IndexesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesOutput)
}





type IndexesArrayInput interface {
	pulumi.Input

	ToIndexesArrayOutput() IndexesArrayOutput
	ToIndexesArrayOutputWithContext(context.Context) IndexesArrayOutput
}

type IndexesArray []IndexesInput

func (IndexesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Indexes)(nil)).Elem()
}

func (i IndexesArray) ToIndexesArrayOutput() IndexesArrayOutput {
	return i.ToIndexesArrayOutputWithContext(context.Background())
}

func (i IndexesArray) ToIndexesArrayOutputWithContext(ctx context.Context) IndexesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesArrayOutput)
}

type IndexesOutput struct{ *pulumi.OutputState }

func (IndexesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Indexes)(nil)).Elem()
}

func (o IndexesOutput) ToIndexesOutput() IndexesOutput {
	return o
}

func (o IndexesOutput) ToIndexesOutputWithContext(ctx context.Context) IndexesOutput {
	return o
}

func (o IndexesOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Indexes) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o IndexesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Indexes) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IndexesOutput) Precision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Indexes) *int { return v.Precision }).(pulumi.IntPtrOutput)
}

type IndexesArrayOutput struct{ *pulumi.OutputState }

func (IndexesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Indexes)(nil)).Elem()
}

func (o IndexesArrayOutput) ToIndexesArrayOutput() IndexesArrayOutput {
	return o
}

func (o IndexesArrayOutput) ToIndexesArrayOutputWithContext(ctx context.Context) IndexesArrayOutput {
	return o
}

func (o IndexesArrayOutput) Index(i pulumi.IntInput) IndexesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Indexes {
		return vs[0].([]Indexes)[vs[1].(int)]
	}).(IndexesOutput)
}

type IndexesResponse struct {
	DataType  *string `pulumi:"dataType"`
	Kind      *string `pulumi:"kind"`
	Precision *int    `pulumi:"precision"`
}


func (val *IndexesResponse) Defaults() *IndexesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataType) {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}

type IndexesResponseOutput struct{ *pulumi.OutputState }

func (IndexesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexesResponse)(nil)).Elem()
}

func (o IndexesResponseOutput) ToIndexesResponseOutput() IndexesResponseOutput {
	return o
}

func (o IndexesResponseOutput) ToIndexesResponseOutputWithContext(ctx context.Context) IndexesResponseOutput {
	return o
}

func (o IndexesResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o IndexesResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IndexesResponseOutput) Precision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *int { return v.Precision }).(pulumi.IntPtrOutput)
}

type IndexesResponseArrayOutput struct{ *pulumi.OutputState }

func (IndexesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IndexesResponse)(nil)).Elem()
}

func (o IndexesResponseArrayOutput) ToIndexesResponseArrayOutput() IndexesResponseArrayOutput {
	return o
}

func (o IndexesResponseArrayOutput) ToIndexesResponseArrayOutputWithContext(ctx context.Context) IndexesResponseArrayOutput {
	return o
}

func (o IndexesResponseArrayOutput) Index(i pulumi.IntInput) IndexesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IndexesResponse {
		return vs[0].([]IndexesResponse)[vs[1].(int)]
	}).(IndexesResponseOutput)
}

type IndexingPolicy struct {
	Automatic     *bool          `pulumi:"automatic"`
	ExcludedPaths []ExcludedPath `pulumi:"excludedPaths"`
	IncludedPaths []IncludedPath `pulumi:"includedPaths"`
	IndexingMode  *string        `pulumi:"indexingMode"`
}


func (val *IndexingPolicy) Defaults() *IndexingPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IndexingMode) {
		indexingMode_ := "Consistent"
		tmp.IndexingMode = &indexingMode_
	}
	return &tmp
}





type IndexingPolicyInput interface {
	pulumi.Input

	ToIndexingPolicyOutput() IndexingPolicyOutput
	ToIndexingPolicyOutputWithContext(context.Context) IndexingPolicyOutput
}

type IndexingPolicyArgs struct {
	Automatic     pulumi.BoolPtrInput    `pulumi:"automatic"`
	ExcludedPaths ExcludedPathArrayInput `pulumi:"excludedPaths"`
	IncludedPaths IncludedPathArrayInput `pulumi:"includedPaths"`
	IndexingMode  pulumi.StringPtrInput  `pulumi:"indexingMode"`
}

func (IndexingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicy)(nil)).Elem()
}

func (i IndexingPolicyArgs) ToIndexingPolicyOutput() IndexingPolicyOutput {
	return i.ToIndexingPolicyOutputWithContext(context.Background())
}

func (i IndexingPolicyArgs) ToIndexingPolicyOutputWithContext(ctx context.Context) IndexingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyOutput)
}

func (i IndexingPolicyArgs) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return i.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (i IndexingPolicyArgs) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyOutput).ToIndexingPolicyPtrOutputWithContext(ctx)
}









type IndexingPolicyPtrInput interface {
	pulumi.Input

	ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput
	ToIndexingPolicyPtrOutputWithContext(context.Context) IndexingPolicyPtrOutput
}

type indexingPolicyPtrType IndexingPolicyArgs

func IndexingPolicyPtr(v *IndexingPolicyArgs) IndexingPolicyPtrInput {
	return (*indexingPolicyPtrType)(v)
}

func (*indexingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicy)(nil)).Elem()
}

func (i *indexingPolicyPtrType) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return i.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (i *indexingPolicyPtrType) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyPtrOutput)
}

type IndexingPolicyOutput struct{ *pulumi.OutputState }

func (IndexingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicy)(nil)).Elem()
}

func (o IndexingPolicyOutput) ToIndexingPolicyOutput() IndexingPolicyOutput {
	return o
}

func (o IndexingPolicyOutput) ToIndexingPolicyOutputWithContext(ctx context.Context) IndexingPolicyOutput {
	return o
}

func (o IndexingPolicyOutput) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return o.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (o IndexingPolicyOutput) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexingPolicy) *IndexingPolicy {
		return &v
	}).(IndexingPolicyPtrOutput)
}

func (o IndexingPolicyOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IndexingPolicy) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyOutput) ExcludedPaths() ExcludedPathArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []ExcludedPath { return v.ExcludedPaths }).(ExcludedPathArrayOutput)
}

func (o IndexingPolicyOutput) IncludedPaths() IncludedPathArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []IncludedPath { return v.IncludedPaths }).(IncludedPathArrayOutput)
}

func (o IndexingPolicyOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexingPolicy) *string { return v.IndexingMode }).(pulumi.StringPtrOutput)
}

type IndexingPolicyPtrOutput struct{ *pulumi.OutputState }

func (IndexingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicy)(nil)).Elem()
}

func (o IndexingPolicyPtrOutput) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return o
}

func (o IndexingPolicyPtrOutput) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return o
}

func (o IndexingPolicyPtrOutput) Elem() IndexingPolicyOutput {
	return o.ApplyT(func(v *IndexingPolicy) IndexingPolicy {
		if v != nil {
			return *v
		}
		var ret IndexingPolicy
		return ret
	}).(IndexingPolicyOutput)
}

func (o IndexingPolicyPtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IndexingPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Automatic
	}).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyPtrOutput) ExcludedPaths() ExcludedPathArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []ExcludedPath {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(ExcludedPathArrayOutput)
}

func (o IndexingPolicyPtrOutput) IncludedPaths() IncludedPathArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []IncludedPath {
		if v == nil {
			return nil
		}
		return v.IncludedPaths
	}).(IncludedPathArrayOutput)
}

func (o IndexingPolicyPtrOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexingPolicy) *string {
		if v == nil {
			return nil
		}
		return v.IndexingMode
	}).(pulumi.StringPtrOutput)
}

type IndexingPolicyResponse struct {
	Automatic     *bool                  `pulumi:"automatic"`
	ExcludedPaths []ExcludedPathResponse `pulumi:"excludedPaths"`
	IncludedPaths []IncludedPathResponse `pulumi:"includedPaths"`
	IndexingMode  *string                `pulumi:"indexingMode"`
}


func (val *IndexingPolicyResponse) Defaults() *IndexingPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IndexingMode) {
		indexingMode_ := "Consistent"
		tmp.IndexingMode = &indexingMode_
	}
	return &tmp
}

type IndexingPolicyResponseOutput struct{ *pulumi.OutputState }

func (IndexingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicyResponse)(nil)).Elem()
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponseOutput() IndexingPolicyResponseOutput {
	return o
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponseOutputWithContext(ctx context.Context) IndexingPolicyResponseOutput {
	return o
}

func (o IndexingPolicyResponseOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyResponseOutput) ExcludedPaths() ExcludedPathResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []ExcludedPathResponse { return v.ExcludedPaths }).(ExcludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponseOutput) IncludedPaths() IncludedPathResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []IncludedPathResponse { return v.IncludedPaths }).(IncludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponseOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) *string { return v.IndexingMode }).(pulumi.StringPtrOutput)
}

type IndexingPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (IndexingPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicyResponse)(nil)).Elem()
}

func (o IndexingPolicyResponsePtrOutput) ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput {
	return o
}

func (o IndexingPolicyResponsePtrOutput) ToIndexingPolicyResponsePtrOutputWithContext(ctx context.Context) IndexingPolicyResponsePtrOutput {
	return o
}

func (o IndexingPolicyResponsePtrOutput) Elem() IndexingPolicyResponseOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) IndexingPolicyResponse {
		if v != nil {
			return *v
		}
		var ret IndexingPolicyResponse
		return ret
	}).(IndexingPolicyResponseOutput)
}

func (o IndexingPolicyResponsePtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Automatic
	}).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyResponsePtrOutput) ExcludedPaths() ExcludedPathResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []ExcludedPathResponse {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(ExcludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) IncludedPaths() IncludedPathResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []IncludedPathResponse {
		if v == nil {
			return nil
		}
		return v.IncludedPaths
	}).(IncludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.IndexingMode
	}).(pulumi.StringPtrOutput)
}

type Location struct {
	FailoverPriority *int    `pulumi:"failoverPriority"`
	IsZoneRedundant  *bool   `pulumi:"isZoneRedundant"`
	LocationName     *string `pulumi:"locationName"`
}





type LocationInput interface {
	pulumi.Input

	ToLocationOutput() LocationOutput
	ToLocationOutputWithContext(context.Context) LocationOutput
}

type LocationArgs struct {
	FailoverPriority pulumi.IntPtrInput    `pulumi:"failoverPriority"`
	IsZoneRedundant  pulumi.BoolPtrInput   `pulumi:"isZoneRedundant"`
	LocationName     pulumi.StringPtrInput `pulumi:"locationName"`
}

func (LocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (i LocationArgs) ToLocationOutput() LocationOutput {
	return i.ToLocationOutputWithContext(context.Background())
}

func (i LocationArgs) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationOutput)
}





type LocationArrayInput interface {
	pulumi.Input

	ToLocationArrayOutput() LocationArrayOutput
	ToLocationArrayOutputWithContext(context.Context) LocationArrayOutput
}

type LocationArray []LocationInput

func (LocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Location)(nil)).Elem()
}

func (i LocationArray) ToLocationArrayOutput() LocationArrayOutput {
	return i.ToLocationArrayOutputWithContext(context.Background())
}

func (i LocationArray) ToLocationArrayOutputWithContext(ctx context.Context) LocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationArrayOutput)
}

type LocationOutput struct{ *pulumi.OutputState }

func (LocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (o LocationOutput) ToLocationOutput() LocationOutput {
	return o
}

func (o LocationOutput) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return o
}

func (o LocationOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Location) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o LocationOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Location) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

func (o LocationOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Location) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

type LocationArrayOutput struct{ *pulumi.OutputState }

func (LocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Location)(nil)).Elem()
}

func (o LocationArrayOutput) ToLocationArrayOutput() LocationArrayOutput {
	return o
}

func (o LocationArrayOutput) ToLocationArrayOutputWithContext(ctx context.Context) LocationArrayOutput {
	return o
}

func (o LocationArrayOutput) Index(i pulumi.IntInput) LocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Location {
		return vs[0].([]Location)[vs[1].(int)]
	}).(LocationOutput)
}

type LocationResponse struct {
	DocumentEndpoint  string  `pulumi:"documentEndpoint"`
	FailoverPriority  *int    `pulumi:"failoverPriority"`
	Id                string  `pulumi:"id"`
	IsZoneRedundant   *bool   `pulumi:"isZoneRedundant"`
	LocationName      *string `pulumi:"locationName"`
	ProvisioningState string  `pulumi:"provisioningState"`
}

type LocationResponseOutput struct{ *pulumi.OutputState }

func (LocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationResponse)(nil)).Elem()
}

func (o LocationResponseOutput) ToLocationResponseOutput() LocationResponseOutput {
	return o
}

func (o LocationResponseOutput) ToLocationResponseOutputWithContext(ctx context.Context) LocationResponseOutput {
	return o
}

func (o LocationResponseOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o LocationResponseOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LocationResponse) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o LocationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o LocationResponseOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LocationResponse) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

func (o LocationResponseOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationResponse) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

func (o LocationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type LocationResponseArrayOutput struct{ *pulumi.OutputState }

func (LocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocationResponse)(nil)).Elem()
}

func (o LocationResponseArrayOutput) ToLocationResponseArrayOutput() LocationResponseArrayOutput {
	return o
}

func (o LocationResponseArrayOutput) ToLocationResponseArrayOutputWithContext(ctx context.Context) LocationResponseArrayOutput {
	return o
}

func (o LocationResponseArrayOutput) Index(i pulumi.IntInput) LocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocationResponse {
		return vs[0].([]LocationResponse)[vs[1].(int)]
	}).(LocationResponseOutput)
}

type MongoDBCollectionResource struct {
	Id       string            `pulumi:"id"`
	Indexes  []MongoIndex      `pulumi:"indexes"`
	ShardKey map[string]string `pulumi:"shardKey"`
}





type MongoDBCollectionResourceInput interface {
	pulumi.Input

	ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput
	ToMongoDBCollectionResourceOutputWithContext(context.Context) MongoDBCollectionResourceOutput
}

type MongoDBCollectionResourceArgs struct {
	Id       pulumi.StringInput    `pulumi:"id"`
	Indexes  MongoIndexArrayInput  `pulumi:"indexes"`
	ShardKey pulumi.StringMapInput `pulumi:"shardKey"`
}

func (MongoDBCollectionResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionResource)(nil)).Elem()
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput {
	return i.ToMongoDBCollectionResourceOutputWithContext(context.Background())
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourceOutputWithContext(ctx context.Context) MongoDBCollectionResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionResourceOutput)
}

type MongoDBCollectionResourceOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionResource)(nil)).Elem()
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput {
	return o
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourceOutputWithContext(ctx context.Context) MongoDBCollectionResourceOutput {
	return o
}

func (o MongoDBCollectionResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBCollectionResourceOutput) Indexes() MongoIndexArrayOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) []MongoIndex { return v.Indexes }).(MongoIndexArrayOutput)
}

func (o MongoDBCollectionResourceOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) map[string]string { return v.ShardKey }).(pulumi.StringMapOutput)
}

type MongoDBDatabaseResource struct {
	Id string `pulumi:"id"`
}





type MongoDBDatabaseResourceInput interface {
	pulumi.Input

	ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput
	ToMongoDBDatabaseResourceOutputWithContext(context.Context) MongoDBDatabaseResourceOutput
}

type MongoDBDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MongoDBDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseResource)(nil)).Elem()
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput {
	return i.ToMongoDBDatabaseResourceOutputWithContext(context.Background())
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseResourceOutput)
}

type MongoDBDatabaseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseResource)(nil)).Elem()
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput {
	return o
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseResourceOutput {
	return o
}

func (o MongoDBDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type MongoIndex struct {
	Key     *MongoIndexKeys    `pulumi:"key"`
	Options *MongoIndexOptions `pulumi:"options"`
}





type MongoIndexInput interface {
	pulumi.Input

	ToMongoIndexOutput() MongoIndexOutput
	ToMongoIndexOutputWithContext(context.Context) MongoIndexOutput
}

type MongoIndexArgs struct {
	Key     MongoIndexKeysPtrInput    `pulumi:"key"`
	Options MongoIndexOptionsPtrInput `pulumi:"options"`
}

func (MongoIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndex)(nil)).Elem()
}

func (i MongoIndexArgs) ToMongoIndexOutput() MongoIndexOutput {
	return i.ToMongoIndexOutputWithContext(context.Background())
}

func (i MongoIndexArgs) ToMongoIndexOutputWithContext(ctx context.Context) MongoIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOutput)
}





type MongoIndexArrayInput interface {
	pulumi.Input

	ToMongoIndexArrayOutput() MongoIndexArrayOutput
	ToMongoIndexArrayOutputWithContext(context.Context) MongoIndexArrayOutput
}

type MongoIndexArray []MongoIndexInput

func (MongoIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndex)(nil)).Elem()
}

func (i MongoIndexArray) ToMongoIndexArrayOutput() MongoIndexArrayOutput {
	return i.ToMongoIndexArrayOutputWithContext(context.Background())
}

func (i MongoIndexArray) ToMongoIndexArrayOutputWithContext(ctx context.Context) MongoIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexArrayOutput)
}

type MongoIndexOutput struct{ *pulumi.OutputState }

func (MongoIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndex)(nil)).Elem()
}

func (o MongoIndexOutput) ToMongoIndexOutput() MongoIndexOutput {
	return o
}

func (o MongoIndexOutput) ToMongoIndexOutputWithContext(ctx context.Context) MongoIndexOutput {
	return o
}

func (o MongoIndexOutput) Key() MongoIndexKeysPtrOutput {
	return o.ApplyT(func(v MongoIndex) *MongoIndexKeys { return v.Key }).(MongoIndexKeysPtrOutput)
}

func (o MongoIndexOutput) Options() MongoIndexOptionsPtrOutput {
	return o.ApplyT(func(v MongoIndex) *MongoIndexOptions { return v.Options }).(MongoIndexOptionsPtrOutput)
}

type MongoIndexArrayOutput struct{ *pulumi.OutputState }

func (MongoIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndex)(nil)).Elem()
}

func (o MongoIndexArrayOutput) ToMongoIndexArrayOutput() MongoIndexArrayOutput {
	return o
}

func (o MongoIndexArrayOutput) ToMongoIndexArrayOutputWithContext(ctx context.Context) MongoIndexArrayOutput {
	return o
}

func (o MongoIndexArrayOutput) Index(i pulumi.IntInput) MongoIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MongoIndex {
		return vs[0].([]MongoIndex)[vs[1].(int)]
	}).(MongoIndexOutput)
}

type MongoIndexKeys struct {
	Keys []string `pulumi:"keys"`
}





type MongoIndexKeysInput interface {
	pulumi.Input

	ToMongoIndexKeysOutput() MongoIndexKeysOutput
	ToMongoIndexKeysOutputWithContext(context.Context) MongoIndexKeysOutput
}

type MongoIndexKeysArgs struct {
	Keys pulumi.StringArrayInput `pulumi:"keys"`
}

func (MongoIndexKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeys)(nil)).Elem()
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysOutput() MongoIndexKeysOutput {
	return i.ToMongoIndexKeysOutputWithContext(context.Background())
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysOutputWithContext(ctx context.Context) MongoIndexKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysOutput)
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return i.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysOutput).ToMongoIndexKeysPtrOutputWithContext(ctx)
}









type MongoIndexKeysPtrInput interface {
	pulumi.Input

	ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput
	ToMongoIndexKeysPtrOutputWithContext(context.Context) MongoIndexKeysPtrOutput
}

type mongoIndexKeysPtrType MongoIndexKeysArgs

func MongoIndexKeysPtr(v *MongoIndexKeysArgs) MongoIndexKeysPtrInput {
	return (*mongoIndexKeysPtrType)(v)
}

func (*mongoIndexKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeys)(nil)).Elem()
}

func (i *mongoIndexKeysPtrType) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return i.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (i *mongoIndexKeysPtrType) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysPtrOutput)
}

type MongoIndexKeysOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeys)(nil)).Elem()
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysOutput() MongoIndexKeysOutput {
	return o
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysOutputWithContext(ctx context.Context) MongoIndexKeysOutput {
	return o
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return o.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexKeys) *MongoIndexKeys {
		return &v
	}).(MongoIndexKeysPtrOutput)
}

func (o MongoIndexKeysOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoIndexKeys) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

type MongoIndexKeysPtrOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeys)(nil)).Elem()
}

func (o MongoIndexKeysPtrOutput) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return o
}

func (o MongoIndexKeysPtrOutput) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return o
}

func (o MongoIndexKeysPtrOutput) Elem() MongoIndexKeysOutput {
	return o.ApplyT(func(v *MongoIndexKeys) MongoIndexKeys {
		if v != nil {
			return *v
		}
		var ret MongoIndexKeys
		return ret
	}).(MongoIndexKeysOutput)
}

func (o MongoIndexKeysPtrOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoIndexKeys) []string {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.StringArrayOutput)
}

type MongoIndexKeysResponse struct {
	Keys []string `pulumi:"keys"`
}

type MongoIndexKeysResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeysResponse)(nil)).Elem()
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponseOutput() MongoIndexKeysResponseOutput {
	return o
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponseOutputWithContext(ctx context.Context) MongoIndexKeysResponseOutput {
	return o
}

func (o MongoIndexKeysResponseOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoIndexKeysResponse) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

type MongoIndexKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeysResponse)(nil)).Elem()
}

func (o MongoIndexKeysResponsePtrOutput) ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput {
	return o
}

func (o MongoIndexKeysResponsePtrOutput) ToMongoIndexKeysResponsePtrOutputWithContext(ctx context.Context) MongoIndexKeysResponsePtrOutput {
	return o
}

func (o MongoIndexKeysResponsePtrOutput) Elem() MongoIndexKeysResponseOutput {
	return o.ApplyT(func(v *MongoIndexKeysResponse) MongoIndexKeysResponse {
		if v != nil {
			return *v
		}
		var ret MongoIndexKeysResponse
		return ret
	}).(MongoIndexKeysResponseOutput)
}

func (o MongoIndexKeysResponsePtrOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoIndexKeysResponse) []string {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.StringArrayOutput)
}

type MongoIndexOptions struct {
	ExpireAfterSeconds *int  `pulumi:"expireAfterSeconds"`
	Unique             *bool `pulumi:"unique"`
}





type MongoIndexOptionsInput interface {
	pulumi.Input

	ToMongoIndexOptionsOutput() MongoIndexOptionsOutput
	ToMongoIndexOptionsOutputWithContext(context.Context) MongoIndexOptionsOutput
}

type MongoIndexOptionsArgs struct {
	ExpireAfterSeconds pulumi.IntPtrInput  `pulumi:"expireAfterSeconds"`
	Unique             pulumi.BoolPtrInput `pulumi:"unique"`
}

func (MongoIndexOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptions)(nil)).Elem()
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsOutput() MongoIndexOptionsOutput {
	return i.ToMongoIndexOptionsOutputWithContext(context.Background())
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsOutputWithContext(ctx context.Context) MongoIndexOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsOutput)
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return i.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsOutput).ToMongoIndexOptionsPtrOutputWithContext(ctx)
}









type MongoIndexOptionsPtrInput interface {
	pulumi.Input

	ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput
	ToMongoIndexOptionsPtrOutputWithContext(context.Context) MongoIndexOptionsPtrOutput
}

type mongoIndexOptionsPtrType MongoIndexOptionsArgs

func MongoIndexOptionsPtr(v *MongoIndexOptionsArgs) MongoIndexOptionsPtrInput {
	return (*mongoIndexOptionsPtrType)(v)
}

func (*mongoIndexOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptions)(nil)).Elem()
}

func (i *mongoIndexOptionsPtrType) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return i.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (i *mongoIndexOptionsPtrType) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsPtrOutput)
}

type MongoIndexOptionsOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptions)(nil)).Elem()
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsOutput() MongoIndexOptionsOutput {
	return o
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsOutputWithContext(ctx context.Context) MongoIndexOptionsOutput {
	return o
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return o.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexOptions) *MongoIndexOptions {
		return &v
	}).(MongoIndexOptionsPtrOutput)
}

func (o MongoIndexOptionsOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoIndexOptions) *int { return v.ExpireAfterSeconds }).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MongoIndexOptions) *bool { return v.Unique }).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptions)(nil)).Elem()
}

func (o MongoIndexOptionsPtrOutput) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return o
}

func (o MongoIndexOptionsPtrOutput) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return o
}

func (o MongoIndexOptionsPtrOutput) Elem() MongoIndexOptionsOutput {
	return o.ApplyT(func(v *MongoIndexOptions) MongoIndexOptions {
		if v != nil {
			return *v
		}
		var ret MongoIndexOptions
		return ret
	}).(MongoIndexOptionsOutput)
}

func (o MongoIndexOptionsPtrOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptions) *int {
		if v == nil {
			return nil
		}
		return v.ExpireAfterSeconds
	}).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsPtrOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptions) *bool {
		if v == nil {
			return nil
		}
		return v.Unique
	}).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsResponse struct {
	ExpireAfterSeconds *int  `pulumi:"expireAfterSeconds"`
	Unique             *bool `pulumi:"unique"`
}

type MongoIndexOptionsResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptionsResponse)(nil)).Elem()
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponseOutput() MongoIndexOptionsResponseOutput {
	return o
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponseOutputWithContext(ctx context.Context) MongoIndexOptionsResponseOutput {
	return o
}

func (o MongoIndexOptionsResponseOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoIndexOptionsResponse) *int { return v.ExpireAfterSeconds }).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsResponseOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MongoIndexOptionsResponse) *bool { return v.Unique }).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptionsResponse)(nil)).Elem()
}

func (o MongoIndexOptionsResponsePtrOutput) ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput {
	return o
}

func (o MongoIndexOptionsResponsePtrOutput) ToMongoIndexOptionsResponsePtrOutputWithContext(ctx context.Context) MongoIndexOptionsResponsePtrOutput {
	return o
}

func (o MongoIndexOptionsResponsePtrOutput) Elem() MongoIndexOptionsResponseOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) MongoIndexOptionsResponse {
		if v != nil {
			return *v
		}
		var ret MongoIndexOptionsResponse
		return ret
	}).(MongoIndexOptionsResponseOutput)
}

func (o MongoIndexOptionsResponsePtrOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ExpireAfterSeconds
	}).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsResponsePtrOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Unique
	}).(pulumi.BoolPtrOutput)
}

type MongoIndexResponse struct {
	Key     *MongoIndexKeysResponse    `pulumi:"key"`
	Options *MongoIndexOptionsResponse `pulumi:"options"`
}

type MongoIndexResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexResponse)(nil)).Elem()
}

func (o MongoIndexResponseOutput) ToMongoIndexResponseOutput() MongoIndexResponseOutput {
	return o
}

func (o MongoIndexResponseOutput) ToMongoIndexResponseOutputWithContext(ctx context.Context) MongoIndexResponseOutput {
	return o
}

func (o MongoIndexResponseOutput) Key() MongoIndexKeysResponsePtrOutput {
	return o.ApplyT(func(v MongoIndexResponse) *MongoIndexKeysResponse { return v.Key }).(MongoIndexKeysResponsePtrOutput)
}

func (o MongoIndexResponseOutput) Options() MongoIndexOptionsResponsePtrOutput {
	return o.ApplyT(func(v MongoIndexResponse) *MongoIndexOptionsResponse { return v.Options }).(MongoIndexOptionsResponsePtrOutput)
}

type MongoIndexResponseArrayOutput struct{ *pulumi.OutputState }

func (MongoIndexResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndexResponse)(nil)).Elem()
}

func (o MongoIndexResponseArrayOutput) ToMongoIndexResponseArrayOutput() MongoIndexResponseArrayOutput {
	return o
}

func (o MongoIndexResponseArrayOutput) ToMongoIndexResponseArrayOutputWithContext(ctx context.Context) MongoIndexResponseArrayOutput {
	return o
}

func (o MongoIndexResponseArrayOutput) Index(i pulumi.IntInput) MongoIndexResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MongoIndexResponse {
		return vs[0].([]MongoIndexResponse)[vs[1].(int)]
	}).(MongoIndexResponseOutput)
}

type SqlContainerResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                      `pulumi:"defaultTtl"`
	Id                       string                    `pulumi:"id"`
	IndexingPolicy           *IndexingPolicy           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKey    `pulumi:"partitionKey"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `pulumi:"uniqueKeyPolicy"`
}


func (val *SqlContainerResource) Defaults() *SqlContainerResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type SqlContainerResourceInput interface {
	pulumi.Input

	ToSqlContainerResourceOutput() SqlContainerResourceOutput
	ToSqlContainerResourceOutputWithContext(context.Context) SqlContainerResourceOutput
}

type SqlContainerResourceArgs struct {
	ConflictResolutionPolicy ConflictResolutionPolicyPtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput               `pulumi:"defaultTtl"`
	Id                       pulumi.StringInput               `pulumi:"id"`
	IndexingPolicy           IndexingPolicyPtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyPtrInput    `pulumi:"partitionKey"`
	UniqueKeyPolicy          UniqueKeyPolicyPtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (SqlContainerResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerResource)(nil)).Elem()
}

func (i SqlContainerResourceArgs) ToSqlContainerResourceOutput() SqlContainerResourceOutput {
	return i.ToSqlContainerResourceOutputWithContext(context.Background())
}

func (i SqlContainerResourceArgs) ToSqlContainerResourceOutputWithContext(ctx context.Context) SqlContainerResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerResourceOutput)
}

type SqlContainerResourceOutput struct{ *pulumi.OutputState }

func (SqlContainerResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerResource)(nil)).Elem()
}

func (o SqlContainerResourceOutput) ToSqlContainerResourceOutput() SqlContainerResourceOutput {
	return o
}

func (o SqlContainerResourceOutput) ToSqlContainerResourceOutputWithContext(ctx context.Context) SqlContainerResourceOutput {
	return o
}

func (o SqlContainerResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *ConflictResolutionPolicy { return v.ConflictResolutionPolicy }).(ConflictResolutionPolicyPtrOutput)
}

func (o SqlContainerResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o SqlContainerResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlContainerResourceOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *IndexingPolicy { return v.IndexingPolicy }).(IndexingPolicyPtrOutput)
}

func (o SqlContainerResourceOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *ContainerPartitionKey { return v.PartitionKey }).(ContainerPartitionKeyPtrOutput)
}

func (o SqlContainerResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *UniqueKeyPolicy { return v.UniqueKeyPolicy }).(UniqueKeyPolicyPtrOutput)
}

type SqlDatabaseResource struct {
	Id string `pulumi:"id"`
}





type SqlDatabaseResourceInput interface {
	pulumi.Input

	ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput
	ToSqlDatabaseResourceOutputWithContext(context.Context) SqlDatabaseResourceOutput
}

type SqlDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SqlDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResource)(nil)).Elem()
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput {
	return i.ToSqlDatabaseResourceOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourceOutputWithContext(ctx context.Context) SqlDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceOutput)
}

type SqlDatabaseResourceOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResource)(nil)).Elem()
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput {
	return o
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourceOutputWithContext(ctx context.Context) SqlDatabaseResourceOutput {
	return o
}

func (o SqlDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type TableResource struct {
	Id string `pulumi:"id"`
}





type TableResourceInput interface {
	pulumi.Input

	ToTableResourceOutput() TableResourceOutput
	ToTableResourceOutputWithContext(context.Context) TableResourceOutput
}

type TableResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (TableResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableResource)(nil)).Elem()
}

func (i TableResourceArgs) ToTableResourceOutput() TableResourceOutput {
	return i.ToTableResourceOutputWithContext(context.Background())
}

func (i TableResourceArgs) ToTableResourceOutputWithContext(ctx context.Context) TableResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourceOutput)
}

type TableResourceOutput struct{ *pulumi.OutputState }

func (TableResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableResource)(nil)).Elem()
}

func (o TableResourceOutput) ToTableResourceOutput() TableResourceOutput {
	return o
}

func (o TableResourceOutput) ToTableResourceOutputWithContext(ctx context.Context) TableResourceOutput {
	return o
}

func (o TableResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TableResource) string { return v.Id }).(pulumi.StringOutput)
}

type UniqueKey struct {
	Paths []string `pulumi:"paths"`
}





type UniqueKeyInput interface {
	pulumi.Input

	ToUniqueKeyOutput() UniqueKeyOutput
	ToUniqueKeyOutputWithContext(context.Context) UniqueKeyOutput
}

type UniqueKeyArgs struct {
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (UniqueKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKey)(nil)).Elem()
}

func (i UniqueKeyArgs) ToUniqueKeyOutput() UniqueKeyOutput {
	return i.ToUniqueKeyOutputWithContext(context.Background())
}

func (i UniqueKeyArgs) ToUniqueKeyOutputWithContext(ctx context.Context) UniqueKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyOutput)
}





type UniqueKeyArrayInput interface {
	pulumi.Input

	ToUniqueKeyArrayOutput() UniqueKeyArrayOutput
	ToUniqueKeyArrayOutputWithContext(context.Context) UniqueKeyArrayOutput
}

type UniqueKeyArray []UniqueKeyInput

func (UniqueKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKey)(nil)).Elem()
}

func (i UniqueKeyArray) ToUniqueKeyArrayOutput() UniqueKeyArrayOutput {
	return i.ToUniqueKeyArrayOutputWithContext(context.Background())
}

func (i UniqueKeyArray) ToUniqueKeyArrayOutputWithContext(ctx context.Context) UniqueKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyArrayOutput)
}

type UniqueKeyOutput struct{ *pulumi.OutputState }

func (UniqueKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKey)(nil)).Elem()
}

func (o UniqueKeyOutput) ToUniqueKeyOutput() UniqueKeyOutput {
	return o
}

func (o UniqueKeyOutput) ToUniqueKeyOutputWithContext(ctx context.Context) UniqueKeyOutput {
	return o
}

func (o UniqueKeyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UniqueKey) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type UniqueKeyArrayOutput struct{ *pulumi.OutputState }

func (UniqueKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKey)(nil)).Elem()
}

func (o UniqueKeyArrayOutput) ToUniqueKeyArrayOutput() UniqueKeyArrayOutput {
	return o
}

func (o UniqueKeyArrayOutput) ToUniqueKeyArrayOutputWithContext(ctx context.Context) UniqueKeyArrayOutput {
	return o
}

func (o UniqueKeyArrayOutput) Index(i pulumi.IntInput) UniqueKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UniqueKey {
		return vs[0].([]UniqueKey)[vs[1].(int)]
	}).(UniqueKeyOutput)
}

type UniqueKeyPolicy struct {
	UniqueKeys []UniqueKey `pulumi:"uniqueKeys"`
}





type UniqueKeyPolicyInput interface {
	pulumi.Input

	ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput
	ToUniqueKeyPolicyOutputWithContext(context.Context) UniqueKeyPolicyOutput
}

type UniqueKeyPolicyArgs struct {
	UniqueKeys UniqueKeyArrayInput `pulumi:"uniqueKeys"`
}

func (UniqueKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicy)(nil)).Elem()
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput {
	return i.ToUniqueKeyPolicyOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyOutputWithContext(ctx context.Context) UniqueKeyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyOutput)
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return i.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyOutput).ToUniqueKeyPolicyPtrOutputWithContext(ctx)
}









type UniqueKeyPolicyPtrInput interface {
	pulumi.Input

	ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput
	ToUniqueKeyPolicyPtrOutputWithContext(context.Context) UniqueKeyPolicyPtrOutput
}

type uniqueKeyPolicyPtrType UniqueKeyPolicyArgs

func UniqueKeyPolicyPtr(v *UniqueKeyPolicyArgs) UniqueKeyPolicyPtrInput {
	return (*uniqueKeyPolicyPtrType)(v)
}

func (*uniqueKeyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicy)(nil)).Elem()
}

func (i *uniqueKeyPolicyPtrType) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return i.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (i *uniqueKeyPolicyPtrType) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyPtrOutput)
}

type UniqueKeyPolicyOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicy)(nil)).Elem()
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput {
	return o
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyOutputWithContext(ctx context.Context) UniqueKeyPolicyOutput {
	return o
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return o.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UniqueKeyPolicy) *UniqueKeyPolicy {
		return &v
	}).(UniqueKeyPolicyPtrOutput)
}

func (o UniqueKeyPolicyOutput) UniqueKeys() UniqueKeyArrayOutput {
	return o.ApplyT(func(v UniqueKeyPolicy) []UniqueKey { return v.UniqueKeys }).(UniqueKeyArrayOutput)
}

type UniqueKeyPolicyPtrOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicy)(nil)).Elem()
}

func (o UniqueKeyPolicyPtrOutput) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return o
}

func (o UniqueKeyPolicyPtrOutput) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return o
}

func (o UniqueKeyPolicyPtrOutput) Elem() UniqueKeyPolicyOutput {
	return o.ApplyT(func(v *UniqueKeyPolicy) UniqueKeyPolicy {
		if v != nil {
			return *v
		}
		var ret UniqueKeyPolicy
		return ret
	}).(UniqueKeyPolicyOutput)
}

func (o UniqueKeyPolicyPtrOutput) UniqueKeys() UniqueKeyArrayOutput {
	return o.ApplyT(func(v *UniqueKeyPolicy) []UniqueKey {
		if v == nil {
			return nil
		}
		return v.UniqueKeys
	}).(UniqueKeyArrayOutput)
}

type UniqueKeyPolicyResponse struct {
	UniqueKeys []UniqueKeyResponse `pulumi:"uniqueKeys"`
}

type UniqueKeyPolicyResponseOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicyResponse)(nil)).Elem()
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponseOutput() UniqueKeyPolicyResponseOutput {
	return o
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponseOutputWithContext(ctx context.Context) UniqueKeyPolicyResponseOutput {
	return o
}

func (o UniqueKeyPolicyResponseOutput) UniqueKeys() UniqueKeyResponseArrayOutput {
	return o.ApplyT(func(v UniqueKeyPolicyResponse) []UniqueKeyResponse { return v.UniqueKeys }).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicyResponse)(nil)).Elem()
}

func (o UniqueKeyPolicyResponsePtrOutput) ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput {
	return o
}

func (o UniqueKeyPolicyResponsePtrOutput) ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx context.Context) UniqueKeyPolicyResponsePtrOutput {
	return o
}

func (o UniqueKeyPolicyResponsePtrOutput) Elem() UniqueKeyPolicyResponseOutput {
	return o.ApplyT(func(v *UniqueKeyPolicyResponse) UniqueKeyPolicyResponse {
		if v != nil {
			return *v
		}
		var ret UniqueKeyPolicyResponse
		return ret
	}).(UniqueKeyPolicyResponseOutput)
}

func (o UniqueKeyPolicyResponsePtrOutput) UniqueKeys() UniqueKeyResponseArrayOutput {
	return o.ApplyT(func(v *UniqueKeyPolicyResponse) []UniqueKeyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeys
	}).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyResponse struct {
	Paths []string `pulumi:"paths"`
}

type UniqueKeyResponseOutput struct{ *pulumi.OutputState }

func (UniqueKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyResponse)(nil)).Elem()
}

func (o UniqueKeyResponseOutput) ToUniqueKeyResponseOutput() UniqueKeyResponseOutput {
	return o
}

func (o UniqueKeyResponseOutput) ToUniqueKeyResponseOutputWithContext(ctx context.Context) UniqueKeyResponseOutput {
	return o
}

func (o UniqueKeyResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UniqueKeyResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type UniqueKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (UniqueKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKeyResponse)(nil)).Elem()
}

func (o UniqueKeyResponseArrayOutput) ToUniqueKeyResponseArrayOutput() UniqueKeyResponseArrayOutput {
	return o
}

func (o UniqueKeyResponseArrayOutput) ToUniqueKeyResponseArrayOutputWithContext(ctx context.Context) UniqueKeyResponseArrayOutput {
	return o
}

func (o UniqueKeyResponseArrayOutput) Index(i pulumi.IntInput) UniqueKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UniqueKeyResponse {
		return vs[0].([]UniqueKeyResponse)[vs[1].(int)]
	}).(UniqueKeyResponseOutput)
}

type VirtualNetworkRule struct {
	Id                               *string `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint *bool   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Id                               pulumi.StringPtrInput `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVNetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *bool { return v.IgnoreMissingVNetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Id                               *string `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint *bool   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) IgnoreMissingVNetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *bool { return v.IgnoreMissingVNetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CapabilityOutput{})
	pulumi.RegisterOutputType(CapabilityArrayOutput{})
	pulumi.RegisterOutputType(CapabilityResponseOutput{})
	pulumi.RegisterOutputType(CapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceResourceOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyArrayOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyResponseOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CassandraSchemaOutput{})
	pulumi.RegisterOutputType(CassandraSchemaPtrOutput{})
	pulumi.RegisterOutputType(CassandraSchemaResponseOutput{})
	pulumi.RegisterOutputType(CassandraSchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(CassandraTableResourceOutput{})
	pulumi.RegisterOutputType(ClusterKeyOutput{})
	pulumi.RegisterOutputType(ClusterKeyArrayOutput{})
	pulumi.RegisterOutputType(ClusterKeyResponseOutput{})
	pulumi.RegisterOutputType(ClusterKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ColumnOutput{})
	pulumi.RegisterOutputType(ColumnArrayOutput{})
	pulumi.RegisterOutputType(ColumnResponseOutput{})
	pulumi.RegisterOutputType(ColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyResponseOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyResponseOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyPtrOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(ExcludedPathOutput{})
	pulumi.RegisterOutputType(ExcludedPathArrayOutput{})
	pulumi.RegisterOutputType(ExcludedPathResponseOutput{})
	pulumi.RegisterOutputType(ExcludedPathResponseArrayOutput{})
	pulumi.RegisterOutputType(FailoverPolicyResponseOutput{})
	pulumi.RegisterOutputType(FailoverPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseResourceOutput{})
	pulumi.RegisterOutputType(GremlinGraphResourceOutput{})
	pulumi.RegisterOutputType(IncludedPathOutput{})
	pulumi.RegisterOutputType(IncludedPathArrayOutput{})
	pulumi.RegisterOutputType(IncludedPathResponseOutput{})
	pulumi.RegisterOutputType(IncludedPathResponseArrayOutput{})
	pulumi.RegisterOutputType(IndexesOutput{})
	pulumi.RegisterOutputType(IndexesArrayOutput{})
	pulumi.RegisterOutputType(IndexesResponseOutput{})
	pulumi.RegisterOutputType(IndexesResponseArrayOutput{})
	pulumi.RegisterOutputType(IndexingPolicyOutput{})
	pulumi.RegisterOutputType(IndexingPolicyPtrOutput{})
	pulumi.RegisterOutputType(IndexingPolicyResponseOutput{})
	pulumi.RegisterOutputType(IndexingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(LocationOutput{})
	pulumi.RegisterOutputType(LocationArrayOutput{})
	pulumi.RegisterOutputType(LocationResponseOutput{})
	pulumi.RegisterOutputType(LocationResponseArrayOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionResourceOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseResourceOutput{})
	pulumi.RegisterOutputType(MongoIndexOutput{})
	pulumi.RegisterOutputType(MongoIndexArrayOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysPtrOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexResponseArrayOutput{})
	pulumi.RegisterOutputType(SqlContainerResourceOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceOutput{})
	pulumi.RegisterOutputType(TableResourceOutput{})
	pulumi.RegisterOutputType(UniqueKeyOutput{})
	pulumi.RegisterOutputType(UniqueKeyArrayOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyPtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyResponseOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyResponseOutput{})
	pulumi.RegisterOutputType(UniqueKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
