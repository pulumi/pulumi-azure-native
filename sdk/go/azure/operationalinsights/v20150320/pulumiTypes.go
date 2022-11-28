


package v20150320

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccount struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Key pulumi.StringInput `pulumi:"key"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Key }).(pulumi.StringOutput)
}

type StorageAccountResponse struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Key }).(pulumi.StringOutput)
}

type StorageInsightStatusResponse struct {
	Description *string `pulumi:"description"`
	State       string  `pulumi:"state"`
}

type StorageInsightStatusResponseOutput struct{ *pulumi.OutputState }

func (StorageInsightStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageInsightStatusResponse)(nil)).Elem()
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutput() StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutputWithContext(ctx context.Context) StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StorageInsightStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type Tag struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(context.Context) TagOutput
}

type TagArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (i TagArgs) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i TagArgs) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}





type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
	ToTagArrayOutputWithContext(context.Context) TagArrayOutput
}

type TagArray []TagInput

func (TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (i TagArray) ToTagArrayOutput() TagArrayOutput {
	return i.ToTagArrayOutputWithContext(context.Background())
}

func (i TagArray) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagArrayOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func (o TagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Tag) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v Tag) string { return v.Value }).(pulumi.StringOutput)
}

type TagArrayOutput struct{ *pulumi.OutputState }

func (TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (o TagArrayOutput) ToTagArrayOutput() TagArrayOutput {
	return o
}

func (o TagArrayOutput) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return o
}

func (o TagArrayOutput) Index(i pulumi.IntInput) TagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Tag {
		return vs[0].([]Tag)[vs[1].(int)]
	}).(TagOutput)
}

type TagResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type TagResponseOutput struct{ *pulumi.OutputState }

func (TagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagResponse)(nil)).Elem()
}

func (o TagResponseOutput) ToTagResponseOutput() TagResponseOutput {
	return o
}

func (o TagResponseOutput) ToTagResponseOutputWithContext(ctx context.Context) TagResponseOutput {
	return o
}

func (o TagResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TagResponseArrayOutput struct{ *pulumi.OutputState }

func (TagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagResponse)(nil)).Elem()
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutput() TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutputWithContext(ctx context.Context) TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) Index(i pulumi.IntInput) TagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagResponse {
		return vs[0].([]TagResponse)[vs[1].(int)]
	}).(TagResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageInsightStatusResponseOutput{})
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagResponseOutput{})
	pulumi.RegisterOutputType(TagResponseArrayOutput{})
}
