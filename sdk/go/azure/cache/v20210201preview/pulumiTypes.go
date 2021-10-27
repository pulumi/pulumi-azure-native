


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabasePropertiesGeoReplication struct {
	GroupNickname   *string          `pulumi:"groupNickname"`
	LinkedDatabases []LinkedDatabase `pulumi:"linkedDatabases"`
}





type DatabasePropertiesGeoReplicationInput interface {
	pulumi.Input

	ToDatabasePropertiesGeoReplicationOutput() DatabasePropertiesGeoReplicationOutput
	ToDatabasePropertiesGeoReplicationOutputWithContext(context.Context) DatabasePropertiesGeoReplicationOutput
}

type DatabasePropertiesGeoReplicationArgs struct {
	GroupNickname   pulumi.StringPtrInput    `pulumi:"groupNickname"`
	LinkedDatabases LinkedDatabaseArrayInput `pulumi:"linkedDatabases"`
}

func (DatabasePropertiesGeoReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePropertiesGeoReplication)(nil)).Elem()
}

func (i DatabasePropertiesGeoReplicationArgs) ToDatabasePropertiesGeoReplicationOutput() DatabasePropertiesGeoReplicationOutput {
	return i.ToDatabasePropertiesGeoReplicationOutputWithContext(context.Background())
}

func (i DatabasePropertiesGeoReplicationArgs) ToDatabasePropertiesGeoReplicationOutputWithContext(ctx context.Context) DatabasePropertiesGeoReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePropertiesGeoReplicationOutput)
}

func (i DatabasePropertiesGeoReplicationArgs) ToDatabasePropertiesGeoReplicationPtrOutput() DatabasePropertiesGeoReplicationPtrOutput {
	return i.ToDatabasePropertiesGeoReplicationPtrOutputWithContext(context.Background())
}

func (i DatabasePropertiesGeoReplicationArgs) ToDatabasePropertiesGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesGeoReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePropertiesGeoReplicationOutput).ToDatabasePropertiesGeoReplicationPtrOutputWithContext(ctx)
}









type DatabasePropertiesGeoReplicationPtrInput interface {
	pulumi.Input

	ToDatabasePropertiesGeoReplicationPtrOutput() DatabasePropertiesGeoReplicationPtrOutput
	ToDatabasePropertiesGeoReplicationPtrOutputWithContext(context.Context) DatabasePropertiesGeoReplicationPtrOutput
}

type databasePropertiesGeoReplicationPtrType DatabasePropertiesGeoReplicationArgs

func DatabasePropertiesGeoReplicationPtr(v *DatabasePropertiesGeoReplicationArgs) DatabasePropertiesGeoReplicationPtrInput {
	return (*databasePropertiesGeoReplicationPtrType)(v)
}

func (*databasePropertiesGeoReplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePropertiesGeoReplication)(nil)).Elem()
}

func (i *databasePropertiesGeoReplicationPtrType) ToDatabasePropertiesGeoReplicationPtrOutput() DatabasePropertiesGeoReplicationPtrOutput {
	return i.ToDatabasePropertiesGeoReplicationPtrOutputWithContext(context.Background())
}

func (i *databasePropertiesGeoReplicationPtrType) ToDatabasePropertiesGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesGeoReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePropertiesGeoReplicationPtrOutput)
}

type DatabasePropertiesGeoReplicationOutput struct{ *pulumi.OutputState }

func (DatabasePropertiesGeoReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePropertiesGeoReplication)(nil)).Elem()
}

func (o DatabasePropertiesGeoReplicationOutput) ToDatabasePropertiesGeoReplicationOutput() DatabasePropertiesGeoReplicationOutput {
	return o
}

func (o DatabasePropertiesGeoReplicationOutput) ToDatabasePropertiesGeoReplicationOutputWithContext(ctx context.Context) DatabasePropertiesGeoReplicationOutput {
	return o
}

func (o DatabasePropertiesGeoReplicationOutput) ToDatabasePropertiesGeoReplicationPtrOutput() DatabasePropertiesGeoReplicationPtrOutput {
	return o.ToDatabasePropertiesGeoReplicationPtrOutputWithContext(context.Background())
}

func (o DatabasePropertiesGeoReplicationOutput) ToDatabasePropertiesGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesGeoReplicationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabasePropertiesGeoReplication) *DatabasePropertiesGeoReplication {
		return &v
	}).(DatabasePropertiesGeoReplicationPtrOutput)
}

func (o DatabasePropertiesGeoReplicationOutput) GroupNickname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePropertiesGeoReplication) *string { return v.GroupNickname }).(pulumi.StringPtrOutput)
}

func (o DatabasePropertiesGeoReplicationOutput) LinkedDatabases() LinkedDatabaseArrayOutput {
	return o.ApplyT(func(v DatabasePropertiesGeoReplication) []LinkedDatabase { return v.LinkedDatabases }).(LinkedDatabaseArrayOutput)
}

type DatabasePropertiesGeoReplicationPtrOutput struct{ *pulumi.OutputState }

func (DatabasePropertiesGeoReplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePropertiesGeoReplication)(nil)).Elem()
}

func (o DatabasePropertiesGeoReplicationPtrOutput) ToDatabasePropertiesGeoReplicationPtrOutput() DatabasePropertiesGeoReplicationPtrOutput {
	return o
}

func (o DatabasePropertiesGeoReplicationPtrOutput) ToDatabasePropertiesGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesGeoReplicationPtrOutput {
	return o
}

func (o DatabasePropertiesGeoReplicationPtrOutput) Elem() DatabasePropertiesGeoReplicationOutput {
	return o.ApplyT(func(v *DatabasePropertiesGeoReplication) DatabasePropertiesGeoReplication {
		if v != nil {
			return *v
		}
		var ret DatabasePropertiesGeoReplication
		return ret
	}).(DatabasePropertiesGeoReplicationOutput)
}

