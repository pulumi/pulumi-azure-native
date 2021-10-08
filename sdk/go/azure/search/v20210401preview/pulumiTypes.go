


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataPlaneAadOrApiKeyAuthOption struct {
	AadAuthFailureMode *AadAuthFailureMode `pulumi:"aadAuthFailureMode"`
}





type DataPlaneAadOrApiKeyAuthOptionInput interface {
	pulumi.Input

	ToDataPlaneAadOrApiKeyAuthOptionOutput() DataPlaneAadOrApiKeyAuthOptionOutput
	ToDataPlaneAadOrApiKeyAuthOptionOutputWithContext(context.Context) DataPlaneAadOrApiKeyAuthOptionOutput
}

type DataPlaneAadOrApiKeyAuthOptionArgs struct {
	AadAuthFailureMode AadAuthFailureModePtrInput `pulumi:"aadAuthFailureMode"`
}

func (DataPlaneAadOrApiKeyAuthOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOption)(nil)).Elem()
}

func (i DataPlaneAadOrApiKeyAuthOptionArgs) ToDataPlaneAadOrApiKeyAuthOptionOutput() DataPlaneAadOrApiKeyAuthOptionOutput {
	return i.ToDataPlaneAadOrApiKeyAuthOptionOutputWithContext(context.Background())
}

func (i DataPlaneAadOrApiKeyAuthOptionArgs) ToDataPlaneAadOrApiKeyAuthOptionOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAadOrApiKeyAuthOptionOutput)
}

func (i DataPlaneAadOrApiKeyAuthOptionArgs) ToDataPlaneAadOrApiKeyAuthOptionPtrOutput() DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return i.ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(context.Background())
}

func (i DataPlaneAadOrApiKeyAuthOptionArgs) ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAadOrApiKeyAuthOptionOutput).ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(ctx)
}









type DataPlaneAadOrApiKeyAuthOptionPtrInput interface {
	pulumi.Input

	ToDataPlaneAadOrApiKeyAuthOptionPtrOutput() DataPlaneAadOrApiKeyAuthOptionPtrOutput
	ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(context.Context) DataPlaneAadOrApiKeyAuthOptionPtrOutput
}

type dataPlaneAadOrApiKeyAuthOptionPtrType DataPlaneAadOrApiKeyAuthOptionArgs

func DataPlaneAadOrApiKeyAuthOptionPtr(v *DataPlaneAadOrApiKeyAuthOptionArgs) DataPlaneAadOrApiKeyAuthOptionPtrInput {
	return (*dataPlaneAadOrApiKeyAuthOptionPtrType)(v)
}

func (*dataPlaneAadOrApiKeyAuthOptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAadOrApiKeyAuthOption)(nil)).Elem()
}

func (i *dataPlaneAadOrApiKeyAuthOptionPtrType) ToDataPlaneAadOrApiKeyAuthOptionPtrOutput() DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return i.ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(context.Background())
}

func (i *dataPlaneAadOrApiKeyAuthOptionPtrType) ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAadOrApiKeyAuthOptionPtrOutput)
}

type DataPlaneAadOrApiKeyAuthOptionOutput struct{ *pulumi.OutputState }

func (DataPlaneAadOrApiKeyAuthOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOption)(nil)).Elem()
}

func (o DataPlaneAadOrApiKeyAuthOptionOutput) ToDataPlaneAadOrApiKeyAuthOptionOutput() DataPlaneAadOrApiKeyAuthOptionOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionOutput) ToDataPlaneAadOrApiKeyAuthOptionOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionOutput) ToDataPlaneAadOrApiKeyAuthOptionPtrOutput() DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return o.ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(context.Background())
}

func (o DataPlaneAadOrApiKeyAuthOptionOutput) ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataPlaneAadOrApiKeyAuthOption) *DataPlaneAadOrApiKeyAuthOption {
		return &v
	}).(DataPlaneAadOrApiKeyAuthOptionPtrOutput)
}

func (o DataPlaneAadOrApiKeyAuthOptionOutput) AadAuthFailureMode() AadAuthFailureModePtrOutput {
	return o.ApplyT(func(v DataPlaneAadOrApiKeyAuthOption) *AadAuthFailureMode { return v.AadAuthFailureMode }).(AadAuthFailureModePtrOutput)
}

type DataPlaneAadOrApiKeyAuthOptionPtrOutput struct{ *pulumi.OutputState }

func (DataPlaneAadOrApiKeyAuthOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAadOrApiKeyAuthOption)(nil)).Elem()
}

func (o DataPlaneAadOrApiKeyAuthOptionPtrOutput) ToDataPlaneAadOrApiKeyAuthOptionPtrOutput() DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionPtrOutput) ToDataPlaneAadOrApiKeyAuthOptionPtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionPtrOutput) Elem() DataPlaneAadOrApiKeyAuthOptionOutput {
	return o.ApplyT(func(v *DataPlaneAadOrApiKeyAuthOption) DataPlaneAadOrApiKeyAuthOption {
		if v != nil {
			return *v
		}
		var ret DataPlaneAadOrApiKeyAuthOption
		return ret
	}).(DataPlaneAadOrApiKeyAuthOptionOutput)
}

func (o DataPlaneAadOrApiKeyAuthOptionPtrOutput) AadAuthFailureMode() AadAuthFailureModePtrOutput {
	return o.ApplyT(func(v *DataPlaneAadOrApiKeyAuthOption) *AadAuthFailureMode {
		if v == nil {
			return nil
		}
		return v.AadAuthFailureMode
	}).(AadAuthFailureModePtrOutput)
}

type DataPlaneAadOrApiKeyAuthOptionResponse struct {
	AadAuthFailureMode *string `pulumi:"aadAuthFailureMode"`
}





type DataPlaneAadOrApiKeyAuthOptionResponseInput interface {
	pulumi.Input

	ToDataPlaneAadOrApiKeyAuthOptionResponseOutput() DataPlaneAadOrApiKeyAuthOptionResponseOutput
	ToDataPlaneAadOrApiKeyAuthOptionResponseOutputWithContext(context.Context) DataPlaneAadOrApiKeyAuthOptionResponseOutput
}

type DataPlaneAadOrApiKeyAuthOptionResponseArgs struct {
	AadAuthFailureMode pulumi.StringPtrInput `pulumi:"aadAuthFailureMode"`
}

func (DataPlaneAadOrApiKeyAuthOptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOptionResponse)(nil)).Elem()
}

func (i DataPlaneAadOrApiKeyAuthOptionResponseArgs) ToDataPlaneAadOrApiKeyAuthOptionResponseOutput() DataPlaneAadOrApiKeyAuthOptionResponseOutput {
	return i.ToDataPlaneAadOrApiKeyAuthOptionResponseOutputWithContext(context.Background())
}

func (i DataPlaneAadOrApiKeyAuthOptionResponseArgs) ToDataPlaneAadOrApiKeyAuthOptionResponseOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAadOrApiKeyAuthOptionResponseOutput)
}

