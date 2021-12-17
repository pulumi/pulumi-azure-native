


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CaptureDescription struct {
	Destination       *Destination                `pulumi:"destination"`
	Enabled           *bool                       `pulumi:"enabled"`
	Encoding          *EncodingCaptureDescription `pulumi:"encoding"`
	IntervalInSeconds *int                        `pulumi:"intervalInSeconds"`
	SizeLimitInBytes  *int                        `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives *bool                       `pulumi:"skipEmptyArchives"`
}





type CaptureDescriptionInput interface {
	pulumi.Input

	ToCaptureDescriptionOutput() CaptureDescriptionOutput
	ToCaptureDescriptionOutputWithContext(context.Context) CaptureDescriptionOutput
}

type CaptureDescriptionArgs struct {
	Destination       DestinationPtrInput                `pulumi:"destination"`
	Enabled           pulumi.BoolPtrInput                `pulumi:"enabled"`
	Encoding          EncodingCaptureDescriptionPtrInput `pulumi:"encoding"`
	IntervalInSeconds pulumi.IntPtrInput                 `pulumi:"intervalInSeconds"`
	SizeLimitInBytes  pulumi.IntPtrInput                 `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives pulumi.BoolPtrInput                `pulumi:"skipEmptyArchives"`
}

func (CaptureDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CaptureDescription)(nil)).Elem()
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionOutput() CaptureDescriptionOutput {
	return i.ToCaptureDescriptionOutputWithContext(context.Background())
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionOutputWithContext(ctx context.Context) CaptureDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaptureDescriptionOutput)
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return i.ToCaptureDescriptionPtrOutputWithContext(context.Background())
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaptureDescriptionOutput).ToCaptureDescriptionPtrOutputWithContext(ctx)
}









type CaptureDescriptionPtrInput interface {
	pulumi.Input

	ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput
	ToCaptureDescriptionPtrOutputWithContext(context.Context) CaptureDescriptionPtrOutput
}

type captureDescriptionPtrType CaptureDescriptionArgs

func CaptureDescriptionPtr(v *CaptureDescriptionArgs) CaptureDescriptionPtrInput {
	return (*captureDescriptionPtrType)(v)
}

func (*captureDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CaptureDescription)(nil)).Elem()
}

func (i *captureDescriptionPtrType) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return i.ToCaptureDescriptionPtrOutputWithContext(context.Background())
}

func (i *captureDescriptionPtrType) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaptureDescriptionPtrOutput)
}

type CaptureDescriptionOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaptureDescription)(nil)).Elem()
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionOutput() CaptureDescriptionOutput {
	return o
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionOutputWithContext(ctx context.Context) CaptureDescriptionOutput {
	return o
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return o.ToCaptureDescriptionPtrOutputWithContext(context.Background())
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CaptureDescription) *CaptureDescription {
		return &v
	}).(CaptureDescriptionPtrOutput)
}

func (o CaptureDescriptionOutput) Destination() DestinationPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *Destination { return v.Destination }).(DestinationPtrOutput)
}

func (o CaptureDescriptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionOutput) Encoding() EncodingCaptureDescriptionPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *EncodingCaptureDescription { return v.Encoding }).(EncodingCaptureDescriptionPtrOutput)
}

func (o CaptureDescriptionOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *int { return v.SizeLimitInBytes }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *bool { return v.SkipEmptyArchives }).(pulumi.BoolPtrOutput)
}

type CaptureDescriptionPtrOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaptureDescription)(nil)).Elem()
}

func (o CaptureDescriptionPtrOutput) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return o
}

func (o CaptureDescriptionPtrOutput) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return o
}

func (o CaptureDescriptionPtrOutput) Elem() CaptureDescriptionOutput {
	return o.ApplyT(func(v *CaptureDescription) CaptureDescription {
		if v != nil {
			return *v
		}
		var ret CaptureDescription
		return ret
	}).(CaptureDescriptionOutput)
}

func (o CaptureDescriptionPtrOutput) Destination() DestinationPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *Destination {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(DestinationPtrOutput)
}