func (o DatabasePropertiesGeoReplicationPtrOutput) GroupNickname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabasePropertiesGeoReplication) *string {
		if v == nil {
			return nil
		}
		return v.GroupNickname
	}).(pulumi.StringPtrOutput)
}

func (o DatabasePropertiesGeoReplicationPtrOutput) LinkedDatabases() LinkedDatabaseArrayOutput {
	return o.ApplyT(func(v *DatabasePropertiesGeoReplication) []LinkedDatabase {
		if v == nil {
			return nil
		}
		return v.LinkedDatabases
	}).(LinkedDatabaseArrayOutput)
}

type DatabasePropertiesResponseGeoReplication struct {
	GroupNickname   *string                  `pulumi:"groupNickname"`
	LinkedDatabases []LinkedDatabaseResponse `pulumi:"linkedDatabases"`
}





type DatabasePropertiesResponseGeoReplicationInput interface {
	pulumi.Input

	ToDatabasePropertiesResponseGeoReplicationOutput() DatabasePropertiesResponseGeoReplicationOutput
	ToDatabasePropertiesResponseGeoReplicationOutputWithContext(context.Context) DatabasePropertiesResponseGeoReplicationOutput
}

type DatabasePropertiesResponseGeoReplicationArgs struct {
	GroupNickname   pulumi.StringPtrInput            `pulumi:"groupNickname"`
	LinkedDatabases LinkedDatabaseResponseArrayInput `pulumi:"linkedDatabases"`
}

func (DatabasePropertiesResponseGeoReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePropertiesResponseGeoReplication)(nil)).Elem()
}

func (i DatabasePropertiesResponseGeoReplicationArgs) ToDatabasePropertiesResponseGeoReplicationOutput() DatabasePropertiesResponseGeoReplicationOutput {
	return i.ToDatabasePropertiesResponseGeoReplicationOutputWithContext(context.Background())
}

func (i DatabasePropertiesResponseGeoReplicationArgs) ToDatabasePropertiesResponseGeoReplicationOutputWithContext(ctx context.Context) DatabasePropertiesResponseGeoReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePropertiesResponseGeoReplicationOutput)
}

func (i DatabasePropertiesResponseGeoReplicationArgs) ToDatabasePropertiesResponseGeoReplicationPtrOutput() DatabasePropertiesResponseGeoReplicationPtrOutput {
	return i.ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(context.Background())
}

func (i DatabasePropertiesResponseGeoReplicationArgs) ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesResponseGeoReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePropertiesResponseGeoReplicationOutput).ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(ctx)
}









type DatabasePropertiesResponseGeoReplicationPtrInput interface {
	pulumi.Input

	ToDatabasePropertiesResponseGeoReplicationPtrOutput() DatabasePropertiesResponseGeoReplicationPtrOutput
	ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(context.Context) DatabasePropertiesResponseGeoReplicationPtrOutput
}

type databasePropertiesResponseGeoReplicationPtrType DatabasePropertiesResponseGeoReplicationArgs

func DatabasePropertiesResponseGeoReplicationPtr(v *DatabasePropertiesResponseGeoReplicationArgs) DatabasePropertiesResponseGeoReplicationPtrInput {
	return (*databasePropertiesResponseGeoReplicationPtrType)(v)
}

func (*databasePropertiesResponseGeoReplicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePropertiesResponseGeoReplication)(nil)).Elem()
}

func (i *databasePropertiesResponseGeoReplicationPtrType) ToDatabasePropertiesResponseGeoReplicationPtrOutput() DatabasePropertiesResponseGeoReplicationPtrOutput {
	return i.ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(context.Background())
}

func (i *databasePropertiesResponseGeoReplicationPtrType) ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesResponseGeoReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePropertiesResponseGeoReplicationPtrOutput)
}

type DatabasePropertiesResponseGeoReplicationOutput struct{ *pulumi.OutputState }

func (DatabasePropertiesResponseGeoReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePropertiesResponseGeoReplication)(nil)).Elem()
}

func (o DatabasePropertiesResponseGeoReplicationOutput) ToDatabasePropertiesResponseGeoReplicationOutput() DatabasePropertiesResponseGeoReplicationOutput {
	return o
}

func (o DatabasePropertiesResponseGeoReplicationOutput) ToDatabasePropertiesResponseGeoReplicationOutputWithContext(ctx context.Context) DatabasePropertiesResponseGeoReplicationOutput {
	return o
}

func (o DatabasePropertiesResponseGeoReplicationOutput) ToDatabasePropertiesResponseGeoReplicationPtrOutput() DatabasePropertiesResponseGeoReplicationPtrOutput {
	return o.ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(context.Background())
}

func (o DatabasePropertiesResponseGeoReplicationOutput) ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesResponseGeoReplicationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabasePropertiesResponseGeoReplication) *DatabasePropertiesResponseGeoReplication {
		return &v
	}).(DatabasePropertiesResponseGeoReplicationPtrOutput)
}

func (o DatabasePropertiesResponseGeoReplicationOutput) GroupNickname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePropertiesResponseGeoReplication) *string { return v.GroupNickname }).(pulumi.StringPtrOutput)
}

func (o DatabasePropertiesResponseGeoReplicationOutput) LinkedDatabases() LinkedDatabaseResponseArrayOutput {
	return o.ApplyT(func(v DatabasePropertiesResponseGeoReplication) []LinkedDatabaseResponse { return v.LinkedDatabases }).(LinkedDatabaseResponseArrayOutput)
}

type DatabasePropertiesResponseGeoReplicationPtrOutput struct{ *pulumi.OutputState }

func (DatabasePropertiesResponseGeoReplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePropertiesResponseGeoReplication)(nil)).Elem()
}

func (o DatabasePropertiesResponseGeoReplicationPtrOutput) ToDatabasePropertiesResponseGeoReplicationPtrOutput() DatabasePropertiesResponseGeoReplicationPtrOutput {
	return o
}

func (o DatabasePropertiesResponseGeoReplicationPtrOutput) ToDatabasePropertiesResponseGeoReplicationPtrOutputWithContext(ctx context.Context) DatabasePropertiesResponseGeoReplicationPtrOutput {
	return o
}

func (o DatabasePropertiesResponseGeoReplicationPtrOutput) Elem() DatabasePropertiesResponseGeoReplicationOutput {
	return o.ApplyT(func(v *DatabasePropertiesResponseGeoReplication) DatabasePropertiesResponseGeoReplication {
		if v != nil {
			return *v
		}
		var ret DatabasePropertiesResponseGeoReplication
		return ret
	}).(DatabasePropertiesResponseGeoReplicationOutput)
}

func (o DatabasePropertiesResponseGeoReplicationPtrOutput) GroupNickname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabasePropertiesResponseGeoReplication) *string {
		if v == nil {
			return nil
		}
		return v.GroupNickname
	}).(pulumi.StringPtrOutput)
}

func (o DatabasePropertiesResponseGeoReplicationPtrOutput) LinkedDatabases() LinkedDatabaseResponseArrayOutput {
	return o.ApplyT(func(v *DatabasePropertiesResponseGeoReplication) []LinkedDatabaseResponse {
		if v == nil {
			return nil
		}
		return v.LinkedDatabases
	}).(LinkedDatabaseResponseArrayOutput)
}

type EnterpriseSku struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type EnterpriseSkuInput interface {
	pulumi.Input

	ToEnterpriseSkuOutput() EnterpriseSkuOutput
	ToEnterpriseSkuOutputWithContext(context.Context) EnterpriseSkuOutput
}

type EnterpriseSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (EnterpriseSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseSku)(nil)).Elem()
}

func (i EnterpriseSkuArgs) ToEnterpriseSkuOutput() EnterpriseSkuOutput {
	return i.ToEnterpriseSkuOutputWithContext(context.Background())
}

func (i EnterpriseSkuArgs) ToEnterpriseSkuOutputWithContext(ctx context.Context) EnterpriseSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSkuOutput)
}

func (i EnterpriseSkuArgs) ToEnterpriseSkuPtrOutput() EnterpriseSkuPtrOutput {
	return i.ToEnterpriseSkuPtrOutputWithContext(context.Background())
}

func (i EnterpriseSkuArgs) ToEnterpriseSkuPtrOutputWithContext(ctx context.Context) EnterpriseSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSkuOutput).ToEnterpriseSkuPtrOutputWithContext(ctx)
}









type EnterpriseSkuPtrInput interface {
	pulumi.Input

	ToEnterpriseSkuPtrOutput() EnterpriseSkuPtrOutput
	ToEnterpriseSkuPtrOutputWithContext(context.Context) EnterpriseSkuPtrOutput
}

type enterpriseSkuPtrType EnterpriseSkuArgs

func EnterpriseSkuPtr(v *EnterpriseSkuArgs) EnterpriseSkuPtrInput {
	return (*enterpriseSkuPtrType)(v)
}

func (*enterpriseSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSku)(nil)).Elem()
}

func (i *enterpriseSkuPtrType) ToEnterpriseSkuPtrOutput() EnterpriseSkuPtrOutput {
	return i.ToEnterpriseSkuPtrOutputWithContext(context.Background())
}

func (i *enterpriseSkuPtrType) ToEnterpriseSkuPtrOutputWithContext(ctx context.Context) EnterpriseSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSkuPtrOutput)
}

type EnterpriseSkuOutput struct{ *pulumi.OutputState }

func (EnterpriseSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseSku)(nil)).Elem()
}

func (o EnterpriseSkuOutput) ToEnterpriseSkuOutput() EnterpriseSkuOutput {
	return o
}

func (o EnterpriseSkuOutput) ToEnterpriseSkuOutputWithContext(ctx context.Context) EnterpriseSkuOutput {
	return o
}

func (o EnterpriseSkuOutput) ToEnterpriseSkuPtrOutput() EnterpriseSkuPtrOutput {
	return o.ToEnterpriseSkuPtrOutputWithContext(context.Background())
}

func (o EnterpriseSkuOutput) ToEnterpriseSkuPtrOutputWithContext(ctx context.Context) EnterpriseSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterpriseSku) *EnterpriseSku {
		return &v
	}).(EnterpriseSkuPtrOutput)
}

func (o EnterpriseSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EnterpriseSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o EnterpriseSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseSku) string { return v.Name }).(pulumi.StringOutput)
}

type EnterpriseSkuPtrOutput struct{ *pulumi.OutputState }

func (EnterpriseSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSku)(nil)).Elem()
}

func (o EnterpriseSkuPtrOutput) ToEnterpriseSkuPtrOutput() EnterpriseSkuPtrOutput {
	return o
}

func (o EnterpriseSkuPtrOutput) ToEnterpriseSkuPtrOutputWithContext(ctx context.Context) EnterpriseSkuPtrOutput {
	return o
}

func (o EnterpriseSkuPtrOutput) Elem() EnterpriseSkuOutput {
	return o.ApplyT(func(v *EnterpriseSku) EnterpriseSku {
		if v != nil {
			return *v
		}
		var ret EnterpriseSku
		return ret
	}).(EnterpriseSkuOutput)
}

