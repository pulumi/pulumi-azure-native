


package v20210101preview

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

type ConnectionState struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type ConnectionStateInput interface {
	pulumi.Input

	ToConnectionStateOutput() ConnectionStateOutput
	ToConnectionStateOutputWithContext(context.Context) ConnectionStateOutput
}

type ConnectionStateArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (i ConnectionStateArgs) ToConnectionStateOutput() ConnectionStateOutput {
	return i.ToConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput)
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput).ToConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionStatePtrOutput() ConnectionStatePtrOutput
	ToConnectionStatePtrOutputWithContext(context.Context) ConnectionStatePtrOutput
}

type connectionStatePtrType ConnectionStateArgs

func ConnectionStatePtr(v *ConnectionStateArgs) ConnectionStatePtrInput {
	return (*connectionStatePtrType)(v)
}

func (*connectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatePtrOutput)
}

type ConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (o ConnectionStateOutput) ToConnectionStateOutput() ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionState) *ConnectionState {
		return &v
	}).(ConnectionStatePtrOutput)
}

func (o ConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) Elem() ConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionState) ConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionState
		return ret
	}).(ConnectionStateOutput)
}

func (o ConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionStateResponse struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}

type ConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) Elem() ConnectionStateResponseOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) ConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStateResponse
		return ret
	}).(ConnectionStateResponseOutput)
}

func (o ConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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

type Encryption struct {
	KeySource                       *KeySource           `pulumi:"keySource"`
	KeyVaultProperties              []KeyVaultProperties `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool                `pulumi:"requireInfrastructureEncryption"`
}


func (val *Encryption) Defaults() *Encryption {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := KeySource("Microsoft.KeyVault")
		tmp.KeySource = &keySource_
	}
	return &tmp
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource                       KeySourcePtrInput            `pulumi:"keySource"`
	KeyVaultProperties              KeyVaultPropertiesArrayInput `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption pulumi.BoolPtrInput          `pulumi:"requireInfrastructureEncryption"`
}


func (val *EncryptionArgs) Defaults() *EncryptionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		tmp.KeySource = KeySource("Microsoft.KeyVault")
	}
	return &tmp
}
func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() KeySourcePtrOutput {
	return o.ApplyT(func(v Encryption) *KeySource { return v.KeySource }).(KeySourcePtrOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesArrayOutput {
	return o.ApplyT(func(v Encryption) []KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesArrayOutput)
}

func (o EncryptionOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Encryption) *bool { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() KeySourcePtrOutput {
	return o.ApplyT(func(v *Encryption) *KeySource {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(KeySourcePtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesArrayOutput {
	return o.ApplyT(func(v *Encryption) []KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesArrayOutput)
}

func (o EncryptionPtrOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Encryption) *bool {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(pulumi.BoolPtrOutput)
}

type EncryptionResponse struct {
	KeySource                       *string                      `pulumi:"keySource"`
	KeyVaultProperties              []KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool                        `pulumi:"requireInfrastructureEncryption"`
}


func (val *EncryptionResponse) Defaults() *EncryptionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := "Microsoft.KeyVault"
		tmp.KeySource = &keySource_
	}
	return &tmp
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v EncryptionResponse) []KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponseArrayOutput)
}

func (o EncryptionResponseOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *bool { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EncryptionResponse) []KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponseArrayOutput)
}

func (o EncryptionResponsePtrOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(pulumi.BoolPtrOutput)
}

type Identity struct {
	Type                   *ManagedServiceIdentityType `pulumi:"type"`
	UserAssignedIdentities map[string]interface{}      `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ManagedServiceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput                    `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ManagedServiceIdentityType { return v.Type }).(ManagedServiceIdentityTypePtrOutput)
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

func (o IdentityPtrOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ManagedServiceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ManagedServiceIdentityTypePtrOutput)
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
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   *string                                           `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
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
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
}

type KeyVaultProperties struct {
	Identity    *UserAssignedIdentityProperties `pulumi:"identity"`
	KeyName     *string                         `pulumi:"keyName"`
	KeyVaultUri *string                         `pulumi:"keyVaultUri"`
	KeyVersion  *string                         `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	Identity    UserAssignedIdentityPropertiesPtrInput `pulumi:"identity"`
	KeyName     pulumi.StringPtrInput                  `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput                  `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput                  `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}





type KeyVaultPropertiesArrayInput interface {
	pulumi.Input

	ToKeyVaultPropertiesArrayOutput() KeyVaultPropertiesArrayOutput
	ToKeyVaultPropertiesArrayOutputWithContext(context.Context) KeyVaultPropertiesArrayOutput
}

type KeyVaultPropertiesArray []KeyVaultPropertiesInput

func (KeyVaultPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArray) ToKeyVaultPropertiesArrayOutput() KeyVaultPropertiesArrayOutput {
	return i.ToKeyVaultPropertiesArrayOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArray) ToKeyVaultPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesArrayOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) Identity() UserAssignedIdentityPropertiesPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *UserAssignedIdentityProperties { return v.Identity }).(UserAssignedIdentityPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesArrayOutput) ToKeyVaultPropertiesArrayOutput() KeyVaultPropertiesArrayOutput {
	return o
}

func (o KeyVaultPropertiesArrayOutput) ToKeyVaultPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultPropertiesArrayOutput {
	return o
}

func (o KeyVaultPropertiesArrayOutput) Index(i pulumi.IntInput) KeyVaultPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultProperties {
		return vs[0].([]KeyVaultProperties)[vs[1].(int)]
	}).(KeyVaultPropertiesOutput)
}

type KeyVaultPropertiesResponse struct {
	Identity    *UserAssignedIdentityPropertiesResponse `pulumi:"identity"`
	KeyName     *string                                 `pulumi:"keyName"`
	KeyVaultUri *string                                 `pulumi:"keyVaultUri"`
	KeyVersion  *string                                 `pulumi:"keyVersion"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) Identity() UserAssignedIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *UserAssignedIdentityPropertiesResponse { return v.Identity }).(UserAssignedIdentityPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseArrayOutput) ToKeyVaultPropertiesResponseArrayOutput() KeyVaultPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultPropertiesResponseArrayOutput) ToKeyVaultPropertiesResponseArrayOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultPropertiesResponseArrayOutput) Index(i pulumi.IntInput) KeyVaultPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultPropertiesResponse {
		return vs[0].([]KeyVaultPropertiesResponse)[vs[1].(int)]
	}).(KeyVaultPropertiesResponseOutput)
}