func (i DataPlaneAadOrApiKeyAuthOptionResponseArgs) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutput() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return i.ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(context.Background())
}

func (i DataPlaneAadOrApiKeyAuthOptionResponseArgs) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAadOrApiKeyAuthOptionResponseOutput).ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(ctx)
}









type DataPlaneAadOrApiKeyAuthOptionResponsePtrInput interface {
	pulumi.Input

	ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutput() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput
	ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(context.Context) DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput
}

type dataPlaneAadOrApiKeyAuthOptionResponsePtrType DataPlaneAadOrApiKeyAuthOptionResponseArgs

func DataPlaneAadOrApiKeyAuthOptionResponsePtr(v *DataPlaneAadOrApiKeyAuthOptionResponseArgs) DataPlaneAadOrApiKeyAuthOptionResponsePtrInput {
	return (*dataPlaneAadOrApiKeyAuthOptionResponsePtrType)(v)
}

func (*dataPlaneAadOrApiKeyAuthOptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAadOrApiKeyAuthOptionResponse)(nil)).Elem()
}

func (i *dataPlaneAadOrApiKeyAuthOptionResponsePtrType) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutput() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return i.ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(context.Background())
}

func (i *dataPlaneAadOrApiKeyAuthOptionResponsePtrType) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput)
}

type DataPlaneAadOrApiKeyAuthOptionResponseOutput struct{ *pulumi.OutputState }

func (DataPlaneAadOrApiKeyAuthOptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOptionResponse)(nil)).Elem()
}

func (o DataPlaneAadOrApiKeyAuthOptionResponseOutput) ToDataPlaneAadOrApiKeyAuthOptionResponseOutput() DataPlaneAadOrApiKeyAuthOptionResponseOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionResponseOutput) ToDataPlaneAadOrApiKeyAuthOptionResponseOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionResponseOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionResponseOutput) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutput() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return o.ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(context.Background())
}

func (o DataPlaneAadOrApiKeyAuthOptionResponseOutput) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataPlaneAadOrApiKeyAuthOptionResponse) *DataPlaneAadOrApiKeyAuthOptionResponse {
		return &v
	}).(DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput)
}

func (o DataPlaneAadOrApiKeyAuthOptionResponseOutput) AadAuthFailureMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataPlaneAadOrApiKeyAuthOptionResponse) *string { return v.AadAuthFailureMode }).(pulumi.StringPtrOutput)
}

type DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput struct{ *pulumi.OutputState }

func (DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAadOrApiKeyAuthOptionResponse)(nil)).Elem()
}

func (o DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutput() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput) ToDataPlaneAadOrApiKeyAuthOptionResponsePtrOutputWithContext(ctx context.Context) DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return o
}

func (o DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput) Elem() DataPlaneAadOrApiKeyAuthOptionResponseOutput {
	return o.ApplyT(func(v *DataPlaneAadOrApiKeyAuthOptionResponse) DataPlaneAadOrApiKeyAuthOptionResponse {
		if v != nil {
			return *v
		}
		var ret DataPlaneAadOrApiKeyAuthOptionResponse
		return ret
	}).(DataPlaneAadOrApiKeyAuthOptionResponseOutput)
}

func (o DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput) AadAuthFailureMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataPlaneAadOrApiKeyAuthOptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadAuthFailureMode
	}).(pulumi.StringPtrOutput)
}

type DataPlaneAuthOptions struct {
	AadOrApiKey *DataPlaneAadOrApiKeyAuthOption `pulumi:"aadOrApiKey"`
	ApiKeyOnly  interface{}                     `pulumi:"apiKeyOnly"`
}





type DataPlaneAuthOptionsInput interface {
	pulumi.Input

	ToDataPlaneAuthOptionsOutput() DataPlaneAuthOptionsOutput
	ToDataPlaneAuthOptionsOutputWithContext(context.Context) DataPlaneAuthOptionsOutput
}

type DataPlaneAuthOptionsArgs struct {
	AadOrApiKey DataPlaneAadOrApiKeyAuthOptionPtrInput `pulumi:"aadOrApiKey"`
	ApiKeyOnly  pulumi.Input                           `pulumi:"apiKeyOnly"`
}

func (DataPlaneAuthOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAuthOptions)(nil)).Elem()
}

func (i DataPlaneAuthOptionsArgs) ToDataPlaneAuthOptionsOutput() DataPlaneAuthOptionsOutput {
	return i.ToDataPlaneAuthOptionsOutputWithContext(context.Background())
}

func (i DataPlaneAuthOptionsArgs) ToDataPlaneAuthOptionsOutputWithContext(ctx context.Context) DataPlaneAuthOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAuthOptionsOutput)
}

func (i DataPlaneAuthOptionsArgs) ToDataPlaneAuthOptionsPtrOutput() DataPlaneAuthOptionsPtrOutput {
	return i.ToDataPlaneAuthOptionsPtrOutputWithContext(context.Background())
}

func (i DataPlaneAuthOptionsArgs) ToDataPlaneAuthOptionsPtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAuthOptionsOutput).ToDataPlaneAuthOptionsPtrOutputWithContext(ctx)
}









type DataPlaneAuthOptionsPtrInput interface {
	pulumi.Input

	ToDataPlaneAuthOptionsPtrOutput() DataPlaneAuthOptionsPtrOutput
	ToDataPlaneAuthOptionsPtrOutputWithContext(context.Context) DataPlaneAuthOptionsPtrOutput
}

type dataPlaneAuthOptionsPtrType DataPlaneAuthOptionsArgs

func DataPlaneAuthOptionsPtr(v *DataPlaneAuthOptionsArgs) DataPlaneAuthOptionsPtrInput {
	return (*dataPlaneAuthOptionsPtrType)(v)
}

func (*dataPlaneAuthOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAuthOptions)(nil)).Elem()
}

func (i *dataPlaneAuthOptionsPtrType) ToDataPlaneAuthOptionsPtrOutput() DataPlaneAuthOptionsPtrOutput {
	return i.ToDataPlaneAuthOptionsPtrOutputWithContext(context.Background())
}

func (i *dataPlaneAuthOptionsPtrType) ToDataPlaneAuthOptionsPtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAuthOptionsPtrOutput)
}

type DataPlaneAuthOptionsOutput struct{ *pulumi.OutputState }

func (DataPlaneAuthOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAuthOptions)(nil)).Elem()
}

func (o DataPlaneAuthOptionsOutput) ToDataPlaneAuthOptionsOutput() DataPlaneAuthOptionsOutput {
	return o
}

func (o DataPlaneAuthOptionsOutput) ToDataPlaneAuthOptionsOutputWithContext(ctx context.Context) DataPlaneAuthOptionsOutput {
	return o
}

func (o DataPlaneAuthOptionsOutput) ToDataPlaneAuthOptionsPtrOutput() DataPlaneAuthOptionsPtrOutput {
	return o.ToDataPlaneAuthOptionsPtrOutputWithContext(context.Background())
}