func (o EnterpriseSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EnterpriseSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o EnterpriseSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type EnterpriseSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type EnterpriseSkuResponseInput interface {
	pulumi.Input

	ToEnterpriseSkuResponseOutput() EnterpriseSkuResponseOutput
	ToEnterpriseSkuResponseOutputWithContext(context.Context) EnterpriseSkuResponseOutput
}

type EnterpriseSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (EnterpriseSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseSkuResponse)(nil)).Elem()
}

func (i EnterpriseSkuResponseArgs) ToEnterpriseSkuResponseOutput() EnterpriseSkuResponseOutput {
	return i.ToEnterpriseSkuResponseOutputWithContext(context.Background())
}

func (i EnterpriseSkuResponseArgs) ToEnterpriseSkuResponseOutputWithContext(ctx context.Context) EnterpriseSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSkuResponseOutput)
}

func (i EnterpriseSkuResponseArgs) ToEnterpriseSkuResponsePtrOutput() EnterpriseSkuResponsePtrOutput {
	return i.ToEnterpriseSkuResponsePtrOutputWithContext(context.Background())
}

func (i EnterpriseSkuResponseArgs) ToEnterpriseSkuResponsePtrOutputWithContext(ctx context.Context) EnterpriseSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSkuResponseOutput).ToEnterpriseSkuResponsePtrOutputWithContext(ctx)
}









type EnterpriseSkuResponsePtrInput interface {
	pulumi.Input

	ToEnterpriseSkuResponsePtrOutput() EnterpriseSkuResponsePtrOutput
	ToEnterpriseSkuResponsePtrOutputWithContext(context.Context) EnterpriseSkuResponsePtrOutput
}

type enterpriseSkuResponsePtrType EnterpriseSkuResponseArgs

func EnterpriseSkuResponsePtr(v *EnterpriseSkuResponseArgs) EnterpriseSkuResponsePtrInput {
	return (*enterpriseSkuResponsePtrType)(v)
}

func (*enterpriseSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSkuResponse)(nil)).Elem()
}

func (i *enterpriseSkuResponsePtrType) ToEnterpriseSkuResponsePtrOutput() EnterpriseSkuResponsePtrOutput {
	return i.ToEnterpriseSkuResponsePtrOutputWithContext(context.Background())
}

func (i *enterpriseSkuResponsePtrType) ToEnterpriseSkuResponsePtrOutputWithContext(ctx context.Context) EnterpriseSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSkuResponsePtrOutput)
}

type EnterpriseSkuResponseOutput struct{ *pulumi.OutputState }

func (EnterpriseSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseSkuResponse)(nil)).Elem()
}

func (o EnterpriseSkuResponseOutput) ToEnterpriseSkuResponseOutput() EnterpriseSkuResponseOutput {
	return o
}

func (o EnterpriseSkuResponseOutput) ToEnterpriseSkuResponseOutputWithContext(ctx context.Context) EnterpriseSkuResponseOutput {
	return o
}

func (o EnterpriseSkuResponseOutput) ToEnterpriseSkuResponsePtrOutput() EnterpriseSkuResponsePtrOutput {
	return o.ToEnterpriseSkuResponsePtrOutputWithContext(context.Background())
}

func (o EnterpriseSkuResponseOutput) ToEnterpriseSkuResponsePtrOutputWithContext(ctx context.Context) EnterpriseSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterpriseSkuResponse) *EnterpriseSkuResponse {
		return &v
	}).(EnterpriseSkuResponsePtrOutput)
}

func (o EnterpriseSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EnterpriseSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o EnterpriseSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type EnterpriseSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (EnterpriseSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSkuResponse)(nil)).Elem()
}

func (o EnterpriseSkuResponsePtrOutput) ToEnterpriseSkuResponsePtrOutput() EnterpriseSkuResponsePtrOutput {
	return o
}

func (o EnterpriseSkuResponsePtrOutput) ToEnterpriseSkuResponsePtrOutputWithContext(ctx context.Context) EnterpriseSkuResponsePtrOutput {
	return o
}

func (o EnterpriseSkuResponsePtrOutput) Elem() EnterpriseSkuResponseOutput {
	return o.ApplyT(func(v *EnterpriseSkuResponse) EnterpriseSkuResponse {
		if v != nil {
			return *v
		}
		var ret EnterpriseSkuResponse
		return ret
	}).(EnterpriseSkuResponseOutput)
}

func (o EnterpriseSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EnterpriseSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o EnterpriseSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type LinkedDatabase struct {
	Id *string `pulumi:"id"`
}





type LinkedDatabaseInput interface {
	pulumi.Input

	ToLinkedDatabaseOutput() LinkedDatabaseOutput
	ToLinkedDatabaseOutputWithContext(context.Context) LinkedDatabaseOutput
}

type LinkedDatabaseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (LinkedDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedDatabase)(nil)).Elem()
}

func (i LinkedDatabaseArgs) ToLinkedDatabaseOutput() LinkedDatabaseOutput {
	return i.ToLinkedDatabaseOutputWithContext(context.Background())
}

func (i LinkedDatabaseArgs) ToLinkedDatabaseOutputWithContext(ctx context.Context) LinkedDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedDatabaseOutput)
}





type LinkedDatabaseArrayInput interface {
	pulumi.Input

	ToLinkedDatabaseArrayOutput() LinkedDatabaseArrayOutput
	ToLinkedDatabaseArrayOutputWithContext(context.Context) LinkedDatabaseArrayOutput
}

type LinkedDatabaseArray []LinkedDatabaseInput

func (LinkedDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedDatabase)(nil)).Elem()
}

func (i LinkedDatabaseArray) ToLinkedDatabaseArrayOutput() LinkedDatabaseArrayOutput {
	return i.ToLinkedDatabaseArrayOutputWithContext(context.Background())
}

