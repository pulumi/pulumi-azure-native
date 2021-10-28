


package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssetItem struct {
	Id           *string                `pulumi:"id"`
	InputPorts   map[string]InputPort   `pulumi:"inputPorts"`
	LocationInfo BlobLocation           `pulumi:"locationInfo"`
	Metadata     map[string]string      `pulumi:"metadata"`
	Name         string                 `pulumi:"name"`
	OutputPorts  map[string]OutputPort  `pulumi:"outputPorts"`
	Parameters   []ModuleAssetParameter `pulumi:"parameters"`
	Type         string                 `pulumi:"type"`
}





type AssetItemInput interface {
	pulumi.Input

	ToAssetItemOutput() AssetItemOutput
	ToAssetItemOutputWithContext(context.Context) AssetItemOutput
}

type AssetItemArgs struct {
	Id           pulumi.StringPtrInput          `pulumi:"id"`
	InputPorts   InputPortMapInput              `pulumi:"inputPorts"`
	LocationInfo BlobLocationInput              `pulumi:"locationInfo"`
	Metadata     pulumi.StringMapInput          `pulumi:"metadata"`
	Name         pulumi.StringInput             `pulumi:"name"`
	OutputPorts  OutputPortMapInput             `pulumi:"outputPorts"`
	Parameters   ModuleAssetParameterArrayInput `pulumi:"parameters"`
	Type         pulumi.StringInput             `pulumi:"type"`
}

func (AssetItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetItem)(nil)).Elem()
}

func (i AssetItemArgs) ToAssetItemOutput() AssetItemOutput {
	return i.ToAssetItemOutputWithContext(context.Background())
}

func (i AssetItemArgs) ToAssetItemOutputWithContext(ctx context.Context) AssetItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetItemOutput)
}





type AssetItemMapInput interface {
	pulumi.Input

	ToAssetItemMapOutput() AssetItemMapOutput
	ToAssetItemMapOutputWithContext(context.Context) AssetItemMapOutput
}

type AssetItemMap map[string]AssetItemInput

func (AssetItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AssetItem)(nil)).Elem()
}

func (i AssetItemMap) ToAssetItemMapOutput() AssetItemMapOutput {
	return i.ToAssetItemMapOutputWithContext(context.Background())
}

func (i AssetItemMap) ToAssetItemMapOutputWithContext(ctx context.Context) AssetItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetItemMapOutput)
}

type AssetItemOutput struct{ *pulumi.OutputState }

func (AssetItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetItem)(nil)).Elem()
}

func (o AssetItemOutput) ToAssetItemOutput() AssetItemOutput {
	return o
}

func (o AssetItemOutput) ToAssetItemOutputWithContext(ctx context.Context) AssetItemOutput {
	return o
}

func (o AssetItemOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssetItem) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o AssetItemOutput) InputPorts() InputPortMapOutput {
	return o.ApplyT(func(v AssetItem) map[string]InputPort { return v.InputPorts }).(InputPortMapOutput)
}

func (o AssetItemOutput) LocationInfo() BlobLocationOutput {
	return o.ApplyT(func(v AssetItem) BlobLocation { return v.LocationInfo }).(BlobLocationOutput)
}

func (o AssetItemOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v AssetItem) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o AssetItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AssetItem) string { return v.Name }).(pulumi.StringOutput)
}

func (o AssetItemOutput) OutputPorts() OutputPortMapOutput {
	return o.ApplyT(func(v AssetItem) map[string]OutputPort { return v.OutputPorts }).(OutputPortMapOutput)
}

func (o AssetItemOutput) Parameters() ModuleAssetParameterArrayOutput {
	return o.ApplyT(func(v AssetItem) []ModuleAssetParameter { return v.Parameters }).(ModuleAssetParameterArrayOutput)
}

func (o AssetItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AssetItem) string { return v.Type }).(pulumi.StringOutput)
}

type AssetItemMapOutput struct{ *pulumi.OutputState }

func (AssetItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AssetItem)(nil)).Elem()
}

func (o AssetItemMapOutput) ToAssetItemMapOutput() AssetItemMapOutput {
	return o
}

func (o AssetItemMapOutput) ToAssetItemMapOutputWithContext(ctx context.Context) AssetItemMapOutput {
	return o
}

func (o AssetItemMapOutput) MapIndex(k pulumi.StringInput) AssetItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AssetItem {
		return vs[0].(map[string]AssetItem)[vs[1].(string)]
	}).(AssetItemOutput)
}

type AssetItemResponse struct {
	Id           *string                        `pulumi:"id"`
	InputPorts   map[string]InputPortResponse   `pulumi:"inputPorts"`
	LocationInfo BlobLocationResponse           `pulumi:"locationInfo"`
	Metadata     map[string]string              `pulumi:"metadata"`
	Name         string                         `pulumi:"name"`
	OutputPorts  map[string]OutputPortResponse  `pulumi:"outputPorts"`
	Parameters   []ModuleAssetParameterResponse `pulumi:"parameters"`
	Type         string                         `pulumi:"type"`
}





type AssetItemResponseInput interface {
	pulumi.Input

	ToAssetItemResponseOutput() AssetItemResponseOutput
	ToAssetItemResponseOutputWithContext(context.Context) AssetItemResponseOutput
}

type AssetItemResponseArgs struct {
	Id           pulumi.StringPtrInput                  `pulumi:"id"`
	InputPorts   InputPortResponseMapInput              `pulumi:"inputPorts"`
	LocationInfo BlobLocationResponseInput              `pulumi:"locationInfo"`
	Metadata     pulumi.StringMapInput                  `pulumi:"metadata"`
	Name         pulumi.StringInput                     `pulumi:"name"`
	OutputPorts  OutputPortResponseMapInput             `pulumi:"outputPorts"`
	Parameters   ModuleAssetParameterResponseArrayInput `pulumi:"parameters"`
	Type         pulumi.StringInput                     `pulumi:"type"`
}

func (AssetItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetItemResponse)(nil)).Elem()
}

func (i AssetItemResponseArgs) ToAssetItemResponseOutput() AssetItemResponseOutput {
	return i.ToAssetItemResponseOutputWithContext(context.Background())
}

func (i AssetItemResponseArgs) ToAssetItemResponseOutputWithContext(ctx context.Context) AssetItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetItemResponseOutput)
}





type AssetItemResponseMapInput interface {
	pulumi.Input

	ToAssetItemResponseMapOutput() AssetItemResponseMapOutput
	ToAssetItemResponseMapOutputWithContext(context.Context) AssetItemResponseMapOutput
}

type AssetItemResponseMap map[string]AssetItemResponseInput

func (AssetItemResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AssetItemResponse)(nil)).Elem()
}

func (i AssetItemResponseMap) ToAssetItemResponseMapOutput() AssetItemResponseMapOutput {
	return i.ToAssetItemResponseMapOutputWithContext(context.Background())
}

func (i AssetItemResponseMap) ToAssetItemResponseMapOutputWithContext(ctx context.Context) AssetItemResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetItemResponseMapOutput)
}

type AssetItemResponseOutput struct{ *pulumi.OutputState }

func (AssetItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetItemResponse)(nil)).Elem()
}

func (o AssetItemResponseOutput) ToAssetItemResponseOutput() AssetItemResponseOutput {
	return o
}

func (o AssetItemResponseOutput) ToAssetItemResponseOutputWithContext(ctx context.Context) AssetItemResponseOutput {
	return o
}

func (o AssetItemResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssetItemResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o AssetItemResponseOutput) InputPorts() InputPortResponseMapOutput {
	return o.ApplyT(func(v AssetItemResponse) map[string]InputPortResponse { return v.InputPorts }).(InputPortResponseMapOutput)
}

func (o AssetItemResponseOutput) LocationInfo() BlobLocationResponseOutput {
	return o.ApplyT(func(v AssetItemResponse) BlobLocationResponse { return v.LocationInfo }).(BlobLocationResponseOutput)
}

func (o AssetItemResponseOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v AssetItemResponse) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o AssetItemResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AssetItemResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AssetItemResponseOutput) OutputPorts() OutputPortResponseMapOutput {
	return o.ApplyT(func(v AssetItemResponse) map[string]OutputPortResponse { return v.OutputPorts }).(OutputPortResponseMapOutput)
}

func (o AssetItemResponseOutput) Parameters() ModuleAssetParameterResponseArrayOutput {
	return o.ApplyT(func(v AssetItemResponse) []ModuleAssetParameterResponse { return v.Parameters }).(ModuleAssetParameterResponseArrayOutput)
}

func (o AssetItemResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AssetItemResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AssetItemResponseMapOutput struct{ *pulumi.OutputState }

func (AssetItemResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AssetItemResponse)(nil)).Elem()
}

func (o AssetItemResponseMapOutput) ToAssetItemResponseMapOutput() AssetItemResponseMapOutput {
	return o
}

func (o AssetItemResponseMapOutput) ToAssetItemResponseMapOutputWithContext(ctx context.Context) AssetItemResponseMapOutput {
	return o
}

func (o AssetItemResponseMapOutput) MapIndex(k pulumi.StringInput) AssetItemResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AssetItemResponse {
		return vs[0].(map[string]AssetItemResponse)[vs[1].(string)]
	}).(AssetItemResponseOutput)
}

type BlobLocation struct {
	Credentials *string `pulumi:"credentials"`
	Uri         string  `pulumi:"uri"`
}





type BlobLocationInput interface {
	pulumi.Input

	ToBlobLocationOutput() BlobLocationOutput
	ToBlobLocationOutputWithContext(context.Context) BlobLocationOutput
}

type BlobLocationArgs struct {
	Credentials pulumi.StringPtrInput `pulumi:"credentials"`
	Uri         pulumi.StringInput    `pulumi:"uri"`
}

func (BlobLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobLocation)(nil)).Elem()
}

func (i BlobLocationArgs) ToBlobLocationOutput() BlobLocationOutput {
	return i.ToBlobLocationOutputWithContext(context.Background())
}

func (i BlobLocationArgs) ToBlobLocationOutputWithContext(ctx context.Context) BlobLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobLocationOutput)
}

func (i BlobLocationArgs) ToBlobLocationPtrOutput() BlobLocationPtrOutput {
	return i.ToBlobLocationPtrOutputWithContext(context.Background())
}

func (i BlobLocationArgs) ToBlobLocationPtrOutputWithContext(ctx context.Context) BlobLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobLocationOutput).ToBlobLocationPtrOutputWithContext(ctx)
}









type BlobLocationPtrInput interface {
	pulumi.Input

	ToBlobLocationPtrOutput() BlobLocationPtrOutput
	ToBlobLocationPtrOutputWithContext(context.Context) BlobLocationPtrOutput
}

type blobLocationPtrType BlobLocationArgs

func BlobLocationPtr(v *BlobLocationArgs) BlobLocationPtrInput {
	return (*blobLocationPtrType)(v)
}

func (*blobLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobLocation)(nil)).Elem()
}

func (i *blobLocationPtrType) ToBlobLocationPtrOutput() BlobLocationPtrOutput {
	return i.ToBlobLocationPtrOutputWithContext(context.Background())
}

func (i *blobLocationPtrType) ToBlobLocationPtrOutputWithContext(ctx context.Context) BlobLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobLocationPtrOutput)
}

type BlobLocationOutput struct{ *pulumi.OutputState }

func (BlobLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobLocation)(nil)).Elem()
}

func (o BlobLocationOutput) ToBlobLocationOutput() BlobLocationOutput {
	return o
}

func (o BlobLocationOutput) ToBlobLocationOutputWithContext(ctx context.Context) BlobLocationOutput {
	return o
}

func (o BlobLocationOutput) ToBlobLocationPtrOutput() BlobLocationPtrOutput {
	return o.ToBlobLocationPtrOutputWithContext(context.Background())
}

func (o BlobLocationOutput) ToBlobLocationPtrOutputWithContext(ctx context.Context) BlobLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobLocation) *BlobLocation {
		return &v
	}).(BlobLocationPtrOutput)
}

func (o BlobLocationOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobLocation) *string { return v.Credentials }).(pulumi.StringPtrOutput)
}

func (o BlobLocationOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v BlobLocation) string { return v.Uri }).(pulumi.StringOutput)
}

type BlobLocationPtrOutput struct{ *pulumi.OutputState }

func (BlobLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobLocation)(nil)).Elem()
}

func (o BlobLocationPtrOutput) ToBlobLocationPtrOutput() BlobLocationPtrOutput {
	return o
}

func (o BlobLocationPtrOutput) ToBlobLocationPtrOutputWithContext(ctx context.Context) BlobLocationPtrOutput {
	return o
}

func (o BlobLocationPtrOutput) Elem() BlobLocationOutput {
	return o.ApplyT(func(v *BlobLocation) BlobLocation {
		if v != nil {
			return *v
		}
		var ret BlobLocation
		return ret
	}).(BlobLocationOutput)
}

func (o BlobLocationPtrOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobLocation) *string {
		if v == nil {
			return nil
		}
		return v.Credentials
	}).(pulumi.StringPtrOutput)
}

func (o BlobLocationPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobLocation) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type BlobLocationResponse struct {
	Credentials *string `pulumi:"credentials"`
	Uri         string  `pulumi:"uri"`
}





type BlobLocationResponseInput interface {
	pulumi.Input

	ToBlobLocationResponseOutput() BlobLocationResponseOutput
	ToBlobLocationResponseOutputWithContext(context.Context) BlobLocationResponseOutput
}

type BlobLocationResponseArgs struct {
	Credentials pulumi.StringPtrInput `pulumi:"credentials"`
	Uri         pulumi.StringInput    `pulumi:"uri"`
}

func (BlobLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobLocationResponse)(nil)).Elem()
}

func (i BlobLocationResponseArgs) ToBlobLocationResponseOutput() BlobLocationResponseOutput {
	return i.ToBlobLocationResponseOutputWithContext(context.Background())
}

func (i BlobLocationResponseArgs) ToBlobLocationResponseOutputWithContext(ctx context.Context) BlobLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobLocationResponseOutput)
}

func (i BlobLocationResponseArgs) ToBlobLocationResponsePtrOutput() BlobLocationResponsePtrOutput {
	return i.ToBlobLocationResponsePtrOutputWithContext(context.Background())
}

func (i BlobLocationResponseArgs) ToBlobLocationResponsePtrOutputWithContext(ctx context.Context) BlobLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobLocationResponseOutput).ToBlobLocationResponsePtrOutputWithContext(ctx)
}









type BlobLocationResponsePtrInput interface {
	pulumi.Input

	ToBlobLocationResponsePtrOutput() BlobLocationResponsePtrOutput
	ToBlobLocationResponsePtrOutputWithContext(context.Context) BlobLocationResponsePtrOutput
}

type blobLocationResponsePtrType BlobLocationResponseArgs

func BlobLocationResponsePtr(v *BlobLocationResponseArgs) BlobLocationResponsePtrInput {
	return (*blobLocationResponsePtrType)(v)
}

func (*blobLocationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobLocationResponse)(nil)).Elem()
}

func (i *blobLocationResponsePtrType) ToBlobLocationResponsePtrOutput() BlobLocationResponsePtrOutput {
	return i.ToBlobLocationResponsePtrOutputWithContext(context.Background())
}

func (i *blobLocationResponsePtrType) ToBlobLocationResponsePtrOutputWithContext(ctx context.Context) BlobLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobLocationResponsePtrOutput)
}

type BlobLocationResponseOutput struct{ *pulumi.OutputState }

func (BlobLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobLocationResponse)(nil)).Elem()
}

func (o BlobLocationResponseOutput) ToBlobLocationResponseOutput() BlobLocationResponseOutput {
	return o
}

func (o BlobLocationResponseOutput) ToBlobLocationResponseOutputWithContext(ctx context.Context) BlobLocationResponseOutput {
	return o
}

func (o BlobLocationResponseOutput) ToBlobLocationResponsePtrOutput() BlobLocationResponsePtrOutput {
	return o.ToBlobLocationResponsePtrOutputWithContext(context.Background())
}

func (o BlobLocationResponseOutput) ToBlobLocationResponsePtrOutputWithContext(ctx context.Context) BlobLocationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobLocationResponse) *BlobLocationResponse {
		return &v
	}).(BlobLocationResponsePtrOutput)
}

func (o BlobLocationResponseOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobLocationResponse) *string { return v.Credentials }).(pulumi.StringPtrOutput)
}

func (o BlobLocationResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v BlobLocationResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type BlobLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobLocationResponse)(nil)).Elem()
}

func (o BlobLocationResponsePtrOutput) ToBlobLocationResponsePtrOutput() BlobLocationResponsePtrOutput {
	return o
}

func (o BlobLocationResponsePtrOutput) ToBlobLocationResponsePtrOutputWithContext(ctx context.Context) BlobLocationResponsePtrOutput {
	return o
}

func (o BlobLocationResponsePtrOutput) Elem() BlobLocationResponseOutput {
	return o.ApplyT(func(v *BlobLocationResponse) BlobLocationResponse {
		if v != nil {
			return *v
		}
		var ret BlobLocationResponse
		return ret
	}).(BlobLocationResponseOutput)
}

func (o BlobLocationResponsePtrOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Credentials
	}).(pulumi.StringPtrOutput)
}

func (o BlobLocationResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobLocationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type ColumnSpecification struct {
	Enum          []interface{} `pulumi:"enum"`
	Format        *string       `pulumi:"format"`
	Type          string        `pulumi:"type"`
	XMsIsnullable *bool         `pulumi:"xMsIsnullable"`
	XMsIsordered  *bool         `pulumi:"xMsIsordered"`
}





type ColumnSpecificationInput interface {
	pulumi.Input

	ToColumnSpecificationOutput() ColumnSpecificationOutput
	ToColumnSpecificationOutputWithContext(context.Context) ColumnSpecificationOutput
}

type ColumnSpecificationArgs struct {
	Enum          pulumi.ArrayInput     `pulumi:"enum"`
	Format        pulumi.StringPtrInput `pulumi:"format"`
	Type          pulumi.StringInput    `pulumi:"type"`
	XMsIsnullable pulumi.BoolPtrInput   `pulumi:"xMsIsnullable"`
	XMsIsordered  pulumi.BoolPtrInput   `pulumi:"xMsIsordered"`
}

func (ColumnSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnSpecification)(nil)).Elem()
}

func (i ColumnSpecificationArgs) ToColumnSpecificationOutput() ColumnSpecificationOutput {
	return i.ToColumnSpecificationOutputWithContext(context.Background())
}

func (i ColumnSpecificationArgs) ToColumnSpecificationOutputWithContext(ctx context.Context) ColumnSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnSpecificationOutput)
}





type ColumnSpecificationMapInput interface {
	pulumi.Input

	ToColumnSpecificationMapOutput() ColumnSpecificationMapOutput
	ToColumnSpecificationMapOutputWithContext(context.Context) ColumnSpecificationMapOutput
}

type ColumnSpecificationMap map[string]ColumnSpecificationInput

func (ColumnSpecificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ColumnSpecification)(nil)).Elem()
}

func (i ColumnSpecificationMap) ToColumnSpecificationMapOutput() ColumnSpecificationMapOutput {
	return i.ToColumnSpecificationMapOutputWithContext(context.Background())
}

func (i ColumnSpecificationMap) ToColumnSpecificationMapOutputWithContext(ctx context.Context) ColumnSpecificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnSpecificationMapOutput)
}

type ColumnSpecificationOutput struct{ *pulumi.OutputState }

func (ColumnSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnSpecification)(nil)).Elem()
}

func (o ColumnSpecificationOutput) ToColumnSpecificationOutput() ColumnSpecificationOutput {
	return o
}

func (o ColumnSpecificationOutput) ToColumnSpecificationOutputWithContext(ctx context.Context) ColumnSpecificationOutput {
	return o
}

func (o ColumnSpecificationOutput) Enum() pulumi.ArrayOutput {
	return o.ApplyT(func(v ColumnSpecification) []interface{} { return v.Enum }).(pulumi.ArrayOutput)
}

func (o ColumnSpecificationOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnSpecification) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ColumnSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ColumnSpecification) string { return v.Type }).(pulumi.StringOutput)
}

func (o ColumnSpecificationOutput) XMsIsnullable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ColumnSpecification) *bool { return v.XMsIsnullable }).(pulumi.BoolPtrOutput)
}

func (o ColumnSpecificationOutput) XMsIsordered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ColumnSpecification) *bool { return v.XMsIsordered }).(pulumi.BoolPtrOutput)
}

type ColumnSpecificationMapOutput struct{ *pulumi.OutputState }

func (ColumnSpecificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ColumnSpecification)(nil)).Elem()
}

func (o ColumnSpecificationMapOutput) ToColumnSpecificationMapOutput() ColumnSpecificationMapOutput {
	return o
}

func (o ColumnSpecificationMapOutput) ToColumnSpecificationMapOutputWithContext(ctx context.Context) ColumnSpecificationMapOutput {
	return o
}

func (o ColumnSpecificationMapOutput) MapIndex(k pulumi.StringInput) ColumnSpecificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ColumnSpecification {
		return vs[0].(map[string]ColumnSpecification)[vs[1].(string)]
	}).(ColumnSpecificationOutput)
}

type ColumnSpecificationResponse struct {
	Enum          []interface{} `pulumi:"enum"`
	Format        *string       `pulumi:"format"`
	Type          string        `pulumi:"type"`
	XMsIsnullable *bool         `pulumi:"xMsIsnullable"`
	XMsIsordered  *bool         `pulumi:"xMsIsordered"`
}





type ColumnSpecificationResponseInput interface {
	pulumi.Input

	ToColumnSpecificationResponseOutput() ColumnSpecificationResponseOutput
	ToColumnSpecificationResponseOutputWithContext(context.Context) ColumnSpecificationResponseOutput
}

type ColumnSpecificationResponseArgs struct {
	Enum          pulumi.ArrayInput     `pulumi:"enum"`
	Format        pulumi.StringPtrInput `pulumi:"format"`
	Type          pulumi.StringInput    `pulumi:"type"`
	XMsIsnullable pulumi.BoolPtrInput   `pulumi:"xMsIsnullable"`
	XMsIsordered  pulumi.BoolPtrInput   `pulumi:"xMsIsordered"`
}

func (ColumnSpecificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnSpecificationResponse)(nil)).Elem()
}

func (i ColumnSpecificationResponseArgs) ToColumnSpecificationResponseOutput() ColumnSpecificationResponseOutput {
	return i.ToColumnSpecificationResponseOutputWithContext(context.Background())
}

func (i ColumnSpecificationResponseArgs) ToColumnSpecificationResponseOutputWithContext(ctx context.Context) ColumnSpecificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnSpecificationResponseOutput)
}





type ColumnSpecificationResponseMapInput interface {
	pulumi.Input

	ToColumnSpecificationResponseMapOutput() ColumnSpecificationResponseMapOutput
	ToColumnSpecificationResponseMapOutputWithContext(context.Context) ColumnSpecificationResponseMapOutput
}

type ColumnSpecificationResponseMap map[string]ColumnSpecificationResponseInput

func (ColumnSpecificationResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ColumnSpecificationResponse)(nil)).Elem()
}

func (i ColumnSpecificationResponseMap) ToColumnSpecificationResponseMapOutput() ColumnSpecificationResponseMapOutput {
	return i.ToColumnSpecificationResponseMapOutputWithContext(context.Background())
}

func (i ColumnSpecificationResponseMap) ToColumnSpecificationResponseMapOutputWithContext(ctx context.Context) ColumnSpecificationResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnSpecificationResponseMapOutput)
}

type ColumnSpecificationResponseOutput struct{ *pulumi.OutputState }

func (ColumnSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnSpecificationResponse)(nil)).Elem()
}

func (o ColumnSpecificationResponseOutput) ToColumnSpecificationResponseOutput() ColumnSpecificationResponseOutput {
	return o
}

func (o ColumnSpecificationResponseOutput) ToColumnSpecificationResponseOutputWithContext(ctx context.Context) ColumnSpecificationResponseOutput {
	return o
}

func (o ColumnSpecificationResponseOutput) Enum() pulumi.ArrayOutput {
	return o.ApplyT(func(v ColumnSpecificationResponse) []interface{} { return v.Enum }).(pulumi.ArrayOutput)
}

func (o ColumnSpecificationResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnSpecificationResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ColumnSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ColumnSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ColumnSpecificationResponseOutput) XMsIsnullable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ColumnSpecificationResponse) *bool { return v.XMsIsnullable }).(pulumi.BoolPtrOutput)
}

func (o ColumnSpecificationResponseOutput) XMsIsordered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ColumnSpecificationResponse) *bool { return v.XMsIsordered }).(pulumi.BoolPtrOutput)
}

type ColumnSpecificationResponseMapOutput struct{ *pulumi.OutputState }

func (ColumnSpecificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ColumnSpecificationResponse)(nil)).Elem()
}

func (o ColumnSpecificationResponseMapOutput) ToColumnSpecificationResponseMapOutput() ColumnSpecificationResponseMapOutput {
	return o
}

func (o ColumnSpecificationResponseMapOutput) ToColumnSpecificationResponseMapOutputWithContext(ctx context.Context) ColumnSpecificationResponseMapOutput {
	return o
}

func (o ColumnSpecificationResponseMapOutput) MapIndex(k pulumi.StringInput) ColumnSpecificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ColumnSpecificationResponse {
		return vs[0].(map[string]ColumnSpecificationResponse)[vs[1].(string)]
	}).(ColumnSpecificationResponseOutput)
}

type CommitmentPlanType struct {
	Id string `pulumi:"id"`
}





type CommitmentPlanTypeInput interface {
	pulumi.Input

	ToCommitmentPlanTypeOutput() CommitmentPlanTypeOutput
	ToCommitmentPlanTypeOutputWithContext(context.Context) CommitmentPlanTypeOutput
}

type CommitmentPlanTypeArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (CommitmentPlanTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanType)(nil)).Elem()
}

func (i CommitmentPlanTypeArgs) ToCommitmentPlanTypeOutput() CommitmentPlanTypeOutput {
	return i.ToCommitmentPlanTypeOutputWithContext(context.Background())
}

func (i CommitmentPlanTypeArgs) ToCommitmentPlanTypeOutputWithContext(ctx context.Context) CommitmentPlanTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanTypeOutput)
}

func (i CommitmentPlanTypeArgs) ToCommitmentPlanTypePtrOutput() CommitmentPlanTypePtrOutput {
	return i.ToCommitmentPlanTypePtrOutputWithContext(context.Background())
}