func (o DataPlaneAuthOptionsOutput) ToDataPlaneAuthOptionsPtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataPlaneAuthOptions) *DataPlaneAuthOptions {
		return &v
	}).(DataPlaneAuthOptionsPtrOutput)
}

func (o DataPlaneAuthOptionsOutput) AadOrApiKey() DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return o.ApplyT(func(v DataPlaneAuthOptions) *DataPlaneAadOrApiKeyAuthOption { return v.AadOrApiKey }).(DataPlaneAadOrApiKeyAuthOptionPtrOutput)
}

func (o DataPlaneAuthOptionsOutput) ApiKeyOnly() pulumi.AnyOutput {
	return o.ApplyT(func(v DataPlaneAuthOptions) interface{} { return v.ApiKeyOnly }).(pulumi.AnyOutput)
}

type DataPlaneAuthOptionsPtrOutput struct{ *pulumi.OutputState }

func (DataPlaneAuthOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAuthOptions)(nil)).Elem()
}

func (o DataPlaneAuthOptionsPtrOutput) ToDataPlaneAuthOptionsPtrOutput() DataPlaneAuthOptionsPtrOutput {
	return o
}

func (o DataPlaneAuthOptionsPtrOutput) ToDataPlaneAuthOptionsPtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsPtrOutput {
	return o
}

func (o DataPlaneAuthOptionsPtrOutput) Elem() DataPlaneAuthOptionsOutput {
	return o.ApplyT(func(v *DataPlaneAuthOptions) DataPlaneAuthOptions {
		if v != nil {
			return *v
		}
		var ret DataPlaneAuthOptions
		return ret
	}).(DataPlaneAuthOptionsOutput)
}

func (o DataPlaneAuthOptionsPtrOutput) AadOrApiKey() DataPlaneAadOrApiKeyAuthOptionPtrOutput {
	return o.ApplyT(func(v *DataPlaneAuthOptions) *DataPlaneAadOrApiKeyAuthOption {
		if v == nil {
			return nil
		}
		return v.AadOrApiKey
	}).(DataPlaneAadOrApiKeyAuthOptionPtrOutput)
}

func (o DataPlaneAuthOptionsPtrOutput) ApiKeyOnly() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataPlaneAuthOptions) interface{} {
		if v == nil {
			return nil
		}
		return v.ApiKeyOnly
	}).(pulumi.AnyOutput)
}

type DataPlaneAuthOptionsResponse struct {
	AadOrApiKey *DataPlaneAadOrApiKeyAuthOptionResponse `pulumi:"aadOrApiKey"`
	ApiKeyOnly  interface{}                             `pulumi:"apiKeyOnly"`
}





type DataPlaneAuthOptionsResponseInput interface {
	pulumi.Input

	ToDataPlaneAuthOptionsResponseOutput() DataPlaneAuthOptionsResponseOutput
	ToDataPlaneAuthOptionsResponseOutputWithContext(context.Context) DataPlaneAuthOptionsResponseOutput
}

type DataPlaneAuthOptionsResponseArgs struct {
	AadOrApiKey DataPlaneAadOrApiKeyAuthOptionResponsePtrInput `pulumi:"aadOrApiKey"`
	ApiKeyOnly  pulumi.Input                                   `pulumi:"apiKeyOnly"`
}

func (DataPlaneAuthOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAuthOptionsResponse)(nil)).Elem()
}

func (i DataPlaneAuthOptionsResponseArgs) ToDataPlaneAuthOptionsResponseOutput() DataPlaneAuthOptionsResponseOutput {
	return i.ToDataPlaneAuthOptionsResponseOutputWithContext(context.Background())
}

func (i DataPlaneAuthOptionsResponseArgs) ToDataPlaneAuthOptionsResponseOutputWithContext(ctx context.Context) DataPlaneAuthOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAuthOptionsResponseOutput)
}

func (i DataPlaneAuthOptionsResponseArgs) ToDataPlaneAuthOptionsResponsePtrOutput() DataPlaneAuthOptionsResponsePtrOutput {
	return i.ToDataPlaneAuthOptionsResponsePtrOutputWithContext(context.Background())
}

func (i DataPlaneAuthOptionsResponseArgs) ToDataPlaneAuthOptionsResponsePtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAuthOptionsResponseOutput).ToDataPlaneAuthOptionsResponsePtrOutputWithContext(ctx)
}









type DataPlaneAuthOptionsResponsePtrInput interface {
	pulumi.Input

	ToDataPlaneAuthOptionsResponsePtrOutput() DataPlaneAuthOptionsResponsePtrOutput
	ToDataPlaneAuthOptionsResponsePtrOutputWithContext(context.Context) DataPlaneAuthOptionsResponsePtrOutput
}

type dataPlaneAuthOptionsResponsePtrType DataPlaneAuthOptionsResponseArgs

func DataPlaneAuthOptionsResponsePtr(v *DataPlaneAuthOptionsResponseArgs) DataPlaneAuthOptionsResponsePtrInput {
	return (*dataPlaneAuthOptionsResponsePtrType)(v)
}

func (*dataPlaneAuthOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAuthOptionsResponse)(nil)).Elem()
}

func (i *dataPlaneAuthOptionsResponsePtrType) ToDataPlaneAuthOptionsResponsePtrOutput() DataPlaneAuthOptionsResponsePtrOutput {
	return i.ToDataPlaneAuthOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *dataPlaneAuthOptionsResponsePtrType) ToDataPlaneAuthOptionsResponsePtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataPlaneAuthOptionsResponsePtrOutput)
}

type DataPlaneAuthOptionsResponseOutput struct{ *pulumi.OutputState }

func (DataPlaneAuthOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPlaneAuthOptionsResponse)(nil)).Elem()
}

func (o DataPlaneAuthOptionsResponseOutput) ToDataPlaneAuthOptionsResponseOutput() DataPlaneAuthOptionsResponseOutput {
	return o
}

func (o DataPlaneAuthOptionsResponseOutput) ToDataPlaneAuthOptionsResponseOutputWithContext(ctx context.Context) DataPlaneAuthOptionsResponseOutput {
	return o
}

func (o DataPlaneAuthOptionsResponseOutput) ToDataPlaneAuthOptionsResponsePtrOutput() DataPlaneAuthOptionsResponsePtrOutput {
	return o.ToDataPlaneAuthOptionsResponsePtrOutputWithContext(context.Background())
}

func (o DataPlaneAuthOptionsResponseOutput) ToDataPlaneAuthOptionsResponsePtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataPlaneAuthOptionsResponse) *DataPlaneAuthOptionsResponse {
		return &v
	}).(DataPlaneAuthOptionsResponsePtrOutput)
}