func (i LinkedDatabaseArray) ToLinkedDatabaseArrayOutputWithContext(ctx context.Context) LinkedDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedDatabaseArrayOutput)
}

type LinkedDatabaseOutput struct{ *pulumi.OutputState }

func (LinkedDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedDatabase)(nil)).Elem()
}

func (o LinkedDatabaseOutput) ToLinkedDatabaseOutput() LinkedDatabaseOutput {
	return o
}

func (o LinkedDatabaseOutput) ToLinkedDatabaseOutputWithContext(ctx context.Context) LinkedDatabaseOutput {
	return o
}

func (o LinkedDatabaseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedDatabase) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type LinkedDatabaseArrayOutput struct{ *pulumi.OutputState }

func (LinkedDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedDatabase)(nil)).Elem()
}

func (o LinkedDatabaseArrayOutput) ToLinkedDatabaseArrayOutput() LinkedDatabaseArrayOutput {
	return o
}

func (o LinkedDatabaseArrayOutput) ToLinkedDatabaseArrayOutputWithContext(ctx context.Context) LinkedDatabaseArrayOutput {
	return o
}

func (o LinkedDatabaseArrayOutput) Index(i pulumi.IntInput) LinkedDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedDatabase {
		return vs[0].([]LinkedDatabase)[vs[1].(int)]
	}).(LinkedDatabaseOutput)
}

type LinkedDatabaseResponse struct {
	Id    *string `pulumi:"id"`
	State string  `pulumi:"state"`
}





type LinkedDatabaseResponseInput interface {
	pulumi.Input

	ToLinkedDatabaseResponseOutput() LinkedDatabaseResponseOutput
	ToLinkedDatabaseResponseOutputWithContext(context.Context) LinkedDatabaseResponseOutput
}

type LinkedDatabaseResponseArgs struct {
	Id    pulumi.StringPtrInput `pulumi:"id"`
	State pulumi.StringInput    `pulumi:"state"`
}

func (LinkedDatabaseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedDatabaseResponse)(nil)).Elem()
}

func (i LinkedDatabaseResponseArgs) ToLinkedDatabaseResponseOutput() LinkedDatabaseResponseOutput {
	return i.ToLinkedDatabaseResponseOutputWithContext(context.Background())
}

func (i LinkedDatabaseResponseArgs) ToLinkedDatabaseResponseOutputWithContext(ctx context.Context) LinkedDatabaseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedDatabaseResponseOutput)
}





type LinkedDatabaseResponseArrayInput interface {
	pulumi.Input

	ToLinkedDatabaseResponseArrayOutput() LinkedDatabaseResponseArrayOutput
	ToLinkedDatabaseResponseArrayOutputWithContext(context.Context) LinkedDatabaseResponseArrayOutput
}

type LinkedDatabaseResponseArray []LinkedDatabaseResponseInput

func (LinkedDatabaseResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedDatabaseResponse)(nil)).Elem()
}

func (i LinkedDatabaseResponseArray) ToLinkedDatabaseResponseArrayOutput() LinkedDatabaseResponseArrayOutput {
	return i.ToLinkedDatabaseResponseArrayOutputWithContext(context.Background())
}

func (i LinkedDatabaseResponseArray) ToLinkedDatabaseResponseArrayOutputWithContext(ctx context.Context) LinkedDatabaseResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedDatabaseResponseArrayOutput)
}

type LinkedDatabaseResponseOutput struct{ *pulumi.OutputState }

func (LinkedDatabaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedDatabaseResponse)(nil)).Elem()
}

func (o LinkedDatabaseResponseOutput) ToLinkedDatabaseResponseOutput() LinkedDatabaseResponseOutput {
	return o
}

func (o LinkedDatabaseResponseOutput) ToLinkedDatabaseResponseOutputWithContext(ctx context.Context) LinkedDatabaseResponseOutput {
	return o
}

func (o LinkedDatabaseResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedDatabaseResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LinkedDatabaseResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedDatabaseResponse) string { return v.State }).(pulumi.StringOutput)
}

type LinkedDatabaseResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkedDatabaseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedDatabaseResponse)(nil)).Elem()
}

func (o LinkedDatabaseResponseArrayOutput) ToLinkedDatabaseResponseArrayOutput() LinkedDatabaseResponseArrayOutput {
	return o
}

func (o LinkedDatabaseResponseArrayOutput) ToLinkedDatabaseResponseArrayOutputWithContext(ctx context.Context) LinkedDatabaseResponseArrayOutput {
	return o
}

func (o LinkedDatabaseResponseArrayOutput) Index(i pulumi.IntInput) LinkedDatabaseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedDatabaseResponse {
		return vs[0].([]LinkedDatabaseResponse)[vs[1].(int)]
	}).(LinkedDatabaseResponseOutput)
}

type Module struct {
	Args *string `pulumi:"args"`
	Name string  `pulumi:"name"`
}





type ModuleInput interface {
	pulumi.Input

	ToModuleOutput() ModuleOutput
	ToModuleOutputWithContext(context.Context) ModuleOutput
}

type ModuleArgs struct {
	Args pulumi.StringPtrInput `pulumi:"args"`
	Name pulumi.StringInput    `pulumi:"name"`
}

func (ModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Module)(nil)).Elem()
}

func (i ModuleArgs) ToModuleOutput() ModuleOutput {
	return i.ToModuleOutputWithContext(context.Background())
}

func (i ModuleArgs) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleOutput)
}





type ModuleArrayInput interface {
	pulumi.Input

	ToModuleArrayOutput() ModuleArrayOutput
	ToModuleArrayOutputWithContext(context.Context) ModuleArrayOutput
}

type ModuleArray []ModuleInput

func (ModuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Module)(nil)).Elem()
}