func (i CommitmentPlanTypeArgs) ToCommitmentPlanTypePtrOutputWithContext(ctx context.Context) CommitmentPlanTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanTypeOutput).ToCommitmentPlanTypePtrOutputWithContext(ctx)
}









type CommitmentPlanTypePtrInput interface {
	pulumi.Input

	ToCommitmentPlanTypePtrOutput() CommitmentPlanTypePtrOutput
	ToCommitmentPlanTypePtrOutputWithContext(context.Context) CommitmentPlanTypePtrOutput
}

type commitmentPlanTypePtrType CommitmentPlanTypeArgs

func CommitmentPlanTypePtr(v *CommitmentPlanTypeArgs) CommitmentPlanTypePtrInput {
	return (*commitmentPlanTypePtrType)(v)
}

func (*commitmentPlanTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanType)(nil)).Elem()
}

func (i *commitmentPlanTypePtrType) ToCommitmentPlanTypePtrOutput() CommitmentPlanTypePtrOutput {
	return i.ToCommitmentPlanTypePtrOutputWithContext(context.Background())
}

func (i *commitmentPlanTypePtrType) ToCommitmentPlanTypePtrOutputWithContext(ctx context.Context) CommitmentPlanTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanTypePtrOutput)
}

type CommitmentPlanTypeOutput struct{ *pulumi.OutputState }

func (CommitmentPlanTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanType)(nil)).Elem()
}

func (o CommitmentPlanTypeOutput) ToCommitmentPlanTypeOutput() CommitmentPlanTypeOutput {
	return o
}

func (o CommitmentPlanTypeOutput) ToCommitmentPlanTypeOutputWithContext(ctx context.Context) CommitmentPlanTypeOutput {
	return o
}

func (o CommitmentPlanTypeOutput) ToCommitmentPlanTypePtrOutput() CommitmentPlanTypePtrOutput {
	return o.ToCommitmentPlanTypePtrOutputWithContext(context.Background())
}

func (o CommitmentPlanTypeOutput) ToCommitmentPlanTypePtrOutputWithContext(ctx context.Context) CommitmentPlanTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPlanType) *CommitmentPlanType {
		return &v
	}).(CommitmentPlanTypePtrOutput)
}

func (o CommitmentPlanTypeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CommitmentPlanType) string { return v.Id }).(pulumi.StringOutput)
}

type CommitmentPlanTypePtrOutput struct{ *pulumi.OutputState }

func (CommitmentPlanTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanType)(nil)).Elem()
}

func (o CommitmentPlanTypePtrOutput) ToCommitmentPlanTypePtrOutput() CommitmentPlanTypePtrOutput {
	return o
}

func (o CommitmentPlanTypePtrOutput) ToCommitmentPlanTypePtrOutputWithContext(ctx context.Context) CommitmentPlanTypePtrOutput {
	return o
}

func (o CommitmentPlanTypePtrOutput) Elem() CommitmentPlanTypeOutput {
	return o.ApplyT(func(v *CommitmentPlanType) CommitmentPlanType {
		if v != nil {
			return *v
		}
		var ret CommitmentPlanType
		return ret
	}).(CommitmentPlanTypeOutput)
}

func (o CommitmentPlanTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanType) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type CommitmentPlanPropertiesResponse struct {
	ChargeForOverage      bool                            `pulumi:"chargeForOverage"`
	ChargeForPlan         bool                            `pulumi:"chargeForPlan"`
	CreationDate          string                          `pulumi:"creationDate"`
	IncludedQuantities    map[string]PlanQuantityResponse `pulumi:"includedQuantities"`
	MaxAssociationLimit   int                             `pulumi:"maxAssociationLimit"`
	MaxCapacityLimit      int                             `pulumi:"maxCapacityLimit"`
	MinCapacityLimit      int                             `pulumi:"minCapacityLimit"`
	PlanMeter             string                          `pulumi:"planMeter"`
	RefillFrequencyInDays int                             `pulumi:"refillFrequencyInDays"`
	SuspendPlanOnOverage  bool                            `pulumi:"suspendPlanOnOverage"`
}





type CommitmentPlanPropertiesResponseInput interface {
	pulumi.Input

	ToCommitmentPlanPropertiesResponseOutput() CommitmentPlanPropertiesResponseOutput
	ToCommitmentPlanPropertiesResponseOutputWithContext(context.Context) CommitmentPlanPropertiesResponseOutput
}

type CommitmentPlanPropertiesResponseArgs struct {
	ChargeForOverage      pulumi.BoolInput             `pulumi:"chargeForOverage"`
	ChargeForPlan         pulumi.BoolInput             `pulumi:"chargeForPlan"`
	CreationDate          pulumi.StringInput           `pulumi:"creationDate"`
	IncludedQuantities    PlanQuantityResponseMapInput `pulumi:"includedQuantities"`
	MaxAssociationLimit   pulumi.IntInput              `pulumi:"maxAssociationLimit"`
	MaxCapacityLimit      pulumi.IntInput              `pulumi:"maxCapacityLimit"`
	MinCapacityLimit      pulumi.IntInput              `pulumi:"minCapacityLimit"`
	PlanMeter             pulumi.StringInput           `pulumi:"planMeter"`
	RefillFrequencyInDays pulumi.IntInput              `pulumi:"refillFrequencyInDays"`
	SuspendPlanOnOverage  pulumi.BoolInput             `pulumi:"suspendPlanOnOverage"`
}

func (CommitmentPlanPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponseOutput() CommitmentPlanPropertiesResponseOutput {
	return i.ToCommitmentPlanPropertiesResponseOutputWithContext(context.Background())
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponseOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesResponseOutput)
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return i.ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CommitmentPlanPropertiesResponseArgs) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesResponseOutput).ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx)
}









type CommitmentPlanPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput
	ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Context) CommitmentPlanPropertiesResponsePtrOutput
}

type commitmentPlanPropertiesResponsePtrType CommitmentPlanPropertiesResponseArgs

func CommitmentPlanPropertiesResponsePtr(v *CommitmentPlanPropertiesResponseArgs) CommitmentPlanPropertiesResponsePtrInput {
	return (*commitmentPlanPropertiesResponsePtrType)(v)
}

func (*commitmentPlanPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (i *commitmentPlanPropertiesResponsePtrType) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return i.ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *commitmentPlanPropertiesResponsePtrType) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanPropertiesResponsePtrOutput)
}

type CommitmentPlanPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CommitmentPlanPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponseOutput() CommitmentPlanPropertiesResponseOutput {
	return o
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponseOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponseOutput {
	return o
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return o.ToCommitmentPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CommitmentPlanPropertiesResponseOutput) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPlanPropertiesResponse) *CommitmentPlanPropertiesResponse {
		return &v
	}).(CommitmentPlanPropertiesResponsePtrOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) ChargeForOverage() pulumi.BoolOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) bool { return v.ChargeForOverage }).(pulumi.BoolOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) ChargeForPlan() pulumi.BoolOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) bool { return v.ChargeForPlan }).(pulumi.BoolOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) IncludedQuantities() PlanQuantityResponseMapOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) map[string]PlanQuantityResponse { return v.IncludedQuantities }).(PlanQuantityResponseMapOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) MaxAssociationLimit() pulumi.IntOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) int { return v.MaxAssociationLimit }).(pulumi.IntOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) MaxCapacityLimit() pulumi.IntOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) int { return v.MaxCapacityLimit }).(pulumi.IntOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) MinCapacityLimit() pulumi.IntOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) int { return v.MinCapacityLimit }).(pulumi.IntOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) PlanMeter() pulumi.StringOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) string { return v.PlanMeter }).(pulumi.StringOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) RefillFrequencyInDays() pulumi.IntOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) int { return v.RefillFrequencyInDays }).(pulumi.IntOutput)
}

func (o CommitmentPlanPropertiesResponseOutput) SuspendPlanOnOverage() pulumi.BoolOutput {
	return o.ApplyT(func(v CommitmentPlanPropertiesResponse) bool { return v.SuspendPlanOnOverage }).(pulumi.BoolOutput)
}

type CommitmentPlanPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CommitmentPlanPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanPropertiesResponse)(nil)).Elem()
}

func (o CommitmentPlanPropertiesResponsePtrOutput) ToCommitmentPlanPropertiesResponsePtrOutput() CommitmentPlanPropertiesResponsePtrOutput {
	return o
}

func (o CommitmentPlanPropertiesResponsePtrOutput) ToCommitmentPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanPropertiesResponsePtrOutput {
	return o
}

func (o CommitmentPlanPropertiesResponsePtrOutput) Elem() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) CommitmentPlanPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CommitmentPlanPropertiesResponse
		return ret
	}).(CommitmentPlanPropertiesResponseOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) ChargeForOverage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ChargeForOverage
	}).(pulumi.BoolPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) ChargeForPlan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ChargeForPlan
	}).(pulumi.BoolPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreationDate
	}).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) IncludedQuantities() PlanQuantityResponseMapOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) map[string]PlanQuantityResponse {
		if v == nil {
			return nil
		}
		return v.IncludedQuantities
	}).(PlanQuantityResponseMapOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) MaxAssociationLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxAssociationLimit
	}).(pulumi.IntPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) MaxCapacityLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxCapacityLimit
	}).(pulumi.IntPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) MinCapacityLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MinCapacityLimit
	}).(pulumi.IntPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) PlanMeter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PlanMeter
	}).(pulumi.StringPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) RefillFrequencyInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RefillFrequencyInDays
	}).(pulumi.IntPtrOutput)
}

func (o CommitmentPlanPropertiesResponsePtrOutput) SuspendPlanOnOverage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendPlanOnOverage
	}).(pulumi.BoolPtrOutput)
}

type CommitmentPlanResponse struct {
	Id string `pulumi:"id"`
}





type CommitmentPlanResponseInput interface {
	pulumi.Input

	ToCommitmentPlanResponseOutput() CommitmentPlanResponseOutput
	ToCommitmentPlanResponseOutputWithContext(context.Context) CommitmentPlanResponseOutput
}

type CommitmentPlanResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (CommitmentPlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanResponse)(nil)).Elem()
}

func (i CommitmentPlanResponseArgs) ToCommitmentPlanResponseOutput() CommitmentPlanResponseOutput {
	return i.ToCommitmentPlanResponseOutputWithContext(context.Background())
}

func (i CommitmentPlanResponseArgs) ToCommitmentPlanResponseOutputWithContext(ctx context.Context) CommitmentPlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanResponseOutput)
}

func (i CommitmentPlanResponseArgs) ToCommitmentPlanResponsePtrOutput() CommitmentPlanResponsePtrOutput {
	return i.ToCommitmentPlanResponsePtrOutputWithContext(context.Background())
}

func (i CommitmentPlanResponseArgs) ToCommitmentPlanResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanResponseOutput).ToCommitmentPlanResponsePtrOutputWithContext(ctx)
}









type CommitmentPlanResponsePtrInput interface {
	pulumi.Input

	ToCommitmentPlanResponsePtrOutput() CommitmentPlanResponsePtrOutput
	ToCommitmentPlanResponsePtrOutputWithContext(context.Context) CommitmentPlanResponsePtrOutput
}

type commitmentPlanResponsePtrType CommitmentPlanResponseArgs

func CommitmentPlanResponsePtr(v *CommitmentPlanResponseArgs) CommitmentPlanResponsePtrInput {
	return (*commitmentPlanResponsePtrType)(v)
}

func (*commitmentPlanResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanResponse)(nil)).Elem()
}

func (i *commitmentPlanResponsePtrType) ToCommitmentPlanResponsePtrOutput() CommitmentPlanResponsePtrOutput {
	return i.ToCommitmentPlanResponsePtrOutputWithContext(context.Background())
}

func (i *commitmentPlanResponsePtrType) ToCommitmentPlanResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommitmentPlanResponsePtrOutput)
}

type CommitmentPlanResponseOutput struct{ *pulumi.OutputState }

func (CommitmentPlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommitmentPlanResponse)(nil)).Elem()
}

func (o CommitmentPlanResponseOutput) ToCommitmentPlanResponseOutput() CommitmentPlanResponseOutput {
	return o
}

func (o CommitmentPlanResponseOutput) ToCommitmentPlanResponseOutputWithContext(ctx context.Context) CommitmentPlanResponseOutput {
	return o
}

func (o CommitmentPlanResponseOutput) ToCommitmentPlanResponsePtrOutput() CommitmentPlanResponsePtrOutput {
	return o.ToCommitmentPlanResponsePtrOutputWithContext(context.Background())
}

func (o CommitmentPlanResponseOutput) ToCommitmentPlanResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommitmentPlanResponse) *CommitmentPlanResponse {
		return &v
	}).(CommitmentPlanResponsePtrOutput)
}

func (o CommitmentPlanResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CommitmentPlanResponse) string { return v.Id }).(pulumi.StringOutput)
}

type CommitmentPlanResponsePtrOutput struct{ *pulumi.OutputState }

func (CommitmentPlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommitmentPlanResponse)(nil)).Elem()
}

func (o CommitmentPlanResponsePtrOutput) ToCommitmentPlanResponsePtrOutput() CommitmentPlanResponsePtrOutput {
	return o
}

func (o CommitmentPlanResponsePtrOutput) ToCommitmentPlanResponsePtrOutputWithContext(ctx context.Context) CommitmentPlanResponsePtrOutput {
	return o
}

func (o CommitmentPlanResponsePtrOutput) Elem() CommitmentPlanResponseOutput {
	return o.ApplyT(func(v *CommitmentPlanResponse) CommitmentPlanResponse {
		if v != nil {
			return *v
		}
		var ret CommitmentPlanResponse
		return ret
	}).(CommitmentPlanResponseOutput)
}

func (o CommitmentPlanResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommitmentPlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type DiagnosticsConfiguration struct {
	Expiry *string `pulumi:"expiry"`
	Level  string  `pulumi:"level"`
}





type DiagnosticsConfigurationInput interface {
	pulumi.Input

	ToDiagnosticsConfigurationOutput() DiagnosticsConfigurationOutput
	ToDiagnosticsConfigurationOutputWithContext(context.Context) DiagnosticsConfigurationOutput
}

type DiagnosticsConfigurationArgs struct {
	Expiry pulumi.StringPtrInput `pulumi:"expiry"`
	Level  pulumi.StringInput    `pulumi:"level"`
}

func (DiagnosticsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsConfiguration)(nil)).Elem()
}

func (i DiagnosticsConfigurationArgs) ToDiagnosticsConfigurationOutput() DiagnosticsConfigurationOutput {
	return i.ToDiagnosticsConfigurationOutputWithContext(context.Background())
}

func (i DiagnosticsConfigurationArgs) ToDiagnosticsConfigurationOutputWithContext(ctx context.Context) DiagnosticsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsConfigurationOutput)
}

func (i DiagnosticsConfigurationArgs) ToDiagnosticsConfigurationPtrOutput() DiagnosticsConfigurationPtrOutput {
	return i.ToDiagnosticsConfigurationPtrOutputWithContext(context.Background())
}

func (i DiagnosticsConfigurationArgs) ToDiagnosticsConfigurationPtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsConfigurationOutput).ToDiagnosticsConfigurationPtrOutputWithContext(ctx)
}









type DiagnosticsConfigurationPtrInput interface {
	pulumi.Input

	ToDiagnosticsConfigurationPtrOutput() DiagnosticsConfigurationPtrOutput
	ToDiagnosticsConfigurationPtrOutputWithContext(context.Context) DiagnosticsConfigurationPtrOutput
}

type diagnosticsConfigurationPtrType DiagnosticsConfigurationArgs

func DiagnosticsConfigurationPtr(v *DiagnosticsConfigurationArgs) DiagnosticsConfigurationPtrInput {
	return (*diagnosticsConfigurationPtrType)(v)
}

func (*diagnosticsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsConfiguration)(nil)).Elem()
}

func (i *diagnosticsConfigurationPtrType) ToDiagnosticsConfigurationPtrOutput() DiagnosticsConfigurationPtrOutput {
	return i.ToDiagnosticsConfigurationPtrOutputWithContext(context.Background())
}

func (i *diagnosticsConfigurationPtrType) ToDiagnosticsConfigurationPtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsConfigurationPtrOutput)
}

type DiagnosticsConfigurationOutput struct{ *pulumi.OutputState }

func (DiagnosticsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsConfiguration)(nil)).Elem()
}

func (o DiagnosticsConfigurationOutput) ToDiagnosticsConfigurationOutput() DiagnosticsConfigurationOutput {
	return o
}

func (o DiagnosticsConfigurationOutput) ToDiagnosticsConfigurationOutputWithContext(ctx context.Context) DiagnosticsConfigurationOutput {
	return o
}

func (o DiagnosticsConfigurationOutput) ToDiagnosticsConfigurationPtrOutput() DiagnosticsConfigurationPtrOutput {
	return o.ToDiagnosticsConfigurationPtrOutputWithContext(context.Background())
}

func (o DiagnosticsConfigurationOutput) ToDiagnosticsConfigurationPtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsConfiguration) *DiagnosticsConfiguration {
		return &v
	}).(DiagnosticsConfigurationPtrOutput)
}

func (o DiagnosticsConfigurationOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticsConfiguration) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsConfigurationOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsConfiguration) string { return v.Level }).(pulumi.StringOutput)
}

type DiagnosticsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsConfiguration)(nil)).Elem()
}

func (o DiagnosticsConfigurationPtrOutput) ToDiagnosticsConfigurationPtrOutput() DiagnosticsConfigurationPtrOutput {
	return o
}

func (o DiagnosticsConfigurationPtrOutput) ToDiagnosticsConfigurationPtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationPtrOutput {
	return o
}

func (o DiagnosticsConfigurationPtrOutput) Elem() DiagnosticsConfigurationOutput {
	return o.ApplyT(func(v *DiagnosticsConfiguration) DiagnosticsConfiguration {
		if v != nil {
			return *v
		}
		var ret DiagnosticsConfiguration
		return ret
	}).(DiagnosticsConfigurationOutput)
}

func (o DiagnosticsConfigurationPtrOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Expiry
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsConfigurationPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

type DiagnosticsConfigurationResponse struct {
	Expiry *string `pulumi:"expiry"`
	Level  string  `pulumi:"level"`
}





type DiagnosticsConfigurationResponseInput interface {
	pulumi.Input

	ToDiagnosticsConfigurationResponseOutput() DiagnosticsConfigurationResponseOutput
	ToDiagnosticsConfigurationResponseOutputWithContext(context.Context) DiagnosticsConfigurationResponseOutput
}

type DiagnosticsConfigurationResponseArgs struct {
	Expiry pulumi.StringPtrInput `pulumi:"expiry"`
	Level  pulumi.StringInput    `pulumi:"level"`
}

func (DiagnosticsConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsConfigurationResponse)(nil)).Elem()
}

func (i DiagnosticsConfigurationResponseArgs) ToDiagnosticsConfigurationResponseOutput() DiagnosticsConfigurationResponseOutput {
	return i.ToDiagnosticsConfigurationResponseOutputWithContext(context.Background())
}

func (i DiagnosticsConfigurationResponseArgs) ToDiagnosticsConfigurationResponseOutputWithContext(ctx context.Context) DiagnosticsConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsConfigurationResponseOutput)
}

func (i DiagnosticsConfigurationResponseArgs) ToDiagnosticsConfigurationResponsePtrOutput() DiagnosticsConfigurationResponsePtrOutput {
	return i.ToDiagnosticsConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i DiagnosticsConfigurationResponseArgs) ToDiagnosticsConfigurationResponsePtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsConfigurationResponseOutput).ToDiagnosticsConfigurationResponsePtrOutputWithContext(ctx)
}









type DiagnosticsConfigurationResponsePtrInput interface {
	pulumi.Input

	ToDiagnosticsConfigurationResponsePtrOutput() DiagnosticsConfigurationResponsePtrOutput
	ToDiagnosticsConfigurationResponsePtrOutputWithContext(context.Context) DiagnosticsConfigurationResponsePtrOutput
}

type diagnosticsConfigurationResponsePtrType DiagnosticsConfigurationResponseArgs

func DiagnosticsConfigurationResponsePtr(v *DiagnosticsConfigurationResponseArgs) DiagnosticsConfigurationResponsePtrInput {
	return (*diagnosticsConfigurationResponsePtrType)(v)
}

func (*diagnosticsConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsConfigurationResponse)(nil)).Elem()
}

func (i *diagnosticsConfigurationResponsePtrType) ToDiagnosticsConfigurationResponsePtrOutput() DiagnosticsConfigurationResponsePtrOutput {
	return i.ToDiagnosticsConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *diagnosticsConfigurationResponsePtrType) ToDiagnosticsConfigurationResponsePtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsConfigurationResponsePtrOutput)
}

type DiagnosticsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsConfigurationResponse)(nil)).Elem()
}

func (o DiagnosticsConfigurationResponseOutput) ToDiagnosticsConfigurationResponseOutput() DiagnosticsConfigurationResponseOutput {
	return o
}

func (o DiagnosticsConfigurationResponseOutput) ToDiagnosticsConfigurationResponseOutputWithContext(ctx context.Context) DiagnosticsConfigurationResponseOutput {
	return o
}

func (o DiagnosticsConfigurationResponseOutput) ToDiagnosticsConfigurationResponsePtrOutput() DiagnosticsConfigurationResponsePtrOutput {
	return o.ToDiagnosticsConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o DiagnosticsConfigurationResponseOutput) ToDiagnosticsConfigurationResponsePtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsConfigurationResponse) *DiagnosticsConfigurationResponse {
		return &v
	}).(DiagnosticsConfigurationResponsePtrOutput)
}

func (o DiagnosticsConfigurationResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticsConfigurationResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsConfigurationResponseOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsConfigurationResponse) string { return v.Level }).(pulumi.StringOutput)
}

type DiagnosticsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsConfigurationResponse)(nil)).Elem()
}

func (o DiagnosticsConfigurationResponsePtrOutput) ToDiagnosticsConfigurationResponsePtrOutput() DiagnosticsConfigurationResponsePtrOutput {
	return o
}

func (o DiagnosticsConfigurationResponsePtrOutput) ToDiagnosticsConfigurationResponsePtrOutputWithContext(ctx context.Context) DiagnosticsConfigurationResponsePtrOutput {
	return o
}

func (o DiagnosticsConfigurationResponsePtrOutput) Elem() DiagnosticsConfigurationResponseOutput {
	return o.ApplyT(func(v *DiagnosticsConfigurationResponse) DiagnosticsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsConfigurationResponse
		return ret
	}).(DiagnosticsConfigurationResponseOutput)
}

func (o DiagnosticsConfigurationResponsePtrOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Expiry
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticsConfigurationResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

type ExampleRequest struct {
	GlobalParameters map[string]interface{}     `pulumi:"globalParameters"`
	Inputs           map[string][][]interface{} `pulumi:"inputs"`
}





type ExampleRequestInput interface {
	pulumi.Input

	ToExampleRequestOutput() ExampleRequestOutput
	ToExampleRequestOutputWithContext(context.Context) ExampleRequestOutput
}

type ExampleRequestArgs struct {
	GlobalParameters pulumi.MapInput           `pulumi:"globalParameters"`
	Inputs           pulumi.ArrayArrayMapInput `pulumi:"inputs"`
}

func (ExampleRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleRequest)(nil)).Elem()
}

func (i ExampleRequestArgs) ToExampleRequestOutput() ExampleRequestOutput {
	return i.ToExampleRequestOutputWithContext(context.Background())
}

func (i ExampleRequestArgs) ToExampleRequestOutputWithContext(ctx context.Context) ExampleRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleRequestOutput)
}

func (i ExampleRequestArgs) ToExampleRequestPtrOutput() ExampleRequestPtrOutput {
	return i.ToExampleRequestPtrOutputWithContext(context.Background())
}

func (i ExampleRequestArgs) ToExampleRequestPtrOutputWithContext(ctx context.Context) ExampleRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleRequestOutput).ToExampleRequestPtrOutputWithContext(ctx)
}









type ExampleRequestPtrInput interface {
	pulumi.Input

	ToExampleRequestPtrOutput() ExampleRequestPtrOutput
	ToExampleRequestPtrOutputWithContext(context.Context) ExampleRequestPtrOutput
}

type exampleRequestPtrType ExampleRequestArgs

func ExampleRequestPtr(v *ExampleRequestArgs) ExampleRequestPtrInput {
	return (*exampleRequestPtrType)(v)
}

func (*exampleRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleRequest)(nil)).Elem()
}

func (i *exampleRequestPtrType) ToExampleRequestPtrOutput() ExampleRequestPtrOutput {
	return i.ToExampleRequestPtrOutputWithContext(context.Background())
}

func (i *exampleRequestPtrType) ToExampleRequestPtrOutputWithContext(ctx context.Context) ExampleRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleRequestPtrOutput)
}

type ExampleRequestOutput struct{ *pulumi.OutputState }

func (ExampleRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleRequest)(nil)).Elem()
}

func (o ExampleRequestOutput) ToExampleRequestOutput() ExampleRequestOutput {
	return o
}

func (o ExampleRequestOutput) ToExampleRequestOutputWithContext(ctx context.Context) ExampleRequestOutput {
	return o
}

func (o ExampleRequestOutput) ToExampleRequestPtrOutput() ExampleRequestPtrOutput {
	return o.ToExampleRequestPtrOutputWithContext(context.Background())
}

func (o ExampleRequestOutput) ToExampleRequestPtrOutputWithContext(ctx context.Context) ExampleRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExampleRequest) *ExampleRequest {
		return &v
	}).(ExampleRequestPtrOutput)
}

func (o ExampleRequestOutput) GlobalParameters() pulumi.MapOutput {
	return o.ApplyT(func(v ExampleRequest) map[string]interface{} { return v.GlobalParameters }).(pulumi.MapOutput)
}

func (o ExampleRequestOutput) Inputs() pulumi.ArrayArrayMapOutput {
	return o.ApplyT(func(v ExampleRequest) map[string][][]interface{} { return v.Inputs }).(pulumi.ArrayArrayMapOutput)
}

type ExampleRequestPtrOutput struct{ *pulumi.OutputState }

func (ExampleRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleRequest)(nil)).Elem()
}

func (o ExampleRequestPtrOutput) ToExampleRequestPtrOutput() ExampleRequestPtrOutput {
	return o
}

func (o ExampleRequestPtrOutput) ToExampleRequestPtrOutputWithContext(ctx context.Context) ExampleRequestPtrOutput {
	return o
}

func (o ExampleRequestPtrOutput) Elem() ExampleRequestOutput {
	return o.ApplyT(func(v *ExampleRequest) ExampleRequest {
		if v != nil {
			return *v
		}
		var ret ExampleRequest
		return ret
	}).(ExampleRequestOutput)
}