func (o DataPlaneAuthOptionsResponseOutput) AadOrApiKey() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return o.ApplyT(func(v DataPlaneAuthOptionsResponse) *DataPlaneAadOrApiKeyAuthOptionResponse { return v.AadOrApiKey }).(DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput)
}

func (o DataPlaneAuthOptionsResponseOutput) ApiKeyOnly() pulumi.AnyOutput {
	return o.ApplyT(func(v DataPlaneAuthOptionsResponse) interface{} { return v.ApiKeyOnly }).(pulumi.AnyOutput)
}

type DataPlaneAuthOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (DataPlaneAuthOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPlaneAuthOptionsResponse)(nil)).Elem()
}

func (o DataPlaneAuthOptionsResponsePtrOutput) ToDataPlaneAuthOptionsResponsePtrOutput() DataPlaneAuthOptionsResponsePtrOutput {
	return o
}

func (o DataPlaneAuthOptionsResponsePtrOutput) ToDataPlaneAuthOptionsResponsePtrOutputWithContext(ctx context.Context) DataPlaneAuthOptionsResponsePtrOutput {
	return o
}

func (o DataPlaneAuthOptionsResponsePtrOutput) Elem() DataPlaneAuthOptionsResponseOutput {
	return o.ApplyT(func(v *DataPlaneAuthOptionsResponse) DataPlaneAuthOptionsResponse {
		if v != nil {
			return *v
		}
		var ret DataPlaneAuthOptionsResponse
		return ret
	}).(DataPlaneAuthOptionsResponseOutput)
}

func (o DataPlaneAuthOptionsResponsePtrOutput) AadOrApiKey() DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput {
	return o.ApplyT(func(v *DataPlaneAuthOptionsResponse) *DataPlaneAadOrApiKeyAuthOptionResponse {
		if v == nil {
			return nil
		}
		return v.AadOrApiKey
	}).(DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput)
}

func (o DataPlaneAuthOptionsResponsePtrOutput) ApiKeyOnly() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataPlaneAuthOptionsResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ApiKeyOnly
	}).(pulumi.AnyOutput)
}

type EncryptionWithCmk struct {
	Enforcement *SearchEncryptionWithCmk `pulumi:"enforcement"`
}





type EncryptionWithCmkInput interface {
	pulumi.Input

	ToEncryptionWithCmkOutput() EncryptionWithCmkOutput
	ToEncryptionWithCmkOutputWithContext(context.Context) EncryptionWithCmkOutput
}

type EncryptionWithCmkArgs struct {
	Enforcement SearchEncryptionWithCmkPtrInput `pulumi:"enforcement"`
}

func (EncryptionWithCmkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionWithCmk)(nil)).Elem()
}

func (i EncryptionWithCmkArgs) ToEncryptionWithCmkOutput() EncryptionWithCmkOutput {
	return i.ToEncryptionWithCmkOutputWithContext(context.Background())
}

func (i EncryptionWithCmkArgs) ToEncryptionWithCmkOutputWithContext(ctx context.Context) EncryptionWithCmkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionWithCmkOutput)
}

func (i EncryptionWithCmkArgs) ToEncryptionWithCmkPtrOutput() EncryptionWithCmkPtrOutput {
	return i.ToEncryptionWithCmkPtrOutputWithContext(context.Background())
}

func (i EncryptionWithCmkArgs) ToEncryptionWithCmkPtrOutputWithContext(ctx context.Context) EncryptionWithCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionWithCmkOutput).ToEncryptionWithCmkPtrOutputWithContext(ctx)
}









type EncryptionWithCmkPtrInput interface {
	pulumi.Input

	ToEncryptionWithCmkPtrOutput() EncryptionWithCmkPtrOutput
	ToEncryptionWithCmkPtrOutputWithContext(context.Context) EncryptionWithCmkPtrOutput
}

type encryptionWithCmkPtrType EncryptionWithCmkArgs

func EncryptionWithCmkPtr(v *EncryptionWithCmkArgs) EncryptionWithCmkPtrInput {
	return (*encryptionWithCmkPtrType)(v)
}

func (*encryptionWithCmkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionWithCmk)(nil)).Elem()
}

func (i *encryptionWithCmkPtrType) ToEncryptionWithCmkPtrOutput() EncryptionWithCmkPtrOutput {
	return i.ToEncryptionWithCmkPtrOutputWithContext(context.Background())
}

func (i *encryptionWithCmkPtrType) ToEncryptionWithCmkPtrOutputWithContext(ctx context.Context) EncryptionWithCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionWithCmkPtrOutput)
}

type EncryptionWithCmkOutput struct{ *pulumi.OutputState }

func (EncryptionWithCmkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionWithCmk)(nil)).Elem()
}

func (o EncryptionWithCmkOutput) ToEncryptionWithCmkOutput() EncryptionWithCmkOutput {
	return o
}

func (o EncryptionWithCmkOutput) ToEncryptionWithCmkOutputWithContext(ctx context.Context) EncryptionWithCmkOutput {
	return o
}

func (o EncryptionWithCmkOutput) ToEncryptionWithCmkPtrOutput() EncryptionWithCmkPtrOutput {
	return o.ToEncryptionWithCmkPtrOutputWithContext(context.Background())
}

func (o EncryptionWithCmkOutput) ToEncryptionWithCmkPtrOutputWithContext(ctx context.Context) EncryptionWithCmkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionWithCmk) *EncryptionWithCmk {
		return &v
	}).(EncryptionWithCmkPtrOutput)
}

func (o EncryptionWithCmkOutput) Enforcement() SearchEncryptionWithCmkPtrOutput {
	return o.ApplyT(func(v EncryptionWithCmk) *SearchEncryptionWithCmk { return v.Enforcement }).(SearchEncryptionWithCmkPtrOutput)
}

type EncryptionWithCmkPtrOutput struct{ *pulumi.OutputState }

func (EncryptionWithCmkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionWithCmk)(nil)).Elem()
}

func (o EncryptionWithCmkPtrOutput) ToEncryptionWithCmkPtrOutput() EncryptionWithCmkPtrOutput {
	return o
}

func (o EncryptionWithCmkPtrOutput) ToEncryptionWithCmkPtrOutputWithContext(ctx context.Context) EncryptionWithCmkPtrOutput {
	return o
}

func (o EncryptionWithCmkPtrOutput) Elem() EncryptionWithCmkOutput {
	return o.ApplyT(func(v *EncryptionWithCmk) EncryptionWithCmk {
		if v != nil {
			return *v
		}
		var ret EncryptionWithCmk
		return ret
	}).(EncryptionWithCmkOutput)
}

func (o EncryptionWithCmkPtrOutput) Enforcement() SearchEncryptionWithCmkPtrOutput {
	return o.ApplyT(func(v *EncryptionWithCmk) *SearchEncryptionWithCmk {
		if v == nil {
			return nil
		}
		return v.Enforcement
	}).(SearchEncryptionWithCmkPtrOutput)
}