func (i ModuleArray) ToModuleArrayOutput() ModuleArrayOutput {
	return i.ToModuleArrayOutputWithContext(context.Background())
}

func (i ModuleArray) ToModuleArrayOutputWithContext(ctx context.Context) ModuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleArrayOutput)
}

type ModuleOutput struct{ *pulumi.OutputState }

func (ModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Module)(nil)).Elem()
}

func (o ModuleOutput) ToModuleOutput() ModuleOutput {
	return o
}

func (o ModuleOutput) ToModuleOutputWithContext(ctx context.Context) ModuleOutput {
	return o
}

func (o ModuleOutput) Args() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Module) *string { return v.Args }).(pulumi.StringPtrOutput)
}

func (o ModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Module) string { return v.Name }).(pulumi.StringOutput)
}

type ModuleArrayOutput struct{ *pulumi.OutputState }

func (ModuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Module)(nil)).Elem()
}

func (o ModuleArrayOutput) ToModuleArrayOutput() ModuleArrayOutput {
	return o
}

func (o ModuleArrayOutput) ToModuleArrayOutputWithContext(ctx context.Context) ModuleArrayOutput {
	return o
}

func (o ModuleArrayOutput) Index(i pulumi.IntInput) ModuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Module {
		return vs[0].([]Module)[vs[1].(int)]
	}).(ModuleOutput)
}

type ModuleResponse struct {
	Args    *string `pulumi:"args"`
	Name    string  `pulumi:"name"`
	Version string  `pulumi:"version"`
}





type ModuleResponseInput interface {
	pulumi.Input

	ToModuleResponseOutput() ModuleResponseOutput
	ToModuleResponseOutputWithContext(context.Context) ModuleResponseOutput
}