type NWRuleSetIpRules struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
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

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionType struct {
	PrivateEndpoint                   *PrivateEndpoint `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string          `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput   `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() ConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *ConnectionState { return v.PrivateLinkServiceConnectionState }).(ConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                   `pulumi:"id"`
	Name                              string                   `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                  `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse       `pulumi:"systemData"`
	Type                              string                   `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *ConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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
	Id *string `pulumi:"id"`
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

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
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
		return v.Id
	}).(pulumi.StringPtrOutput)
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
	Id *string `pulumi:"id"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
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

func (o SubnetOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.Id }).(pulumi.StringPtrOutput)
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
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubnetResponse struct {
	Id *string `pulumi:"id"`
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

func (o SubnetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
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
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityProperties struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type UserAssignedIdentityPropertiesInput interface {
	pulumi.Input

	ToUserAssignedIdentityPropertiesOutput() UserAssignedIdentityPropertiesOutput
	ToUserAssignedIdentityPropertiesOutputWithContext(context.Context) UserAssignedIdentityPropertiesOutput
}

type UserAssignedIdentityPropertiesArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (UserAssignedIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityProperties)(nil)).Elem()
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesOutput() UserAssignedIdentityPropertiesOutput {
	return i.ToUserAssignedIdentityPropertiesOutputWithContext(context.Background())
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertiesOutput)
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return i.ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i UserAssignedIdentityPropertiesArgs) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertiesOutput).ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx)
}









type UserAssignedIdentityPropertiesPtrInput interface {
	pulumi.Input

	ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput
	ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Context) UserAssignedIdentityPropertiesPtrOutput
}

type userAssignedIdentityPropertiesPtrType UserAssignedIdentityPropertiesArgs

func UserAssignedIdentityPropertiesPtr(v *UserAssignedIdentityPropertiesArgs) UserAssignedIdentityPropertiesPtrInput {
	return (*userAssignedIdentityPropertiesPtrType)(v)
}

func (*userAssignedIdentityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentityProperties)(nil)).Elem()
}

func (i *userAssignedIdentityPropertiesPtrType) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return i.ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *userAssignedIdentityPropertiesPtrType) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityPropertiesPtrOutput)
}

type UserAssignedIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityProperties)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesOutput() UserAssignedIdentityPropertiesOutput {
	return o
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesOutput {
	return o
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return o.ToUserAssignedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o UserAssignedIdentityPropertiesOutput) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAssignedIdentityProperties) *UserAssignedIdentityProperties {
		return &v
	}).(UserAssignedIdentityPropertiesPtrOutput)
}

func (o UserAssignedIdentityPropertiesOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityProperties) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentityProperties)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesPtrOutput) ToUserAssignedIdentityPropertiesPtrOutput() UserAssignedIdentityPropertiesPtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesPtrOutput) ToUserAssignedIdentityPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesPtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesPtrOutput) Elem() UserAssignedIdentityPropertiesOutput {
	return o.ApplyT(func(v *UserAssignedIdentityProperties) UserAssignedIdentityProperties {
		if v != nil {
			return *v
		}
		var ret UserAssignedIdentityProperties
		return ret
	}).(UserAssignedIdentityPropertiesOutput)
}

func (o UserAssignedIdentityPropertiesPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertiesResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type UserAssignedIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesResponseOutput) ToUserAssignedIdentityPropertiesResponseOutput() UserAssignedIdentityPropertiesResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponseOutput) ToUserAssignedIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityPropertiesResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) ToUserAssignedIdentityPropertiesResponsePtrOutput() UserAssignedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) ToUserAssignedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) UserAssignedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) Elem() UserAssignedIdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *UserAssignedIdentityPropertiesResponse) UserAssignedIdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserAssignedIdentityPropertiesResponse
		return ret
	}).(UserAssignedIdentityPropertiesResponseOutput)
}

func (o UserAssignedIdentityPropertiesResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CaptureDescriptionOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionPtrOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionResponseOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(DestinationOutput{})
	pulumi.RegisterOutputType(DestinationPtrOutput{})
	pulumi.RegisterOutputType(DestinationResponseOutput{})
	pulumi.RegisterOutputType(DestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetPtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertiesResponsePtrOutput{})
}