func (o ExampleRequestPtrOutput) GlobalParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *ExampleRequest) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.GlobalParameters
	}).(pulumi.MapOutput)
}

func (o ExampleRequestPtrOutput) Inputs() pulumi.ArrayArrayMapOutput {
	return o.ApplyT(func(v *ExampleRequest) map[string][][]interface{} {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(pulumi.ArrayArrayMapOutput)
}

type ExampleRequestResponse struct {
	GlobalParameters map[string]interface{}     `pulumi:"globalParameters"`
	Inputs           map[string][][]interface{} `pulumi:"inputs"`
}





type ExampleRequestResponseInput interface {
	pulumi.Input

	ToExampleRequestResponseOutput() ExampleRequestResponseOutput
	ToExampleRequestResponseOutputWithContext(context.Context) ExampleRequestResponseOutput
}

type ExampleRequestResponseArgs struct {
	GlobalParameters pulumi.MapInput           `pulumi:"globalParameters"`
	Inputs           pulumi.ArrayArrayMapInput `pulumi:"inputs"`
}

func (ExampleRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleRequestResponse)(nil)).Elem()
}

func (i ExampleRequestResponseArgs) ToExampleRequestResponseOutput() ExampleRequestResponseOutput {
	return i.ToExampleRequestResponseOutputWithContext(context.Background())
}

func (i ExampleRequestResponseArgs) ToExampleRequestResponseOutputWithContext(ctx context.Context) ExampleRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleRequestResponseOutput)
}

func (i ExampleRequestResponseArgs) ToExampleRequestResponsePtrOutput() ExampleRequestResponsePtrOutput {
	return i.ToExampleRequestResponsePtrOutputWithContext(context.Background())
}

func (i ExampleRequestResponseArgs) ToExampleRequestResponsePtrOutputWithContext(ctx context.Context) ExampleRequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleRequestResponseOutput).ToExampleRequestResponsePtrOutputWithContext(ctx)
}









type ExampleRequestResponsePtrInput interface {
	pulumi.Input

	ToExampleRequestResponsePtrOutput() ExampleRequestResponsePtrOutput
	ToExampleRequestResponsePtrOutputWithContext(context.Context) ExampleRequestResponsePtrOutput
}

type exampleRequestResponsePtrType ExampleRequestResponseArgs

func ExampleRequestResponsePtr(v *ExampleRequestResponseArgs) ExampleRequestResponsePtrInput {
	return (*exampleRequestResponsePtrType)(v)
}

func (*exampleRequestResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleRequestResponse)(nil)).Elem()
}

func (i *exampleRequestResponsePtrType) ToExampleRequestResponsePtrOutput() ExampleRequestResponsePtrOutput {
	return i.ToExampleRequestResponsePtrOutputWithContext(context.Background())
}

func (i *exampleRequestResponsePtrType) ToExampleRequestResponsePtrOutputWithContext(ctx context.Context) ExampleRequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExampleRequestResponsePtrOutput)
}

type ExampleRequestResponseOutput struct{ *pulumi.OutputState }

func (ExampleRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExampleRequestResponse)(nil)).Elem()
}

func (o ExampleRequestResponseOutput) ToExampleRequestResponseOutput() ExampleRequestResponseOutput {
	return o
}

func (o ExampleRequestResponseOutput) ToExampleRequestResponseOutputWithContext(ctx context.Context) ExampleRequestResponseOutput {
	return o
}

func (o ExampleRequestResponseOutput) ToExampleRequestResponsePtrOutput() ExampleRequestResponsePtrOutput {
	return o.ToExampleRequestResponsePtrOutputWithContext(context.Background())
}

func (o ExampleRequestResponseOutput) ToExampleRequestResponsePtrOutputWithContext(ctx context.Context) ExampleRequestResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExampleRequestResponse) *ExampleRequestResponse {
		return &v
	}).(ExampleRequestResponsePtrOutput)
}

func (o ExampleRequestResponseOutput) GlobalParameters() pulumi.MapOutput {
	return o.ApplyT(func(v ExampleRequestResponse) map[string]interface{} { return v.GlobalParameters }).(pulumi.MapOutput)
}

func (o ExampleRequestResponseOutput) Inputs() pulumi.ArrayArrayMapOutput {
	return o.ApplyT(func(v ExampleRequestResponse) map[string][][]interface{} { return v.Inputs }).(pulumi.ArrayArrayMapOutput)
}

type ExampleRequestResponsePtrOutput struct{ *pulumi.OutputState }

func (ExampleRequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExampleRequestResponse)(nil)).Elem()
}

func (o ExampleRequestResponsePtrOutput) ToExampleRequestResponsePtrOutput() ExampleRequestResponsePtrOutput {
	return o
}

func (o ExampleRequestResponsePtrOutput) ToExampleRequestResponsePtrOutputWithContext(ctx context.Context) ExampleRequestResponsePtrOutput {
	return o
}

func (o ExampleRequestResponsePtrOutput) Elem() ExampleRequestResponseOutput {
	return o.ApplyT(func(v *ExampleRequestResponse) ExampleRequestResponse {
		if v != nil {
			return *v
		}
		var ret ExampleRequestResponse
		return ret
	}).(ExampleRequestResponseOutput)
}

func (o ExampleRequestResponsePtrOutput) GlobalParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *ExampleRequestResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.GlobalParameters
	}).(pulumi.MapOutput)
}

func (o ExampleRequestResponsePtrOutput) Inputs() pulumi.ArrayArrayMapOutput {
	return o.ApplyT(func(v *ExampleRequestResponse) map[string][][]interface{} {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(pulumi.ArrayArrayMapOutput)
}

type GraphEdge struct {
	SourceNodeId *string `pulumi:"sourceNodeId"`
	SourcePortId *string `pulumi:"sourcePortId"`
	TargetNodeId *string `pulumi:"targetNodeId"`
	TargetPortId *string `pulumi:"targetPortId"`
}





type GraphEdgeInput interface {
	pulumi.Input

	ToGraphEdgeOutput() GraphEdgeOutput
	ToGraphEdgeOutputWithContext(context.Context) GraphEdgeOutput
}

type GraphEdgeArgs struct {
	SourceNodeId pulumi.StringPtrInput `pulumi:"sourceNodeId"`
	SourcePortId pulumi.StringPtrInput `pulumi:"sourcePortId"`
	TargetNodeId pulumi.StringPtrInput `pulumi:"targetNodeId"`
	TargetPortId pulumi.StringPtrInput `pulumi:"targetPortId"`
}

func (GraphEdgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphEdge)(nil)).Elem()
}

func (i GraphEdgeArgs) ToGraphEdgeOutput() GraphEdgeOutput {
	return i.ToGraphEdgeOutputWithContext(context.Background())
}

func (i GraphEdgeArgs) ToGraphEdgeOutputWithContext(ctx context.Context) GraphEdgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphEdgeOutput)
}





type GraphEdgeArrayInput interface {
	pulumi.Input

	ToGraphEdgeArrayOutput() GraphEdgeArrayOutput
	ToGraphEdgeArrayOutputWithContext(context.Context) GraphEdgeArrayOutput
}

type GraphEdgeArray []GraphEdgeInput

func (GraphEdgeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphEdge)(nil)).Elem()
}

func (i GraphEdgeArray) ToGraphEdgeArrayOutput() GraphEdgeArrayOutput {
	return i.ToGraphEdgeArrayOutputWithContext(context.Background())
}

func (i GraphEdgeArray) ToGraphEdgeArrayOutputWithContext(ctx context.Context) GraphEdgeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphEdgeArrayOutput)
}

type GraphEdgeOutput struct{ *pulumi.OutputState }

func (GraphEdgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphEdge)(nil)).Elem()
}

func (o GraphEdgeOutput) ToGraphEdgeOutput() GraphEdgeOutput {
	return o
}

func (o GraphEdgeOutput) ToGraphEdgeOutputWithContext(ctx context.Context) GraphEdgeOutput {
	return o
}

func (o GraphEdgeOutput) SourceNodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdge) *string { return v.SourceNodeId }).(pulumi.StringPtrOutput)
}

func (o GraphEdgeOutput) SourcePortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdge) *string { return v.SourcePortId }).(pulumi.StringPtrOutput)
}

func (o GraphEdgeOutput) TargetNodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdge) *string { return v.TargetNodeId }).(pulumi.StringPtrOutput)
}

func (o GraphEdgeOutput) TargetPortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdge) *string { return v.TargetPortId }).(pulumi.StringPtrOutput)
}

type GraphEdgeArrayOutput struct{ *pulumi.OutputState }

func (GraphEdgeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphEdge)(nil)).Elem()
}

func (o GraphEdgeArrayOutput) ToGraphEdgeArrayOutput() GraphEdgeArrayOutput {
	return o
}

func (o GraphEdgeArrayOutput) ToGraphEdgeArrayOutputWithContext(ctx context.Context) GraphEdgeArrayOutput {
	return o
}

func (o GraphEdgeArrayOutput) Index(i pulumi.IntInput) GraphEdgeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GraphEdge {
		return vs[0].([]GraphEdge)[vs[1].(int)]
	}).(GraphEdgeOutput)
}

type GraphEdgeResponse struct {
	SourceNodeId *string `pulumi:"sourceNodeId"`
	SourcePortId *string `pulumi:"sourcePortId"`
	TargetNodeId *string `pulumi:"targetNodeId"`
	TargetPortId *string `pulumi:"targetPortId"`
}





type GraphEdgeResponseInput interface {
	pulumi.Input

	ToGraphEdgeResponseOutput() GraphEdgeResponseOutput
	ToGraphEdgeResponseOutputWithContext(context.Context) GraphEdgeResponseOutput
}

type GraphEdgeResponseArgs struct {
	SourceNodeId pulumi.StringPtrInput `pulumi:"sourceNodeId"`
	SourcePortId pulumi.StringPtrInput `pulumi:"sourcePortId"`
	TargetNodeId pulumi.StringPtrInput `pulumi:"targetNodeId"`
	TargetPortId pulumi.StringPtrInput `pulumi:"targetPortId"`
}

func (GraphEdgeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphEdgeResponse)(nil)).Elem()
}

func (i GraphEdgeResponseArgs) ToGraphEdgeResponseOutput() GraphEdgeResponseOutput {
	return i.ToGraphEdgeResponseOutputWithContext(context.Background())
}

func (i GraphEdgeResponseArgs) ToGraphEdgeResponseOutputWithContext(ctx context.Context) GraphEdgeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphEdgeResponseOutput)
}





type GraphEdgeResponseArrayInput interface {
	pulumi.Input

	ToGraphEdgeResponseArrayOutput() GraphEdgeResponseArrayOutput
	ToGraphEdgeResponseArrayOutputWithContext(context.Context) GraphEdgeResponseArrayOutput
}

type GraphEdgeResponseArray []GraphEdgeResponseInput

func (GraphEdgeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphEdgeResponse)(nil)).Elem()
}

func (i GraphEdgeResponseArray) ToGraphEdgeResponseArrayOutput() GraphEdgeResponseArrayOutput {
	return i.ToGraphEdgeResponseArrayOutputWithContext(context.Background())
}

func (i GraphEdgeResponseArray) ToGraphEdgeResponseArrayOutputWithContext(ctx context.Context) GraphEdgeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphEdgeResponseArrayOutput)
}

type GraphEdgeResponseOutput struct{ *pulumi.OutputState }

func (GraphEdgeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphEdgeResponse)(nil)).Elem()
}

func (o GraphEdgeResponseOutput) ToGraphEdgeResponseOutput() GraphEdgeResponseOutput {
	return o
}

func (o GraphEdgeResponseOutput) ToGraphEdgeResponseOutputWithContext(ctx context.Context) GraphEdgeResponseOutput {
	return o
}

func (o GraphEdgeResponseOutput) SourceNodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdgeResponse) *string { return v.SourceNodeId }).(pulumi.StringPtrOutput)
}

func (o GraphEdgeResponseOutput) SourcePortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdgeResponse) *string { return v.SourcePortId }).(pulumi.StringPtrOutput)
}

func (o GraphEdgeResponseOutput) TargetNodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdgeResponse) *string { return v.TargetNodeId }).(pulumi.StringPtrOutput)
}

func (o GraphEdgeResponseOutput) TargetPortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphEdgeResponse) *string { return v.TargetPortId }).(pulumi.StringPtrOutput)
}

type GraphEdgeResponseArrayOutput struct{ *pulumi.OutputState }

func (GraphEdgeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphEdgeResponse)(nil)).Elem()
}

func (o GraphEdgeResponseArrayOutput) ToGraphEdgeResponseArrayOutput() GraphEdgeResponseArrayOutput {
	return o
}

func (o GraphEdgeResponseArrayOutput) ToGraphEdgeResponseArrayOutputWithContext(ctx context.Context) GraphEdgeResponseArrayOutput {
	return o
}

func (o GraphEdgeResponseArrayOutput) Index(i pulumi.IntInput) GraphEdgeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GraphEdgeResponse {
		return vs[0].([]GraphEdgeResponse)[vs[1].(int)]
	}).(GraphEdgeResponseOutput)
}

type GraphNode struct {
	AssetId    *string                        `pulumi:"assetId"`
	InputId    *string                        `pulumi:"inputId"`
	OutputId   *string                        `pulumi:"outputId"`
	Parameters map[string]WebServiceParameter `pulumi:"parameters"`
}





type GraphNodeInput interface {
	pulumi.Input

	ToGraphNodeOutput() GraphNodeOutput
	ToGraphNodeOutputWithContext(context.Context) GraphNodeOutput
}

type GraphNodeArgs struct {
	AssetId    pulumi.StringPtrInput       `pulumi:"assetId"`
	InputId    pulumi.StringPtrInput       `pulumi:"inputId"`
	OutputId   pulumi.StringPtrInput       `pulumi:"outputId"`
	Parameters WebServiceParameterMapInput `pulumi:"parameters"`
}

func (GraphNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphNode)(nil)).Elem()
}

func (i GraphNodeArgs) ToGraphNodeOutput() GraphNodeOutput {
	return i.ToGraphNodeOutputWithContext(context.Background())
}

func (i GraphNodeArgs) ToGraphNodeOutputWithContext(ctx context.Context) GraphNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphNodeOutput)
}





type GraphNodeMapInput interface {
	pulumi.Input

	ToGraphNodeMapOutput() GraphNodeMapOutput
	ToGraphNodeMapOutputWithContext(context.Context) GraphNodeMapOutput
}

type GraphNodeMap map[string]GraphNodeInput

func (GraphNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphNode)(nil)).Elem()
}

func (i GraphNodeMap) ToGraphNodeMapOutput() GraphNodeMapOutput {
	return i.ToGraphNodeMapOutputWithContext(context.Background())
}

func (i GraphNodeMap) ToGraphNodeMapOutputWithContext(ctx context.Context) GraphNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphNodeMapOutput)
}

type GraphNodeOutput struct{ *pulumi.OutputState }

func (GraphNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphNode)(nil)).Elem()
}

func (o GraphNodeOutput) ToGraphNodeOutput() GraphNodeOutput {
	return o
}

func (o GraphNodeOutput) ToGraphNodeOutputWithContext(ctx context.Context) GraphNodeOutput {
	return o
}

func (o GraphNodeOutput) AssetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphNode) *string { return v.AssetId }).(pulumi.StringPtrOutput)
}

func (o GraphNodeOutput) InputId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphNode) *string { return v.InputId }).(pulumi.StringPtrOutput)
}

func (o GraphNodeOutput) OutputId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphNode) *string { return v.OutputId }).(pulumi.StringPtrOutput)
}

func (o GraphNodeOutput) Parameters() WebServiceParameterMapOutput {
	return o.ApplyT(func(v GraphNode) map[string]WebServiceParameter { return v.Parameters }).(WebServiceParameterMapOutput)
}

type GraphNodeMapOutput struct{ *pulumi.OutputState }

func (GraphNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphNode)(nil)).Elem()
}

func (o GraphNodeMapOutput) ToGraphNodeMapOutput() GraphNodeMapOutput {
	return o
}

func (o GraphNodeMapOutput) ToGraphNodeMapOutputWithContext(ctx context.Context) GraphNodeMapOutput {
	return o
}

func (o GraphNodeMapOutput) MapIndex(k pulumi.StringInput) GraphNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GraphNode {
		return vs[0].(map[string]GraphNode)[vs[1].(string)]
	}).(GraphNodeOutput)
}

type GraphNodeResponse struct {
	AssetId    *string                                `pulumi:"assetId"`
	InputId    *string                                `pulumi:"inputId"`
	OutputId   *string                                `pulumi:"outputId"`
	Parameters map[string]WebServiceParameterResponse `pulumi:"parameters"`
}





type GraphNodeResponseInput interface {
	pulumi.Input

	ToGraphNodeResponseOutput() GraphNodeResponseOutput
	ToGraphNodeResponseOutputWithContext(context.Context) GraphNodeResponseOutput
}

type GraphNodeResponseArgs struct {
	AssetId    pulumi.StringPtrInput               `pulumi:"assetId"`
	InputId    pulumi.StringPtrInput               `pulumi:"inputId"`
	OutputId   pulumi.StringPtrInput               `pulumi:"outputId"`
	Parameters WebServiceParameterResponseMapInput `pulumi:"parameters"`
}

func (GraphNodeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphNodeResponse)(nil)).Elem()
}

func (i GraphNodeResponseArgs) ToGraphNodeResponseOutput() GraphNodeResponseOutput {
	return i.ToGraphNodeResponseOutputWithContext(context.Background())
}

func (i GraphNodeResponseArgs) ToGraphNodeResponseOutputWithContext(ctx context.Context) GraphNodeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphNodeResponseOutput)
}





type GraphNodeResponseMapInput interface {
	pulumi.Input

	ToGraphNodeResponseMapOutput() GraphNodeResponseMapOutput
	ToGraphNodeResponseMapOutputWithContext(context.Context) GraphNodeResponseMapOutput
}

type GraphNodeResponseMap map[string]GraphNodeResponseInput

func (GraphNodeResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphNodeResponse)(nil)).Elem()
}

func (i GraphNodeResponseMap) ToGraphNodeResponseMapOutput() GraphNodeResponseMapOutput {
	return i.ToGraphNodeResponseMapOutputWithContext(context.Background())
}

func (i GraphNodeResponseMap) ToGraphNodeResponseMapOutputWithContext(ctx context.Context) GraphNodeResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphNodeResponseMapOutput)
}

type GraphNodeResponseOutput struct{ *pulumi.OutputState }

func (GraphNodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphNodeResponse)(nil)).Elem()
}

func (o GraphNodeResponseOutput) ToGraphNodeResponseOutput() GraphNodeResponseOutput {
	return o
}

func (o GraphNodeResponseOutput) ToGraphNodeResponseOutputWithContext(ctx context.Context) GraphNodeResponseOutput {
	return o
}

func (o GraphNodeResponseOutput) AssetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphNodeResponse) *string { return v.AssetId }).(pulumi.StringPtrOutput)
}

func (o GraphNodeResponseOutput) InputId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphNodeResponse) *string { return v.InputId }).(pulumi.StringPtrOutput)
}

func (o GraphNodeResponseOutput) OutputId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphNodeResponse) *string { return v.OutputId }).(pulumi.StringPtrOutput)
}

func (o GraphNodeResponseOutput) Parameters() WebServiceParameterResponseMapOutput {
	return o.ApplyT(func(v GraphNodeResponse) map[string]WebServiceParameterResponse { return v.Parameters }).(WebServiceParameterResponseMapOutput)
}

type GraphNodeResponseMapOutput struct{ *pulumi.OutputState }

func (GraphNodeResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphNodeResponse)(nil)).Elem()
}

func (o GraphNodeResponseMapOutput) ToGraphNodeResponseMapOutput() GraphNodeResponseMapOutput {
	return o
}

func (o GraphNodeResponseMapOutput) ToGraphNodeResponseMapOutputWithContext(ctx context.Context) GraphNodeResponseMapOutput {
	return o
}

func (o GraphNodeResponseMapOutput) MapIndex(k pulumi.StringInput) GraphNodeResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GraphNodeResponse {
		return vs[0].(map[string]GraphNodeResponse)[vs[1].(string)]
	}).(GraphNodeResponseOutput)
}

type GraphPackage struct {
	Edges           []GraphEdge               `pulumi:"edges"`
	GraphParameters map[string]GraphParameter `pulumi:"graphParameters"`
	Nodes           map[string]GraphNode      `pulumi:"nodes"`
}





type GraphPackageInput interface {
	pulumi.Input

	ToGraphPackageOutput() GraphPackageOutput
	ToGraphPackageOutputWithContext(context.Context) GraphPackageOutput
}

type GraphPackageArgs struct {
	Edges           GraphEdgeArrayInput    `pulumi:"edges"`
	GraphParameters GraphParameterMapInput `pulumi:"graphParameters"`
	Nodes           GraphNodeMapInput      `pulumi:"nodes"`
}

func (GraphPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphPackage)(nil)).Elem()
}

func (i GraphPackageArgs) ToGraphPackageOutput() GraphPackageOutput {
	return i.ToGraphPackageOutputWithContext(context.Background())
}

func (i GraphPackageArgs) ToGraphPackageOutputWithContext(ctx context.Context) GraphPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphPackageOutput)
}

func (i GraphPackageArgs) ToGraphPackagePtrOutput() GraphPackagePtrOutput {
	return i.ToGraphPackagePtrOutputWithContext(context.Background())
}

func (i GraphPackageArgs) ToGraphPackagePtrOutputWithContext(ctx context.Context) GraphPackagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphPackageOutput).ToGraphPackagePtrOutputWithContext(ctx)
}









type GraphPackagePtrInput interface {
	pulumi.Input

	ToGraphPackagePtrOutput() GraphPackagePtrOutput
	ToGraphPackagePtrOutputWithContext(context.Context) GraphPackagePtrOutput
}

type graphPackagePtrType GraphPackageArgs

func GraphPackagePtr(v *GraphPackageArgs) GraphPackagePtrInput {
	return (*graphPackagePtrType)(v)
}

func (*graphPackagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphPackage)(nil)).Elem()
}

func (i *graphPackagePtrType) ToGraphPackagePtrOutput() GraphPackagePtrOutput {
	return i.ToGraphPackagePtrOutputWithContext(context.Background())
}

func (i *graphPackagePtrType) ToGraphPackagePtrOutputWithContext(ctx context.Context) GraphPackagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphPackagePtrOutput)
}

type GraphPackageOutput struct{ *pulumi.OutputState }

func (GraphPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphPackage)(nil)).Elem()
}

func (o GraphPackageOutput) ToGraphPackageOutput() GraphPackageOutput {
	return o
}

func (o GraphPackageOutput) ToGraphPackageOutputWithContext(ctx context.Context) GraphPackageOutput {
	return o
}

func (o GraphPackageOutput) ToGraphPackagePtrOutput() GraphPackagePtrOutput {
	return o.ToGraphPackagePtrOutputWithContext(context.Background())
}

func (o GraphPackageOutput) ToGraphPackagePtrOutputWithContext(ctx context.Context) GraphPackagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GraphPackage) *GraphPackage {
		return &v
	}).(GraphPackagePtrOutput)
}

func (o GraphPackageOutput) Edges() GraphEdgeArrayOutput {
	return o.ApplyT(func(v GraphPackage) []GraphEdge { return v.Edges }).(GraphEdgeArrayOutput)
}

func (o GraphPackageOutput) GraphParameters() GraphParameterMapOutput {
	return o.ApplyT(func(v GraphPackage) map[string]GraphParameter { return v.GraphParameters }).(GraphParameterMapOutput)
}

func (o GraphPackageOutput) Nodes() GraphNodeMapOutput {
	return o.ApplyT(func(v GraphPackage) map[string]GraphNode { return v.Nodes }).(GraphNodeMapOutput)
}

type GraphPackagePtrOutput struct{ *pulumi.OutputState }

func (GraphPackagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphPackage)(nil)).Elem()
}

func (o GraphPackagePtrOutput) ToGraphPackagePtrOutput() GraphPackagePtrOutput {
	return o
}

func (o GraphPackagePtrOutput) ToGraphPackagePtrOutputWithContext(ctx context.Context) GraphPackagePtrOutput {
	return o
}

func (o GraphPackagePtrOutput) Elem() GraphPackageOutput {
	return o.ApplyT(func(v *GraphPackage) GraphPackage {
		if v != nil {
			return *v
		}
		var ret GraphPackage
		return ret
	}).(GraphPackageOutput)
}

func (o GraphPackagePtrOutput) Edges() GraphEdgeArrayOutput {
	return o.ApplyT(func(v *GraphPackage) []GraphEdge {
		if v == nil {
			return nil
		}
		return v.Edges
	}).(GraphEdgeArrayOutput)
}

func (o GraphPackagePtrOutput) GraphParameters() GraphParameterMapOutput {
	return o.ApplyT(func(v *GraphPackage) map[string]GraphParameter {
		if v == nil {
			return nil
		}
		return v.GraphParameters
	}).(GraphParameterMapOutput)
}

func (o GraphPackagePtrOutput) Nodes() GraphNodeMapOutput {
	return o.ApplyT(func(v *GraphPackage) map[string]GraphNode {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(GraphNodeMapOutput)
}

type GraphPackageResponse struct {
	Edges           []GraphEdgeResponse               `pulumi:"edges"`
	GraphParameters map[string]GraphParameterResponse `pulumi:"graphParameters"`
	Nodes           map[string]GraphNodeResponse      `pulumi:"nodes"`
}





type GraphPackageResponseInput interface {
	pulumi.Input

	ToGraphPackageResponseOutput() GraphPackageResponseOutput
	ToGraphPackageResponseOutputWithContext(context.Context) GraphPackageResponseOutput
}

type GraphPackageResponseArgs struct {
	Edges           GraphEdgeResponseArrayInput    `pulumi:"edges"`
	GraphParameters GraphParameterResponseMapInput `pulumi:"graphParameters"`
	Nodes           GraphNodeResponseMapInput      `pulumi:"nodes"`
}

func (GraphPackageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphPackageResponse)(nil)).Elem()
}

func (i GraphPackageResponseArgs) ToGraphPackageResponseOutput() GraphPackageResponseOutput {
	return i.ToGraphPackageResponseOutputWithContext(context.Background())
}

func (i GraphPackageResponseArgs) ToGraphPackageResponseOutputWithContext(ctx context.Context) GraphPackageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphPackageResponseOutput)
}

func (i GraphPackageResponseArgs) ToGraphPackageResponsePtrOutput() GraphPackageResponsePtrOutput {
	return i.ToGraphPackageResponsePtrOutputWithContext(context.Background())
}