type EncryptionWithCmkResponse struct {
	EncryptionComplianceStatus string  `pulumi:"encryptionComplianceStatus"`
	Enforcement                *string `pulumi:"enforcement"`
}





type EncryptionWithCmkResponseInput interface {
	pulumi.Input

	ToEncryptionWithCmkResponseOutput() EncryptionWithCmkResponseOutput
	ToEncryptionWithCmkResponseOutputWithContext(context.Context) EncryptionWithCmkResponseOutput
}

type EncryptionWithCmkResponseArgs struct {
	EncryptionComplianceStatus pulumi.StringInput    `pulumi:"encryptionComplianceStatus"`
	Enforcement                pulumi.StringPtrInput `pulumi:"enforcement"`
}

func (EncryptionWithCmkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionWithCmkResponse)(nil)).Elem()
}

func (i EncryptionWithCmkResponseArgs) ToEncryptionWithCmkResponseOutput() EncryptionWithCmkResponseOutput {
	return i.ToEncryptionWithCmkResponseOutputWithContext(context.Background())
}

func (i EncryptionWithCmkResponseArgs) ToEncryptionWithCmkResponseOutputWithContext(ctx context.Context) EncryptionWithCmkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionWithCmkResponseOutput)
}

func (i EncryptionWithCmkResponseArgs) ToEncryptionWithCmkResponsePtrOutput() EncryptionWithCmkResponsePtrOutput {
	return i.ToEncryptionWithCmkResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionWithCmkResponseArgs) ToEncryptionWithCmkResponsePtrOutputWithContext(ctx context.Context) EncryptionWithCmkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionWithCmkResponseOutput).ToEncryptionWithCmkResponsePtrOutputWithContext(ctx)
}









type EncryptionWithCmkResponsePtrInput interface {
	pulumi.Input

	ToEncryptionWithCmkResponsePtrOutput() EncryptionWithCmkResponsePtrOutput
	ToEncryptionWithCmkResponsePtrOutputWithContext(context.Context) EncryptionWithCmkResponsePtrOutput
}

type encryptionWithCmkResponsePtrType EncryptionWithCmkResponseArgs

func EncryptionWithCmkResponsePtr(v *EncryptionWithCmkResponseArgs) EncryptionWithCmkResponsePtrInput {
	return (*encryptionWithCmkResponsePtrType)(v)
}

func (*encryptionWithCmkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionWithCmkResponse)(nil)).Elem()
}

func (i *encryptionWithCmkResponsePtrType) ToEncryptionWithCmkResponsePtrOutput() EncryptionWithCmkResponsePtrOutput {
	return i.ToEncryptionWithCmkResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionWithCmkResponsePtrType) ToEncryptionWithCmkResponsePtrOutputWithContext(ctx context.Context) EncryptionWithCmkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionWithCmkResponsePtrOutput)
}

type EncryptionWithCmkResponseOutput struct{ *pulumi.OutputState }

func (EncryptionWithCmkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionWithCmkResponse)(nil)).Elem()
}

func (o EncryptionWithCmkResponseOutput) ToEncryptionWithCmkResponseOutput() EncryptionWithCmkResponseOutput {
	return o
}

func (o EncryptionWithCmkResponseOutput) ToEncryptionWithCmkResponseOutputWithContext(ctx context.Context) EncryptionWithCmkResponseOutput {
	return o
}

func (o EncryptionWithCmkResponseOutput) ToEncryptionWithCmkResponsePtrOutput() EncryptionWithCmkResponsePtrOutput {
	return o.ToEncryptionWithCmkResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionWithCmkResponseOutput) ToEncryptionWithCmkResponsePtrOutputWithContext(ctx context.Context) EncryptionWithCmkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionWithCmkResponse) *EncryptionWithCmkResponse {
		return &v
	}).(EncryptionWithCmkResponsePtrOutput)
}

func (o EncryptionWithCmkResponseOutput) EncryptionComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionWithCmkResponse) string { return v.EncryptionComplianceStatus }).(pulumi.StringOutput)
}

func (o EncryptionWithCmkResponseOutput) Enforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionWithCmkResponse) *string { return v.Enforcement }).(pulumi.StringPtrOutput)
}

type EncryptionWithCmkResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionWithCmkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionWithCmkResponse)(nil)).Elem()
}

func (o EncryptionWithCmkResponsePtrOutput) ToEncryptionWithCmkResponsePtrOutput() EncryptionWithCmkResponsePtrOutput {
	return o
}

func (o EncryptionWithCmkResponsePtrOutput) ToEncryptionWithCmkResponsePtrOutputWithContext(ctx context.Context) EncryptionWithCmkResponsePtrOutput {
	return o
}

func (o EncryptionWithCmkResponsePtrOutput) Elem() EncryptionWithCmkResponseOutput {
	return o.ApplyT(func(v *EncryptionWithCmkResponse) EncryptionWithCmkResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionWithCmkResponse
		return ret
	}).(EncryptionWithCmkResponseOutput)
}

func (o EncryptionWithCmkResponsePtrOutput) EncryptionComplianceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionWithCmkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionComplianceStatus
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionWithCmkResponsePtrOutput) Enforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionWithCmkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enforcement
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Identity) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                         `pulumi:"principalId"`
	TenantId               string                                         `pulumi:"tenantId"`
	Type                   string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedManagedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                          `pulumi:"principalId"`
	TenantId               pulumi.StringInput                          `pulumi:"tenantId"`
	Type                   pulumi.StringInput                          `pulumi:"type"`
	UserAssignedIdentities UserAssignedManagedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type IpRule struct {
	Value *string `pulumi:"value"`
}





type IpRuleInput interface {
	pulumi.Input

	ToIpRuleOutput() IpRuleOutput
	ToIpRuleOutputWithContext(context.Context) IpRuleOutput
}

type IpRuleArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (IpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (i IpRuleArgs) ToIpRuleOutput() IpRuleOutput {
	return i.ToIpRuleOutputWithContext(context.Background())
}

func (i IpRuleArgs) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleOutput)
}





type IpRuleArrayInput interface {
	pulumi.Input

	ToIpRuleArrayOutput() IpRuleArrayOutput
	ToIpRuleArrayOutputWithContext(context.Context) IpRuleArrayOutput
}

type IpRuleArray []IpRuleInput

func (IpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (i IpRuleArray) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return i.ToIpRuleArrayOutputWithContext(context.Background())
}

func (i IpRuleArray) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleArrayOutput)
}

type IpRuleOutput struct{ *pulumi.OutputState }

func (IpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (o IpRuleOutput) ToIpRuleOutput() IpRuleOutput {
	return o
}

func (o IpRuleOutput) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return o
}

func (o IpRuleOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpRule) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IpRuleArrayOutput struct{ *pulumi.OutputState }

func (IpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) Index(i pulumi.IntInput) IpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRule {
		return vs[0].([]IpRule)[vs[1].(int)]
	}).(IpRuleOutput)
}