func (o CaptureDescriptionPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionPtrOutput) Encoding() EncodingCaptureDescriptionPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *EncodingCaptureDescription {
		if v == nil {
			return nil
		}
		return v.Encoding
	}).(EncodingCaptureDescriptionPtrOutput)
}

func (o CaptureDescriptionPtrOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *int {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionPtrOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *int {
		if v == nil {
			return nil
		}
		return v.SizeLimitInBytes
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionPtrOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *bool {
		if v == nil {
			return nil
		}
		return v.SkipEmptyArchives
	}).(pulumi.BoolPtrOutput)
}

type CaptureDescriptionResponse struct {
	Destination       *DestinationResponse `pulumi:"destination"`
	Enabled           *bool                `pulumi:"enabled"`
	Encoding          *string              `pulumi:"encoding"`
	IntervalInSeconds *int                 `pulumi:"intervalInSeconds"`
	SizeLimitInBytes  *int                 `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives *bool                `pulumi:"skipEmptyArchives"`
}

type CaptureDescriptionResponseOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaptureDescriptionResponse)(nil)).Elem()
}

func (o CaptureDescriptionResponseOutput) ToCaptureDescriptionResponseOutput() CaptureDescriptionResponseOutput {
	return o
}

func (o CaptureDescriptionResponseOutput) ToCaptureDescriptionResponseOutputWithContext(ctx context.Context) CaptureDescriptionResponseOutput {
	return o
}

func (o CaptureDescriptionResponseOutput) Destination() DestinationResponsePtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *DestinationResponse { return v.Destination }).(DestinationResponsePtrOutput)
}

func (o CaptureDescriptionResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionResponseOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o CaptureDescriptionResponseOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponseOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *int { return v.SizeLimitInBytes }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponseOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *bool { return v.SkipEmptyArchives }).(pulumi.BoolPtrOutput)
}

type CaptureDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaptureDescriptionResponse)(nil)).Elem()
}

func (o CaptureDescriptionResponsePtrOutput) ToCaptureDescriptionResponsePtrOutput() CaptureDescriptionResponsePtrOutput {
	return o
}

func (o CaptureDescriptionResponsePtrOutput) ToCaptureDescriptionResponsePtrOutputWithContext(ctx context.Context) CaptureDescriptionResponsePtrOutput {
	return o
}

func (o CaptureDescriptionResponsePtrOutput) Elem() CaptureDescriptionResponseOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) CaptureDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret CaptureDescriptionResponse
		return ret
	}).(CaptureDescriptionResponseOutput)
}

func (o CaptureDescriptionResponsePtrOutput) Destination() DestinationResponsePtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *DestinationResponse {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(DestinationResponsePtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Encoding
	}).(pulumi.StringPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.SizeLimitInBytes
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SkipEmptyArchives
	}).(pulumi.BoolPtrOutput)
}

type Destination struct {
	ArchiveNameFormat        *string `pulumi:"archiveNameFormat"`
	BlobContainer            *string `pulumi:"blobContainer"`
	Name                     *string `pulumi:"name"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}





type DestinationInput interface {
	pulumi.Input

	ToDestinationOutput() DestinationOutput
	ToDestinationOutputWithContext(context.Context) DestinationOutput
}

type DestinationArgs struct {
	ArchiveNameFormat        pulumi.StringPtrInput `pulumi:"archiveNameFormat"`
	BlobContainer            pulumi.StringPtrInput `pulumi:"blobContainer"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
}

func (DestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Destination)(nil)).Elem()
}

func (i DestinationArgs) ToDestinationOutput() DestinationOutput {
	return i.ToDestinationOutputWithContext(context.Background())
}

func (i DestinationArgs) ToDestinationOutputWithContext(ctx context.Context) DestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationOutput)
}

func (i DestinationArgs) ToDestinationPtrOutput() DestinationPtrOutput {
	return i.ToDestinationPtrOutputWithContext(context.Background())
}

func (i DestinationArgs) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationOutput).ToDestinationPtrOutputWithContext(ctx)
}









type DestinationPtrInput interface {
	pulumi.Input

	ToDestinationPtrOutput() DestinationPtrOutput
	ToDestinationPtrOutputWithContext(context.Context) DestinationPtrOutput
}

type destinationPtrType DestinationArgs

func DestinationPtr(v *DestinationArgs) DestinationPtrInput {
	return (*destinationPtrType)(v)
}

func (*destinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Destination)(nil)).Elem()
}

func (i *destinationPtrType) ToDestinationPtrOutput() DestinationPtrOutput {
	return i.ToDestinationPtrOutputWithContext(context.Background())
}

func (i *destinationPtrType) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationPtrOutput)
}

type DestinationOutput struct{ *pulumi.OutputState }

func (DestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Destination)(nil)).Elem()
}

func (o DestinationOutput) ToDestinationOutput() DestinationOutput {
	return o
}

func (o DestinationOutput) ToDestinationOutputWithContext(ctx context.Context) DestinationOutput {
	return o
}

func (o DestinationOutput) ToDestinationPtrOutput() DestinationPtrOutput {
	return o.ToDestinationPtrOutputWithContext(context.Background())
}

func (o DestinationOutput) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Destination) *Destination {
		return &v
	}).(DestinationPtrOutput)
}

func (o DestinationOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.ArchiveNameFormat }).(pulumi.StringPtrOutput)
}

func (o DestinationOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.BlobContainer }).(pulumi.StringPtrOutput)
}

func (o DestinationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DestinationOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type DestinationPtrOutput struct{ *pulumi.OutputState }

func (DestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Destination)(nil)).Elem()
}

func (o DestinationPtrOutput) ToDestinationPtrOutput() DestinationPtrOutput {
	return o
}

func (o DestinationPtrOutput) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return o
}

func (o DestinationPtrOutput) Elem() DestinationOutput {
	return o.ApplyT(func(v *Destination) Destination {
		if v != nil {
			return *v
		}
		var ret Destination
		return ret
	}).(DestinationOutput)
}

func (o DestinationPtrOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.ArchiveNameFormat
	}).(pulumi.StringPtrOutput)
}

func (o DestinationPtrOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainer
	}).(pulumi.StringPtrOutput)
}

func (o DestinationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DestinationPtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type DestinationResponse struct {
	ArchiveNameFormat        *string `pulumi:"archiveNameFormat"`
	BlobContainer            *string `pulumi:"blobContainer"`
	Name                     *string `pulumi:"name"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}

type DestinationResponseOutput struct{ *pulumi.OutputState }

func (DestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationResponse)(nil)).Elem()
}