func (i GraphPackageResponseArgs) ToGraphPackageResponsePtrOutputWithContext(ctx context.Context) GraphPackageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphPackageResponseOutput).ToGraphPackageResponsePtrOutputWithContext(ctx)
}









type GraphPackageResponsePtrInput interface {
	pulumi.Input

	ToGraphPackageResponsePtrOutput() GraphPackageResponsePtrOutput
	ToGraphPackageResponsePtrOutputWithContext(context.Context) GraphPackageResponsePtrOutput
}

type graphPackageResponsePtrType GraphPackageResponseArgs

func GraphPackageResponsePtr(v *GraphPackageResponseArgs) GraphPackageResponsePtrInput {
	return (*graphPackageResponsePtrType)(v)
}

func (*graphPackageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphPackageResponse)(nil)).Elem()
}

func (i *graphPackageResponsePtrType) ToGraphPackageResponsePtrOutput() GraphPackageResponsePtrOutput {
	return i.ToGraphPackageResponsePtrOutputWithContext(context.Background())
}

func (i *graphPackageResponsePtrType) ToGraphPackageResponsePtrOutputWithContext(ctx context.Context) GraphPackageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphPackageResponsePtrOutput)
}

type GraphPackageResponseOutput struct{ *pulumi.OutputState }

func (GraphPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphPackageResponse)(nil)).Elem()
}

func (o GraphPackageResponseOutput) ToGraphPackageResponseOutput() GraphPackageResponseOutput {
	return o
}

func (o GraphPackageResponseOutput) ToGraphPackageResponseOutputWithContext(ctx context.Context) GraphPackageResponseOutput {
	return o
}

func (o GraphPackageResponseOutput) ToGraphPackageResponsePtrOutput() GraphPackageResponsePtrOutput {
	return o.ToGraphPackageResponsePtrOutputWithContext(context.Background())
}

func (o GraphPackageResponseOutput) ToGraphPackageResponsePtrOutputWithContext(ctx context.Context) GraphPackageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GraphPackageResponse) *GraphPackageResponse {
		return &v
	}).(GraphPackageResponsePtrOutput)
}

func (o GraphPackageResponseOutput) Edges() GraphEdgeResponseArrayOutput {
	return o.ApplyT(func(v GraphPackageResponse) []GraphEdgeResponse { return v.Edges }).(GraphEdgeResponseArrayOutput)
}

func (o GraphPackageResponseOutput) GraphParameters() GraphParameterResponseMapOutput {
	return o.ApplyT(func(v GraphPackageResponse) map[string]GraphParameterResponse { return v.GraphParameters }).(GraphParameterResponseMapOutput)
}

func (o GraphPackageResponseOutput) Nodes() GraphNodeResponseMapOutput {
	return o.ApplyT(func(v GraphPackageResponse) map[string]GraphNodeResponse { return v.Nodes }).(GraphNodeResponseMapOutput)
}

type GraphPackageResponsePtrOutput struct{ *pulumi.OutputState }

func (GraphPackageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphPackageResponse)(nil)).Elem()
}

func (o GraphPackageResponsePtrOutput) ToGraphPackageResponsePtrOutput() GraphPackageResponsePtrOutput {
	return o
}

func (o GraphPackageResponsePtrOutput) ToGraphPackageResponsePtrOutputWithContext(ctx context.Context) GraphPackageResponsePtrOutput {
	return o
}

func (o GraphPackageResponsePtrOutput) Elem() GraphPackageResponseOutput {
	return o.ApplyT(func(v *GraphPackageResponse) GraphPackageResponse {
		if v != nil {
			return *v
		}
		var ret GraphPackageResponse
		return ret
	}).(GraphPackageResponseOutput)
}

func (o GraphPackageResponsePtrOutput) Edges() GraphEdgeResponseArrayOutput {
	return o.ApplyT(func(v *GraphPackageResponse) []GraphEdgeResponse {
		if v == nil {
			return nil
		}
		return v.Edges
	}).(GraphEdgeResponseArrayOutput)
}

func (o GraphPackageResponsePtrOutput) GraphParameters() GraphParameterResponseMapOutput {
	return o.ApplyT(func(v *GraphPackageResponse) map[string]GraphParameterResponse {
		if v == nil {
			return nil
		}
		return v.GraphParameters
	}).(GraphParameterResponseMapOutput)
}

func (o GraphPackageResponsePtrOutput) Nodes() GraphNodeResponseMapOutput {
	return o.ApplyT(func(v *GraphPackageResponse) map[string]GraphNodeResponse {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(GraphNodeResponseMapOutput)
}

type GraphParameter struct {
	Description *string              `pulumi:"description"`
	Links       []GraphParameterLink `pulumi:"links"`
	Type        string               `pulumi:"type"`
}





type GraphParameterInput interface {
	pulumi.Input

	ToGraphParameterOutput() GraphParameterOutput
	ToGraphParameterOutputWithContext(context.Context) GraphParameterOutput
}

type GraphParameterArgs struct {
	Description pulumi.StringPtrInput        `pulumi:"description"`
	Links       GraphParameterLinkArrayInput `pulumi:"links"`
	Type        pulumi.StringInput           `pulumi:"type"`
}

func (GraphParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameter)(nil)).Elem()
}

func (i GraphParameterArgs) ToGraphParameterOutput() GraphParameterOutput {
	return i.ToGraphParameterOutputWithContext(context.Background())
}

func (i GraphParameterArgs) ToGraphParameterOutputWithContext(ctx context.Context) GraphParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterOutput)
}





type GraphParameterMapInput interface {
	pulumi.Input

	ToGraphParameterMapOutput() GraphParameterMapOutput
	ToGraphParameterMapOutputWithContext(context.Context) GraphParameterMapOutput
}

type GraphParameterMap map[string]GraphParameterInput

func (GraphParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphParameter)(nil)).Elem()
}

func (i GraphParameterMap) ToGraphParameterMapOutput() GraphParameterMapOutput {
	return i.ToGraphParameterMapOutputWithContext(context.Background())
}

func (i GraphParameterMap) ToGraphParameterMapOutputWithContext(ctx context.Context) GraphParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterMapOutput)
}

type GraphParameterOutput struct{ *pulumi.OutputState }

func (GraphParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameter)(nil)).Elem()
}

func (o GraphParameterOutput) ToGraphParameterOutput() GraphParameterOutput {
	return o
}

func (o GraphParameterOutput) ToGraphParameterOutputWithContext(ctx context.Context) GraphParameterOutput {
	return o
}

func (o GraphParameterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphParameter) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GraphParameterOutput) Links() GraphParameterLinkArrayOutput {
	return o.ApplyT(func(v GraphParameter) []GraphParameterLink { return v.Links }).(GraphParameterLinkArrayOutput)
}

func (o GraphParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GraphParameter) string { return v.Type }).(pulumi.StringOutput)
}

type GraphParameterMapOutput struct{ *pulumi.OutputState }

func (GraphParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphParameter)(nil)).Elem()
}

func (o GraphParameterMapOutput) ToGraphParameterMapOutput() GraphParameterMapOutput {
	return o
}

func (o GraphParameterMapOutput) ToGraphParameterMapOutputWithContext(ctx context.Context) GraphParameterMapOutput {
	return o
}

func (o GraphParameterMapOutput) MapIndex(k pulumi.StringInput) GraphParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GraphParameter {
		return vs[0].(map[string]GraphParameter)[vs[1].(string)]
	}).(GraphParameterOutput)
}

type GraphParameterLink struct {
	NodeId       string `pulumi:"nodeId"`
	ParameterKey string `pulumi:"parameterKey"`
}





type GraphParameterLinkInput interface {
	pulumi.Input

	ToGraphParameterLinkOutput() GraphParameterLinkOutput
	ToGraphParameterLinkOutputWithContext(context.Context) GraphParameterLinkOutput
}

type GraphParameterLinkArgs struct {
	NodeId       pulumi.StringInput `pulumi:"nodeId"`
	ParameterKey pulumi.StringInput `pulumi:"parameterKey"`
}

func (GraphParameterLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameterLink)(nil)).Elem()
}

func (i GraphParameterLinkArgs) ToGraphParameterLinkOutput() GraphParameterLinkOutput {
	return i.ToGraphParameterLinkOutputWithContext(context.Background())
}

func (i GraphParameterLinkArgs) ToGraphParameterLinkOutputWithContext(ctx context.Context) GraphParameterLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterLinkOutput)
}





type GraphParameterLinkArrayInput interface {
	pulumi.Input

	ToGraphParameterLinkArrayOutput() GraphParameterLinkArrayOutput
	ToGraphParameterLinkArrayOutputWithContext(context.Context) GraphParameterLinkArrayOutput
}

type GraphParameterLinkArray []GraphParameterLinkInput

func (GraphParameterLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphParameterLink)(nil)).Elem()
}

func (i GraphParameterLinkArray) ToGraphParameterLinkArrayOutput() GraphParameterLinkArrayOutput {
	return i.ToGraphParameterLinkArrayOutputWithContext(context.Background())
}

func (i GraphParameterLinkArray) ToGraphParameterLinkArrayOutputWithContext(ctx context.Context) GraphParameterLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterLinkArrayOutput)
}

type GraphParameterLinkOutput struct{ *pulumi.OutputState }

func (GraphParameterLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameterLink)(nil)).Elem()
}

func (o GraphParameterLinkOutput) ToGraphParameterLinkOutput() GraphParameterLinkOutput {
	return o
}

func (o GraphParameterLinkOutput) ToGraphParameterLinkOutputWithContext(ctx context.Context) GraphParameterLinkOutput {
	return o
}

func (o GraphParameterLinkOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GraphParameterLink) string { return v.NodeId }).(pulumi.StringOutput)
}

func (o GraphParameterLinkOutput) ParameterKey() pulumi.StringOutput {
	return o.ApplyT(func(v GraphParameterLink) string { return v.ParameterKey }).(pulumi.StringOutput)
}

type GraphParameterLinkArrayOutput struct{ *pulumi.OutputState }

func (GraphParameterLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphParameterLink)(nil)).Elem()
}

func (o GraphParameterLinkArrayOutput) ToGraphParameterLinkArrayOutput() GraphParameterLinkArrayOutput {
	return o
}

func (o GraphParameterLinkArrayOutput) ToGraphParameterLinkArrayOutputWithContext(ctx context.Context) GraphParameterLinkArrayOutput {
	return o
}

func (o GraphParameterLinkArrayOutput) Index(i pulumi.IntInput) GraphParameterLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GraphParameterLink {
		return vs[0].([]GraphParameterLink)[vs[1].(int)]
	}).(GraphParameterLinkOutput)
}

type GraphParameterLinkResponse struct {
	NodeId       string `pulumi:"nodeId"`
	ParameterKey string `pulumi:"parameterKey"`
}





type GraphParameterLinkResponseInput interface {
	pulumi.Input

	ToGraphParameterLinkResponseOutput() GraphParameterLinkResponseOutput
	ToGraphParameterLinkResponseOutputWithContext(context.Context) GraphParameterLinkResponseOutput
}

type GraphParameterLinkResponseArgs struct {
	NodeId       pulumi.StringInput `pulumi:"nodeId"`
	ParameterKey pulumi.StringInput `pulumi:"parameterKey"`
}

func (GraphParameterLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameterLinkResponse)(nil)).Elem()
}

func (i GraphParameterLinkResponseArgs) ToGraphParameterLinkResponseOutput() GraphParameterLinkResponseOutput {
	return i.ToGraphParameterLinkResponseOutputWithContext(context.Background())
}

func (i GraphParameterLinkResponseArgs) ToGraphParameterLinkResponseOutputWithContext(ctx context.Context) GraphParameterLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterLinkResponseOutput)
}





type GraphParameterLinkResponseArrayInput interface {
	pulumi.Input

	ToGraphParameterLinkResponseArrayOutput() GraphParameterLinkResponseArrayOutput
	ToGraphParameterLinkResponseArrayOutputWithContext(context.Context) GraphParameterLinkResponseArrayOutput
}

type GraphParameterLinkResponseArray []GraphParameterLinkResponseInput

func (GraphParameterLinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphParameterLinkResponse)(nil)).Elem()
}

func (i GraphParameterLinkResponseArray) ToGraphParameterLinkResponseArrayOutput() GraphParameterLinkResponseArrayOutput {
	return i.ToGraphParameterLinkResponseArrayOutputWithContext(context.Background())
}

func (i GraphParameterLinkResponseArray) ToGraphParameterLinkResponseArrayOutputWithContext(ctx context.Context) GraphParameterLinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterLinkResponseArrayOutput)
}

type GraphParameterLinkResponseOutput struct{ *pulumi.OutputState }

func (GraphParameterLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameterLinkResponse)(nil)).Elem()
}

func (o GraphParameterLinkResponseOutput) ToGraphParameterLinkResponseOutput() GraphParameterLinkResponseOutput {
	return o
}

func (o GraphParameterLinkResponseOutput) ToGraphParameterLinkResponseOutputWithContext(ctx context.Context) GraphParameterLinkResponseOutput {
	return o
}

func (o GraphParameterLinkResponseOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GraphParameterLinkResponse) string { return v.NodeId }).(pulumi.StringOutput)
}

func (o GraphParameterLinkResponseOutput) ParameterKey() pulumi.StringOutput {
	return o.ApplyT(func(v GraphParameterLinkResponse) string { return v.ParameterKey }).(pulumi.StringOutput)
}

type GraphParameterLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (GraphParameterLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphParameterLinkResponse)(nil)).Elem()
}

func (o GraphParameterLinkResponseArrayOutput) ToGraphParameterLinkResponseArrayOutput() GraphParameterLinkResponseArrayOutput {
	return o
}

func (o GraphParameterLinkResponseArrayOutput) ToGraphParameterLinkResponseArrayOutputWithContext(ctx context.Context) GraphParameterLinkResponseArrayOutput {
	return o
}

func (o GraphParameterLinkResponseArrayOutput) Index(i pulumi.IntInput) GraphParameterLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GraphParameterLinkResponse {
		return vs[0].([]GraphParameterLinkResponse)[vs[1].(int)]
	}).(GraphParameterLinkResponseOutput)
}

type GraphParameterResponse struct {
	Description *string                      `pulumi:"description"`
	Links       []GraphParameterLinkResponse `pulumi:"links"`
	Type        string                       `pulumi:"type"`
}





type GraphParameterResponseInput interface {
	pulumi.Input

	ToGraphParameterResponseOutput() GraphParameterResponseOutput
	ToGraphParameterResponseOutputWithContext(context.Context) GraphParameterResponseOutput
}

type GraphParameterResponseArgs struct {
	Description pulumi.StringPtrInput                `pulumi:"description"`
	Links       GraphParameterLinkResponseArrayInput `pulumi:"links"`
	Type        pulumi.StringInput                   `pulumi:"type"`
}

func (GraphParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameterResponse)(nil)).Elem()
}

func (i GraphParameterResponseArgs) ToGraphParameterResponseOutput() GraphParameterResponseOutput {
	return i.ToGraphParameterResponseOutputWithContext(context.Background())
}

func (i GraphParameterResponseArgs) ToGraphParameterResponseOutputWithContext(ctx context.Context) GraphParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterResponseOutput)
}





type GraphParameterResponseMapInput interface {
	pulumi.Input

	ToGraphParameterResponseMapOutput() GraphParameterResponseMapOutput
	ToGraphParameterResponseMapOutputWithContext(context.Context) GraphParameterResponseMapOutput
}

type GraphParameterResponseMap map[string]GraphParameterResponseInput

func (GraphParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphParameterResponse)(nil)).Elem()
}

func (i GraphParameterResponseMap) ToGraphParameterResponseMapOutput() GraphParameterResponseMapOutput {
	return i.ToGraphParameterResponseMapOutputWithContext(context.Background())
}

func (i GraphParameterResponseMap) ToGraphParameterResponseMapOutputWithContext(ctx context.Context) GraphParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphParameterResponseMapOutput)
}

type GraphParameterResponseOutput struct{ *pulumi.OutputState }

func (GraphParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphParameterResponse)(nil)).Elem()
}

func (o GraphParameterResponseOutput) ToGraphParameterResponseOutput() GraphParameterResponseOutput {
	return o
}

func (o GraphParameterResponseOutput) ToGraphParameterResponseOutputWithContext(ctx context.Context) GraphParameterResponseOutput {
	return o
}

func (o GraphParameterResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphParameterResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GraphParameterResponseOutput) Links() GraphParameterLinkResponseArrayOutput {
	return o.ApplyT(func(v GraphParameterResponse) []GraphParameterLinkResponse { return v.Links }).(GraphParameterLinkResponseArrayOutput)
}

func (o GraphParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GraphParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

type GraphParameterResponseMapOutput struct{ *pulumi.OutputState }

func (GraphParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphParameterResponse)(nil)).Elem()
}

func (o GraphParameterResponseMapOutput) ToGraphParameterResponseMapOutput() GraphParameterResponseMapOutput {
	return o
}

func (o GraphParameterResponseMapOutput) ToGraphParameterResponseMapOutputWithContext(ctx context.Context) GraphParameterResponseMapOutput {
	return o
}

func (o GraphParameterResponseMapOutput) MapIndex(k pulumi.StringInput) GraphParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GraphParameterResponse {
		return vs[0].(map[string]GraphParameterResponse)[vs[1].(string)]
	}).(GraphParameterResponseOutput)
}

type InputPort struct {
	Type *string `pulumi:"type"`
}





type InputPortInput interface {
	pulumi.Input

	ToInputPortOutput() InputPortOutput
	ToInputPortOutputWithContext(context.Context) InputPortOutput
}

type InputPortArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (InputPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPort)(nil)).Elem()
}

func (i InputPortArgs) ToInputPortOutput() InputPortOutput {
	return i.ToInputPortOutputWithContext(context.Background())
}

func (i InputPortArgs) ToInputPortOutputWithContext(ctx context.Context) InputPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPortOutput)
}





type InputPortMapInput interface {
	pulumi.Input

	ToInputPortMapOutput() InputPortMapOutput
	ToInputPortMapOutputWithContext(context.Context) InputPortMapOutput
}

type InputPortMap map[string]InputPortInput

func (InputPortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputPort)(nil)).Elem()
}

func (i InputPortMap) ToInputPortMapOutput() InputPortMapOutput {
	return i.ToInputPortMapOutputWithContext(context.Background())
}

func (i InputPortMap) ToInputPortMapOutputWithContext(ctx context.Context) InputPortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPortMapOutput)
}

type InputPortOutput struct{ *pulumi.OutputState }

func (InputPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPort)(nil)).Elem()
}

func (o InputPortOutput) ToInputPortOutput() InputPortOutput {
	return o
}

func (o InputPortOutput) ToInputPortOutputWithContext(ctx context.Context) InputPortOutput {
	return o
}

func (o InputPortOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputPort) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type InputPortMapOutput struct{ *pulumi.OutputState }

func (InputPortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputPort)(nil)).Elem()
}

func (o InputPortMapOutput) ToInputPortMapOutput() InputPortMapOutput {
	return o
}

func (o InputPortMapOutput) ToInputPortMapOutputWithContext(ctx context.Context) InputPortMapOutput {
	return o
}

func (o InputPortMapOutput) MapIndex(k pulumi.StringInput) InputPortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InputPort {
		return vs[0].(map[string]InputPort)[vs[1].(string)]
	}).(InputPortOutput)
}

type InputPortResponse struct {
	Type *string `pulumi:"type"`
}





type InputPortResponseInput interface {
	pulumi.Input

	ToInputPortResponseOutput() InputPortResponseOutput
	ToInputPortResponseOutputWithContext(context.Context) InputPortResponseOutput
}

type InputPortResponseArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (InputPortResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPortResponse)(nil)).Elem()
}

func (i InputPortResponseArgs) ToInputPortResponseOutput() InputPortResponseOutput {
	return i.ToInputPortResponseOutputWithContext(context.Background())
}

func (i InputPortResponseArgs) ToInputPortResponseOutputWithContext(ctx context.Context) InputPortResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPortResponseOutput)
}





type InputPortResponseMapInput interface {
	pulumi.Input

	ToInputPortResponseMapOutput() InputPortResponseMapOutput
	ToInputPortResponseMapOutputWithContext(context.Context) InputPortResponseMapOutput
}

type InputPortResponseMap map[string]InputPortResponseInput

func (InputPortResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputPortResponse)(nil)).Elem()
}

func (i InputPortResponseMap) ToInputPortResponseMapOutput() InputPortResponseMapOutput {
	return i.ToInputPortResponseMapOutputWithContext(context.Background())
}

func (i InputPortResponseMap) ToInputPortResponseMapOutputWithContext(ctx context.Context) InputPortResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPortResponseMapOutput)
}

type InputPortResponseOutput struct{ *pulumi.OutputState }

func (InputPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPortResponse)(nil)).Elem()
}

func (o InputPortResponseOutput) ToInputPortResponseOutput() InputPortResponseOutput {
	return o
}

func (o InputPortResponseOutput) ToInputPortResponseOutputWithContext(ctx context.Context) InputPortResponseOutput {
	return o
}

func (o InputPortResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputPortResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type InputPortResponseMapOutput struct{ *pulumi.OutputState }

func (InputPortResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InputPortResponse)(nil)).Elem()
}

func (o InputPortResponseMapOutput) ToInputPortResponseMapOutput() InputPortResponseMapOutput {
	return o
}

func (o InputPortResponseMapOutput) ToInputPortResponseMapOutputWithContext(ctx context.Context) InputPortResponseMapOutput {
	return o
}

func (o InputPortResponseMapOutput) MapIndex(k pulumi.StringInput) InputPortResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InputPortResponse {
		return vs[0].(map[string]InputPortResponse)[vs[1].(string)]
	}).(InputPortResponseOutput)
}

type MachineLearningWorkspace struct {
	Id string `pulumi:"id"`
}





type MachineLearningWorkspaceInput interface {
	pulumi.Input

	ToMachineLearningWorkspaceOutput() MachineLearningWorkspaceOutput
	ToMachineLearningWorkspaceOutputWithContext(context.Context) MachineLearningWorkspaceOutput
}

type MachineLearningWorkspaceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MachineLearningWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningWorkspace)(nil)).Elem()
}

func (i MachineLearningWorkspaceArgs) ToMachineLearningWorkspaceOutput() MachineLearningWorkspaceOutput {
	return i.ToMachineLearningWorkspaceOutputWithContext(context.Background())
}

func (i MachineLearningWorkspaceArgs) ToMachineLearningWorkspaceOutputWithContext(ctx context.Context) MachineLearningWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningWorkspaceOutput)
}

func (i MachineLearningWorkspaceArgs) ToMachineLearningWorkspacePtrOutput() MachineLearningWorkspacePtrOutput {
	return i.ToMachineLearningWorkspacePtrOutputWithContext(context.Background())
}

func (i MachineLearningWorkspaceArgs) ToMachineLearningWorkspacePtrOutputWithContext(ctx context.Context) MachineLearningWorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningWorkspaceOutput).ToMachineLearningWorkspacePtrOutputWithContext(ctx)
}









type MachineLearningWorkspacePtrInput interface {
	pulumi.Input

	ToMachineLearningWorkspacePtrOutput() MachineLearningWorkspacePtrOutput
	ToMachineLearningWorkspacePtrOutputWithContext(context.Context) MachineLearningWorkspacePtrOutput
}

type machineLearningWorkspacePtrType MachineLearningWorkspaceArgs

func MachineLearningWorkspacePtr(v *MachineLearningWorkspaceArgs) MachineLearningWorkspacePtrInput {
	return (*machineLearningWorkspacePtrType)(v)
}

func (*machineLearningWorkspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningWorkspace)(nil)).Elem()
}

func (i *machineLearningWorkspacePtrType) ToMachineLearningWorkspacePtrOutput() MachineLearningWorkspacePtrOutput {
	return i.ToMachineLearningWorkspacePtrOutputWithContext(context.Background())
}

func (i *machineLearningWorkspacePtrType) ToMachineLearningWorkspacePtrOutputWithContext(ctx context.Context) MachineLearningWorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningWorkspacePtrOutput)
}

type MachineLearningWorkspaceOutput struct{ *pulumi.OutputState }

func (MachineLearningWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningWorkspace)(nil)).Elem()
}

func (o MachineLearningWorkspaceOutput) ToMachineLearningWorkspaceOutput() MachineLearningWorkspaceOutput {
	return o
}

func (o MachineLearningWorkspaceOutput) ToMachineLearningWorkspaceOutputWithContext(ctx context.Context) MachineLearningWorkspaceOutput {
	return o
}

func (o MachineLearningWorkspaceOutput) ToMachineLearningWorkspacePtrOutput() MachineLearningWorkspacePtrOutput {
	return o.ToMachineLearningWorkspacePtrOutputWithContext(context.Background())
}

func (o MachineLearningWorkspaceOutput) ToMachineLearningWorkspacePtrOutputWithContext(ctx context.Context) MachineLearningWorkspacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineLearningWorkspace) *MachineLearningWorkspace {
		return &v
	}).(MachineLearningWorkspacePtrOutput)
}

func (o MachineLearningWorkspaceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineLearningWorkspace) string { return v.Id }).(pulumi.StringOutput)
}

type MachineLearningWorkspacePtrOutput struct{ *pulumi.OutputState }

func (MachineLearningWorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningWorkspace)(nil)).Elem()
}

func (o MachineLearningWorkspacePtrOutput) ToMachineLearningWorkspacePtrOutput() MachineLearningWorkspacePtrOutput {
	return o
}

func (o MachineLearningWorkspacePtrOutput) ToMachineLearningWorkspacePtrOutputWithContext(ctx context.Context) MachineLearningWorkspacePtrOutput {
	return o
}

func (o MachineLearningWorkspacePtrOutput) Elem() MachineLearningWorkspaceOutput {
	return o.ApplyT(func(v *MachineLearningWorkspace) MachineLearningWorkspace {
		if v != nil {
			return *v
		}
		var ret MachineLearningWorkspace
		return ret
	}).(MachineLearningWorkspaceOutput)
}

func (o MachineLearningWorkspacePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningWorkspace) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type MachineLearningWorkspaceResponse struct {
	Id string `pulumi:"id"`
}





type MachineLearningWorkspaceResponseInput interface {
	pulumi.Input

	ToMachineLearningWorkspaceResponseOutput() MachineLearningWorkspaceResponseOutput
	ToMachineLearningWorkspaceResponseOutputWithContext(context.Context) MachineLearningWorkspaceResponseOutput
}

type MachineLearningWorkspaceResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MachineLearningWorkspaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningWorkspaceResponse)(nil)).Elem()
}

func (i MachineLearningWorkspaceResponseArgs) ToMachineLearningWorkspaceResponseOutput() MachineLearningWorkspaceResponseOutput {
	return i.ToMachineLearningWorkspaceResponseOutputWithContext(context.Background())
}

func (i MachineLearningWorkspaceResponseArgs) ToMachineLearningWorkspaceResponseOutputWithContext(ctx context.Context) MachineLearningWorkspaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningWorkspaceResponseOutput)
}

func (i MachineLearningWorkspaceResponseArgs) ToMachineLearningWorkspaceResponsePtrOutput() MachineLearningWorkspaceResponsePtrOutput {
	return i.ToMachineLearningWorkspaceResponsePtrOutputWithContext(context.Background())
}

func (i MachineLearningWorkspaceResponseArgs) ToMachineLearningWorkspaceResponsePtrOutputWithContext(ctx context.Context) MachineLearningWorkspaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningWorkspaceResponseOutput).ToMachineLearningWorkspaceResponsePtrOutputWithContext(ctx)
}









type MachineLearningWorkspaceResponsePtrInput interface {
	pulumi.Input

	ToMachineLearningWorkspaceResponsePtrOutput() MachineLearningWorkspaceResponsePtrOutput
	ToMachineLearningWorkspaceResponsePtrOutputWithContext(context.Context) MachineLearningWorkspaceResponsePtrOutput
}

type machineLearningWorkspaceResponsePtrType MachineLearningWorkspaceResponseArgs

func MachineLearningWorkspaceResponsePtr(v *MachineLearningWorkspaceResponseArgs) MachineLearningWorkspaceResponsePtrInput {
	return (*machineLearningWorkspaceResponsePtrType)(v)
}

func (*machineLearningWorkspaceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningWorkspaceResponse)(nil)).Elem()
}

func (i *machineLearningWorkspaceResponsePtrType) ToMachineLearningWorkspaceResponsePtrOutput() MachineLearningWorkspaceResponsePtrOutput {
	return i.ToMachineLearningWorkspaceResponsePtrOutputWithContext(context.Background())
}

func (i *machineLearningWorkspaceResponsePtrType) ToMachineLearningWorkspaceResponsePtrOutputWithContext(ctx context.Context) MachineLearningWorkspaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningWorkspaceResponsePtrOutput)
}

type MachineLearningWorkspaceResponseOutput struct{ *pulumi.OutputState }

func (MachineLearningWorkspaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningWorkspaceResponse)(nil)).Elem()
}

func (o MachineLearningWorkspaceResponseOutput) ToMachineLearningWorkspaceResponseOutput() MachineLearningWorkspaceResponseOutput {
	return o
}

func (o MachineLearningWorkspaceResponseOutput) ToMachineLearningWorkspaceResponseOutputWithContext(ctx context.Context) MachineLearningWorkspaceResponseOutput {
	return o
}

func (o MachineLearningWorkspaceResponseOutput) ToMachineLearningWorkspaceResponsePtrOutput() MachineLearningWorkspaceResponsePtrOutput {
	return o.ToMachineLearningWorkspaceResponsePtrOutputWithContext(context.Background())
}

func (o MachineLearningWorkspaceResponseOutput) ToMachineLearningWorkspaceResponsePtrOutputWithContext(ctx context.Context) MachineLearningWorkspaceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineLearningWorkspaceResponse) *MachineLearningWorkspaceResponse {
		return &v
	}).(MachineLearningWorkspaceResponsePtrOutput)
}

func (o MachineLearningWorkspaceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineLearningWorkspaceResponse) string { return v.Id }).(pulumi.StringOutput)
}

type MachineLearningWorkspaceResponsePtrOutput struct{ *pulumi.OutputState }

func (MachineLearningWorkspaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningWorkspaceResponse)(nil)).Elem()
}

func (o MachineLearningWorkspaceResponsePtrOutput) ToMachineLearningWorkspaceResponsePtrOutput() MachineLearningWorkspaceResponsePtrOutput {
	return o
}

func (o MachineLearningWorkspaceResponsePtrOutput) ToMachineLearningWorkspaceResponsePtrOutputWithContext(ctx context.Context) MachineLearningWorkspaceResponsePtrOutput {
	return o
}

func (o MachineLearningWorkspaceResponsePtrOutput) Elem() MachineLearningWorkspaceResponseOutput {
	return o.ApplyT(func(v *MachineLearningWorkspaceResponse) MachineLearningWorkspaceResponse {
		if v != nil {
			return *v
		}
		var ret MachineLearningWorkspaceResponse
		return ret
	}).(MachineLearningWorkspaceResponseOutput)
}

func (o MachineLearningWorkspaceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningWorkspaceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ModeValueInfo struct {
	InterfaceString *string                `pulumi:"interfaceString"`
	Parameters      []ModuleAssetParameter `pulumi:"parameters"`
}





type ModeValueInfoInput interface {
	pulumi.Input

	ToModeValueInfoOutput() ModeValueInfoOutput
	ToModeValueInfoOutputWithContext(context.Context) ModeValueInfoOutput
}

type ModeValueInfoArgs struct {
	InterfaceString pulumi.StringPtrInput          `pulumi:"interfaceString"`
	Parameters      ModuleAssetParameterArrayInput `pulumi:"parameters"`
}

func (ModeValueInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModeValueInfo)(nil)).Elem()
}

func (i ModeValueInfoArgs) ToModeValueInfoOutput() ModeValueInfoOutput {
	return i.ToModeValueInfoOutputWithContext(context.Background())
}

func (i ModeValueInfoArgs) ToModeValueInfoOutputWithContext(ctx context.Context) ModeValueInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModeValueInfoOutput)
}





type ModeValueInfoMapInput interface {
	pulumi.Input

	ToModeValueInfoMapOutput() ModeValueInfoMapOutput
	ToModeValueInfoMapOutputWithContext(context.Context) ModeValueInfoMapOutput
}

type ModeValueInfoMap map[string]ModeValueInfoInput

func (ModeValueInfoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ModeValueInfo)(nil)).Elem()
}

func (i ModeValueInfoMap) ToModeValueInfoMapOutput() ModeValueInfoMapOutput {
	return i.ToModeValueInfoMapOutputWithContext(context.Background())
}

func (i ModeValueInfoMap) ToModeValueInfoMapOutputWithContext(ctx context.Context) ModeValueInfoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModeValueInfoMapOutput)
}

type ModeValueInfoOutput struct{ *pulumi.OutputState }

func (ModeValueInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModeValueInfo)(nil)).Elem()
}

func (o ModeValueInfoOutput) ToModeValueInfoOutput() ModeValueInfoOutput {
	return o
}

func (o ModeValueInfoOutput) ToModeValueInfoOutputWithContext(ctx context.Context) ModeValueInfoOutput {
	return o
}

func (o ModeValueInfoOutput) InterfaceString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModeValueInfo) *string { return v.InterfaceString }).(pulumi.StringPtrOutput)
}

func (o ModeValueInfoOutput) Parameters() ModuleAssetParameterArrayOutput {
	return o.ApplyT(func(v ModeValueInfo) []ModuleAssetParameter { return v.Parameters }).(ModuleAssetParameterArrayOutput)
}

type ModeValueInfoMapOutput struct{ *pulumi.OutputState }

func (ModeValueInfoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ModeValueInfo)(nil)).Elem()
}

func (o ModeValueInfoMapOutput) ToModeValueInfoMapOutput() ModeValueInfoMapOutput {
	return o
}

func (o ModeValueInfoMapOutput) ToModeValueInfoMapOutputWithContext(ctx context.Context) ModeValueInfoMapOutput {
	return o
}

func (o ModeValueInfoMapOutput) MapIndex(k pulumi.StringInput) ModeValueInfoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ModeValueInfo {
		return vs[0].(map[string]ModeValueInfo)[vs[1].(string)]
	}).(ModeValueInfoOutput)
}

type ModeValueInfoResponse struct {
	InterfaceString *string                        `pulumi:"interfaceString"`
	Parameters      []ModuleAssetParameterResponse `pulumi:"parameters"`
}





type ModeValueInfoResponseInput interface {
	pulumi.Input

	ToModeValueInfoResponseOutput() ModeValueInfoResponseOutput
	ToModeValueInfoResponseOutputWithContext(context.Context) ModeValueInfoResponseOutput
}

type ModeValueInfoResponseArgs struct {
	InterfaceString pulumi.StringPtrInput                  `pulumi:"interfaceString"`
	Parameters      ModuleAssetParameterResponseArrayInput `pulumi:"parameters"`
}

func (ModeValueInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModeValueInfoResponse)(nil)).Elem()
}

func (i ModeValueInfoResponseArgs) ToModeValueInfoResponseOutput() ModeValueInfoResponseOutput {
	return i.ToModeValueInfoResponseOutputWithContext(context.Background())
}

func (i ModeValueInfoResponseArgs) ToModeValueInfoResponseOutputWithContext(ctx context.Context) ModeValueInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModeValueInfoResponseOutput)
}





type ModeValueInfoResponseMapInput interface {
	pulumi.Input

	ToModeValueInfoResponseMapOutput() ModeValueInfoResponseMapOutput
	ToModeValueInfoResponseMapOutputWithContext(context.Context) ModeValueInfoResponseMapOutput
}

type ModeValueInfoResponseMap map[string]ModeValueInfoResponseInput

func (ModeValueInfoResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ModeValueInfoResponse)(nil)).Elem()
}

func (i ModeValueInfoResponseMap) ToModeValueInfoResponseMapOutput() ModeValueInfoResponseMapOutput {
	return i.ToModeValueInfoResponseMapOutputWithContext(context.Background())
}

func (i ModeValueInfoResponseMap) ToModeValueInfoResponseMapOutputWithContext(ctx context.Context) ModeValueInfoResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModeValueInfoResponseMapOutput)
}

type ModeValueInfoResponseOutput struct{ *pulumi.OutputState }

func (ModeValueInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModeValueInfoResponse)(nil)).Elem()
}

func (o ModeValueInfoResponseOutput) ToModeValueInfoResponseOutput() ModeValueInfoResponseOutput {
	return o
}

func (o ModeValueInfoResponseOutput) ToModeValueInfoResponseOutputWithContext(ctx context.Context) ModeValueInfoResponseOutput {
	return o
}

func (o ModeValueInfoResponseOutput) InterfaceString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModeValueInfoResponse) *string { return v.InterfaceString }).(pulumi.StringPtrOutput)
}

func (o ModeValueInfoResponseOutput) Parameters() ModuleAssetParameterResponseArrayOutput {
	return o.ApplyT(func(v ModeValueInfoResponse) []ModuleAssetParameterResponse { return v.Parameters }).(ModuleAssetParameterResponseArrayOutput)
}

type ModeValueInfoResponseMapOutput struct{ *pulumi.OutputState }

func (ModeValueInfoResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ModeValueInfoResponse)(nil)).Elem()
}

func (o ModeValueInfoResponseMapOutput) ToModeValueInfoResponseMapOutput() ModeValueInfoResponseMapOutput {
	return o
}

func (o ModeValueInfoResponseMapOutput) ToModeValueInfoResponseMapOutputWithContext(ctx context.Context) ModeValueInfoResponseMapOutput {
	return o
}

func (o ModeValueInfoResponseMapOutput) MapIndex(k pulumi.StringInput) ModeValueInfoResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ModeValueInfoResponse {
		return vs[0].(map[string]ModeValueInfoResponse)[vs[1].(string)]
	}).(ModeValueInfoResponseOutput)
}

type ModuleAssetParameter struct {
	ModeValuesInfo map[string]ModeValueInfo `pulumi:"modeValuesInfo"`
	Name           *string                  `pulumi:"name"`
	ParameterType  *string                  `pulumi:"parameterType"`
}





type ModuleAssetParameterInput interface {
	pulumi.Input

	ToModuleAssetParameterOutput() ModuleAssetParameterOutput
	ToModuleAssetParameterOutputWithContext(context.Context) ModuleAssetParameterOutput
}