type IpRuleResponse struct {
	Value *string `pulumi:"value"`
}





type IpRuleResponseInput interface {
	pulumi.Input

	ToIpRuleResponseOutput() IpRuleResponseOutput
	ToIpRuleResponseOutputWithContext(context.Context) IpRuleResponseOutput
}

type IpRuleResponseArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (IpRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (i IpRuleResponseArgs) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return i.ToIpRuleResponseOutputWithContext(context.Background())
}

func (i IpRuleResponseArgs) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleResponseOutput)
}





type IpRuleResponseArrayInput interface {
	pulumi.Input

	ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput
	ToIpRuleResponseArrayOutputWithContext(context.Context) IpRuleResponseArrayOutput
}

type IpRuleResponseArray []IpRuleResponseInput

func (IpRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (i IpRuleResponseArray) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return i.ToIpRuleResponseArrayOutputWithContext(context.Background())
}

func (i IpRuleResponseArray) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleResponseArrayOutput)
}

type IpRuleResponseOutput struct{ *pulumi.OutputState }

func (IpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpRuleResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) Index(i pulumi.IntInput) IpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRuleResponse {
		return vs[0].([]IpRuleResponse)[vs[1].(int)]
	}).(IpRuleResponseOutput)
}

type NetworkRuleSet struct {
	Bypass  *string  `pulumi:"bypass"`
	IpRules []IpRule `pulumi:"ipRules"`
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	Bypass  pulumi.StringPtrInput `pulumi:"bypass"`
	IpRules IpRuleArrayInput      `pulumi:"ipRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IpRule { return v.IpRules }).(IpRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IpRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	Bypass  *string          `pulumi:"bypass"`
	IpRules []IpRuleResponse `pulumi:"ipRules"`
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	Bypass  pulumi.StringPtrInput    `pulumi:"bypass"`
	IpRules IpRuleResponseArrayInput `pulumi:"ipRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IpRuleResponse { return v.IpRules }).(IpRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IpRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleResponseArrayOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateEndpoint                   *PrivateEndpointConnectionPropertiesPrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateEndpoint                   PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateEndpoint {
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointConnectionPropertiesPrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateEndpointOutput() PrivateEndpointConnectionPropertiesPrivateEndpointOutput
	ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointOutput
}

type PrivateEndpointConnectionPropertiesPrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutput() PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateEndpointOutput)
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateEndpointOutput).ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput
	ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput
}

type privateEndpointConnectionPropertiesPrivateEndpointPtrType PrivateEndpointConnectionPropertiesPrivateEndpointArgs

func PrivateEndpointConnectionPropertiesPrivateEndpointPtr(v *PrivateEndpointConnectionPropertiesPrivateEndpointArgs) PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput {
	return (*privateEndpointConnectionPropertiesPrivateEndpointPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPrivateEndpointPtrType) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPrivateEndpointPtrType) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutput() PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesPrivateEndpoint) *PrivateEndpointConnectionPropertiesPrivateEndpoint {
		return &v
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) Elem() PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateEndpoint) PrivateEndpointConnectionPropertiesPrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesPrivateEndpoint
		return ret
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState struct {
	ActionsRequired *string                             `pulumi:"actionsRequired"`
	Description     *string                             `pulumi:"description"`
	Status          *PrivateLinkServiceConnectionStatus `pulumi:"status"`
}