func (o DestinationResponseOutput) ToDestinationResponseOutput() DestinationResponseOutput {
	return o
}

func (o DestinationResponseOutput) ToDestinationResponseOutputWithContext(ctx context.Context) DestinationResponseOutput {
	return o
}

func (o DestinationResponseOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.ArchiveNameFormat }).(pulumi.StringPtrOutput)
}

func (o DestinationResponseOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.BlobContainer }).(pulumi.StringPtrOutput)
}

func (o DestinationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DestinationResponseOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type DestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (DestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationResponse)(nil)).Elem()
}

func (o DestinationResponsePtrOutput) ToDestinationResponsePtrOutput() DestinationResponsePtrOutput {
	return o
}

func (o DestinationResponsePtrOutput) ToDestinationResponsePtrOutputWithContext(ctx context.Context) DestinationResponsePtrOutput {
	return o
}

func (o DestinationResponsePtrOutput) Elem() DestinationResponseOutput {
	return o.ApplyT(func(v *DestinationResponse) DestinationResponse {
		if v != nil {
			return *v
		}
		var ret DestinationResponse
		return ret
	}).(DestinationResponseOutput)
}

func (o DestinationResponsePtrOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArchiveNameFormat
	}).(pulumi.StringPtrOutput)
}

func (o DestinationResponsePtrOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainer
	}).(pulumi.StringPtrOutput)
}