type ModuleAssetParameterArgs struct {
	ModeValuesInfo ModeValueInfoMapInput `pulumi:"modeValuesInfo"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	ParameterType  pulumi.StringPtrInput `pulumi:"parameterType"`
}

func (ModuleAssetParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleAssetParameter)(nil)).Elem()
}

func (i ModuleAssetParameterArgs) ToModuleAssetParameterOutput() ModuleAssetParameterOutput {
	return i.ToModuleAssetParameterOutputWithContext(context.Background())
}

func (i ModuleAssetParameterArgs) ToModuleAssetParameterOutputWithContext(ctx context.Context) ModuleAssetParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleAssetParameterOutput)
}





type ModuleAssetParameterArrayInput interface {
	pulumi.Input

	ToModuleAssetParameterArrayOutput() ModuleAssetParameterArrayOutput
	ToModuleAssetParameterArrayOutputWithContext(context.Context) ModuleAssetParameterArrayOutput
}

type ModuleAssetParameterArray []ModuleAssetParameterInput

func (ModuleAssetParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModuleAssetParameter)(nil)).Elem()
}

func (i ModuleAssetParameterArray) ToModuleAssetParameterArrayOutput() ModuleAssetParameterArrayOutput {
	return i.ToModuleAssetParameterArrayOutputWithContext(context.Background())
}

func (i ModuleAssetParameterArray) ToModuleAssetParameterArrayOutputWithContext(ctx context.Context) ModuleAssetParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleAssetParameterArrayOutput)
}

type ModuleAssetParameterOutput struct{ *pulumi.OutputState }

func (ModuleAssetParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleAssetParameter)(nil)).Elem()
}

func (o ModuleAssetParameterOutput) ToModuleAssetParameterOutput() ModuleAssetParameterOutput {
	return o
}

func (o ModuleAssetParameterOutput) ToModuleAssetParameterOutputWithContext(ctx context.Context) ModuleAssetParameterOutput {
	return o
}

func (o ModuleAssetParameterOutput) ModeValuesInfo() ModeValueInfoMapOutput {
	return o.ApplyT(func(v ModuleAssetParameter) map[string]ModeValueInfo { return v.ModeValuesInfo }).(ModeValueInfoMapOutput)
}

func (o ModuleAssetParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleAssetParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ModuleAssetParameterOutput) ParameterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleAssetParameter) *string { return v.ParameterType }).(pulumi.StringPtrOutput)
}

type ModuleAssetParameterArrayOutput struct{ *pulumi.OutputState }

func (ModuleAssetParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModuleAssetParameter)(nil)).Elem()
}

func (o ModuleAssetParameterArrayOutput) ToModuleAssetParameterArrayOutput() ModuleAssetParameterArrayOutput {
	return o
}

func (o ModuleAssetParameterArrayOutput) ToModuleAssetParameterArrayOutputWithContext(ctx context.Context) ModuleAssetParameterArrayOutput {
	return o
}

func (o ModuleAssetParameterArrayOutput) Index(i pulumi.IntInput) ModuleAssetParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ModuleAssetParameter {
		return vs[0].([]ModuleAssetParameter)[vs[1].(int)]
	}).(ModuleAssetParameterOutput)
}

type ModuleAssetParameterResponse struct {
	ModeValuesInfo map[string]ModeValueInfoResponse `pulumi:"modeValuesInfo"`
	Name           *string                          `pulumi:"name"`
	ParameterType  *string                          `pulumi:"parameterType"`
}





type ModuleAssetParameterResponseInput interface {
	pulumi.Input

	ToModuleAssetParameterResponseOutput() ModuleAssetParameterResponseOutput
	ToModuleAssetParameterResponseOutputWithContext(context.Context) ModuleAssetParameterResponseOutput
}

type ModuleAssetParameterResponseArgs struct {
	ModeValuesInfo ModeValueInfoResponseMapInput `pulumi:"modeValuesInfo"`
	Name           pulumi.StringPtrInput         `pulumi:"name"`
	ParameterType  pulumi.StringPtrInput         `pulumi:"parameterType"`
}

func (ModuleAssetParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleAssetParameterResponse)(nil)).Elem()
}

func (i ModuleAssetParameterResponseArgs) ToModuleAssetParameterResponseOutput() ModuleAssetParameterResponseOutput {
	return i.ToModuleAssetParameterResponseOutputWithContext(context.Background())
}

func (i ModuleAssetParameterResponseArgs) ToModuleAssetParameterResponseOutputWithContext(ctx context.Context) ModuleAssetParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleAssetParameterResponseOutput)
}





type ModuleAssetParameterResponseArrayInput interface {
	pulumi.Input

	ToModuleAssetParameterResponseArrayOutput() ModuleAssetParameterResponseArrayOutput
	ToModuleAssetParameterResponseArrayOutputWithContext(context.Context) ModuleAssetParameterResponseArrayOutput
}

type ModuleAssetParameterResponseArray []ModuleAssetParameterResponseInput

func (ModuleAssetParameterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModuleAssetParameterResponse)(nil)).Elem()
}

func (i ModuleAssetParameterResponseArray) ToModuleAssetParameterResponseArrayOutput() ModuleAssetParameterResponseArrayOutput {
	return i.ToModuleAssetParameterResponseArrayOutputWithContext(context.Background())
}

func (i ModuleAssetParameterResponseArray) ToModuleAssetParameterResponseArrayOutputWithContext(ctx context.Context) ModuleAssetParameterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleAssetParameterResponseArrayOutput)
}

type ModuleAssetParameterResponseOutput struct{ *pulumi.OutputState }

func (ModuleAssetParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleAssetParameterResponse)(nil)).Elem()
}

func (o ModuleAssetParameterResponseOutput) ToModuleAssetParameterResponseOutput() ModuleAssetParameterResponseOutput {
	return o
}

func (o ModuleAssetParameterResponseOutput) ToModuleAssetParameterResponseOutputWithContext(ctx context.Context) ModuleAssetParameterResponseOutput {
	return o
}

func (o ModuleAssetParameterResponseOutput) ModeValuesInfo() ModeValueInfoResponseMapOutput {
	return o.ApplyT(func(v ModuleAssetParameterResponse) map[string]ModeValueInfoResponse { return v.ModeValuesInfo }).(ModeValueInfoResponseMapOutput)
}

func (o ModuleAssetParameterResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleAssetParameterResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ModuleAssetParameterResponseOutput) ParameterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleAssetParameterResponse) *string { return v.ParameterType }).(pulumi.StringPtrOutput)
}

type ModuleAssetParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ModuleAssetParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModuleAssetParameterResponse)(nil)).Elem()
}

func (o ModuleAssetParameterResponseArrayOutput) ToModuleAssetParameterResponseArrayOutput() ModuleAssetParameterResponseArrayOutput {
	return o
}

func (o ModuleAssetParameterResponseArrayOutput) ToModuleAssetParameterResponseArrayOutputWithContext(ctx context.Context) ModuleAssetParameterResponseArrayOutput {
	return o
}

func (o ModuleAssetParameterResponseArrayOutput) Index(i pulumi.IntInput) ModuleAssetParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ModuleAssetParameterResponse {
		return vs[0].([]ModuleAssetParameterResponse)[vs[1].(int)]
	}).(ModuleAssetParameterResponseOutput)
}

type OutputPort struct {
	Type *string `pulumi:"type"`
}





type OutputPortInput interface {
	pulumi.Input

	ToOutputPortOutput() OutputPortOutput
	ToOutputPortOutputWithContext(context.Context) OutputPortOutput
}

type OutputPortArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (OutputPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputPort)(nil)).Elem()
}

func (i OutputPortArgs) ToOutputPortOutput() OutputPortOutput {
	return i.ToOutputPortOutputWithContext(context.Background())
}

func (i OutputPortArgs) ToOutputPortOutputWithContext(ctx context.Context) OutputPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputPortOutput)
}





type OutputPortMapInput interface {
	pulumi.Input

	ToOutputPortMapOutput() OutputPortMapOutput
	ToOutputPortMapOutputWithContext(context.Context) OutputPortMapOutput
}

type OutputPortMap map[string]OutputPortInput

func (OutputPortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputPort)(nil)).Elem()
}

func (i OutputPortMap) ToOutputPortMapOutput() OutputPortMapOutput {
	return i.ToOutputPortMapOutputWithContext(context.Background())
}

func (i OutputPortMap) ToOutputPortMapOutputWithContext(ctx context.Context) OutputPortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputPortMapOutput)
}

type OutputPortOutput struct{ *pulumi.OutputState }

func (OutputPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputPort)(nil)).Elem()
}

func (o OutputPortOutput) ToOutputPortOutput() OutputPortOutput {
	return o
}

func (o OutputPortOutput) ToOutputPortOutputWithContext(ctx context.Context) OutputPortOutput {
	return o
}

func (o OutputPortOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputPort) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OutputPortMapOutput struct{ *pulumi.OutputState }

func (OutputPortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputPort)(nil)).Elem()
}

func (o OutputPortMapOutput) ToOutputPortMapOutput() OutputPortMapOutput {
	return o
}

func (o OutputPortMapOutput) ToOutputPortMapOutputWithContext(ctx context.Context) OutputPortMapOutput {
	return o
}

func (o OutputPortMapOutput) MapIndex(k pulumi.StringInput) OutputPortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputPort {
		return vs[0].(map[string]OutputPort)[vs[1].(string)]
	}).(OutputPortOutput)
}

type OutputPortResponse struct {
	Type *string `pulumi:"type"`
}





type OutputPortResponseInput interface {
	pulumi.Input

	ToOutputPortResponseOutput() OutputPortResponseOutput
	ToOutputPortResponseOutputWithContext(context.Context) OutputPortResponseOutput
}

type OutputPortResponseArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (OutputPortResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputPortResponse)(nil)).Elem()
}

func (i OutputPortResponseArgs) ToOutputPortResponseOutput() OutputPortResponseOutput {
	return i.ToOutputPortResponseOutputWithContext(context.Background())
}

func (i OutputPortResponseArgs) ToOutputPortResponseOutputWithContext(ctx context.Context) OutputPortResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputPortResponseOutput)
}





type OutputPortResponseMapInput interface {
	pulumi.Input

	ToOutputPortResponseMapOutput() OutputPortResponseMapOutput
	ToOutputPortResponseMapOutputWithContext(context.Context) OutputPortResponseMapOutput
}

type OutputPortResponseMap map[string]OutputPortResponseInput

func (OutputPortResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputPortResponse)(nil)).Elem()
}

func (i OutputPortResponseMap) ToOutputPortResponseMapOutput() OutputPortResponseMapOutput {
	return i.ToOutputPortResponseMapOutputWithContext(context.Background())
}

func (i OutputPortResponseMap) ToOutputPortResponseMapOutputWithContext(ctx context.Context) OutputPortResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputPortResponseMapOutput)
}

type OutputPortResponseOutput struct{ *pulumi.OutputState }

func (OutputPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputPortResponse)(nil)).Elem()
}

func (o OutputPortResponseOutput) ToOutputPortResponseOutput() OutputPortResponseOutput {
	return o
}

func (o OutputPortResponseOutput) ToOutputPortResponseOutputWithContext(ctx context.Context) OutputPortResponseOutput {
	return o
}

func (o OutputPortResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputPortResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OutputPortResponseMapOutput struct{ *pulumi.OutputState }

func (OutputPortResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputPortResponse)(nil)).Elem()
}

func (o OutputPortResponseMapOutput) ToOutputPortResponseMapOutput() OutputPortResponseMapOutput {
	return o
}

func (o OutputPortResponseMapOutput) ToOutputPortResponseMapOutputWithContext(ctx context.Context) OutputPortResponseMapOutput {
	return o
}

func (o OutputPortResponseMapOutput) MapIndex(k pulumi.StringInput) OutputPortResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputPortResponse {
		return vs[0].(map[string]OutputPortResponse)[vs[1].(string)]
	}).(OutputPortResponseOutput)
}

type PlanQuantityResponse struct {
	Allowance             float64 `pulumi:"allowance"`
	Amount                float64 `pulumi:"amount"`
	IncludedQuantityMeter string  `pulumi:"includedQuantityMeter"`
	OverageMeter          string  `pulumi:"overageMeter"`
}





type PlanQuantityResponseInput interface {
	pulumi.Input

	ToPlanQuantityResponseOutput() PlanQuantityResponseOutput
	ToPlanQuantityResponseOutputWithContext(context.Context) PlanQuantityResponseOutput
}

type PlanQuantityResponseArgs struct {
	Allowance             pulumi.Float64Input `pulumi:"allowance"`
	Amount                pulumi.Float64Input `pulumi:"amount"`
	IncludedQuantityMeter pulumi.StringInput  `pulumi:"includedQuantityMeter"`
	OverageMeter          pulumi.StringInput  `pulumi:"overageMeter"`
}

func (PlanQuantityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanQuantityResponse)(nil)).Elem()
}

func (i PlanQuantityResponseArgs) ToPlanQuantityResponseOutput() PlanQuantityResponseOutput {
	return i.ToPlanQuantityResponseOutputWithContext(context.Background())
}

func (i PlanQuantityResponseArgs) ToPlanQuantityResponseOutputWithContext(ctx context.Context) PlanQuantityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanQuantityResponseOutput)
}





type PlanQuantityResponseMapInput interface {
	pulumi.Input

	ToPlanQuantityResponseMapOutput() PlanQuantityResponseMapOutput
	ToPlanQuantityResponseMapOutputWithContext(context.Context) PlanQuantityResponseMapOutput
}

type PlanQuantityResponseMap map[string]PlanQuantityResponseInput

func (PlanQuantityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PlanQuantityResponse)(nil)).Elem()
}

func (i PlanQuantityResponseMap) ToPlanQuantityResponseMapOutput() PlanQuantityResponseMapOutput {
	return i.ToPlanQuantityResponseMapOutputWithContext(context.Background())
}

func (i PlanQuantityResponseMap) ToPlanQuantityResponseMapOutputWithContext(ctx context.Context) PlanQuantityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanQuantityResponseMapOutput)
}

type PlanQuantityResponseOutput struct{ *pulumi.OutputState }

func (PlanQuantityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanQuantityResponse)(nil)).Elem()
}

func (o PlanQuantityResponseOutput) ToPlanQuantityResponseOutput() PlanQuantityResponseOutput {
	return o
}

func (o PlanQuantityResponseOutput) ToPlanQuantityResponseOutputWithContext(ctx context.Context) PlanQuantityResponseOutput {
	return o
}

func (o PlanQuantityResponseOutput) Allowance() pulumi.Float64Output {
	return o.ApplyT(func(v PlanQuantityResponse) float64 { return v.Allowance }).(pulumi.Float64Output)
}

func (o PlanQuantityResponseOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v PlanQuantityResponse) float64 { return v.Amount }).(pulumi.Float64Output)
}

func (o PlanQuantityResponseOutput) IncludedQuantityMeter() pulumi.StringOutput {
	return o.ApplyT(func(v PlanQuantityResponse) string { return v.IncludedQuantityMeter }).(pulumi.StringOutput)
}

func (o PlanQuantityResponseOutput) OverageMeter() pulumi.StringOutput {
	return o.ApplyT(func(v PlanQuantityResponse) string { return v.OverageMeter }).(pulumi.StringOutput)
}

type PlanQuantityResponseMapOutput struct{ *pulumi.OutputState }

func (PlanQuantityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PlanQuantityResponse)(nil)).Elem()
}

func (o PlanQuantityResponseMapOutput) ToPlanQuantityResponseMapOutput() PlanQuantityResponseMapOutput {
	return o
}

func (o PlanQuantityResponseMapOutput) ToPlanQuantityResponseMapOutputWithContext(ctx context.Context) PlanQuantityResponseMapOutput {
	return o
}

func (o PlanQuantityResponseMapOutput) MapIndex(k pulumi.StringInput) PlanQuantityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PlanQuantityResponse {
		return vs[0].(map[string]PlanQuantityResponse)[vs[1].(string)]
	}).(PlanQuantityResponseOutput)
}

type RealtimeConfiguration struct {
	MaxConcurrentCalls *int `pulumi:"maxConcurrentCalls"`
}





type RealtimeConfigurationInput interface {
	pulumi.Input

	ToRealtimeConfigurationOutput() RealtimeConfigurationOutput
	ToRealtimeConfigurationOutputWithContext(context.Context) RealtimeConfigurationOutput
}

type RealtimeConfigurationArgs struct {
	MaxConcurrentCalls pulumi.IntPtrInput `pulumi:"maxConcurrentCalls"`
}

func (RealtimeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RealtimeConfiguration)(nil)).Elem()
}

func (i RealtimeConfigurationArgs) ToRealtimeConfigurationOutput() RealtimeConfigurationOutput {
	return i.ToRealtimeConfigurationOutputWithContext(context.Background())
}

func (i RealtimeConfigurationArgs) ToRealtimeConfigurationOutputWithContext(ctx context.Context) RealtimeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeConfigurationOutput)
}

func (i RealtimeConfigurationArgs) ToRealtimeConfigurationPtrOutput() RealtimeConfigurationPtrOutput {
	return i.ToRealtimeConfigurationPtrOutputWithContext(context.Background())
}

func (i RealtimeConfigurationArgs) ToRealtimeConfigurationPtrOutputWithContext(ctx context.Context) RealtimeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeConfigurationOutput).ToRealtimeConfigurationPtrOutputWithContext(ctx)
}









type RealtimeConfigurationPtrInput interface {
	pulumi.Input

	ToRealtimeConfigurationPtrOutput() RealtimeConfigurationPtrOutput
	ToRealtimeConfigurationPtrOutputWithContext(context.Context) RealtimeConfigurationPtrOutput
}

type realtimeConfigurationPtrType RealtimeConfigurationArgs

func RealtimeConfigurationPtr(v *RealtimeConfigurationArgs) RealtimeConfigurationPtrInput {
	return (*realtimeConfigurationPtrType)(v)
}

func (*realtimeConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeConfiguration)(nil)).Elem()
}

func (i *realtimeConfigurationPtrType) ToRealtimeConfigurationPtrOutput() RealtimeConfigurationPtrOutput {
	return i.ToRealtimeConfigurationPtrOutputWithContext(context.Background())
}

func (i *realtimeConfigurationPtrType) ToRealtimeConfigurationPtrOutputWithContext(ctx context.Context) RealtimeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeConfigurationPtrOutput)
}

type RealtimeConfigurationOutput struct{ *pulumi.OutputState }

func (RealtimeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RealtimeConfiguration)(nil)).Elem()
}

func (o RealtimeConfigurationOutput) ToRealtimeConfigurationOutput() RealtimeConfigurationOutput {
	return o
}

func (o RealtimeConfigurationOutput) ToRealtimeConfigurationOutputWithContext(ctx context.Context) RealtimeConfigurationOutput {
	return o
}

func (o RealtimeConfigurationOutput) ToRealtimeConfigurationPtrOutput() RealtimeConfigurationPtrOutput {
	return o.ToRealtimeConfigurationPtrOutputWithContext(context.Background())
}

func (o RealtimeConfigurationOutput) ToRealtimeConfigurationPtrOutputWithContext(ctx context.Context) RealtimeConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RealtimeConfiguration) *RealtimeConfiguration {
		return &v
	}).(RealtimeConfigurationPtrOutput)
}

func (o RealtimeConfigurationOutput) MaxConcurrentCalls() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RealtimeConfiguration) *int { return v.MaxConcurrentCalls }).(pulumi.IntPtrOutput)
}

type RealtimeConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RealtimeConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeConfiguration)(nil)).Elem()
}

func (o RealtimeConfigurationPtrOutput) ToRealtimeConfigurationPtrOutput() RealtimeConfigurationPtrOutput {
	return o
}

func (o RealtimeConfigurationPtrOutput) ToRealtimeConfigurationPtrOutputWithContext(ctx context.Context) RealtimeConfigurationPtrOutput {
	return o
}

func (o RealtimeConfigurationPtrOutput) Elem() RealtimeConfigurationOutput {
	return o.ApplyT(func(v *RealtimeConfiguration) RealtimeConfiguration {
		if v != nil {
			return *v
		}
		var ret RealtimeConfiguration
		return ret
	}).(RealtimeConfigurationOutput)
}

func (o RealtimeConfigurationPtrOutput) MaxConcurrentCalls() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealtimeConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.MaxConcurrentCalls
	}).(pulumi.IntPtrOutput)
}

type RealtimeConfigurationResponse struct {
	MaxConcurrentCalls *int `pulumi:"maxConcurrentCalls"`
}





type RealtimeConfigurationResponseInput interface {
	pulumi.Input

	ToRealtimeConfigurationResponseOutput() RealtimeConfigurationResponseOutput
	ToRealtimeConfigurationResponseOutputWithContext(context.Context) RealtimeConfigurationResponseOutput
}

type RealtimeConfigurationResponseArgs struct {
	MaxConcurrentCalls pulumi.IntPtrInput `pulumi:"maxConcurrentCalls"`
}

func (RealtimeConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RealtimeConfigurationResponse)(nil)).Elem()
}

func (i RealtimeConfigurationResponseArgs) ToRealtimeConfigurationResponseOutput() RealtimeConfigurationResponseOutput {
	return i.ToRealtimeConfigurationResponseOutputWithContext(context.Background())
}

func (i RealtimeConfigurationResponseArgs) ToRealtimeConfigurationResponseOutputWithContext(ctx context.Context) RealtimeConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeConfigurationResponseOutput)
}

func (i RealtimeConfigurationResponseArgs) ToRealtimeConfigurationResponsePtrOutput() RealtimeConfigurationResponsePtrOutput {
	return i.ToRealtimeConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i RealtimeConfigurationResponseArgs) ToRealtimeConfigurationResponsePtrOutputWithContext(ctx context.Context) RealtimeConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeConfigurationResponseOutput).ToRealtimeConfigurationResponsePtrOutputWithContext(ctx)
}









type RealtimeConfigurationResponsePtrInput interface {
	pulumi.Input

	ToRealtimeConfigurationResponsePtrOutput() RealtimeConfigurationResponsePtrOutput
	ToRealtimeConfigurationResponsePtrOutputWithContext(context.Context) RealtimeConfigurationResponsePtrOutput
}

type realtimeConfigurationResponsePtrType RealtimeConfigurationResponseArgs

func RealtimeConfigurationResponsePtr(v *RealtimeConfigurationResponseArgs) RealtimeConfigurationResponsePtrInput {
	return (*realtimeConfigurationResponsePtrType)(v)
}

func (*realtimeConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeConfigurationResponse)(nil)).Elem()
}

func (i *realtimeConfigurationResponsePtrType) ToRealtimeConfigurationResponsePtrOutput() RealtimeConfigurationResponsePtrOutput {
	return i.ToRealtimeConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *realtimeConfigurationResponsePtrType) ToRealtimeConfigurationResponsePtrOutputWithContext(ctx context.Context) RealtimeConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeConfigurationResponsePtrOutput)
}

type RealtimeConfigurationResponseOutput struct{ *pulumi.OutputState }

func (RealtimeConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RealtimeConfigurationResponse)(nil)).Elem()
}

func (o RealtimeConfigurationResponseOutput) ToRealtimeConfigurationResponseOutput() RealtimeConfigurationResponseOutput {
	return o
}

func (o RealtimeConfigurationResponseOutput) ToRealtimeConfigurationResponseOutputWithContext(ctx context.Context) RealtimeConfigurationResponseOutput {
	return o
}

func (o RealtimeConfigurationResponseOutput) ToRealtimeConfigurationResponsePtrOutput() RealtimeConfigurationResponsePtrOutput {
	return o.ToRealtimeConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o RealtimeConfigurationResponseOutput) ToRealtimeConfigurationResponsePtrOutputWithContext(ctx context.Context) RealtimeConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RealtimeConfigurationResponse) *RealtimeConfigurationResponse {
		return &v
	}).(RealtimeConfigurationResponsePtrOutput)
}

func (o RealtimeConfigurationResponseOutput) MaxConcurrentCalls() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RealtimeConfigurationResponse) *int { return v.MaxConcurrentCalls }).(pulumi.IntPtrOutput)
}

type RealtimeConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (RealtimeConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeConfigurationResponse)(nil)).Elem()
}

func (o RealtimeConfigurationResponsePtrOutput) ToRealtimeConfigurationResponsePtrOutput() RealtimeConfigurationResponsePtrOutput {
	return o
}

func (o RealtimeConfigurationResponsePtrOutput) ToRealtimeConfigurationResponsePtrOutputWithContext(ctx context.Context) RealtimeConfigurationResponsePtrOutput {
	return o
}

func (o RealtimeConfigurationResponsePtrOutput) Elem() RealtimeConfigurationResponseOutput {
	return o.ApplyT(func(v *RealtimeConfigurationResponse) RealtimeConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret RealtimeConfigurationResponse
		return ret
	}).(RealtimeConfigurationResponseOutput)
}

func (o RealtimeConfigurationResponsePtrOutput) MaxConcurrentCalls() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealtimeConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxConcurrentCalls
	}).(pulumi.IntPtrOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuResponseInput interface {
	pulumi.Input

	ToResourceSkuResponseOutput() ResourceSkuResponseOutput
	ToResourceSkuResponseOutputWithContext(context.Context) ResourceSkuResponseOutput
}

type ResourceSkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return i.ToResourceSkuResponseOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput)
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i ResourceSkuResponseArgs) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponseOutput).ToResourceSkuResponsePtrOutputWithContext(ctx)
}









type ResourceSkuResponsePtrInput interface {
	pulumi.Input

	ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput
	ToResourceSkuResponsePtrOutputWithContext(context.Context) ResourceSkuResponsePtrOutput
}

type resourceSkuResponsePtrType ResourceSkuResponseArgs

func ResourceSkuResponsePtr(v *ResourceSkuResponseArgs) ResourceSkuResponsePtrInput {
	return (*resourceSkuResponsePtrType)(v)
}

func (*resourceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return i.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *resourceSkuResponsePtrType) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuResponsePtrOutput)
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o.ToResourceSkuResponsePtrOutputWithContext(context.Background())
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSkuResponse) *ResourceSkuResponse {
		return &v
	}).(ResourceSkuResponsePtrOutput)
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ServiceInputOutputSpecification struct {
	Description *string                       `pulumi:"description"`
	Properties  map[string]TableSpecification `pulumi:"properties"`
	Title       *string                       `pulumi:"title"`
	Type        string                        `pulumi:"type"`
}





type ServiceInputOutputSpecificationInput interface {
	pulumi.Input

	ToServiceInputOutputSpecificationOutput() ServiceInputOutputSpecificationOutput
	ToServiceInputOutputSpecificationOutputWithContext(context.Context) ServiceInputOutputSpecificationOutput
}

type ServiceInputOutputSpecificationArgs struct {
	Description pulumi.StringPtrInput      `pulumi:"description"`
	Properties  TableSpecificationMapInput `pulumi:"properties"`
	Title       pulumi.StringPtrInput      `pulumi:"title"`
	Type        pulumi.StringInput         `pulumi:"type"`
}

func (ServiceInputOutputSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceInputOutputSpecification)(nil)).Elem()
}

func (i ServiceInputOutputSpecificationArgs) ToServiceInputOutputSpecificationOutput() ServiceInputOutputSpecificationOutput {
	return i.ToServiceInputOutputSpecificationOutputWithContext(context.Background())
}

func (i ServiceInputOutputSpecificationArgs) ToServiceInputOutputSpecificationOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceInputOutputSpecificationOutput)
}

func (i ServiceInputOutputSpecificationArgs) ToServiceInputOutputSpecificationPtrOutput() ServiceInputOutputSpecificationPtrOutput {
	return i.ToServiceInputOutputSpecificationPtrOutputWithContext(context.Background())
}

func (i ServiceInputOutputSpecificationArgs) ToServiceInputOutputSpecificationPtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceInputOutputSpecificationOutput).ToServiceInputOutputSpecificationPtrOutputWithContext(ctx)
}









type ServiceInputOutputSpecificationPtrInput interface {
	pulumi.Input

	ToServiceInputOutputSpecificationPtrOutput() ServiceInputOutputSpecificationPtrOutput
	ToServiceInputOutputSpecificationPtrOutputWithContext(context.Context) ServiceInputOutputSpecificationPtrOutput
}

type serviceInputOutputSpecificationPtrType ServiceInputOutputSpecificationArgs

func ServiceInputOutputSpecificationPtr(v *ServiceInputOutputSpecificationArgs) ServiceInputOutputSpecificationPtrInput {
	return (*serviceInputOutputSpecificationPtrType)(v)
}

func (*serviceInputOutputSpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceInputOutputSpecification)(nil)).Elem()
}

func (i *serviceInputOutputSpecificationPtrType) ToServiceInputOutputSpecificationPtrOutput() ServiceInputOutputSpecificationPtrOutput {
	return i.ToServiceInputOutputSpecificationPtrOutputWithContext(context.Background())
}

func (i *serviceInputOutputSpecificationPtrType) ToServiceInputOutputSpecificationPtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceInputOutputSpecificationPtrOutput)
}

type ServiceInputOutputSpecificationOutput struct{ *pulumi.OutputState }

func (ServiceInputOutputSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceInputOutputSpecification)(nil)).Elem()
}

func (o ServiceInputOutputSpecificationOutput) ToServiceInputOutputSpecificationOutput() ServiceInputOutputSpecificationOutput {
	return o
}

func (o ServiceInputOutputSpecificationOutput) ToServiceInputOutputSpecificationOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationOutput {
	return o
}

func (o ServiceInputOutputSpecificationOutput) ToServiceInputOutputSpecificationPtrOutput() ServiceInputOutputSpecificationPtrOutput {
	return o.ToServiceInputOutputSpecificationPtrOutputWithContext(context.Background())
}

func (o ServiceInputOutputSpecificationOutput) ToServiceInputOutputSpecificationPtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceInputOutputSpecification) *ServiceInputOutputSpecification {
		return &v
	}).(ServiceInputOutputSpecificationPtrOutput)
}

func (o ServiceInputOutputSpecificationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecification) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationOutput) Properties() TableSpecificationMapOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecification) map[string]TableSpecification { return v.Properties }).(TableSpecificationMapOutput)
}

func (o ServiceInputOutputSpecificationOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecification) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecification) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceInputOutputSpecificationPtrOutput struct{ *pulumi.OutputState }

func (ServiceInputOutputSpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceInputOutputSpecification)(nil)).Elem()
}

func (o ServiceInputOutputSpecificationPtrOutput) ToServiceInputOutputSpecificationPtrOutput() ServiceInputOutputSpecificationPtrOutput {
	return o
}

func (o ServiceInputOutputSpecificationPtrOutput) ToServiceInputOutputSpecificationPtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationPtrOutput {
	return o
}

func (o ServiceInputOutputSpecificationPtrOutput) Elem() ServiceInputOutputSpecificationOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecification) ServiceInputOutputSpecification {
		if v != nil {
			return *v
		}
		var ret ServiceInputOutputSpecification
		return ret
	}).(ServiceInputOutputSpecificationOutput)
}

func (o ServiceInputOutputSpecificationPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecification) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationPtrOutput) Properties() TableSpecificationMapOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecification) map[string]TableSpecification {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(TableSpecificationMapOutput)
}

func (o ServiceInputOutputSpecificationPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecification) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecification) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ServiceInputOutputSpecificationResponse struct {
	Description *string                               `pulumi:"description"`
	Properties  map[string]TableSpecificationResponse `pulumi:"properties"`
	Title       *string                               `pulumi:"title"`
	Type        string                                `pulumi:"type"`
}





type ServiceInputOutputSpecificationResponseInput interface {
	pulumi.Input

	ToServiceInputOutputSpecificationResponseOutput() ServiceInputOutputSpecificationResponseOutput
	ToServiceInputOutputSpecificationResponseOutputWithContext(context.Context) ServiceInputOutputSpecificationResponseOutput
}

type ServiceInputOutputSpecificationResponseArgs struct {
	Description pulumi.StringPtrInput              `pulumi:"description"`
	Properties  TableSpecificationResponseMapInput `pulumi:"properties"`
	Title       pulumi.StringPtrInput              `pulumi:"title"`
	Type        pulumi.StringInput                 `pulumi:"type"`
}

func (ServiceInputOutputSpecificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceInputOutputSpecificationResponse)(nil)).Elem()
}

func (i ServiceInputOutputSpecificationResponseArgs) ToServiceInputOutputSpecificationResponseOutput() ServiceInputOutputSpecificationResponseOutput {
	return i.ToServiceInputOutputSpecificationResponseOutputWithContext(context.Background())
}

func (i ServiceInputOutputSpecificationResponseArgs) ToServiceInputOutputSpecificationResponseOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceInputOutputSpecificationResponseOutput)
}

func (i ServiceInputOutputSpecificationResponseArgs) ToServiceInputOutputSpecificationResponsePtrOutput() ServiceInputOutputSpecificationResponsePtrOutput {
	return i.ToServiceInputOutputSpecificationResponsePtrOutputWithContext(context.Background())
}

func (i ServiceInputOutputSpecificationResponseArgs) ToServiceInputOutputSpecificationResponsePtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceInputOutputSpecificationResponseOutput).ToServiceInputOutputSpecificationResponsePtrOutputWithContext(ctx)
}









type ServiceInputOutputSpecificationResponsePtrInput interface {
	pulumi.Input

	ToServiceInputOutputSpecificationResponsePtrOutput() ServiceInputOutputSpecificationResponsePtrOutput
	ToServiceInputOutputSpecificationResponsePtrOutputWithContext(context.Context) ServiceInputOutputSpecificationResponsePtrOutput
}

type serviceInputOutputSpecificationResponsePtrType ServiceInputOutputSpecificationResponseArgs

func ServiceInputOutputSpecificationResponsePtr(v *ServiceInputOutputSpecificationResponseArgs) ServiceInputOutputSpecificationResponsePtrInput {
	return (*serviceInputOutputSpecificationResponsePtrType)(v)
}

func (*serviceInputOutputSpecificationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceInputOutputSpecificationResponse)(nil)).Elem()
}

func (i *serviceInputOutputSpecificationResponsePtrType) ToServiceInputOutputSpecificationResponsePtrOutput() ServiceInputOutputSpecificationResponsePtrOutput {
	return i.ToServiceInputOutputSpecificationResponsePtrOutputWithContext(context.Background())
}

func (i *serviceInputOutputSpecificationResponsePtrType) ToServiceInputOutputSpecificationResponsePtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceInputOutputSpecificationResponsePtrOutput)
}

type ServiceInputOutputSpecificationResponseOutput struct{ *pulumi.OutputState }

func (ServiceInputOutputSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceInputOutputSpecificationResponse)(nil)).Elem()
}

func (o ServiceInputOutputSpecificationResponseOutput) ToServiceInputOutputSpecificationResponseOutput() ServiceInputOutputSpecificationResponseOutput {
	return o
}

func (o ServiceInputOutputSpecificationResponseOutput) ToServiceInputOutputSpecificationResponseOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationResponseOutput {
	return o
}

func (o ServiceInputOutputSpecificationResponseOutput) ToServiceInputOutputSpecificationResponsePtrOutput() ServiceInputOutputSpecificationResponsePtrOutput {
	return o.ToServiceInputOutputSpecificationResponsePtrOutputWithContext(context.Background())
}

func (o ServiceInputOutputSpecificationResponseOutput) ToServiceInputOutputSpecificationResponsePtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceInputOutputSpecificationResponse) *ServiceInputOutputSpecificationResponse {
		return &v
	}).(ServiceInputOutputSpecificationResponsePtrOutput)
}

func (o ServiceInputOutputSpecificationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecificationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationResponseOutput) Properties() TableSpecificationResponseMapOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecificationResponse) map[string]TableSpecificationResponse {
		return v.Properties
	}).(TableSpecificationResponseMapOutput)
}

func (o ServiceInputOutputSpecificationResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecificationResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceInputOutputSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceInputOutputSpecificationResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceInputOutputSpecificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceInputOutputSpecificationResponse)(nil)).Elem()
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) ToServiceInputOutputSpecificationResponsePtrOutput() ServiceInputOutputSpecificationResponsePtrOutput {
	return o
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) ToServiceInputOutputSpecificationResponsePtrOutputWithContext(ctx context.Context) ServiceInputOutputSpecificationResponsePtrOutput {
	return o
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) Elem() ServiceInputOutputSpecificationResponseOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecificationResponse) ServiceInputOutputSpecificationResponse {
		if v != nil {
			return *v
		}
		var ret ServiceInputOutputSpecificationResponse
		return ret
	}).(ServiceInputOutputSpecificationResponseOutput)
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecificationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) Properties() TableSpecificationResponseMapOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecificationResponse) map[string]TableSpecificationResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(TableSpecificationResponseMapOutput)
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecificationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

func (o ServiceInputOutputSpecificationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceInputOutputSpecificationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	Key  *string `pulumi:"key"`
	Name *string `pulumi:"name"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Key  pulumi.StringPtrInput `pulumi:"key"`
	Name pulumi.StringPtrInput `pulumi:"name"`
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

func (i StorageAccountArgs) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput).ToStorageAccountPtrOutputWithContext(ctx)
}









type StorageAccountPtrInput interface {
	pulumi.Input

	ToStorageAccountPtrOutput() StorageAccountPtrOutput
	ToStorageAccountPtrOutputWithContext(context.Context) StorageAccountPtrOutput
}

type storageAccountPtrType StorageAccountArgs

func StorageAccountPtr(v *StorageAccountArgs) StorageAccountPtrInput {
	return (*storageAccountPtrType)(v)
}

func (*storageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPtrOutput)
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

func (o StorageAccountOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (o StorageAccountOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccount) *StorageAccount {
		return &v
	}).(StorageAccountPtrOutput)
}

func (o StorageAccountOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StorageAccountPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) Elem() StorageAccountOutput {
	return o.ApplyT(func(v *StorageAccount) StorageAccount {
		if v != nil {
			return *v
		}
		var ret StorageAccount
		return ret
	}).(StorageAccountOutput)
}