type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput
	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput                      `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput                      `pulumi:"description"`
	Status          PrivateLinkServiceConnectionStatusPtrInput `pulumi:"status"`
}

func (PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput).ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
}

type privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs

func PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtr(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput {
	return (*privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType)(v)
}

func (*privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		return &v
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Status() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionStatus {
		return v.Status
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState
		return ret
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Status() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointConnectionPropertiesResponsePrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput
	ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput)
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput).ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput
}

type privateEndpointConnectionPropertiesResponsePrivateEndpointPtrType PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs

func PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtr(v *PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrInput {
	return (*privateEndpointConnectionPropertiesResponsePrivateEndpointPtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePrivateEndpointPtrType) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePrivateEndpointPtrType) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) Elem() PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponsePrivateEndpoint
		return ret
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput
	ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput)
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput).ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput
}

type privateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs

func PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtr(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState
		return ret
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         string                                       `pulumi:"id"`
	Name       string                                       `pulumi:"name"`
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id         pulumi.StringInput                                  `pulumi:"id"`
	Name       pulumi.StringInput                                  `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput `pulumi:"properties"`
	Type       pulumi.StringInput                                  `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
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

type QueryKeyResponse struct {
	Key  string `pulumi:"key"`
	Name string `pulumi:"name"`
}





type QueryKeyResponseInput interface {
	pulumi.Input

	ToQueryKeyResponseOutput() QueryKeyResponseOutput
	ToQueryKeyResponseOutputWithContext(context.Context) QueryKeyResponseOutput
}

type QueryKeyResponseArgs struct {
	Key  pulumi.StringInput `pulumi:"key"`
	Name pulumi.StringInput `pulumi:"name"`
}

func (QueryKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryKeyResponse)(nil)).Elem()
}

func (i QueryKeyResponseArgs) ToQueryKeyResponseOutput() QueryKeyResponseOutput {
	return i.ToQueryKeyResponseOutputWithContext(context.Background())
}

func (i QueryKeyResponseArgs) ToQueryKeyResponseOutputWithContext(ctx context.Context) QueryKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryKeyResponseOutput)
}





type QueryKeyResponseArrayInput interface {
	pulumi.Input

	ToQueryKeyResponseArrayOutput() QueryKeyResponseArrayOutput
	ToQueryKeyResponseArrayOutputWithContext(context.Context) QueryKeyResponseArrayOutput
}

type QueryKeyResponseArray []QueryKeyResponseInput

func (QueryKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryKeyResponse)(nil)).Elem()
}

func (i QueryKeyResponseArray) ToQueryKeyResponseArrayOutput() QueryKeyResponseArrayOutput {
	return i.ToQueryKeyResponseArrayOutputWithContext(context.Background())
}

func (i QueryKeyResponseArray) ToQueryKeyResponseArrayOutputWithContext(ctx context.Context) QueryKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryKeyResponseArrayOutput)
}

type QueryKeyResponseOutput struct{ *pulumi.OutputState }

func (QueryKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryKeyResponse)(nil)).Elem()
}

func (o QueryKeyResponseOutput) ToQueryKeyResponseOutput() QueryKeyResponseOutput {
	return o
}

func (o QueryKeyResponseOutput) ToQueryKeyResponseOutputWithContext(ctx context.Context) QueryKeyResponseOutput {
	return o
}

func (o QueryKeyResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v QueryKeyResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o QueryKeyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryKeyResponse) string { return v.Name }).(pulumi.StringOutput)
}

type QueryKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (QueryKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryKeyResponse)(nil)).Elem()
}

func (o QueryKeyResponseArrayOutput) ToQueryKeyResponseArrayOutput() QueryKeyResponseArrayOutput {
	return o
}

func (o QueryKeyResponseArrayOutput) ToQueryKeyResponseArrayOutputWithContext(ctx context.Context) QueryKeyResponseArrayOutput {
	return o
}

func (o QueryKeyResponseArrayOutput) Index(i pulumi.IntInput) QueryKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryKeyResponse {
		return vs[0].([]QueryKeyResponse)[vs[1].(int)]
	}).(QueryKeyResponseOutput)
}

type SharedPrivateLinkResourceProperties struct {
	GroupId               *string `pulumi:"groupId"`
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	ProvisioningState     *string `pulumi:"provisioningState"`
	RequestMessage        *string `pulumi:"requestMessage"`
	ResourceRegion        *string `pulumi:"resourceRegion"`
	Status                *string `pulumi:"status"`
}





type SharedPrivateLinkResourcePropertiesInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourcePropertiesOutput() SharedPrivateLinkResourcePropertiesOutput
	ToSharedPrivateLinkResourcePropertiesOutputWithContext(context.Context) SharedPrivateLinkResourcePropertiesOutput
}

type SharedPrivateLinkResourcePropertiesArgs struct {
	GroupId               pulumi.StringPtrInput `pulumi:"groupId"`
	PrivateLinkResourceId pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	ProvisioningState     pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequestMessage        pulumi.StringPtrInput `pulumi:"requestMessage"`
	ResourceRegion        pulumi.StringPtrInput `pulumi:"resourceRegion"`
	Status                pulumi.StringPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesOutput() SharedPrivateLinkResourcePropertiesOutput {
	return i.ToSharedPrivateLinkResourcePropertiesOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesOutput)
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return i.ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourcePropertiesArgs) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesOutput).ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx)
}









type SharedPrivateLinkResourcePropertiesPtrInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput
	ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Context) SharedPrivateLinkResourcePropertiesPtrOutput
}

type sharedPrivateLinkResourcePropertiesPtrType SharedPrivateLinkResourcePropertiesArgs

func SharedPrivateLinkResourcePropertiesPtr(v *SharedPrivateLinkResourcePropertiesArgs) SharedPrivateLinkResourcePropertiesPtrInput {
	return (*sharedPrivateLinkResourcePropertiesPtrType)(v)
}

func (*sharedPrivateLinkResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (i *sharedPrivateLinkResourcePropertiesPtrType) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return i.ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *sharedPrivateLinkResourcePropertiesPtrType) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesPtrOutput)
}

type SharedPrivateLinkResourcePropertiesOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesOutput() SharedPrivateLinkResourcePropertiesOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return o.ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourcePropertiesOutput) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedPrivateLinkResourceProperties) *SharedPrivateLinkResourceProperties {
		return &v
	}).(SharedPrivateLinkResourcePropertiesPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) ResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.ResourceRegion }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourceProperties)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) ToSharedPrivateLinkResourcePropertiesPtrOutput() SharedPrivateLinkResourcePropertiesPtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) ToSharedPrivateLinkResourcePropertiesPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesPtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) Elem() SharedPrivateLinkResourcePropertiesOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) SharedPrivateLinkResourceProperties {
		if v != nil {
			return *v
		}
		var ret SharedPrivateLinkResourceProperties
		return ret
	}).(SharedPrivateLinkResourcePropertiesOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.GroupId
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.RequestMessage
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) ResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceRegion
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourcePropertiesResponse struct {
	GroupId               *string `pulumi:"groupId"`
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	ProvisioningState     *string `pulumi:"provisioningState"`
	RequestMessage        *string `pulumi:"requestMessage"`
	ResourceRegion        *string `pulumi:"resourceRegion"`
	Status                *string `pulumi:"status"`
}





type SharedPrivateLinkResourcePropertiesResponseInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourcePropertiesResponseOutput() SharedPrivateLinkResourcePropertiesResponseOutput
	ToSharedPrivateLinkResourcePropertiesResponseOutputWithContext(context.Context) SharedPrivateLinkResourcePropertiesResponseOutput
}

type SharedPrivateLinkResourcePropertiesResponseArgs struct {
	GroupId               pulumi.StringPtrInput `pulumi:"groupId"`
	PrivateLinkResourceId pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	ProvisioningState     pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequestMessage        pulumi.StringPtrInput `pulumi:"requestMessage"`
	ResourceRegion        pulumi.StringPtrInput `pulumi:"resourceRegion"`
	Status                pulumi.StringPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourcePropertiesResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourcePropertiesResponseArgs) ToSharedPrivateLinkResourcePropertiesResponseOutput() SharedPrivateLinkResourcePropertiesResponseOutput {
	return i.ToSharedPrivateLinkResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourcePropertiesResponseArgs) ToSharedPrivateLinkResourcePropertiesResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesResponseOutput)
}

func (i SharedPrivateLinkResourcePropertiesResponseArgs) ToSharedPrivateLinkResourcePropertiesResponsePtrOutput() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return i.ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourcePropertiesResponseArgs) ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesResponseOutput).ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type SharedPrivateLinkResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourcePropertiesResponsePtrOutput() SharedPrivateLinkResourcePropertiesResponsePtrOutput
	ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(context.Context) SharedPrivateLinkResourcePropertiesResponsePtrOutput
}

type sharedPrivateLinkResourcePropertiesResponsePtrType SharedPrivateLinkResourcePropertiesResponseArgs

func SharedPrivateLinkResourcePropertiesResponsePtr(v *SharedPrivateLinkResourcePropertiesResponseArgs) SharedPrivateLinkResourcePropertiesResponsePtrInput {
	return (*sharedPrivateLinkResourcePropertiesResponsePtrType)(v)
}

func (*sharedPrivateLinkResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourcePropertiesResponse)(nil)).Elem()
}

func (i *sharedPrivateLinkResourcePropertiesResponsePtrType) ToSharedPrivateLinkResourcePropertiesResponsePtrOutput() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return i.ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *sharedPrivateLinkResourcePropertiesResponsePtrType) ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourcePropertiesResponsePtrOutput)
}

type SharedPrivateLinkResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourcePropertiesResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ToSharedPrivateLinkResourcePropertiesResponseOutput() SharedPrivateLinkResourcePropertiesResponseOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ToSharedPrivateLinkResourcePropertiesResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponseOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ToSharedPrivateLinkResourcePropertiesResponsePtrOutput() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o.ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedPrivateLinkResourcePropertiesResponse) *SharedPrivateLinkResourcePropertiesResponse {
		return &v
	}).(SharedPrivateLinkResourcePropertiesResponsePtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) ResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.ResourceRegion }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourcePropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourcePropertiesResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) ToSharedPrivateLinkResourcePropertiesResponsePtrOutput() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) ToSharedPrivateLinkResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) Elem() SharedPrivateLinkResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) SharedPrivateLinkResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SharedPrivateLinkResourcePropertiesResponse
		return ret
	}).(SharedPrivateLinkResourcePropertiesResponseOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupId
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkResourceId
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestMessage
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) ResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceRegion
	}).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourcePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceResponse struct {
	Id         string                                       `pulumi:"id"`
	Name       string                                       `pulumi:"name"`
	Properties *SharedPrivateLinkResourcePropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}





type SharedPrivateLinkResourceResponseInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput
	ToSharedPrivateLinkResourceResponseOutputWithContext(context.Context) SharedPrivateLinkResourceResponseOutput
}

type SharedPrivateLinkResourceResponseArgs struct {
	Id         pulumi.StringInput                                  `pulumi:"id"`
	Name       pulumi.StringInput                                  `pulumi:"name"`
	Properties SharedPrivateLinkResourcePropertiesResponsePtrInput `pulumi:"properties"`
	Type       pulumi.StringInput                                  `pulumi:"type"`
}

func (SharedPrivateLinkResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return i.ToSharedPrivateLinkResourceResponseOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseOutput)
}





type SharedPrivateLinkResourceResponseArrayInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput
	ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Context) SharedPrivateLinkResourceResponseArrayOutput
}

type SharedPrivateLinkResourceResponseArray []SharedPrivateLinkResourceResponseInput

func (SharedPrivateLinkResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return i.ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseArrayOutput)
}

type SharedPrivateLinkResourceResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Properties() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *SharedPrivateLinkResourcePropertiesResponse {
		return v.Properties
	}).(SharedPrivateLinkResourcePropertiesResponsePtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SharedPrivateLinkResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResourceResponse {
		return vs[0].([]SharedPrivateLinkResourceResponse)[vs[1].(int)]
	}).(SharedPrivateLinkResourceResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type UserAssignedManagedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedManagedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput
	ToUserAssignedManagedIdentityResponseOutputWithContext(context.Context) UserAssignedManagedIdentityResponseOutput
}

type UserAssignedManagedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedManagedIdentityResponseArgs) ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput {
	return i.ToUserAssignedManagedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedManagedIdentityResponseArgs) ToUserAssignedManagedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedManagedIdentityResponseOutput)
}





type UserAssignedManagedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput
	ToUserAssignedManagedIdentityResponseMapOutputWithContext(context.Context) UserAssignedManagedIdentityResponseMapOutput
}

type UserAssignedManagedIdentityResponseMap map[string]UserAssignedManagedIdentityResponseInput

func (UserAssignedManagedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedManagedIdentityResponseMap) ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput {
	return i.ToUserAssignedManagedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedManagedIdentityResponseMap) ToUserAssignedManagedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedManagedIdentityResponseMapOutput)
}

type UserAssignedManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedManagedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedManagedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedManagedIdentityResponse {
		return vs[0].(map[string]UserAssignedManagedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedManagedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOptionInput)(nil)).Elem(), DataPlaneAadOrApiKeyAuthOptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOptionPtrInput)(nil)).Elem(), DataPlaneAadOrApiKeyAuthOptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOptionResponseInput)(nil)).Elem(), DataPlaneAadOrApiKeyAuthOptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAadOrApiKeyAuthOptionResponsePtrInput)(nil)).Elem(), DataPlaneAadOrApiKeyAuthOptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAuthOptionsInput)(nil)).Elem(), DataPlaneAuthOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAuthOptionsPtrInput)(nil)).Elem(), DataPlaneAuthOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAuthOptionsResponseInput)(nil)).Elem(), DataPlaneAuthOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataPlaneAuthOptionsResponsePtrInput)(nil)).Elem(), DataPlaneAuthOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionWithCmkInput)(nil)).Elem(), EncryptionWithCmkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionWithCmkPtrInput)(nil)).Elem(), EncryptionWithCmkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionWithCmkResponseInput)(nil)).Elem(), EncryptionWithCmkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionWithCmkResponsePtrInput)(nil)).Elem(), EncryptionWithCmkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRuleInput)(nil)).Elem(), IpRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRuleArrayInput)(nil)).Elem(), IpRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRuleResponseInput)(nil)).Elem(), IpRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpRuleResponseArrayInput)(nil)).Elem(), IpRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetInput)(nil)).Elem(), NetworkRuleSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetPtrInput)(nil)).Elem(), NetworkRuleSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetResponseInput)(nil)).Elem(), NetworkRuleSetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetResponsePtrInput)(nil)).Elem(), NetworkRuleSetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesPtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateEndpointInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesPrivateEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesPrivateEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponseInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateEndpointInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueryKeyResponseInput)(nil)).Elem(), QueryKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueryKeyResponseArrayInput)(nil)).Elem(), QueryKeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourcePropertiesInput)(nil)).Elem(), SharedPrivateLinkResourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourcePropertiesPtrInput)(nil)).Elem(), SharedPrivateLinkResourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourcePropertiesResponseInput)(nil)).Elem(), SharedPrivateLinkResourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourcePropertiesResponsePtrInput)(nil)).Elem(), SharedPrivateLinkResourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceResponseInput)(nil)).Elem(), SharedPrivateLinkResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceResponseArrayInput)(nil)).Elem(), SharedPrivateLinkResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedManagedIdentityResponseInput)(nil)).Elem(), UserAssignedManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedManagedIdentityResponseMapInput)(nil)).Elem(), UserAssignedManagedIdentityResponseMap{})
	pulumi.RegisterOutputType(DataPlaneAadOrApiKeyAuthOptionOutput{})
	pulumi.RegisterOutputType(DataPlaneAadOrApiKeyAuthOptionPtrOutput{})
	pulumi.RegisterOutputType(DataPlaneAadOrApiKeyAuthOptionResponseOutput{})
	pulumi.RegisterOutputType(DataPlaneAadOrApiKeyAuthOptionResponsePtrOutput{})
	pulumi.RegisterOutputType(DataPlaneAuthOptionsOutput{})
	pulumi.RegisterOutputType(DataPlaneAuthOptionsPtrOutput{})
	pulumi.RegisterOutputType(DataPlaneAuthOptionsResponseOutput{})
	pulumi.RegisterOutputType(DataPlaneAuthOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionWithCmkOutput{})
	pulumi.RegisterOutputType(EncryptionWithCmkPtrOutput{})
	pulumi.RegisterOutputType(EncryptionWithCmkResponseOutput{})
	pulumi.RegisterOutputType(EncryptionWithCmkResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IpRuleOutput{})
	pulumi.RegisterOutputType(IpRuleArrayOutput{})
	pulumi.RegisterOutputType(IpRuleResponseOutput{})
	pulumi.RegisterOutputType(IpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(QueryKeyResponseOutput{})
	pulumi.RegisterOutputType(QueryKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseMapOutput{})
}