func (o DestinationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DestinationResponsePtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRules struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRules) Defaults() *NWRuleSetIpRules {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type NWRuleSetIpRulesInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput
	ToNWRuleSetIpRulesOutputWithContext(context.Context) NWRuleSetIpRulesOutput
}

type NWRuleSetIpRulesArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}

func (NWRuleSetIpRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return i.ToNWRuleSetIpRulesOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesOutput)
}





type NWRuleSetIpRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput
	ToNWRuleSetIpRulesArrayOutputWithContext(context.Context) NWRuleSetIpRulesArrayOutput
}

type NWRuleSetIpRulesArray []NWRuleSetIpRulesInput

func (NWRuleSetIpRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return i.ToNWRuleSetIpRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesArrayOutput)
}

type NWRuleSetIpRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRules {
		return vs[0].([]NWRuleSetIpRules)[vs[1].(int)]
	}).(NWRuleSetIpRulesOutput)
}

type NWRuleSetIpRulesResponse struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRulesResponse) Defaults() *NWRuleSetIpRulesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type NWRuleSetIpRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutput() NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutput() NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRulesResponse {
		return vs[0].([]NWRuleSetIpRulesResponse)[vs[1].(int)]
	}).(NWRuleSetIpRulesResponseOutput)
}

type NWRuleSetVirtualNetworkRules struct {
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *Subnet `pulumi:"subnet"`
}





type NWRuleSetVirtualNetworkRulesInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput
	ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesOutput
}

type NWRuleSetVirtualNetworkRulesArgs struct {
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           SubnetPtrInput      `pulumi:"subnet"`
}

func (NWRuleSetVirtualNetworkRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return i.ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesOutput)
}





type NWRuleSetVirtualNetworkRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput
	ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesArrayOutput
}

type NWRuleSetVirtualNetworkRulesArray []NWRuleSetVirtualNetworkRulesInput

func (NWRuleSetVirtualNetworkRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return i.ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesArrayOutput)
}

type NWRuleSetVirtualNetworkRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesOutput) Subnet() SubnetPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *Subnet { return v.Subnet }).(SubnetPtrOutput)
}

type NWRuleSetVirtualNetworkRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRules {
		return vs[0].([]NWRuleSetVirtualNetworkRules)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesOutput)
}

type NWRuleSetVirtualNetworkRulesResponse struct {
	IgnoreMissingVnetServiceEndpoint *bool           `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *SubnetResponse `pulumi:"subnet"`
}

type NWRuleSetVirtualNetworkRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutput() NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type NWRuleSetVirtualNetworkRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutput() NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRulesResponse {
		return vs[0].([]NWRuleSetVirtualNetworkRulesResponse)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesResponseOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type Subnet struct {
	Id string `pulumi:"id"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i SubnetArgs) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput).ToSubnetPtrOutputWithContext(ctx)
}









type SubnetPtrInput interface {
	pulumi.Input

	ToSubnetPtrOutput() SubnetPtrOutput
	ToSubnetPtrOutputWithContext(context.Context) SubnetPtrOutput
}

type subnetPtrType SubnetArgs

func SubnetPtr(v *SubnetArgs) SubnetPtrInput {
	return (*subnetPtrType)(v)
}

func (*subnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *subnetPtrType) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i *subnetPtrType) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPtrOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o.ToSubnetPtrOutputWithContext(context.Background())
}

func (o SubnetOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Subnet) *Subnet {
		return &v
	}).(SubnetPtrOutput)
}

func (o SubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Subnet) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetPtrOutput struct{ *pulumi.OutputState }

func (SubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetPtrOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) Elem() SubnetOutput {
	return o.ApplyT(func(v *Subnet) Subnet {
		if v != nil {
			return *v
		}
		var ret Subnet
		return ret
	}).(SubnetOutput)
}

func (o SubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SubnetResponse struct {
	Id string `pulumi:"id"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CaptureDescriptionOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionPtrOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionResponseOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(DestinationOutput{})
	pulumi.RegisterOutputType(DestinationPtrOutput{})
	pulumi.RegisterOutputType(DestinationResponseOutput{})
	pulumi.RegisterOutputType(DestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetPtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
}