func (o StorageAccountPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageAccountResponse struct {
	Key  *string `pulumi:"key"`
	Name *string `pulumi:"name"`
}





type StorageAccountResponseInput interface {
	pulumi.Input

	ToStorageAccountResponseOutput() StorageAccountResponseOutput
	ToStorageAccountResponseOutputWithContext(context.Context) StorageAccountResponseOutput
}

type StorageAccountResponseArgs struct {
	Key  pulumi.StringPtrInput `pulumi:"key"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (StorageAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return i.ToStorageAccountResponseOutputWithContext(context.Background())
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseOutput)
}

func (i StorageAccountResponseArgs) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return i.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountResponseArgs) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseOutput).ToStorageAccountResponsePtrOutputWithContext(ctx)
}









type StorageAccountResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput
	ToStorageAccountResponsePtrOutputWithContext(context.Context) StorageAccountResponsePtrOutput
}

type storageAccountResponsePtrType StorageAccountResponseArgs

func StorageAccountResponsePtr(v *StorageAccountResponseArgs) StorageAccountResponsePtrInput {
	return (*storageAccountResponsePtrType)(v)
}

func (*storageAccountResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountResponse)(nil)).Elem()
}

func (i *storageAccountResponsePtrType) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return i.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountResponsePtrType) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponsePtrOutput)
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

func (o StorageAccountResponseOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountResponseOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountResponse) *StorageAccountResponse {
		return &v
	}).(StorageAccountResponsePtrOutput)
}

func (o StorageAccountResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) Elem() StorageAccountResponseOutput {
	return o.ApplyT(func(v *StorageAccountResponse) StorageAccountResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountResponse
		return ret
	}).(StorageAccountResponseOutput)
}

func (o StorageAccountResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type TableSpecification struct {
	Description *string                        `pulumi:"description"`
	Format      *string                        `pulumi:"format"`
	Properties  map[string]ColumnSpecification `pulumi:"properties"`
	Title       *string                        `pulumi:"title"`
	Type        string                         `pulumi:"type"`
}





type TableSpecificationInput interface {
	pulumi.Input

	ToTableSpecificationOutput() TableSpecificationOutput
	ToTableSpecificationOutputWithContext(context.Context) TableSpecificationOutput
}

type TableSpecificationArgs struct {
	Description pulumi.StringPtrInput       `pulumi:"description"`
	Format      pulumi.StringPtrInput       `pulumi:"format"`
	Properties  ColumnSpecificationMapInput `pulumi:"properties"`
	Title       pulumi.StringPtrInput       `pulumi:"title"`
	Type        pulumi.StringInput          `pulumi:"type"`
}

func (TableSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableSpecification)(nil)).Elem()
}

func (i TableSpecificationArgs) ToTableSpecificationOutput() TableSpecificationOutput {
	return i.ToTableSpecificationOutputWithContext(context.Background())
}

func (i TableSpecificationArgs) ToTableSpecificationOutputWithContext(ctx context.Context) TableSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableSpecificationOutput)
}





type TableSpecificationMapInput interface {
	pulumi.Input

	ToTableSpecificationMapOutput() TableSpecificationMapOutput
	ToTableSpecificationMapOutputWithContext(context.Context) TableSpecificationMapOutput
}

type TableSpecificationMap map[string]TableSpecificationInput

func (TableSpecificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TableSpecification)(nil)).Elem()
}

func (i TableSpecificationMap) ToTableSpecificationMapOutput() TableSpecificationMapOutput {
	return i.ToTableSpecificationMapOutputWithContext(context.Background())
}

func (i TableSpecificationMap) ToTableSpecificationMapOutputWithContext(ctx context.Context) TableSpecificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableSpecificationMapOutput)
}

type TableSpecificationOutput struct{ *pulumi.OutputState }

func (TableSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableSpecification)(nil)).Elem()
}

func (o TableSpecificationOutput) ToTableSpecificationOutput() TableSpecificationOutput {
	return o
}

func (o TableSpecificationOutput) ToTableSpecificationOutputWithContext(ctx context.Context) TableSpecificationOutput {
	return o
}

func (o TableSpecificationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableSpecification) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TableSpecificationOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableSpecification) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o TableSpecificationOutput) Properties() ColumnSpecificationMapOutput {
	return o.ApplyT(func(v TableSpecification) map[string]ColumnSpecification { return v.Properties }).(ColumnSpecificationMapOutput)
}

func (o TableSpecificationOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableSpecification) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o TableSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TableSpecification) string { return v.Type }).(pulumi.StringOutput)
}

type TableSpecificationMapOutput struct{ *pulumi.OutputState }

func (TableSpecificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TableSpecification)(nil)).Elem()
}

func (o TableSpecificationMapOutput) ToTableSpecificationMapOutput() TableSpecificationMapOutput {
	return o
}

func (o TableSpecificationMapOutput) ToTableSpecificationMapOutputWithContext(ctx context.Context) TableSpecificationMapOutput {
	return o
}

func (o TableSpecificationMapOutput) MapIndex(k pulumi.StringInput) TableSpecificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TableSpecification {
		return vs[0].(map[string]TableSpecification)[vs[1].(string)]
	}).(TableSpecificationOutput)
}

type TableSpecificationResponse struct {
	Description *string                                `pulumi:"description"`
	Format      *string                                `pulumi:"format"`
	Properties  map[string]ColumnSpecificationResponse `pulumi:"properties"`
	Title       *string                                `pulumi:"title"`
	Type        string                                 `pulumi:"type"`
}





type TableSpecificationResponseInput interface {
	pulumi.Input

	ToTableSpecificationResponseOutput() TableSpecificationResponseOutput
	ToTableSpecificationResponseOutputWithContext(context.Context) TableSpecificationResponseOutput
}

type TableSpecificationResponseArgs struct {
	Description pulumi.StringPtrInput               `pulumi:"description"`
	Format      pulumi.StringPtrInput               `pulumi:"format"`
	Properties  ColumnSpecificationResponseMapInput `pulumi:"properties"`
	Title       pulumi.StringPtrInput               `pulumi:"title"`
	Type        pulumi.StringInput                  `pulumi:"type"`
}

func (TableSpecificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableSpecificationResponse)(nil)).Elem()
}

func (i TableSpecificationResponseArgs) ToTableSpecificationResponseOutput() TableSpecificationResponseOutput {
	return i.ToTableSpecificationResponseOutputWithContext(context.Background())
}

func (i TableSpecificationResponseArgs) ToTableSpecificationResponseOutputWithContext(ctx context.Context) TableSpecificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableSpecificationResponseOutput)
}





type TableSpecificationResponseMapInput interface {
	pulumi.Input

	ToTableSpecificationResponseMapOutput() TableSpecificationResponseMapOutput
	ToTableSpecificationResponseMapOutputWithContext(context.Context) TableSpecificationResponseMapOutput
}

type TableSpecificationResponseMap map[string]TableSpecificationResponseInput

func (TableSpecificationResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TableSpecificationResponse)(nil)).Elem()
}

func (i TableSpecificationResponseMap) ToTableSpecificationResponseMapOutput() TableSpecificationResponseMapOutput {
	return i.ToTableSpecificationResponseMapOutputWithContext(context.Background())
}

func (i TableSpecificationResponseMap) ToTableSpecificationResponseMapOutputWithContext(ctx context.Context) TableSpecificationResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableSpecificationResponseMapOutput)
}

type TableSpecificationResponseOutput struct{ *pulumi.OutputState }

func (TableSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableSpecificationResponse)(nil)).Elem()
}

func (o TableSpecificationResponseOutput) ToTableSpecificationResponseOutput() TableSpecificationResponseOutput {
	return o
}

func (o TableSpecificationResponseOutput) ToTableSpecificationResponseOutputWithContext(ctx context.Context) TableSpecificationResponseOutput {
	return o
}

func (o TableSpecificationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableSpecificationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TableSpecificationResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableSpecificationResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o TableSpecificationResponseOutput) Properties() ColumnSpecificationResponseMapOutput {
	return o.ApplyT(func(v TableSpecificationResponse) map[string]ColumnSpecificationResponse { return v.Properties }).(ColumnSpecificationResponseMapOutput)
}

func (o TableSpecificationResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableSpecificationResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o TableSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TableSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TableSpecificationResponseMapOutput struct{ *pulumi.OutputState }

func (TableSpecificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TableSpecificationResponse)(nil)).Elem()
}

func (o TableSpecificationResponseMapOutput) ToTableSpecificationResponseMapOutput() TableSpecificationResponseMapOutput {
	return o
}

func (o TableSpecificationResponseMapOutput) ToTableSpecificationResponseMapOutputWithContext(ctx context.Context) TableSpecificationResponseMapOutput {
	return o
}

func (o TableSpecificationResponseMapOutput) MapIndex(k pulumi.StringInput) TableSpecificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TableSpecificationResponse {
		return vs[0].(map[string]TableSpecificationResponse)[vs[1].(string)]
	}).(TableSpecificationResponseOutput)
}

type WebServiceKeys struct {
	Primary   *string `pulumi:"primary"`
	Secondary *string `pulumi:"secondary"`
}





type WebServiceKeysInput interface {
	pulumi.Input

	ToWebServiceKeysOutput() WebServiceKeysOutput
	ToWebServiceKeysOutputWithContext(context.Context) WebServiceKeysOutput
}

type WebServiceKeysArgs struct {
	Primary   pulumi.StringPtrInput `pulumi:"primary"`
	Secondary pulumi.StringPtrInput `pulumi:"secondary"`
}

func (WebServiceKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceKeys)(nil)).Elem()
}

func (i WebServiceKeysArgs) ToWebServiceKeysOutput() WebServiceKeysOutput {
	return i.ToWebServiceKeysOutputWithContext(context.Background())
}

func (i WebServiceKeysArgs) ToWebServiceKeysOutputWithContext(ctx context.Context) WebServiceKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceKeysOutput)
}

func (i WebServiceKeysArgs) ToWebServiceKeysPtrOutput() WebServiceKeysPtrOutput {
	return i.ToWebServiceKeysPtrOutputWithContext(context.Background())
}

func (i WebServiceKeysArgs) ToWebServiceKeysPtrOutputWithContext(ctx context.Context) WebServiceKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceKeysOutput).ToWebServiceKeysPtrOutputWithContext(ctx)
}









type WebServiceKeysPtrInput interface {
	pulumi.Input

	ToWebServiceKeysPtrOutput() WebServiceKeysPtrOutput
	ToWebServiceKeysPtrOutputWithContext(context.Context) WebServiceKeysPtrOutput
}

type webServiceKeysPtrType WebServiceKeysArgs

func WebServiceKeysPtr(v *WebServiceKeysArgs) WebServiceKeysPtrInput {
	return (*webServiceKeysPtrType)(v)
}

func (*webServiceKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServiceKeys)(nil)).Elem()
}

func (i *webServiceKeysPtrType) ToWebServiceKeysPtrOutput() WebServiceKeysPtrOutput {
	return i.ToWebServiceKeysPtrOutputWithContext(context.Background())
}

func (i *webServiceKeysPtrType) ToWebServiceKeysPtrOutputWithContext(ctx context.Context) WebServiceKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceKeysPtrOutput)
}

type WebServiceKeysOutput struct{ *pulumi.OutputState }

func (WebServiceKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceKeys)(nil)).Elem()
}

func (o WebServiceKeysOutput) ToWebServiceKeysOutput() WebServiceKeysOutput {
	return o
}

func (o WebServiceKeysOutput) ToWebServiceKeysOutputWithContext(ctx context.Context) WebServiceKeysOutput {
	return o
}

func (o WebServiceKeysOutput) ToWebServiceKeysPtrOutput() WebServiceKeysPtrOutput {
	return o.ToWebServiceKeysPtrOutputWithContext(context.Background())
}

func (o WebServiceKeysOutput) ToWebServiceKeysPtrOutputWithContext(ctx context.Context) WebServiceKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebServiceKeys) *WebServiceKeys {
		return &v
	}).(WebServiceKeysPtrOutput)
}

func (o WebServiceKeysOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServiceKeys) *string { return v.Primary }).(pulumi.StringPtrOutput)
}

func (o WebServiceKeysOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServiceKeys) *string { return v.Secondary }).(pulumi.StringPtrOutput)
}

type WebServiceKeysPtrOutput struct{ *pulumi.OutputState }

func (WebServiceKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServiceKeys)(nil)).Elem()
}

func (o WebServiceKeysPtrOutput) ToWebServiceKeysPtrOutput() WebServiceKeysPtrOutput {
	return o
}

func (o WebServiceKeysPtrOutput) ToWebServiceKeysPtrOutputWithContext(ctx context.Context) WebServiceKeysPtrOutput {
	return o
}

func (o WebServiceKeysPtrOutput) Elem() WebServiceKeysOutput {
	return o.ApplyT(func(v *WebServiceKeys) WebServiceKeys {
		if v != nil {
			return *v
		}
		var ret WebServiceKeys
		return ret
	}).(WebServiceKeysOutput)
}

func (o WebServiceKeysPtrOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServiceKeys) *string {
		if v == nil {
			return nil
		}
		return v.Primary
	}).(pulumi.StringPtrOutput)
}

func (o WebServiceKeysPtrOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServiceKeys) *string {
		if v == nil {
			return nil
		}
		return v.Secondary
	}).(pulumi.StringPtrOutput)
}

type WebServiceKeysResponse struct {
	Primary   *string `pulumi:"primary"`
	Secondary *string `pulumi:"secondary"`
}





type WebServiceKeysResponseInput interface {
	pulumi.Input

	ToWebServiceKeysResponseOutput() WebServiceKeysResponseOutput
	ToWebServiceKeysResponseOutputWithContext(context.Context) WebServiceKeysResponseOutput
}

type WebServiceKeysResponseArgs struct {
	Primary   pulumi.StringPtrInput `pulumi:"primary"`
	Secondary pulumi.StringPtrInput `pulumi:"secondary"`
}

func (WebServiceKeysResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceKeysResponse)(nil)).Elem()
}

func (i WebServiceKeysResponseArgs) ToWebServiceKeysResponseOutput() WebServiceKeysResponseOutput {
	return i.ToWebServiceKeysResponseOutputWithContext(context.Background())
}

func (i WebServiceKeysResponseArgs) ToWebServiceKeysResponseOutputWithContext(ctx context.Context) WebServiceKeysResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceKeysResponseOutput)
}

func (i WebServiceKeysResponseArgs) ToWebServiceKeysResponsePtrOutput() WebServiceKeysResponsePtrOutput {
	return i.ToWebServiceKeysResponsePtrOutputWithContext(context.Background())
}

func (i WebServiceKeysResponseArgs) ToWebServiceKeysResponsePtrOutputWithContext(ctx context.Context) WebServiceKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceKeysResponseOutput).ToWebServiceKeysResponsePtrOutputWithContext(ctx)
}









type WebServiceKeysResponsePtrInput interface {
	pulumi.Input

	ToWebServiceKeysResponsePtrOutput() WebServiceKeysResponsePtrOutput
	ToWebServiceKeysResponsePtrOutputWithContext(context.Context) WebServiceKeysResponsePtrOutput
}

type webServiceKeysResponsePtrType WebServiceKeysResponseArgs

func WebServiceKeysResponsePtr(v *WebServiceKeysResponseArgs) WebServiceKeysResponsePtrInput {
	return (*webServiceKeysResponsePtrType)(v)
}

func (*webServiceKeysResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServiceKeysResponse)(nil)).Elem()
}

func (i *webServiceKeysResponsePtrType) ToWebServiceKeysResponsePtrOutput() WebServiceKeysResponsePtrOutput {
	return i.ToWebServiceKeysResponsePtrOutputWithContext(context.Background())
}

func (i *webServiceKeysResponsePtrType) ToWebServiceKeysResponsePtrOutputWithContext(ctx context.Context) WebServiceKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceKeysResponsePtrOutput)
}

type WebServiceKeysResponseOutput struct{ *pulumi.OutputState }

func (WebServiceKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceKeysResponse)(nil)).Elem()
}

func (o WebServiceKeysResponseOutput) ToWebServiceKeysResponseOutput() WebServiceKeysResponseOutput {
	return o
}

func (o WebServiceKeysResponseOutput) ToWebServiceKeysResponseOutputWithContext(ctx context.Context) WebServiceKeysResponseOutput {
	return o
}

func (o WebServiceKeysResponseOutput) ToWebServiceKeysResponsePtrOutput() WebServiceKeysResponsePtrOutput {
	return o.ToWebServiceKeysResponsePtrOutputWithContext(context.Background())
}

func (o WebServiceKeysResponseOutput) ToWebServiceKeysResponsePtrOutputWithContext(ctx context.Context) WebServiceKeysResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebServiceKeysResponse) *WebServiceKeysResponse {
		return &v
	}).(WebServiceKeysResponsePtrOutput)
}

func (o WebServiceKeysResponseOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServiceKeysResponse) *string { return v.Primary }).(pulumi.StringPtrOutput)
}

func (o WebServiceKeysResponseOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServiceKeysResponse) *string { return v.Secondary }).(pulumi.StringPtrOutput)
}

type WebServiceKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (WebServiceKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServiceKeysResponse)(nil)).Elem()
}

func (o WebServiceKeysResponsePtrOutput) ToWebServiceKeysResponsePtrOutput() WebServiceKeysResponsePtrOutput {
	return o
}

func (o WebServiceKeysResponsePtrOutput) ToWebServiceKeysResponsePtrOutputWithContext(ctx context.Context) WebServiceKeysResponsePtrOutput {
	return o
}

func (o WebServiceKeysResponsePtrOutput) Elem() WebServiceKeysResponseOutput {
	return o.ApplyT(func(v *WebServiceKeysResponse) WebServiceKeysResponse {
		if v != nil {
			return *v
		}
		var ret WebServiceKeysResponse
		return ret
	}).(WebServiceKeysResponseOutput)
}

func (o WebServiceKeysResponsePtrOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServiceKeysResponse) *string {
		if v == nil {
			return nil
		}
		return v.Primary
	}).(pulumi.StringPtrOutput)
}

func (o WebServiceKeysResponsePtrOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServiceKeysResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secondary
	}).(pulumi.StringPtrOutput)
}

type WebServiceParameter struct {
	CertificateThumbprint *string     `pulumi:"certificateThumbprint"`
	Value                 interface{} `pulumi:"value"`
}





type WebServiceParameterInput interface {
	pulumi.Input

	ToWebServiceParameterOutput() WebServiceParameterOutput
	ToWebServiceParameterOutputWithContext(context.Context) WebServiceParameterOutput
}

type WebServiceParameterArgs struct {
	CertificateThumbprint pulumi.StringPtrInput `pulumi:"certificateThumbprint"`
	Value                 pulumi.Input          `pulumi:"value"`
}

func (WebServiceParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceParameter)(nil)).Elem()
}

func (i WebServiceParameterArgs) ToWebServiceParameterOutput() WebServiceParameterOutput {
	return i.ToWebServiceParameterOutputWithContext(context.Background())
}

func (i WebServiceParameterArgs) ToWebServiceParameterOutputWithContext(ctx context.Context) WebServiceParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceParameterOutput)
}





type WebServiceParameterMapInput interface {
	pulumi.Input

	ToWebServiceParameterMapOutput() WebServiceParameterMapOutput
	ToWebServiceParameterMapOutputWithContext(context.Context) WebServiceParameterMapOutput
}

type WebServiceParameterMap map[string]WebServiceParameterInput

func (WebServiceParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebServiceParameter)(nil)).Elem()
}

func (i WebServiceParameterMap) ToWebServiceParameterMapOutput() WebServiceParameterMapOutput {
	return i.ToWebServiceParameterMapOutputWithContext(context.Background())
}

func (i WebServiceParameterMap) ToWebServiceParameterMapOutputWithContext(ctx context.Context) WebServiceParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceParameterMapOutput)
}

type WebServiceParameterOutput struct{ *pulumi.OutputState }

func (WebServiceParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceParameter)(nil)).Elem()
}

func (o WebServiceParameterOutput) ToWebServiceParameterOutput() WebServiceParameterOutput {
	return o
}

func (o WebServiceParameterOutput) ToWebServiceParameterOutputWithContext(ctx context.Context) WebServiceParameterOutput {
	return o
}

func (o WebServiceParameterOutput) CertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServiceParameter) *string { return v.CertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o WebServiceParameterOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WebServiceParameter) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WebServiceParameterMapOutput struct{ *pulumi.OutputState }

func (WebServiceParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebServiceParameter)(nil)).Elem()
}

func (o WebServiceParameterMapOutput) ToWebServiceParameterMapOutput() WebServiceParameterMapOutput {
	return o
}

func (o WebServiceParameterMapOutput) ToWebServiceParameterMapOutputWithContext(ctx context.Context) WebServiceParameterMapOutput {
	return o
}

func (o WebServiceParameterMapOutput) MapIndex(k pulumi.StringInput) WebServiceParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebServiceParameter {
		return vs[0].(map[string]WebServiceParameter)[vs[1].(string)]
	}).(WebServiceParameterOutput)
}

type WebServiceParameterResponse struct {
	CertificateThumbprint *string     `pulumi:"certificateThumbprint"`
	Value                 interface{} `pulumi:"value"`
}





type WebServiceParameterResponseInput interface {
	pulumi.Input

	ToWebServiceParameterResponseOutput() WebServiceParameterResponseOutput
	ToWebServiceParameterResponseOutputWithContext(context.Context) WebServiceParameterResponseOutput
}

type WebServiceParameterResponseArgs struct {
	CertificateThumbprint pulumi.StringPtrInput `pulumi:"certificateThumbprint"`
	Value                 pulumi.Input          `pulumi:"value"`
}

func (WebServiceParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceParameterResponse)(nil)).Elem()
}

func (i WebServiceParameterResponseArgs) ToWebServiceParameterResponseOutput() WebServiceParameterResponseOutput {
	return i.ToWebServiceParameterResponseOutputWithContext(context.Background())
}

func (i WebServiceParameterResponseArgs) ToWebServiceParameterResponseOutputWithContext(ctx context.Context) WebServiceParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceParameterResponseOutput)
}





type WebServiceParameterResponseMapInput interface {
	pulumi.Input

	ToWebServiceParameterResponseMapOutput() WebServiceParameterResponseMapOutput
	ToWebServiceParameterResponseMapOutputWithContext(context.Context) WebServiceParameterResponseMapOutput
}

type WebServiceParameterResponseMap map[string]WebServiceParameterResponseInput

func (WebServiceParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebServiceParameterResponse)(nil)).Elem()
}

func (i WebServiceParameterResponseMap) ToWebServiceParameterResponseMapOutput() WebServiceParameterResponseMapOutput {
	return i.ToWebServiceParameterResponseMapOutputWithContext(context.Background())
}

func (i WebServiceParameterResponseMap) ToWebServiceParameterResponseMapOutputWithContext(ctx context.Context) WebServiceParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceParameterResponseMapOutput)
}

type WebServiceParameterResponseOutput struct{ *pulumi.OutputState }

func (WebServiceParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServiceParameterResponse)(nil)).Elem()
}

func (o WebServiceParameterResponseOutput) ToWebServiceParameterResponseOutput() WebServiceParameterResponseOutput {
	return o
}

func (o WebServiceParameterResponseOutput) ToWebServiceParameterResponseOutputWithContext(ctx context.Context) WebServiceParameterResponseOutput {
	return o
}

func (o WebServiceParameterResponseOutput) CertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServiceParameterResponse) *string { return v.CertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o WebServiceParameterResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WebServiceParameterResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WebServiceParameterResponseMapOutput struct{ *pulumi.OutputState }

func (WebServiceParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebServiceParameterResponse)(nil)).Elem()
}

func (o WebServiceParameterResponseMapOutput) ToWebServiceParameterResponseMapOutput() WebServiceParameterResponseMapOutput {
	return o
}

func (o WebServiceParameterResponseMapOutput) ToWebServiceParameterResponseMapOutputWithContext(ctx context.Context) WebServiceParameterResponseMapOutput {
	return o
}

func (o WebServiceParameterResponseMapOutput) MapIndex(k pulumi.StringInput) WebServiceParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebServiceParameterResponse {
		return vs[0].(map[string]WebServiceParameterResponse)[vs[1].(string)]
	}).(WebServiceParameterResponseOutput)
}

type WebServicePropertiesForGraph struct {
	Assets                   map[string]AssetItem             `pulumi:"assets"`
	CommitmentPlan           *CommitmentPlanType              `pulumi:"commitmentPlan"`
	Description              *string                          `pulumi:"description"`
	Diagnostics              *DiagnosticsConfiguration        `pulumi:"diagnostics"`
	ExampleRequest           *ExampleRequest                  `pulumi:"exampleRequest"`
	ExposeSampleData         *bool                            `pulumi:"exposeSampleData"`
	Input                    *ServiceInputOutputSpecification `pulumi:"input"`
	Keys                     *WebServiceKeys                  `pulumi:"keys"`
	MachineLearningWorkspace *MachineLearningWorkspace        `pulumi:"machineLearningWorkspace"`
	Output                   *ServiceInputOutputSpecification `pulumi:"output"`
	Package                  *GraphPackage                    `pulumi:"package"`
	PackageType              string                           `pulumi:"packageType"`
	Parameters               map[string]WebServiceParameter   `pulumi:"parameters"`
	PayloadsInBlobStorage    *bool                            `pulumi:"payloadsInBlobStorage"`
	PayloadsLocation         *BlobLocation                    `pulumi:"payloadsLocation"`
	ReadOnly                 *bool                            `pulumi:"readOnly"`
	RealtimeConfiguration    *RealtimeConfiguration           `pulumi:"realtimeConfiguration"`
	StorageAccount           *StorageAccount                  `pulumi:"storageAccount"`
	Title                    *string                          `pulumi:"title"`
}





type WebServicePropertiesForGraphInput interface {
	pulumi.Input

	ToWebServicePropertiesForGraphOutput() WebServicePropertiesForGraphOutput
	ToWebServicePropertiesForGraphOutputWithContext(context.Context) WebServicePropertiesForGraphOutput
}

type WebServicePropertiesForGraphArgs struct {
	Assets                   AssetItemMapInput                       `pulumi:"assets"`
	CommitmentPlan           CommitmentPlanTypePtrInput              `pulumi:"commitmentPlan"`
	Description              pulumi.StringPtrInput                   `pulumi:"description"`
	Diagnostics              DiagnosticsConfigurationPtrInput        `pulumi:"diagnostics"`
	ExampleRequest           ExampleRequestPtrInput                  `pulumi:"exampleRequest"`
	ExposeSampleData         pulumi.BoolPtrInput                     `pulumi:"exposeSampleData"`
	Input                    ServiceInputOutputSpecificationPtrInput `pulumi:"input"`
	Keys                     WebServiceKeysPtrInput                  `pulumi:"keys"`
	MachineLearningWorkspace MachineLearningWorkspacePtrInput        `pulumi:"machineLearningWorkspace"`
	Output                   ServiceInputOutputSpecificationPtrInput `pulumi:"output"`
	Package                  GraphPackagePtrInput                    `pulumi:"package"`
	PackageType              pulumi.StringInput                      `pulumi:"packageType"`
	Parameters               WebServiceParameterMapInput             `pulumi:"parameters"`
	PayloadsInBlobStorage    pulumi.BoolPtrInput                     `pulumi:"payloadsInBlobStorage"`
	PayloadsLocation         BlobLocationPtrInput                    `pulumi:"payloadsLocation"`
	ReadOnly                 pulumi.BoolPtrInput                     `pulumi:"readOnly"`
	RealtimeConfiguration    RealtimeConfigurationPtrInput           `pulumi:"realtimeConfiguration"`
	StorageAccount           StorageAccountPtrInput                  `pulumi:"storageAccount"`
	Title                    pulumi.StringPtrInput                   `pulumi:"title"`
}

func (WebServicePropertiesForGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServicePropertiesForGraph)(nil)).Elem()
}

func (i WebServicePropertiesForGraphArgs) ToWebServicePropertiesForGraphOutput() WebServicePropertiesForGraphOutput {
	return i.ToWebServicePropertiesForGraphOutputWithContext(context.Background())
}

func (i WebServicePropertiesForGraphArgs) ToWebServicePropertiesForGraphOutputWithContext(ctx context.Context) WebServicePropertiesForGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServicePropertiesForGraphOutput)
}

func (i WebServicePropertiesForGraphArgs) ToWebServicePropertiesForGraphPtrOutput() WebServicePropertiesForGraphPtrOutput {
	return i.ToWebServicePropertiesForGraphPtrOutputWithContext(context.Background())
}

func (i WebServicePropertiesForGraphArgs) ToWebServicePropertiesForGraphPtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServicePropertiesForGraphOutput).ToWebServicePropertiesForGraphPtrOutputWithContext(ctx)
}









type WebServicePropertiesForGraphPtrInput interface {
	pulumi.Input

	ToWebServicePropertiesForGraphPtrOutput() WebServicePropertiesForGraphPtrOutput
	ToWebServicePropertiesForGraphPtrOutputWithContext(context.Context) WebServicePropertiesForGraphPtrOutput
}

type webServicePropertiesForGraphPtrType WebServicePropertiesForGraphArgs

func WebServicePropertiesForGraphPtr(v *WebServicePropertiesForGraphArgs) WebServicePropertiesForGraphPtrInput {
	return (*webServicePropertiesForGraphPtrType)(v)
}

func (*webServicePropertiesForGraphPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServicePropertiesForGraph)(nil)).Elem()
}

func (i *webServicePropertiesForGraphPtrType) ToWebServicePropertiesForGraphPtrOutput() WebServicePropertiesForGraphPtrOutput {
	return i.ToWebServicePropertiesForGraphPtrOutputWithContext(context.Background())
}

func (i *webServicePropertiesForGraphPtrType) ToWebServicePropertiesForGraphPtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServicePropertiesForGraphPtrOutput)
}

type WebServicePropertiesForGraphOutput struct{ *pulumi.OutputState }

func (WebServicePropertiesForGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServicePropertiesForGraph)(nil)).Elem()
}

func (o WebServicePropertiesForGraphOutput) ToWebServicePropertiesForGraphOutput() WebServicePropertiesForGraphOutput {
	return o
}

func (o WebServicePropertiesForGraphOutput) ToWebServicePropertiesForGraphOutputWithContext(ctx context.Context) WebServicePropertiesForGraphOutput {
	return o
}

func (o WebServicePropertiesForGraphOutput) ToWebServicePropertiesForGraphPtrOutput() WebServicePropertiesForGraphPtrOutput {
	return o.ToWebServicePropertiesForGraphPtrOutputWithContext(context.Background())
}

func (o WebServicePropertiesForGraphOutput) ToWebServicePropertiesForGraphPtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebServicePropertiesForGraph) *WebServicePropertiesForGraph {
		return &v
	}).(WebServicePropertiesForGraphPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Assets() AssetItemMapOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) map[string]AssetItem { return v.Assets }).(AssetItemMapOutput)
}

func (o WebServicePropertiesForGraphOutput) CommitmentPlan() CommitmentPlanTypePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *CommitmentPlanType { return v.CommitmentPlan }).(CommitmentPlanTypePtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Diagnostics() DiagnosticsConfigurationPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *DiagnosticsConfiguration { return v.Diagnostics }).(DiagnosticsConfigurationPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) ExampleRequest() ExampleRequestPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *ExampleRequest { return v.ExampleRequest }).(ExampleRequestPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) ExposeSampleData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *bool { return v.ExposeSampleData }).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Input() ServiceInputOutputSpecificationPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *ServiceInputOutputSpecification { return v.Input }).(ServiceInputOutputSpecificationPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Keys() WebServiceKeysPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *WebServiceKeys { return v.Keys }).(WebServiceKeysPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) MachineLearningWorkspace() MachineLearningWorkspacePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *MachineLearningWorkspace { return v.MachineLearningWorkspace }).(MachineLearningWorkspacePtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Output() ServiceInputOutputSpecificationPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *ServiceInputOutputSpecification { return v.Output }).(ServiceInputOutputSpecificationPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Package() GraphPackagePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *GraphPackage { return v.Package }).(GraphPackagePtrOutput)
}

func (o WebServicePropertiesForGraphOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o WebServicePropertiesForGraphOutput) Parameters() WebServiceParameterMapOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) map[string]WebServiceParameter { return v.Parameters }).(WebServiceParameterMapOutput)
}

func (o WebServicePropertiesForGraphOutput) PayloadsInBlobStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *bool { return v.PayloadsInBlobStorage }).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) PayloadsLocation() BlobLocationPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *BlobLocation { return v.PayloadsLocation }).(BlobLocationPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) RealtimeConfiguration() RealtimeConfigurationPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *RealtimeConfiguration { return v.RealtimeConfiguration }).(RealtimeConfigurationPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) StorageAccount() StorageAccountPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *StorageAccount { return v.StorageAccount }).(StorageAccountPtrOutput)
}

func (o WebServicePropertiesForGraphOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraph) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type WebServicePropertiesForGraphPtrOutput struct{ *pulumi.OutputState }

func (WebServicePropertiesForGraphPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServicePropertiesForGraph)(nil)).Elem()
}

func (o WebServicePropertiesForGraphPtrOutput) ToWebServicePropertiesForGraphPtrOutput() WebServicePropertiesForGraphPtrOutput {
	return o
}

func (o WebServicePropertiesForGraphPtrOutput) ToWebServicePropertiesForGraphPtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphPtrOutput {
	return o
}

func (o WebServicePropertiesForGraphPtrOutput) Elem() WebServicePropertiesForGraphOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) WebServicePropertiesForGraph {
		if v != nil {
			return *v
		}
		var ret WebServicePropertiesForGraph
		return ret
	}).(WebServicePropertiesForGraphOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Assets() AssetItemMapOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) map[string]AssetItem {
		if v == nil {
			return nil
		}
		return v.Assets
	}).(AssetItemMapOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) CommitmentPlan() CommitmentPlanTypePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *CommitmentPlanType {
		if v == nil {
			return nil
		}
		return v.CommitmentPlan
	}).(CommitmentPlanTypePtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Diagnostics() DiagnosticsConfigurationPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *DiagnosticsConfiguration {
		if v == nil {
			return nil
		}
		return v.Diagnostics
	}).(DiagnosticsConfigurationPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) ExampleRequest() ExampleRequestPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *ExampleRequest {
		if v == nil {
			return nil
		}
		return v.ExampleRequest
	}).(ExampleRequestPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) ExposeSampleData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *bool {
		if v == nil {
			return nil
		}
		return v.ExposeSampleData
	}).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Input() ServiceInputOutputSpecificationPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *ServiceInputOutputSpecification {
		if v == nil {
			return nil
		}
		return v.Input
	}).(ServiceInputOutputSpecificationPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Keys() WebServiceKeysPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *WebServiceKeys {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(WebServiceKeysPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) MachineLearningWorkspace() MachineLearningWorkspacePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *MachineLearningWorkspace {
		if v == nil {
			return nil
		}
		return v.MachineLearningWorkspace
	}).(MachineLearningWorkspacePtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Output() ServiceInputOutputSpecificationPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *ServiceInputOutputSpecification {
		if v == nil {
			return nil
		}
		return v.Output
	}).(ServiceInputOutputSpecificationPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Package() GraphPackagePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *GraphPackage {
		if v == nil {
			return nil
		}
		return v.Package
	}).(GraphPackagePtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *string {
		if v == nil {
			return nil
		}
		return &v.PackageType
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Parameters() WebServiceParameterMapOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) map[string]WebServiceParameter {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(WebServiceParameterMapOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) PayloadsInBlobStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *bool {
		if v == nil {
			return nil
		}
		return v.PayloadsInBlobStorage
	}).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) PayloadsLocation() BlobLocationPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *BlobLocation {
		if v == nil {
			return nil
		}
		return v.PayloadsLocation
	}).(BlobLocationPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *bool {
		if v == nil {
			return nil
		}
		return v.ReadOnly
	}).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) RealtimeConfiguration() RealtimeConfigurationPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *RealtimeConfiguration {
		if v == nil {
			return nil
		}
		return v.RealtimeConfiguration
	}).(RealtimeConfigurationPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) StorageAccount() StorageAccountPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *StorageAccount {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(StorageAccountPtrOutput)
}

func (o WebServicePropertiesForGraphPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraph) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type WebServicePropertiesForGraphResponse struct {
	Assets                   map[string]AssetItemResponse             `pulumi:"assets"`
	CommitmentPlan           *CommitmentPlanResponse                  `pulumi:"commitmentPlan"`
	CreatedOn                string                                   `pulumi:"createdOn"`
	Description              *string                                  `pulumi:"description"`
	Diagnostics              *DiagnosticsConfigurationResponse        `pulumi:"diagnostics"`
	ExampleRequest           *ExampleRequestResponse                  `pulumi:"exampleRequest"`
	ExposeSampleData         *bool                                    `pulumi:"exposeSampleData"`
	Input                    *ServiceInputOutputSpecificationResponse `pulumi:"input"`
	Keys                     *WebServiceKeysResponse                  `pulumi:"keys"`
	MachineLearningWorkspace *MachineLearningWorkspaceResponse        `pulumi:"machineLearningWorkspace"`
	ModifiedOn               string                                   `pulumi:"modifiedOn"`
	Output                   *ServiceInputOutputSpecificationResponse `pulumi:"output"`
	Package                  *GraphPackageResponse                    `pulumi:"package"`
	PackageType              string                                   `pulumi:"packageType"`
	Parameters               map[string]WebServiceParameterResponse   `pulumi:"parameters"`
	PayloadsInBlobStorage    *bool                                    `pulumi:"payloadsInBlobStorage"`
	PayloadsLocation         *BlobLocationResponse                    `pulumi:"payloadsLocation"`
	ProvisioningState        string                                   `pulumi:"provisioningState"`
	ReadOnly                 *bool                                    `pulumi:"readOnly"`
	RealtimeConfiguration    *RealtimeConfigurationResponse           `pulumi:"realtimeConfiguration"`
	StorageAccount           *StorageAccountResponse                  `pulumi:"storageAccount"`
	SwaggerLocation          string                                   `pulumi:"swaggerLocation"`
	Title                    *string                                  `pulumi:"title"`
}





type WebServicePropertiesForGraphResponseInput interface {
	pulumi.Input

	ToWebServicePropertiesForGraphResponseOutput() WebServicePropertiesForGraphResponseOutput
	ToWebServicePropertiesForGraphResponseOutputWithContext(context.Context) WebServicePropertiesForGraphResponseOutput
}

type WebServicePropertiesForGraphResponseArgs struct {
	Assets                   AssetItemResponseMapInput                       `pulumi:"assets"`
	CommitmentPlan           CommitmentPlanResponsePtrInput                  `pulumi:"commitmentPlan"`
	CreatedOn                pulumi.StringInput                              `pulumi:"createdOn"`
	Description              pulumi.StringPtrInput                           `pulumi:"description"`
	Diagnostics              DiagnosticsConfigurationResponsePtrInput        `pulumi:"diagnostics"`
	ExampleRequest           ExampleRequestResponsePtrInput                  `pulumi:"exampleRequest"`
	ExposeSampleData         pulumi.BoolPtrInput                             `pulumi:"exposeSampleData"`
	Input                    ServiceInputOutputSpecificationResponsePtrInput `pulumi:"input"`
	Keys                     WebServiceKeysResponsePtrInput                  `pulumi:"keys"`
	MachineLearningWorkspace MachineLearningWorkspaceResponsePtrInput        `pulumi:"machineLearningWorkspace"`
	ModifiedOn               pulumi.StringInput                              `pulumi:"modifiedOn"`
	Output                   ServiceInputOutputSpecificationResponsePtrInput `pulumi:"output"`
	Package                  GraphPackageResponsePtrInput                    `pulumi:"package"`
	PackageType              pulumi.StringInput                              `pulumi:"packageType"`
	Parameters               WebServiceParameterResponseMapInput             `pulumi:"parameters"`
	PayloadsInBlobStorage    pulumi.BoolPtrInput                             `pulumi:"payloadsInBlobStorage"`
	PayloadsLocation         BlobLocationResponsePtrInput                    `pulumi:"payloadsLocation"`
	ProvisioningState        pulumi.StringInput                              `pulumi:"provisioningState"`
	ReadOnly                 pulumi.BoolPtrInput                             `pulumi:"readOnly"`
	RealtimeConfiguration    RealtimeConfigurationResponsePtrInput           `pulumi:"realtimeConfiguration"`
	StorageAccount           StorageAccountResponsePtrInput                  `pulumi:"storageAccount"`
	SwaggerLocation          pulumi.StringInput                              `pulumi:"swaggerLocation"`
	Title                    pulumi.StringPtrInput                           `pulumi:"title"`
}

func (WebServicePropertiesForGraphResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServicePropertiesForGraphResponse)(nil)).Elem()
}

func (i WebServicePropertiesForGraphResponseArgs) ToWebServicePropertiesForGraphResponseOutput() WebServicePropertiesForGraphResponseOutput {
	return i.ToWebServicePropertiesForGraphResponseOutputWithContext(context.Background())
}

func (i WebServicePropertiesForGraphResponseArgs) ToWebServicePropertiesForGraphResponseOutputWithContext(ctx context.Context) WebServicePropertiesForGraphResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServicePropertiesForGraphResponseOutput)
}

func (i WebServicePropertiesForGraphResponseArgs) ToWebServicePropertiesForGraphResponsePtrOutput() WebServicePropertiesForGraphResponsePtrOutput {
	return i.ToWebServicePropertiesForGraphResponsePtrOutputWithContext(context.Background())
}

func (i WebServicePropertiesForGraphResponseArgs) ToWebServicePropertiesForGraphResponsePtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServicePropertiesForGraphResponseOutput).ToWebServicePropertiesForGraphResponsePtrOutputWithContext(ctx)
}









type WebServicePropertiesForGraphResponsePtrInput interface {
	pulumi.Input

	ToWebServicePropertiesForGraphResponsePtrOutput() WebServicePropertiesForGraphResponsePtrOutput
	ToWebServicePropertiesForGraphResponsePtrOutputWithContext(context.Context) WebServicePropertiesForGraphResponsePtrOutput
}

type webServicePropertiesForGraphResponsePtrType WebServicePropertiesForGraphResponseArgs

func WebServicePropertiesForGraphResponsePtr(v *WebServicePropertiesForGraphResponseArgs) WebServicePropertiesForGraphResponsePtrInput {
	return (*webServicePropertiesForGraphResponsePtrType)(v)
}

func (*webServicePropertiesForGraphResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServicePropertiesForGraphResponse)(nil)).Elem()
}

func (i *webServicePropertiesForGraphResponsePtrType) ToWebServicePropertiesForGraphResponsePtrOutput() WebServicePropertiesForGraphResponsePtrOutput {
	return i.ToWebServicePropertiesForGraphResponsePtrOutputWithContext(context.Background())
}

func (i *webServicePropertiesForGraphResponsePtrType) ToWebServicePropertiesForGraphResponsePtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServicePropertiesForGraphResponsePtrOutput)
}

type WebServicePropertiesForGraphResponseOutput struct{ *pulumi.OutputState }

func (WebServicePropertiesForGraphResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebServicePropertiesForGraphResponse)(nil)).Elem()
}

func (o WebServicePropertiesForGraphResponseOutput) ToWebServicePropertiesForGraphResponseOutput() WebServicePropertiesForGraphResponseOutput {
	return o
}

func (o WebServicePropertiesForGraphResponseOutput) ToWebServicePropertiesForGraphResponseOutputWithContext(ctx context.Context) WebServicePropertiesForGraphResponseOutput {
	return o
}

func (o WebServicePropertiesForGraphResponseOutput) ToWebServicePropertiesForGraphResponsePtrOutput() WebServicePropertiesForGraphResponsePtrOutput {
	return o.ToWebServicePropertiesForGraphResponsePtrOutputWithContext(context.Background())
}

func (o WebServicePropertiesForGraphResponseOutput) ToWebServicePropertiesForGraphResponsePtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebServicePropertiesForGraphResponse) *WebServicePropertiesForGraphResponse {
		return &v
	}).(WebServicePropertiesForGraphResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Assets() AssetItemResponseMapOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) map[string]AssetItemResponse { return v.Assets }).(AssetItemResponseMapOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) CommitmentPlan() CommitmentPlanResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *CommitmentPlanResponse { return v.CommitmentPlan }).(CommitmentPlanResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Diagnostics() DiagnosticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *DiagnosticsConfigurationResponse { return v.Diagnostics }).(DiagnosticsConfigurationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) ExampleRequest() ExampleRequestResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *ExampleRequestResponse { return v.ExampleRequest }).(ExampleRequestResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) ExposeSampleData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *bool { return v.ExposeSampleData }).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Input() ServiceInputOutputSpecificationResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *ServiceInputOutputSpecificationResponse { return v.Input }).(ServiceInputOutputSpecificationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Keys() WebServiceKeysResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *WebServiceKeysResponse { return v.Keys }).(WebServiceKeysResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) MachineLearningWorkspace() MachineLearningWorkspaceResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *MachineLearningWorkspaceResponse {
		return v.MachineLearningWorkspace
	}).(MachineLearningWorkspaceResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Output() ServiceInputOutputSpecificationResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *ServiceInputOutputSpecificationResponse { return v.Output }).(ServiceInputOutputSpecificationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Package() GraphPackageResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *GraphPackageResponse { return v.Package }).(GraphPackageResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Parameters() WebServiceParameterResponseMapOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) map[string]WebServiceParameterResponse {
		return v.Parameters
	}).(WebServiceParameterResponseMapOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) PayloadsInBlobStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *bool { return v.PayloadsInBlobStorage }).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) PayloadsLocation() BlobLocationResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *BlobLocationResponse { return v.PayloadsLocation }).(BlobLocationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) RealtimeConfiguration() RealtimeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *RealtimeConfigurationResponse {
		return v.RealtimeConfiguration
	}).(RealtimeConfigurationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) StorageAccount() StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *StorageAccountResponse { return v.StorageAccount }).(StorageAccountResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) SwaggerLocation() pulumi.StringOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) string { return v.SwaggerLocation }).(pulumi.StringOutput)
}

func (o WebServicePropertiesForGraphResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebServicePropertiesForGraphResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type WebServicePropertiesForGraphResponsePtrOutput struct{ *pulumi.OutputState }

func (WebServicePropertiesForGraphResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebServicePropertiesForGraphResponse)(nil)).Elem()
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ToWebServicePropertiesForGraphResponsePtrOutput() WebServicePropertiesForGraphResponsePtrOutput {
	return o
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ToWebServicePropertiesForGraphResponsePtrOutputWithContext(ctx context.Context) WebServicePropertiesForGraphResponsePtrOutput {
	return o
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Elem() WebServicePropertiesForGraphResponseOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) WebServicePropertiesForGraphResponse {
		if v != nil {
			return *v
		}
		var ret WebServicePropertiesForGraphResponse
		return ret
	}).(WebServicePropertiesForGraphResponseOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Assets() AssetItemResponseMapOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) map[string]AssetItemResponse {
		if v == nil {
			return nil
		}
		return v.Assets
	}).(AssetItemResponseMapOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) CommitmentPlan() CommitmentPlanResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *CommitmentPlanResponse {
		if v == nil {
			return nil
		}
		return v.CommitmentPlan
	}).(CommitmentPlanResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) CreatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedOn
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Diagnostics() DiagnosticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *DiagnosticsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Diagnostics
	}).(DiagnosticsConfigurationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ExampleRequest() ExampleRequestResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *ExampleRequestResponse {
		if v == nil {
			return nil
		}
		return v.ExampleRequest
	}).(ExampleRequestResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ExposeSampleData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ExposeSampleData
	}).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Input() ServiceInputOutputSpecificationResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *ServiceInputOutputSpecificationResponse {
		if v == nil {
			return nil
		}
		return v.Input
	}).(ServiceInputOutputSpecificationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Keys() WebServiceKeysResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *WebServiceKeysResponse {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(WebServiceKeysResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) MachineLearningWorkspace() MachineLearningWorkspaceResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *MachineLearningWorkspaceResponse {
		if v == nil {
			return nil
		}
		return v.MachineLearningWorkspace
	}).(MachineLearningWorkspaceResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ModifiedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ModifiedOn
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Output() ServiceInputOutputSpecificationResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *ServiceInputOutputSpecificationResponse {
		if v == nil {
			return nil
		}
		return v.Output
	}).(ServiceInputOutputSpecificationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Package() GraphPackageResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *GraphPackageResponse {
		if v == nil {
			return nil
		}
		return v.Package
	}).(GraphPackageResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PackageType
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Parameters() WebServiceParameterResponseMapOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) map[string]WebServiceParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(WebServiceParameterResponseMapOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) PayloadsInBlobStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PayloadsInBlobStorage
	}).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) PayloadsLocation() BlobLocationResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *BlobLocationResponse {
		if v == nil {
			return nil
		}
		return v.PayloadsLocation
	}).(BlobLocationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ReadOnly
	}).(pulumi.BoolPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) RealtimeConfiguration() RealtimeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *RealtimeConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.RealtimeConfiguration
	}).(RealtimeConfigurationResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) StorageAccount() StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *StorageAccountResponse {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(StorageAccountResponsePtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) SwaggerLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SwaggerLocation
	}).(pulumi.StringPtrOutput)
}

func (o WebServicePropertiesForGraphResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebServicePropertiesForGraphResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetItemOutput{})
	pulumi.RegisterOutputType(AssetItemMapOutput{})
	pulumi.RegisterOutputType(AssetItemResponseOutput{})
	pulumi.RegisterOutputType(AssetItemResponseMapOutput{})
	pulumi.RegisterOutputType(BlobLocationOutput{})
	pulumi.RegisterOutputType(BlobLocationPtrOutput{})
	pulumi.RegisterOutputType(BlobLocationResponseOutput{})
	pulumi.RegisterOutputType(BlobLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(ColumnSpecificationOutput{})
	pulumi.RegisterOutputType(ColumnSpecificationMapOutput{})
	pulumi.RegisterOutputType(ColumnSpecificationResponseOutput{})
	pulumi.RegisterOutputType(ColumnSpecificationResponseMapOutput{})
	pulumi.RegisterOutputType(CommitmentPlanTypeOutput{})
	pulumi.RegisterOutputType(CommitmentPlanTypePtrOutput{})
	pulumi.RegisterOutputType(CommitmentPlanPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CommitmentPlanPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CommitmentPlanResponseOutput{})
	pulumi.RegisterOutputType(CommitmentPlanResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsConfigurationOutput{})
	pulumi.RegisterOutputType(DiagnosticsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ExampleRequestOutput{})
	pulumi.RegisterOutputType(ExampleRequestPtrOutput{})
	pulumi.RegisterOutputType(ExampleRequestResponseOutput{})
	pulumi.RegisterOutputType(ExampleRequestResponsePtrOutput{})
	pulumi.RegisterOutputType(GraphEdgeOutput{})
	pulumi.RegisterOutputType(GraphEdgeArrayOutput{})
	pulumi.RegisterOutputType(GraphEdgeResponseOutput{})
	pulumi.RegisterOutputType(GraphEdgeResponseArrayOutput{})
	pulumi.RegisterOutputType(GraphNodeOutput{})
	pulumi.RegisterOutputType(GraphNodeMapOutput{})
	pulumi.RegisterOutputType(GraphNodeResponseOutput{})
	pulumi.RegisterOutputType(GraphNodeResponseMapOutput{})
	pulumi.RegisterOutputType(GraphPackageOutput{})
	pulumi.RegisterOutputType(GraphPackagePtrOutput{})
	pulumi.RegisterOutputType(GraphPackageResponseOutput{})
	pulumi.RegisterOutputType(GraphPackageResponsePtrOutput{})
	pulumi.RegisterOutputType(GraphParameterOutput{})
	pulumi.RegisterOutputType(GraphParameterMapOutput{})
	pulumi.RegisterOutputType(GraphParameterLinkOutput{})
	pulumi.RegisterOutputType(GraphParameterLinkArrayOutput{})
	pulumi.RegisterOutputType(GraphParameterLinkResponseOutput{})
	pulumi.RegisterOutputType(GraphParameterLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(GraphParameterResponseOutput{})
	pulumi.RegisterOutputType(GraphParameterResponseMapOutput{})
	pulumi.RegisterOutputType(InputPortOutput{})
	pulumi.RegisterOutputType(InputPortMapOutput{})
	pulumi.RegisterOutputType(InputPortResponseOutput{})
	pulumi.RegisterOutputType(InputPortResponseMapOutput{})
	pulumi.RegisterOutputType(MachineLearningWorkspaceOutput{})
	pulumi.RegisterOutputType(MachineLearningWorkspacePtrOutput{})
	pulumi.RegisterOutputType(MachineLearningWorkspaceResponseOutput{})
	pulumi.RegisterOutputType(MachineLearningWorkspaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ModeValueInfoOutput{})
	pulumi.RegisterOutputType(ModeValueInfoMapOutput{})
	pulumi.RegisterOutputType(ModeValueInfoResponseOutput{})
	pulumi.RegisterOutputType(ModeValueInfoResponseMapOutput{})
	pulumi.RegisterOutputType(ModuleAssetParameterOutput{})
	pulumi.RegisterOutputType(ModuleAssetParameterArrayOutput{})
	pulumi.RegisterOutputType(ModuleAssetParameterResponseOutput{})
	pulumi.RegisterOutputType(ModuleAssetParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(OutputPortOutput{})
	pulumi.RegisterOutputType(OutputPortMapOutput{})
	pulumi.RegisterOutputType(OutputPortResponseOutput{})
	pulumi.RegisterOutputType(OutputPortResponseMapOutput{})
	pulumi.RegisterOutputType(PlanQuantityResponseOutput{})
	pulumi.RegisterOutputType(PlanQuantityResponseMapOutput{})
	pulumi.RegisterOutputType(RealtimeConfigurationOutput{})
	pulumi.RegisterOutputType(RealtimeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RealtimeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(RealtimeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceInputOutputSpecificationOutput{})
	pulumi.RegisterOutputType(ServiceInputOutputSpecificationPtrOutput{})
	pulumi.RegisterOutputType(ServiceInputOutputSpecificationResponseOutput{})
	pulumi.RegisterOutputType(ServiceInputOutputSpecificationResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(TableSpecificationOutput{})
	pulumi.RegisterOutputType(TableSpecificationMapOutput{})
	pulumi.RegisterOutputType(TableSpecificationResponseOutput{})
	pulumi.RegisterOutputType(TableSpecificationResponseMapOutput{})
	pulumi.RegisterOutputType(WebServiceKeysOutput{})
	pulumi.RegisterOutputType(WebServiceKeysPtrOutput{})
	pulumi.RegisterOutputType(WebServiceKeysResponseOutput{})
	pulumi.RegisterOutputType(WebServiceKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(WebServiceParameterOutput{})
	pulumi.RegisterOutputType(WebServiceParameterMapOutput{})
	pulumi.RegisterOutputType(WebServiceParameterResponseOutput{})
	pulumi.RegisterOutputType(WebServiceParameterResponseMapOutput{})
	pulumi.RegisterOutputType(WebServicePropertiesForGraphOutput{})
	pulumi.RegisterOutputType(WebServicePropertiesForGraphPtrOutput{})
	pulumi.RegisterOutputType(WebServicePropertiesForGraphResponseOutput{})
	pulumi.RegisterOutputType(WebServicePropertiesForGraphResponsePtrOutput{})
}