type ModuleResponseArgs struct {
	Args    pulumi.StringPtrInput `pulumi:"args"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Version pulumi.StringInput    `pulumi:"version"`
}

func (ModuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleResponse)(nil)).Elem()
}

func (i ModuleResponseArgs) ToModuleResponseOutput() ModuleResponseOutput {
	return i.ToModuleResponseOutputWithContext(context.Background())
}

func (i ModuleResponseArgs) ToModuleResponseOutputWithContext(ctx context.Context) ModuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleResponseOutput)
}





type ModuleResponseArrayInput interface {
	pulumi.Input

	ToModuleResponseArrayOutput() ModuleResponseArrayOutput
	ToModuleResponseArrayOutputWithContext(context.Context) ModuleResponseArrayOutput
}

type ModuleResponseArray []ModuleResponseInput

func (ModuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModuleResponse)(nil)).Elem()
}

func (i ModuleResponseArray) ToModuleResponseArrayOutput() ModuleResponseArrayOutput {
	return i.ToModuleResponseArrayOutputWithContext(context.Background())
}

func (i ModuleResponseArray) ToModuleResponseArrayOutputWithContext(ctx context.Context) ModuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleResponseArrayOutput)
}

type ModuleResponseOutput struct{ *pulumi.OutputState }

func (ModuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleResponse)(nil)).Elem()
}

func (o ModuleResponseOutput) ToModuleResponseOutput() ModuleResponseOutput {
	return o
}

func (o ModuleResponseOutput) ToModuleResponseOutputWithContext(ctx context.Context) ModuleResponseOutput {
	return o
}

func (o ModuleResponseOutput) Args() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleResponse) *string { return v.Args }).(pulumi.StringPtrOutput)
}

func (o ModuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ModuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ModuleResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ModuleResponse) string { return v.Version }).(pulumi.StringOutput)
}

type ModuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ModuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModuleResponse)(nil)).Elem()
}

func (o ModuleResponseArrayOutput) ToModuleResponseArrayOutput() ModuleResponseArrayOutput {
	return o
}

func (o ModuleResponseArrayOutput) ToModuleResponseArrayOutputWithContext(ctx context.Context) ModuleResponseArrayOutput {
	return o
}

func (o ModuleResponseArrayOutput) Index(i pulumi.IntInput) ModuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ModuleResponse {
		return vs[0].([]ModuleResponse)[vs[1].(int)]
	}).(ModuleResponseOutput)
}

type Persistence struct {
	AofEnabled   *bool   `pulumi:"aofEnabled"`
	AofFrequency *string `pulumi:"aofFrequency"`
	RdbEnabled   *bool   `pulumi:"rdbEnabled"`
	RdbFrequency *string `pulumi:"rdbFrequency"`
}





type PersistenceInput interface {
	pulumi.Input

	ToPersistenceOutput() PersistenceOutput
	ToPersistenceOutputWithContext(context.Context) PersistenceOutput
}

type PersistenceArgs struct {
	AofEnabled   pulumi.BoolPtrInput   `pulumi:"aofEnabled"`
	AofFrequency pulumi.StringPtrInput `pulumi:"aofFrequency"`
	RdbEnabled   pulumi.BoolPtrInput   `pulumi:"rdbEnabled"`
	RdbFrequency pulumi.StringPtrInput `pulumi:"rdbFrequency"`
}

func (PersistenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Persistence)(nil)).Elem()
}

func (i PersistenceArgs) ToPersistenceOutput() PersistenceOutput {
	return i.ToPersistenceOutputWithContext(context.Background())
}

func (i PersistenceArgs) ToPersistenceOutputWithContext(ctx context.Context) PersistenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceOutput)
}

func (i PersistenceArgs) ToPersistencePtrOutput() PersistencePtrOutput {
	return i.ToPersistencePtrOutputWithContext(context.Background())
}

func (i PersistenceArgs) ToPersistencePtrOutputWithContext(ctx context.Context) PersistencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceOutput).ToPersistencePtrOutputWithContext(ctx)
}









type PersistencePtrInput interface {
	pulumi.Input

	ToPersistencePtrOutput() PersistencePtrOutput
	ToPersistencePtrOutputWithContext(context.Context) PersistencePtrOutput
}

type persistencePtrType PersistenceArgs

func PersistencePtr(v *PersistenceArgs) PersistencePtrInput {
	return (*persistencePtrType)(v)
}

func (*persistencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Persistence)(nil)).Elem()
}

func (i *persistencePtrType) ToPersistencePtrOutput() PersistencePtrOutput {
	return i.ToPersistencePtrOutputWithContext(context.Background())
}

func (i *persistencePtrType) ToPersistencePtrOutputWithContext(ctx context.Context) PersistencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistencePtrOutput)
}

type PersistenceOutput struct{ *pulumi.OutputState }

func (PersistenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Persistence)(nil)).Elem()
}

func (o PersistenceOutput) ToPersistenceOutput() PersistenceOutput {
	return o
}

func (o PersistenceOutput) ToPersistenceOutputWithContext(ctx context.Context) PersistenceOutput {
	return o
}

func (o PersistenceOutput) ToPersistencePtrOutput() PersistencePtrOutput {
	return o.ToPersistencePtrOutputWithContext(context.Background())
}

func (o PersistenceOutput) ToPersistencePtrOutputWithContext(ctx context.Context) PersistencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Persistence) *Persistence {
		return &v
	}).(PersistencePtrOutput)
}

func (o PersistenceOutput) AofEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Persistence) *bool { return v.AofEnabled }).(pulumi.BoolPtrOutput)
}

func (o PersistenceOutput) AofFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Persistence) *string { return v.AofFrequency }).(pulumi.StringPtrOutput)
}

func (o PersistenceOutput) RdbEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Persistence) *bool { return v.RdbEnabled }).(pulumi.BoolPtrOutput)
}

func (o PersistenceOutput) RdbFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Persistence) *string { return v.RdbFrequency }).(pulumi.StringPtrOutput)
}

type PersistencePtrOutput struct{ *pulumi.OutputState }

func (PersistencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Persistence)(nil)).Elem()
}

func (o PersistencePtrOutput) ToPersistencePtrOutput() PersistencePtrOutput {
	return o
}

func (o PersistencePtrOutput) ToPersistencePtrOutputWithContext(ctx context.Context) PersistencePtrOutput {
	return o
}

func (o PersistencePtrOutput) Elem() PersistenceOutput {
	return o.ApplyT(func(v *Persistence) Persistence {
		if v != nil {
			return *v
		}
		var ret Persistence
		return ret
	}).(PersistenceOutput)
}

func (o PersistencePtrOutput) AofEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Persistence) *bool {
		if v == nil {
			return nil
		}
		return v.AofEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PersistencePtrOutput) AofFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Persistence) *string {
		if v == nil {
			return nil
		}
		return v.AofFrequency
	}).(pulumi.StringPtrOutput)
}

func (o PersistencePtrOutput) RdbEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Persistence) *bool {
		if v == nil {
			return nil
		}
		return v.RdbEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PersistencePtrOutput) RdbFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Persistence) *string {
		if v == nil {
			return nil
		}
		return v.RdbFrequency
	}).(pulumi.StringPtrOutput)
}

type PersistenceResponse struct {
	AofEnabled   *bool   `pulumi:"aofEnabled"`
	AofFrequency *string `pulumi:"aofFrequency"`
	RdbEnabled   *bool   `pulumi:"rdbEnabled"`
	RdbFrequency *string `pulumi:"rdbFrequency"`
}





type PersistenceResponseInput interface {
	pulumi.Input

	ToPersistenceResponseOutput() PersistenceResponseOutput
	ToPersistenceResponseOutputWithContext(context.Context) PersistenceResponseOutput
}

type PersistenceResponseArgs struct {
	AofEnabled   pulumi.BoolPtrInput   `pulumi:"aofEnabled"`
	AofFrequency pulumi.StringPtrInput `pulumi:"aofFrequency"`
	RdbEnabled   pulumi.BoolPtrInput   `pulumi:"rdbEnabled"`
	RdbFrequency pulumi.StringPtrInput `pulumi:"rdbFrequency"`
}

func (PersistenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistenceResponse)(nil)).Elem()
}

func (i PersistenceResponseArgs) ToPersistenceResponseOutput() PersistenceResponseOutput {
	return i.ToPersistenceResponseOutputWithContext(context.Background())
}

func (i PersistenceResponseArgs) ToPersistenceResponseOutputWithContext(ctx context.Context) PersistenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceResponseOutput)
}

func (i PersistenceResponseArgs) ToPersistenceResponsePtrOutput() PersistenceResponsePtrOutput {
	return i.ToPersistenceResponsePtrOutputWithContext(context.Background())
}

func (i PersistenceResponseArgs) ToPersistenceResponsePtrOutputWithContext(ctx context.Context) PersistenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceResponseOutput).ToPersistenceResponsePtrOutputWithContext(ctx)
}









type PersistenceResponsePtrInput interface {
	pulumi.Input

	ToPersistenceResponsePtrOutput() PersistenceResponsePtrOutput
	ToPersistenceResponsePtrOutputWithContext(context.Context) PersistenceResponsePtrOutput
}

type persistenceResponsePtrType PersistenceResponseArgs

func PersistenceResponsePtr(v *PersistenceResponseArgs) PersistenceResponsePtrInput {
	return (*persistenceResponsePtrType)(v)
}

func (*persistenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceResponse)(nil)).Elem()
}

func (i *persistenceResponsePtrType) ToPersistenceResponsePtrOutput() PersistenceResponsePtrOutput {
	return i.ToPersistenceResponsePtrOutputWithContext(context.Background())
}

func (i *persistenceResponsePtrType) ToPersistenceResponsePtrOutputWithContext(ctx context.Context) PersistenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceResponsePtrOutput)
}

type PersistenceResponseOutput struct{ *pulumi.OutputState }

func (PersistenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistenceResponse)(nil)).Elem()
}

func (o PersistenceResponseOutput) ToPersistenceResponseOutput() PersistenceResponseOutput {
	return o
}

func (o PersistenceResponseOutput) ToPersistenceResponseOutputWithContext(ctx context.Context) PersistenceResponseOutput {
	return o
}

func (o PersistenceResponseOutput) ToPersistenceResponsePtrOutput() PersistenceResponsePtrOutput {
	return o.ToPersistenceResponsePtrOutputWithContext(context.Background())
}

func (o PersistenceResponseOutput) ToPersistenceResponsePtrOutputWithContext(ctx context.Context) PersistenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersistenceResponse) *PersistenceResponse {
		return &v
	}).(PersistenceResponsePtrOutput)
}

func (o PersistenceResponseOutput) AofEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PersistenceResponse) *bool { return v.AofEnabled }).(pulumi.BoolPtrOutput)
}

func (o PersistenceResponseOutput) AofFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PersistenceResponse) *string { return v.AofFrequency }).(pulumi.StringPtrOutput)
}

func (o PersistenceResponseOutput) RdbEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PersistenceResponse) *bool { return v.RdbEnabled }).(pulumi.BoolPtrOutput)
}

func (o PersistenceResponseOutput) RdbFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PersistenceResponse) *string { return v.RdbFrequency }).(pulumi.StringPtrOutput)
}

type PersistenceResponsePtrOutput struct{ *pulumi.OutputState }

func (PersistenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceResponse)(nil)).Elem()
}

func (o PersistenceResponsePtrOutput) ToPersistenceResponsePtrOutput() PersistenceResponsePtrOutput {
	return o
}

func (o PersistenceResponsePtrOutput) ToPersistenceResponsePtrOutputWithContext(ctx context.Context) PersistenceResponsePtrOutput {
	return o
}

func (o PersistenceResponsePtrOutput) Elem() PersistenceResponseOutput {
	return o.ApplyT(func(v *PersistenceResponse) PersistenceResponse {
		if v != nil {
			return *v
		}
		var ret PersistenceResponse
		return ret
	}).(PersistenceResponseOutput)
}

func (o PersistenceResponsePtrOutput) AofEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PersistenceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AofEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PersistenceResponsePtrOutput) AofFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.AofFrequency
	}).(pulumi.StringPtrOutput)
}

func (o PersistenceResponsePtrOutput) RdbEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PersistenceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RdbEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PersistenceResponsePtrOutput) RdbFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.RdbFrequency
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePropertiesGeoReplicationInput)(nil)).Elem(), DatabasePropertiesGeoReplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePropertiesGeoReplicationPtrInput)(nil)).Elem(), DatabasePropertiesGeoReplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePropertiesResponseGeoReplicationInput)(nil)).Elem(), DatabasePropertiesResponseGeoReplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePropertiesResponseGeoReplicationPtrInput)(nil)).Elem(), DatabasePropertiesResponseGeoReplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSkuInput)(nil)).Elem(), EnterpriseSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSkuPtrInput)(nil)).Elem(), EnterpriseSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSkuResponseInput)(nil)).Elem(), EnterpriseSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSkuResponsePtrInput)(nil)).Elem(), EnterpriseSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedDatabaseInput)(nil)).Elem(), LinkedDatabaseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedDatabaseArrayInput)(nil)).Elem(), LinkedDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedDatabaseResponseInput)(nil)).Elem(), LinkedDatabaseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedDatabaseResponseArrayInput)(nil)).Elem(), LinkedDatabaseResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleInput)(nil)).Elem(), ModuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleArrayInput)(nil)).Elem(), ModuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleResponseInput)(nil)).Elem(), ModuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleResponseArrayInput)(nil)).Elem(), ModuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceInput)(nil)).Elem(), PersistenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistencePtrInput)(nil)).Elem(), PersistenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceResponseInput)(nil)).Elem(), PersistenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceResponsePtrInput)(nil)).Elem(), PersistenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterOutputType(DatabasePropertiesGeoReplicationOutput{})
	pulumi.RegisterOutputType(DatabasePropertiesGeoReplicationPtrOutput{})
	pulumi.RegisterOutputType(DatabasePropertiesResponseGeoReplicationOutput{})
	pulumi.RegisterOutputType(DatabasePropertiesResponseGeoReplicationPtrOutput{})
	pulumi.RegisterOutputType(EnterpriseSkuOutput{})
	pulumi.RegisterOutputType(EnterpriseSkuPtrOutput{})
	pulumi.RegisterOutputType(EnterpriseSkuResponseOutput{})
	pulumi.RegisterOutputType(EnterpriseSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(LinkedDatabaseOutput{})
	pulumi.RegisterOutputType(LinkedDatabaseArrayOutput{})
	pulumi.RegisterOutputType(LinkedDatabaseResponseOutput{})
	pulumi.RegisterOutputType(LinkedDatabaseResponseArrayOutput{})
	pulumi.RegisterOutputType(ModuleOutput{})
	pulumi.RegisterOutputType(ModuleArrayOutput{})
	pulumi.RegisterOutputType(ModuleResponseOutput{})
	pulumi.RegisterOutputType(ModuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PersistenceOutput{})
	pulumi.RegisterOutputType(PersistencePtrOutput{})
	pulumi.RegisterOutputType(PersistenceResponseOutput{})
	pulumi.RegisterOutputType(PersistenceResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
}
