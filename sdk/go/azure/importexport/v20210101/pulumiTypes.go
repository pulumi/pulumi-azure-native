


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeliveryPackageInformation struct {
	CarrierName    string   `pulumi:"carrierName"`
	DriveCount     *float64 `pulumi:"driveCount"`
	ShipDate       *string  `pulumi:"shipDate"`
	TrackingNumber string   `pulumi:"trackingNumber"`
}





type DeliveryPackageInformationInput interface {
	pulumi.Input

	ToDeliveryPackageInformationOutput() DeliveryPackageInformationOutput
	ToDeliveryPackageInformationOutputWithContext(context.Context) DeliveryPackageInformationOutput
}

type DeliveryPackageInformationArgs struct {
	CarrierName    pulumi.StringInput     `pulumi:"carrierName"`
	DriveCount     pulumi.Float64PtrInput `pulumi:"driveCount"`
	ShipDate       pulumi.StringPtrInput  `pulumi:"shipDate"`
	TrackingNumber pulumi.StringInput     `pulumi:"trackingNumber"`
}

func (DeliveryPackageInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryPackageInformation)(nil)).Elem()
}

func (i DeliveryPackageInformationArgs) ToDeliveryPackageInformationOutput() DeliveryPackageInformationOutput {
	return i.ToDeliveryPackageInformationOutputWithContext(context.Background())
}

func (i DeliveryPackageInformationArgs) ToDeliveryPackageInformationOutputWithContext(ctx context.Context) DeliveryPackageInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPackageInformationOutput)
}

func (i DeliveryPackageInformationArgs) ToDeliveryPackageInformationPtrOutput() DeliveryPackageInformationPtrOutput {
	return i.ToDeliveryPackageInformationPtrOutputWithContext(context.Background())
}

func (i DeliveryPackageInformationArgs) ToDeliveryPackageInformationPtrOutputWithContext(ctx context.Context) DeliveryPackageInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPackageInformationOutput).ToDeliveryPackageInformationPtrOutputWithContext(ctx)
}









type DeliveryPackageInformationPtrInput interface {
	pulumi.Input

	ToDeliveryPackageInformationPtrOutput() DeliveryPackageInformationPtrOutput
	ToDeliveryPackageInformationPtrOutputWithContext(context.Context) DeliveryPackageInformationPtrOutput
}

type deliveryPackageInformationPtrType DeliveryPackageInformationArgs

func DeliveryPackageInformationPtr(v *DeliveryPackageInformationArgs) DeliveryPackageInformationPtrInput {
	return (*deliveryPackageInformationPtrType)(v)
}

func (*deliveryPackageInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPackageInformation)(nil)).Elem()
}

func (i *deliveryPackageInformationPtrType) ToDeliveryPackageInformationPtrOutput() DeliveryPackageInformationPtrOutput {
	return i.ToDeliveryPackageInformationPtrOutputWithContext(context.Background())
}

func (i *deliveryPackageInformationPtrType) ToDeliveryPackageInformationPtrOutputWithContext(ctx context.Context) DeliveryPackageInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPackageInformationPtrOutput)
}

type DeliveryPackageInformationOutput struct{ *pulumi.OutputState }

func (DeliveryPackageInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryPackageInformation)(nil)).Elem()
}

func (o DeliveryPackageInformationOutput) ToDeliveryPackageInformationOutput() DeliveryPackageInformationOutput {
	return o
}

func (o DeliveryPackageInformationOutput) ToDeliveryPackageInformationOutputWithContext(ctx context.Context) DeliveryPackageInformationOutput {
	return o
}

func (o DeliveryPackageInformationOutput) ToDeliveryPackageInformationPtrOutput() DeliveryPackageInformationPtrOutput {
	return o.ToDeliveryPackageInformationPtrOutputWithContext(context.Background())
}

func (o DeliveryPackageInformationOutput) ToDeliveryPackageInformationPtrOutputWithContext(ctx context.Context) DeliveryPackageInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeliveryPackageInformation) *DeliveryPackageInformation {
		return &v
	}).(DeliveryPackageInformationPtrOutput)
}

func (o DeliveryPackageInformationOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryPackageInformation) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o DeliveryPackageInformationOutput) DriveCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DeliveryPackageInformation) *float64 { return v.DriveCount }).(pulumi.Float64PtrOutput)
}

func (o DeliveryPackageInformationOutput) ShipDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeliveryPackageInformation) *string { return v.ShipDate }).(pulumi.StringPtrOutput)
}

func (o DeliveryPackageInformationOutput) TrackingNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryPackageInformation) string { return v.TrackingNumber }).(pulumi.StringOutput)
}

type DeliveryPackageInformationPtrOutput struct{ *pulumi.OutputState }

func (DeliveryPackageInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPackageInformation)(nil)).Elem()
}

func (o DeliveryPackageInformationPtrOutput) ToDeliveryPackageInformationPtrOutput() DeliveryPackageInformationPtrOutput {
	return o
}

func (o DeliveryPackageInformationPtrOutput) ToDeliveryPackageInformationPtrOutputWithContext(ctx context.Context) DeliveryPackageInformationPtrOutput {
	return o
}

func (o DeliveryPackageInformationPtrOutput) Elem() DeliveryPackageInformationOutput {
	return o.ApplyT(func(v *DeliveryPackageInformation) DeliveryPackageInformation {
		if v != nil {
			return *v
		}
		var ret DeliveryPackageInformation
		return ret
	}).(DeliveryPackageInformationOutput)
}

func (o DeliveryPackageInformationPtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformation) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

func (o DeliveryPackageInformationPtrOutput) DriveCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformation) *float64 {
		if v == nil {
			return nil
		}
		return v.DriveCount
	}).(pulumi.Float64PtrOutput)
}

func (o DeliveryPackageInformationPtrOutput) ShipDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformation) *string {
		if v == nil {
			return nil
		}
		return v.ShipDate
	}).(pulumi.StringPtrOutput)
}

func (o DeliveryPackageInformationPtrOutput) TrackingNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformation) *string {
		if v == nil {
			return nil
		}
		return &v.TrackingNumber
	}).(pulumi.StringPtrOutput)
}

type DeliveryPackageInformationResponse struct {
	CarrierName    string   `pulumi:"carrierName"`
	DriveCount     *float64 `pulumi:"driveCount"`
	ShipDate       *string  `pulumi:"shipDate"`
	TrackingNumber string   `pulumi:"trackingNumber"`
}





type DeliveryPackageInformationResponseInput interface {
	pulumi.Input

	ToDeliveryPackageInformationResponseOutput() DeliveryPackageInformationResponseOutput
	ToDeliveryPackageInformationResponseOutputWithContext(context.Context) DeliveryPackageInformationResponseOutput
}

type DeliveryPackageInformationResponseArgs struct {
	CarrierName    pulumi.StringInput     `pulumi:"carrierName"`
	DriveCount     pulumi.Float64PtrInput `pulumi:"driveCount"`
	ShipDate       pulumi.StringPtrInput  `pulumi:"shipDate"`
	TrackingNumber pulumi.StringInput     `pulumi:"trackingNumber"`
}

func (DeliveryPackageInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryPackageInformationResponse)(nil)).Elem()
}

func (i DeliveryPackageInformationResponseArgs) ToDeliveryPackageInformationResponseOutput() DeliveryPackageInformationResponseOutput {
	return i.ToDeliveryPackageInformationResponseOutputWithContext(context.Background())
}

func (i DeliveryPackageInformationResponseArgs) ToDeliveryPackageInformationResponseOutputWithContext(ctx context.Context) DeliveryPackageInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPackageInformationResponseOutput)
}

func (i DeliveryPackageInformationResponseArgs) ToDeliveryPackageInformationResponsePtrOutput() DeliveryPackageInformationResponsePtrOutput {
	return i.ToDeliveryPackageInformationResponsePtrOutputWithContext(context.Background())
}

func (i DeliveryPackageInformationResponseArgs) ToDeliveryPackageInformationResponsePtrOutputWithContext(ctx context.Context) DeliveryPackageInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPackageInformationResponseOutput).ToDeliveryPackageInformationResponsePtrOutputWithContext(ctx)
}









type DeliveryPackageInformationResponsePtrInput interface {
	pulumi.Input

	ToDeliveryPackageInformationResponsePtrOutput() DeliveryPackageInformationResponsePtrOutput
	ToDeliveryPackageInformationResponsePtrOutputWithContext(context.Context) DeliveryPackageInformationResponsePtrOutput
}

type deliveryPackageInformationResponsePtrType DeliveryPackageInformationResponseArgs

func DeliveryPackageInformationResponsePtr(v *DeliveryPackageInformationResponseArgs) DeliveryPackageInformationResponsePtrInput {
	return (*deliveryPackageInformationResponsePtrType)(v)
}

func (*deliveryPackageInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPackageInformationResponse)(nil)).Elem()
}

func (i *deliveryPackageInformationResponsePtrType) ToDeliveryPackageInformationResponsePtrOutput() DeliveryPackageInformationResponsePtrOutput {
	return i.ToDeliveryPackageInformationResponsePtrOutputWithContext(context.Background())
}

func (i *deliveryPackageInformationResponsePtrType) ToDeliveryPackageInformationResponsePtrOutputWithContext(ctx context.Context) DeliveryPackageInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPackageInformationResponsePtrOutput)
}

type DeliveryPackageInformationResponseOutput struct{ *pulumi.OutputState }

func (DeliveryPackageInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryPackageInformationResponse)(nil)).Elem()
}

func (o DeliveryPackageInformationResponseOutput) ToDeliveryPackageInformationResponseOutput() DeliveryPackageInformationResponseOutput {
	return o
}

func (o DeliveryPackageInformationResponseOutput) ToDeliveryPackageInformationResponseOutputWithContext(ctx context.Context) DeliveryPackageInformationResponseOutput {
	return o
}

func (o DeliveryPackageInformationResponseOutput) ToDeliveryPackageInformationResponsePtrOutput() DeliveryPackageInformationResponsePtrOutput {
	return o.ToDeliveryPackageInformationResponsePtrOutputWithContext(context.Background())
}

func (o DeliveryPackageInformationResponseOutput) ToDeliveryPackageInformationResponsePtrOutputWithContext(ctx context.Context) DeliveryPackageInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeliveryPackageInformationResponse) *DeliveryPackageInformationResponse {
		return &v
	}).(DeliveryPackageInformationResponsePtrOutput)
}

func (o DeliveryPackageInformationResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryPackageInformationResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o DeliveryPackageInformationResponseOutput) DriveCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DeliveryPackageInformationResponse) *float64 { return v.DriveCount }).(pulumi.Float64PtrOutput)
}

func (o DeliveryPackageInformationResponseOutput) ShipDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeliveryPackageInformationResponse) *string { return v.ShipDate }).(pulumi.StringPtrOutput)
}

func (o DeliveryPackageInformationResponseOutput) TrackingNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DeliveryPackageInformationResponse) string { return v.TrackingNumber }).(pulumi.StringOutput)
}

type DeliveryPackageInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (DeliveryPackageInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPackageInformationResponse)(nil)).Elem()
}

func (o DeliveryPackageInformationResponsePtrOutput) ToDeliveryPackageInformationResponsePtrOutput() DeliveryPackageInformationResponsePtrOutput {
	return o
}

func (o DeliveryPackageInformationResponsePtrOutput) ToDeliveryPackageInformationResponsePtrOutputWithContext(ctx context.Context) DeliveryPackageInformationResponsePtrOutput {
	return o
}

func (o DeliveryPackageInformationResponsePtrOutput) Elem() DeliveryPackageInformationResponseOutput {
	return o.ApplyT(func(v *DeliveryPackageInformationResponse) DeliveryPackageInformationResponse {
		if v != nil {
			return *v
		}
		var ret DeliveryPackageInformationResponse
		return ret
	}).(DeliveryPackageInformationResponseOutput)
}

func (o DeliveryPackageInformationResponsePtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

func (o DeliveryPackageInformationResponsePtrOutput) DriveCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DriveCount
	}).(pulumi.Float64PtrOutput)
}

func (o DeliveryPackageInformationResponsePtrOutput) ShipDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ShipDate
	}).(pulumi.StringPtrOutput)
}

func (o DeliveryPackageInformationResponsePtrOutput) TrackingNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPackageInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrackingNumber
	}).(pulumi.StringPtrOutput)
}

type DriveBitLockerKeyResponse struct {
	BitLockerKey *string `pulumi:"bitLockerKey"`
	DriveId      *string `pulumi:"driveId"`
}





type DriveBitLockerKeyResponseInput interface {
	pulumi.Input

	ToDriveBitLockerKeyResponseOutput() DriveBitLockerKeyResponseOutput
	ToDriveBitLockerKeyResponseOutputWithContext(context.Context) DriveBitLockerKeyResponseOutput
}

type DriveBitLockerKeyResponseArgs struct {
	BitLockerKey pulumi.StringPtrInput `pulumi:"bitLockerKey"`
	DriveId      pulumi.StringPtrInput `pulumi:"driveId"`
}

func (DriveBitLockerKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DriveBitLockerKeyResponse)(nil)).Elem()
}

func (i DriveBitLockerKeyResponseArgs) ToDriveBitLockerKeyResponseOutput() DriveBitLockerKeyResponseOutput {
	return i.ToDriveBitLockerKeyResponseOutputWithContext(context.Background())
}

func (i DriveBitLockerKeyResponseArgs) ToDriveBitLockerKeyResponseOutputWithContext(ctx context.Context) DriveBitLockerKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriveBitLockerKeyResponseOutput)
}





type DriveBitLockerKeyResponseArrayInput interface {
	pulumi.Input

	ToDriveBitLockerKeyResponseArrayOutput() DriveBitLockerKeyResponseArrayOutput
	ToDriveBitLockerKeyResponseArrayOutputWithContext(context.Context) DriveBitLockerKeyResponseArrayOutput
}

type DriveBitLockerKeyResponseArray []DriveBitLockerKeyResponseInput

func (DriveBitLockerKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DriveBitLockerKeyResponse)(nil)).Elem()
}

func (i DriveBitLockerKeyResponseArray) ToDriveBitLockerKeyResponseArrayOutput() DriveBitLockerKeyResponseArrayOutput {
	return i.ToDriveBitLockerKeyResponseArrayOutputWithContext(context.Background())
}

func (i DriveBitLockerKeyResponseArray) ToDriveBitLockerKeyResponseArrayOutputWithContext(ctx context.Context) DriveBitLockerKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriveBitLockerKeyResponseArrayOutput)
}

type DriveBitLockerKeyResponseOutput struct{ *pulumi.OutputState }

func (DriveBitLockerKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DriveBitLockerKeyResponse)(nil)).Elem()
}

func (o DriveBitLockerKeyResponseOutput) ToDriveBitLockerKeyResponseOutput() DriveBitLockerKeyResponseOutput {
	return o
}

func (o DriveBitLockerKeyResponseOutput) ToDriveBitLockerKeyResponseOutputWithContext(ctx context.Context) DriveBitLockerKeyResponseOutput {
	return o
}

func (o DriveBitLockerKeyResponseOutput) BitLockerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveBitLockerKeyResponse) *string { return v.BitLockerKey }).(pulumi.StringPtrOutput)
}

func (o DriveBitLockerKeyResponseOutput) DriveId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveBitLockerKeyResponse) *string { return v.DriveId }).(pulumi.StringPtrOutput)
}

type DriveBitLockerKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (DriveBitLockerKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DriveBitLockerKeyResponse)(nil)).Elem()
}

func (o DriveBitLockerKeyResponseArrayOutput) ToDriveBitLockerKeyResponseArrayOutput() DriveBitLockerKeyResponseArrayOutput {
	return o
}

func (o DriveBitLockerKeyResponseArrayOutput) ToDriveBitLockerKeyResponseArrayOutputWithContext(ctx context.Context) DriveBitLockerKeyResponseArrayOutput {
	return o
}

func (o DriveBitLockerKeyResponseArrayOutput) Index(i pulumi.IntInput) DriveBitLockerKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DriveBitLockerKeyResponse {
		return vs[0].([]DriveBitLockerKeyResponse)[vs[1].(int)]
	}).(DriveBitLockerKeyResponseOutput)
}

type DriveStatus struct {
	BitLockerKey    *string  `pulumi:"bitLockerKey"`
	BytesSucceeded  *float64 `pulumi:"bytesSucceeded"`
	CopyStatus      *string  `pulumi:"copyStatus"`
	DriveHeaderHash *string  `pulumi:"driveHeaderHash"`
	DriveId         *string  `pulumi:"driveId"`
	ErrorLogUri     *string  `pulumi:"errorLogUri"`
	ManifestFile    *string  `pulumi:"manifestFile"`
	ManifestHash    *string  `pulumi:"manifestHash"`
	ManifestUri     *string  `pulumi:"manifestUri"`
	PercentComplete *float64 `pulumi:"percentComplete"`
	State           *string  `pulumi:"state"`
	VerboseLogUri   *string  `pulumi:"verboseLogUri"`
}





type DriveStatusInput interface {
	pulumi.Input

	ToDriveStatusOutput() DriveStatusOutput
	ToDriveStatusOutputWithContext(context.Context) DriveStatusOutput
}

type DriveStatusArgs struct {
	BitLockerKey    pulumi.StringPtrInput  `pulumi:"bitLockerKey"`
	BytesSucceeded  pulumi.Float64PtrInput `pulumi:"bytesSucceeded"`
	CopyStatus      pulumi.StringPtrInput  `pulumi:"copyStatus"`
	DriveHeaderHash pulumi.StringPtrInput  `pulumi:"driveHeaderHash"`
	DriveId         pulumi.StringPtrInput  `pulumi:"driveId"`
	ErrorLogUri     pulumi.StringPtrInput  `pulumi:"errorLogUri"`
	ManifestFile    pulumi.StringPtrInput  `pulumi:"manifestFile"`
	ManifestHash    pulumi.StringPtrInput  `pulumi:"manifestHash"`
	ManifestUri     pulumi.StringPtrInput  `pulumi:"manifestUri"`
	PercentComplete pulumi.Float64PtrInput `pulumi:"percentComplete"`
	State           pulumi.StringPtrInput  `pulumi:"state"`
	VerboseLogUri   pulumi.StringPtrInput  `pulumi:"verboseLogUri"`
}

func (DriveStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DriveStatus)(nil)).Elem()
}

func (i DriveStatusArgs) ToDriveStatusOutput() DriveStatusOutput {
	return i.ToDriveStatusOutputWithContext(context.Background())
}

func (i DriveStatusArgs) ToDriveStatusOutputWithContext(ctx context.Context) DriveStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriveStatusOutput)
}





type DriveStatusArrayInput interface {
	pulumi.Input

	ToDriveStatusArrayOutput() DriveStatusArrayOutput
	ToDriveStatusArrayOutputWithContext(context.Context) DriveStatusArrayOutput
}

type DriveStatusArray []DriveStatusInput

func (DriveStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DriveStatus)(nil)).Elem()
}

func (i DriveStatusArray) ToDriveStatusArrayOutput() DriveStatusArrayOutput {
	return i.ToDriveStatusArrayOutputWithContext(context.Background())
}

func (i DriveStatusArray) ToDriveStatusArrayOutputWithContext(ctx context.Context) DriveStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriveStatusArrayOutput)
}

type DriveStatusOutput struct{ *pulumi.OutputState }

func (DriveStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DriveStatus)(nil)).Elem()
}

func (o DriveStatusOutput) ToDriveStatusOutput() DriveStatusOutput {
	return o
}

func (o DriveStatusOutput) ToDriveStatusOutputWithContext(ctx context.Context) DriveStatusOutput {
	return o
}

func (o DriveStatusOutput) BitLockerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.BitLockerKey }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) BytesSucceeded() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DriveStatus) *float64 { return v.BytesSucceeded }).(pulumi.Float64PtrOutput)
}

func (o DriveStatusOutput) CopyStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.CopyStatus }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) DriveHeaderHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.DriveHeaderHash }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) DriveId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.DriveId }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) ErrorLogUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.ErrorLogUri }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) ManifestFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.ManifestFile }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) ManifestHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.ManifestHash }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) ManifestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.ManifestUri }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) PercentComplete() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DriveStatus) *float64 { return v.PercentComplete }).(pulumi.Float64PtrOutput)
}

func (o DriveStatusOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o DriveStatusOutput) VerboseLogUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatus) *string { return v.VerboseLogUri }).(pulumi.StringPtrOutput)
}

type DriveStatusArrayOutput struct{ *pulumi.OutputState }

func (DriveStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DriveStatus)(nil)).Elem()
}

func (o DriveStatusArrayOutput) ToDriveStatusArrayOutput() DriveStatusArrayOutput {
	return o
}

func (o DriveStatusArrayOutput) ToDriveStatusArrayOutputWithContext(ctx context.Context) DriveStatusArrayOutput {
	return o
}

func (o DriveStatusArrayOutput) Index(i pulumi.IntInput) DriveStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DriveStatus {
		return vs[0].([]DriveStatus)[vs[1].(int)]
	}).(DriveStatusOutput)
}

type DriveStatusResponse struct {
	BitLockerKey    *string  `pulumi:"bitLockerKey"`
	BytesSucceeded  *float64 `pulumi:"bytesSucceeded"`
	CopyStatus      *string  `pulumi:"copyStatus"`
	DriveHeaderHash *string  `pulumi:"driveHeaderHash"`
	DriveId         *string  `pulumi:"driveId"`
	ErrorLogUri     *string  `pulumi:"errorLogUri"`
	ManifestFile    *string  `pulumi:"manifestFile"`
	ManifestHash    *string  `pulumi:"manifestHash"`
	ManifestUri     *string  `pulumi:"manifestUri"`
	PercentComplete *float64 `pulumi:"percentComplete"`
	State           *string  `pulumi:"state"`
	VerboseLogUri   *string  `pulumi:"verboseLogUri"`
}





type DriveStatusResponseInput interface {
	pulumi.Input

	ToDriveStatusResponseOutput() DriveStatusResponseOutput
	ToDriveStatusResponseOutputWithContext(context.Context) DriveStatusResponseOutput
}

type DriveStatusResponseArgs struct {
	BitLockerKey    pulumi.StringPtrInput  `pulumi:"bitLockerKey"`
	BytesSucceeded  pulumi.Float64PtrInput `pulumi:"bytesSucceeded"`
	CopyStatus      pulumi.StringPtrInput  `pulumi:"copyStatus"`
	DriveHeaderHash pulumi.StringPtrInput  `pulumi:"driveHeaderHash"`
	DriveId         pulumi.StringPtrInput  `pulumi:"driveId"`
	ErrorLogUri     pulumi.StringPtrInput  `pulumi:"errorLogUri"`
	ManifestFile    pulumi.StringPtrInput  `pulumi:"manifestFile"`
	ManifestHash    pulumi.StringPtrInput  `pulumi:"manifestHash"`
	ManifestUri     pulumi.StringPtrInput  `pulumi:"manifestUri"`
	PercentComplete pulumi.Float64PtrInput `pulumi:"percentComplete"`
	State           pulumi.StringPtrInput  `pulumi:"state"`
	VerboseLogUri   pulumi.StringPtrInput  `pulumi:"verboseLogUri"`
}

func (DriveStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DriveStatusResponse)(nil)).Elem()
}

func (i DriveStatusResponseArgs) ToDriveStatusResponseOutput() DriveStatusResponseOutput {
	return i.ToDriveStatusResponseOutputWithContext(context.Background())
}

func (i DriveStatusResponseArgs) ToDriveStatusResponseOutputWithContext(ctx context.Context) DriveStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriveStatusResponseOutput)
}





type DriveStatusResponseArrayInput interface {
	pulumi.Input

	ToDriveStatusResponseArrayOutput() DriveStatusResponseArrayOutput
	ToDriveStatusResponseArrayOutputWithContext(context.Context) DriveStatusResponseArrayOutput
}

type DriveStatusResponseArray []DriveStatusResponseInput

func (DriveStatusResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DriveStatusResponse)(nil)).Elem()
}

func (i DriveStatusResponseArray) ToDriveStatusResponseArrayOutput() DriveStatusResponseArrayOutput {
	return i.ToDriveStatusResponseArrayOutputWithContext(context.Background())
}

func (i DriveStatusResponseArray) ToDriveStatusResponseArrayOutputWithContext(ctx context.Context) DriveStatusResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DriveStatusResponseArrayOutput)
}

type DriveStatusResponseOutput struct{ *pulumi.OutputState }

func (DriveStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DriveStatusResponse)(nil)).Elem()
}

func (o DriveStatusResponseOutput) ToDriveStatusResponseOutput() DriveStatusResponseOutput {
	return o
}

func (o DriveStatusResponseOutput) ToDriveStatusResponseOutputWithContext(ctx context.Context) DriveStatusResponseOutput {
	return o
}

func (o DriveStatusResponseOutput) BitLockerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.BitLockerKey }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) BytesSucceeded() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *float64 { return v.BytesSucceeded }).(pulumi.Float64PtrOutput)
}

func (o DriveStatusResponseOutput) CopyStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.CopyStatus }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) DriveHeaderHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.DriveHeaderHash }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) DriveId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.DriveId }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) ErrorLogUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.ErrorLogUri }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) ManifestFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.ManifestFile }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) ManifestHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.ManifestHash }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) ManifestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.ManifestUri }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) PercentComplete() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *float64 { return v.PercentComplete }).(pulumi.Float64PtrOutput)
}

func (o DriveStatusResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o DriveStatusResponseOutput) VerboseLogUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DriveStatusResponse) *string { return v.VerboseLogUri }).(pulumi.StringPtrOutput)
}

type DriveStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (DriveStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DriveStatusResponse)(nil)).Elem()
}

func (o DriveStatusResponseArrayOutput) ToDriveStatusResponseArrayOutput() DriveStatusResponseArrayOutput {
	return o
}

func (o DriveStatusResponseArrayOutput) ToDriveStatusResponseArrayOutputWithContext(ctx context.Context) DriveStatusResponseArrayOutput {
	return o
}

func (o DriveStatusResponseArrayOutput) Index(i pulumi.IntInput) DriveStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DriveStatusResponse {
		return vs[0].([]DriveStatusResponse)[vs[1].(int)]
	}).(DriveStatusResponseOutput)
}

type EncryptionKeyDetails struct {
	KekType            *string `pulumi:"kekType"`
	KekUrl             *string `pulumi:"kekUrl"`
	KekVaultResourceID *string `pulumi:"kekVaultResourceID"`
}





type EncryptionKeyDetailsInput interface {
	pulumi.Input

	ToEncryptionKeyDetailsOutput() EncryptionKeyDetailsOutput
	ToEncryptionKeyDetailsOutputWithContext(context.Context) EncryptionKeyDetailsOutput
}

type EncryptionKeyDetailsArgs struct {
	KekType            pulumi.StringPtrInput `pulumi:"kekType"`
	KekUrl             pulumi.StringPtrInput `pulumi:"kekUrl"`
	KekVaultResourceID pulumi.StringPtrInput `pulumi:"kekVaultResourceID"`
}

func (EncryptionKeyDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyDetails)(nil)).Elem()
}

func (i EncryptionKeyDetailsArgs) ToEncryptionKeyDetailsOutput() EncryptionKeyDetailsOutput {
	return i.ToEncryptionKeyDetailsOutputWithContext(context.Background())
}

func (i EncryptionKeyDetailsArgs) ToEncryptionKeyDetailsOutputWithContext(ctx context.Context) EncryptionKeyDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyDetailsOutput)
}

func (i EncryptionKeyDetailsArgs) ToEncryptionKeyDetailsPtrOutput() EncryptionKeyDetailsPtrOutput {
	return i.ToEncryptionKeyDetailsPtrOutputWithContext(context.Background())
}

func (i EncryptionKeyDetailsArgs) ToEncryptionKeyDetailsPtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyDetailsOutput).ToEncryptionKeyDetailsPtrOutputWithContext(ctx)
}









type EncryptionKeyDetailsPtrInput interface {
	pulumi.Input

	ToEncryptionKeyDetailsPtrOutput() EncryptionKeyDetailsPtrOutput
	ToEncryptionKeyDetailsPtrOutputWithContext(context.Context) EncryptionKeyDetailsPtrOutput
}

type encryptionKeyDetailsPtrType EncryptionKeyDetailsArgs

func EncryptionKeyDetailsPtr(v *EncryptionKeyDetailsArgs) EncryptionKeyDetailsPtrInput {
	return (*encryptionKeyDetailsPtrType)(v)
}

func (*encryptionKeyDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyDetails)(nil)).Elem()
}

func (i *encryptionKeyDetailsPtrType) ToEncryptionKeyDetailsPtrOutput() EncryptionKeyDetailsPtrOutput {
	return i.ToEncryptionKeyDetailsPtrOutputWithContext(context.Background())
}

func (i *encryptionKeyDetailsPtrType) ToEncryptionKeyDetailsPtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyDetailsPtrOutput)
}

type EncryptionKeyDetailsOutput struct{ *pulumi.OutputState }

func (EncryptionKeyDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyDetails)(nil)).Elem()
}

func (o EncryptionKeyDetailsOutput) ToEncryptionKeyDetailsOutput() EncryptionKeyDetailsOutput {
	return o
}

func (o EncryptionKeyDetailsOutput) ToEncryptionKeyDetailsOutputWithContext(ctx context.Context) EncryptionKeyDetailsOutput {
	return o
}

func (o EncryptionKeyDetailsOutput) ToEncryptionKeyDetailsPtrOutput() EncryptionKeyDetailsPtrOutput {
	return o.ToEncryptionKeyDetailsPtrOutputWithContext(context.Background())
}

func (o EncryptionKeyDetailsOutput) ToEncryptionKeyDetailsPtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyDetails) *EncryptionKeyDetails {
		return &v
	}).(EncryptionKeyDetailsPtrOutput)
}

func (o EncryptionKeyDetailsOutput) KekType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyDetails) *string { return v.KekType }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyDetails) *string { return v.KekUrl }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyDetails) *string { return v.KekVaultResourceID }).(pulumi.StringPtrOutput)
}

type EncryptionKeyDetailsPtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyDetails)(nil)).Elem()
}

func (o EncryptionKeyDetailsPtrOutput) ToEncryptionKeyDetailsPtrOutput() EncryptionKeyDetailsPtrOutput {
	return o
}

func (o EncryptionKeyDetailsPtrOutput) ToEncryptionKeyDetailsPtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsPtrOutput {
	return o
}

func (o EncryptionKeyDetailsPtrOutput) Elem() EncryptionKeyDetailsOutput {
	return o.ApplyT(func(v *EncryptionKeyDetails) EncryptionKeyDetails {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyDetails
		return ret
	}).(EncryptionKeyDetailsOutput)
}

func (o EncryptionKeyDetailsPtrOutput) KekType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyDetails) *string {
		if v == nil {
			return nil
		}
		return v.KekType
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsPtrOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyDetails) *string {
		if v == nil {
			return nil
		}
		return v.KekUrl
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsPtrOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyDetails) *string {
		if v == nil {
			return nil
		}
		return v.KekVaultResourceID
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyDetailsResponse struct {
	KekType            *string `pulumi:"kekType"`
	KekUrl             *string `pulumi:"kekUrl"`
	KekVaultResourceID *string `pulumi:"kekVaultResourceID"`
}





type EncryptionKeyDetailsResponseInput interface {
	pulumi.Input

	ToEncryptionKeyDetailsResponseOutput() EncryptionKeyDetailsResponseOutput
	ToEncryptionKeyDetailsResponseOutputWithContext(context.Context) EncryptionKeyDetailsResponseOutput
}

type EncryptionKeyDetailsResponseArgs struct {
	KekType            pulumi.StringPtrInput `pulumi:"kekType"`
	KekUrl             pulumi.StringPtrInput `pulumi:"kekUrl"`
	KekVaultResourceID pulumi.StringPtrInput `pulumi:"kekVaultResourceID"`
}

func (EncryptionKeyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyDetailsResponse)(nil)).Elem()
}

func (i EncryptionKeyDetailsResponseArgs) ToEncryptionKeyDetailsResponseOutput() EncryptionKeyDetailsResponseOutput {
	return i.ToEncryptionKeyDetailsResponseOutputWithContext(context.Background())
}

func (i EncryptionKeyDetailsResponseArgs) ToEncryptionKeyDetailsResponseOutputWithContext(ctx context.Context) EncryptionKeyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyDetailsResponseOutput)
}

func (i EncryptionKeyDetailsResponseArgs) ToEncryptionKeyDetailsResponsePtrOutput() EncryptionKeyDetailsResponsePtrOutput {
	return i.ToEncryptionKeyDetailsResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionKeyDetailsResponseArgs) ToEncryptionKeyDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyDetailsResponseOutput).ToEncryptionKeyDetailsResponsePtrOutputWithContext(ctx)
}









type EncryptionKeyDetailsResponsePtrInput interface {
	pulumi.Input

	ToEncryptionKeyDetailsResponsePtrOutput() EncryptionKeyDetailsResponsePtrOutput
	ToEncryptionKeyDetailsResponsePtrOutputWithContext(context.Context) EncryptionKeyDetailsResponsePtrOutput
}

type encryptionKeyDetailsResponsePtrType EncryptionKeyDetailsResponseArgs

func EncryptionKeyDetailsResponsePtr(v *EncryptionKeyDetailsResponseArgs) EncryptionKeyDetailsResponsePtrInput {
	return (*encryptionKeyDetailsResponsePtrType)(v)
}

func (*encryptionKeyDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyDetailsResponse)(nil)).Elem()
}

func (i *encryptionKeyDetailsResponsePtrType) ToEncryptionKeyDetailsResponsePtrOutput() EncryptionKeyDetailsResponsePtrOutput {
	return i.ToEncryptionKeyDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionKeyDetailsResponsePtrType) ToEncryptionKeyDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyDetailsResponsePtrOutput)
}

type EncryptionKeyDetailsResponseOutput struct{ *pulumi.OutputState }

func (EncryptionKeyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyDetailsResponse)(nil)).Elem()
}

func (o EncryptionKeyDetailsResponseOutput) ToEncryptionKeyDetailsResponseOutput() EncryptionKeyDetailsResponseOutput {
	return o
}

func (o EncryptionKeyDetailsResponseOutput) ToEncryptionKeyDetailsResponseOutputWithContext(ctx context.Context) EncryptionKeyDetailsResponseOutput {
	return o
}

func (o EncryptionKeyDetailsResponseOutput) ToEncryptionKeyDetailsResponsePtrOutput() EncryptionKeyDetailsResponsePtrOutput {
	return o.ToEncryptionKeyDetailsResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionKeyDetailsResponseOutput) ToEncryptionKeyDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyDetailsResponse) *EncryptionKeyDetailsResponse {
		return &v
	}).(EncryptionKeyDetailsResponsePtrOutput)
}

func (o EncryptionKeyDetailsResponseOutput) KekType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyDetailsResponse) *string { return v.KekType }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsResponseOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyDetailsResponse) *string { return v.KekUrl }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsResponseOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyDetailsResponse) *string { return v.KekVaultResourceID }).(pulumi.StringPtrOutput)
}

type EncryptionKeyDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyDetailsResponse)(nil)).Elem()
}

func (o EncryptionKeyDetailsResponsePtrOutput) ToEncryptionKeyDetailsResponsePtrOutput() EncryptionKeyDetailsResponsePtrOutput {
	return o
}

func (o EncryptionKeyDetailsResponsePtrOutput) ToEncryptionKeyDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyDetailsResponsePtrOutput {
	return o
}

func (o EncryptionKeyDetailsResponsePtrOutput) Elem() EncryptionKeyDetailsResponseOutput {
	return o.ApplyT(func(v *EncryptionKeyDetailsResponse) EncryptionKeyDetailsResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyDetailsResponse
		return ret
	}).(EncryptionKeyDetailsResponseOutput)
}

func (o EncryptionKeyDetailsResponsePtrOutput) KekType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekType
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsResponsePtrOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekUrl
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyDetailsResponsePtrOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekVaultResourceID
	}).(pulumi.StringPtrOutput)
}

type Export struct {
	BlobListBlobPath *string  `pulumi:"blobListBlobPath"`
	BlobPath         []string `pulumi:"blobPath"`
	BlobPathPrefix   []string `pulumi:"blobPathPrefix"`
}





type ExportInput interface {
	pulumi.Input

	ToExportOutput() ExportOutput
	ToExportOutputWithContext(context.Context) ExportOutput
}

type ExportArgs struct {
	BlobListBlobPath pulumi.StringPtrInput   `pulumi:"blobListBlobPath"`
	BlobPath         pulumi.StringArrayInput `pulumi:"blobPath"`
	BlobPathPrefix   pulumi.StringArrayInput `pulumi:"blobPathPrefix"`
}

func (ExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Export)(nil)).Elem()
}

func (i ExportArgs) ToExportOutput() ExportOutput {
	return i.ToExportOutputWithContext(context.Background())
}

func (i ExportArgs) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportOutput)
}

func (i ExportArgs) ToExportPtrOutput() ExportPtrOutput {
	return i.ToExportPtrOutputWithContext(context.Background())
}

func (i ExportArgs) ToExportPtrOutputWithContext(ctx context.Context) ExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportOutput).ToExportPtrOutputWithContext(ctx)
}









type ExportPtrInput interface {
	pulumi.Input

	ToExportPtrOutput() ExportPtrOutput
	ToExportPtrOutputWithContext(context.Context) ExportPtrOutput
}

type exportPtrType ExportArgs

func ExportPtr(v *ExportArgs) ExportPtrInput {
	return (*exportPtrType)(v)
}

func (*exportPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Export)(nil)).Elem()
}

func (i *exportPtrType) ToExportPtrOutput() ExportPtrOutput {
	return i.ToExportPtrOutputWithContext(context.Background())
}

func (i *exportPtrType) ToExportPtrOutputWithContext(ctx context.Context) ExportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPtrOutput)
}

type ExportOutput struct{ *pulumi.OutputState }

func (ExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Export)(nil)).Elem()
}

func (o ExportOutput) ToExportOutput() ExportOutput {
	return o
}

func (o ExportOutput) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return o
}

func (o ExportOutput) ToExportPtrOutput() ExportPtrOutput {
	return o.ToExportPtrOutputWithContext(context.Background())
}

func (o ExportOutput) ToExportPtrOutputWithContext(ctx context.Context) ExportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Export) *Export {
		return &v
	}).(ExportPtrOutput)
}

func (o ExportOutput) BlobListBlobPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Export) *string { return v.BlobListBlobPath }).(pulumi.StringPtrOutput)
}

func (o ExportOutput) BlobPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Export) []string { return v.BlobPath }).(pulumi.StringArrayOutput)
}

func (o ExportOutput) BlobPathPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Export) []string { return v.BlobPathPrefix }).(pulumi.StringArrayOutput)
}

type ExportPtrOutput struct{ *pulumi.OutputState }

func (ExportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Export)(nil)).Elem()
}

func (o ExportPtrOutput) ToExportPtrOutput() ExportPtrOutput {
	return o
}

func (o ExportPtrOutput) ToExportPtrOutputWithContext(ctx context.Context) ExportPtrOutput {
	return o
}

func (o ExportPtrOutput) Elem() ExportOutput {
	return o.ApplyT(func(v *Export) Export {
		if v != nil {
			return *v
		}
		var ret Export
		return ret
	}).(ExportOutput)
}

func (o ExportPtrOutput) BlobListBlobPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Export) *string {
		if v == nil {
			return nil
		}
		return v.BlobListBlobPath
	}).(pulumi.StringPtrOutput)
}

func (o ExportPtrOutput) BlobPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Export) []string {
		if v == nil {
			return nil
		}
		return v.BlobPath
	}).(pulumi.StringArrayOutput)
}

func (o ExportPtrOutput) BlobPathPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Export) []string {
		if v == nil {
			return nil
		}
		return v.BlobPathPrefix
	}).(pulumi.StringArrayOutput)
}

type ExportResponse struct {
	BlobListBlobPath *string  `pulumi:"blobListBlobPath"`
	BlobPath         []string `pulumi:"blobPath"`
	BlobPathPrefix   []string `pulumi:"blobPathPrefix"`
}





type ExportResponseInput interface {
	pulumi.Input

	ToExportResponseOutput() ExportResponseOutput
	ToExportResponseOutputWithContext(context.Context) ExportResponseOutput
}

type ExportResponseArgs struct {
	BlobListBlobPath pulumi.StringPtrInput   `pulumi:"blobListBlobPath"`
	BlobPath         pulumi.StringArrayInput `pulumi:"blobPath"`
	BlobPathPrefix   pulumi.StringArrayInput `pulumi:"blobPathPrefix"`
}

func (ExportResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportResponse)(nil)).Elem()
}

func (i ExportResponseArgs) ToExportResponseOutput() ExportResponseOutput {
	return i.ToExportResponseOutputWithContext(context.Background())
}

func (i ExportResponseArgs) ToExportResponseOutputWithContext(ctx context.Context) ExportResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportResponseOutput)
}

func (i ExportResponseArgs) ToExportResponsePtrOutput() ExportResponsePtrOutput {
	return i.ToExportResponsePtrOutputWithContext(context.Background())
}

func (i ExportResponseArgs) ToExportResponsePtrOutputWithContext(ctx context.Context) ExportResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportResponseOutput).ToExportResponsePtrOutputWithContext(ctx)
}









type ExportResponsePtrInput interface {
	pulumi.Input

	ToExportResponsePtrOutput() ExportResponsePtrOutput
	ToExportResponsePtrOutputWithContext(context.Context) ExportResponsePtrOutput
}

type exportResponsePtrType ExportResponseArgs

func ExportResponsePtr(v *ExportResponseArgs) ExportResponsePtrInput {
	return (*exportResponsePtrType)(v)
}

func (*exportResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportResponse)(nil)).Elem()
}

func (i *exportResponsePtrType) ToExportResponsePtrOutput() ExportResponsePtrOutput {
	return i.ToExportResponsePtrOutputWithContext(context.Background())
}

func (i *exportResponsePtrType) ToExportResponsePtrOutputWithContext(ctx context.Context) ExportResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportResponsePtrOutput)
}

type ExportResponseOutput struct{ *pulumi.OutputState }

func (ExportResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportResponse)(nil)).Elem()
}

func (o ExportResponseOutput) ToExportResponseOutput() ExportResponseOutput {
	return o
}

func (o ExportResponseOutput) ToExportResponseOutputWithContext(ctx context.Context) ExportResponseOutput {
	return o
}

func (o ExportResponseOutput) ToExportResponsePtrOutput() ExportResponsePtrOutput {
	return o.ToExportResponsePtrOutputWithContext(context.Background())
}

func (o ExportResponseOutput) ToExportResponsePtrOutputWithContext(ctx context.Context) ExportResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportResponse) *ExportResponse {
		return &v
	}).(ExportResponsePtrOutput)
}

func (o ExportResponseOutput) BlobListBlobPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportResponse) *string { return v.BlobListBlobPath }).(pulumi.StringPtrOutput)
}

func (o ExportResponseOutput) BlobPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExportResponse) []string { return v.BlobPath }).(pulumi.StringArrayOutput)
}

func (o ExportResponseOutput) BlobPathPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExportResponse) []string { return v.BlobPathPrefix }).(pulumi.StringArrayOutput)
}

type ExportResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportResponse)(nil)).Elem()
}

func (o ExportResponsePtrOutput) ToExportResponsePtrOutput() ExportResponsePtrOutput {
	return o
}

func (o ExportResponsePtrOutput) ToExportResponsePtrOutputWithContext(ctx context.Context) ExportResponsePtrOutput {
	return o
}

func (o ExportResponsePtrOutput) Elem() ExportResponseOutput {
	return o.ApplyT(func(v *ExportResponse) ExportResponse {
		if v != nil {
			return *v
		}
		var ret ExportResponse
		return ret
	}).(ExportResponseOutput)
}

func (o ExportResponsePtrOutput) BlobListBlobPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobListBlobPath
	}).(pulumi.StringPtrOutput)
}

func (o ExportResponsePtrOutput) BlobPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExportResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobPath
	}).(pulumi.StringArrayOutput)
}

func (o ExportResponsePtrOutput) BlobPathPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExportResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobPathPrefix
	}).(pulumi.StringArrayOutput)
}

type IdentityDetailsResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityDetailsResponseInput interface {
	pulumi.Input

	ToIdentityDetailsResponseOutput() IdentityDetailsResponseOutput
	ToIdentityDetailsResponseOutputWithContext(context.Context) IdentityDetailsResponseOutput
}

type IdentityDetailsResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityDetailsResponse)(nil)).Elem()
}

func (i IdentityDetailsResponseArgs) ToIdentityDetailsResponseOutput() IdentityDetailsResponseOutput {
	return i.ToIdentityDetailsResponseOutputWithContext(context.Background())
}

func (i IdentityDetailsResponseArgs) ToIdentityDetailsResponseOutputWithContext(ctx context.Context) IdentityDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDetailsResponseOutput)
}

func (i IdentityDetailsResponseArgs) ToIdentityDetailsResponsePtrOutput() IdentityDetailsResponsePtrOutput {
	return i.ToIdentityDetailsResponsePtrOutputWithContext(context.Background())
}

func (i IdentityDetailsResponseArgs) ToIdentityDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDetailsResponseOutput).ToIdentityDetailsResponsePtrOutputWithContext(ctx)
}









type IdentityDetailsResponsePtrInput interface {
	pulumi.Input

	ToIdentityDetailsResponsePtrOutput() IdentityDetailsResponsePtrOutput
	ToIdentityDetailsResponsePtrOutputWithContext(context.Context) IdentityDetailsResponsePtrOutput
}

type identityDetailsResponsePtrType IdentityDetailsResponseArgs

func IdentityDetailsResponsePtr(v *IdentityDetailsResponseArgs) IdentityDetailsResponsePtrInput {
	return (*identityDetailsResponsePtrType)(v)
}

func (*identityDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityDetailsResponse)(nil)).Elem()
}

func (i *identityDetailsResponsePtrType) ToIdentityDetailsResponsePtrOutput() IdentityDetailsResponsePtrOutput {
	return i.ToIdentityDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *identityDetailsResponsePtrType) ToIdentityDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDetailsResponsePtrOutput)
}

type IdentityDetailsResponseOutput struct{ *pulumi.OutputState }

func (IdentityDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityDetailsResponse)(nil)).Elem()
}

func (o IdentityDetailsResponseOutput) ToIdentityDetailsResponseOutput() IdentityDetailsResponseOutput {
	return o
}

func (o IdentityDetailsResponseOutput) ToIdentityDetailsResponseOutputWithContext(ctx context.Context) IdentityDetailsResponseOutput {
	return o
}

func (o IdentityDetailsResponseOutput) ToIdentityDetailsResponsePtrOutput() IdentityDetailsResponsePtrOutput {
	return o.ToIdentityDetailsResponsePtrOutputWithContext(context.Background())
}

func (o IdentityDetailsResponseOutput) ToIdentityDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityDetailsResponse) *IdentityDetailsResponse {
		return &v
	}).(IdentityDetailsResponsePtrOutput)
}

func (o IdentityDetailsResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDetailsResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityDetailsResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDetailsResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityDetailsResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityDetailsResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityDetailsResponse)(nil)).Elem()
}

func (o IdentityDetailsResponsePtrOutput) ToIdentityDetailsResponsePtrOutput() IdentityDetailsResponsePtrOutput {
	return o
}

func (o IdentityDetailsResponsePtrOutput) ToIdentityDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityDetailsResponsePtrOutput {
	return o
}

func (o IdentityDetailsResponsePtrOutput) Elem() IdentityDetailsResponseOutput {
	return o.ApplyT(func(v *IdentityDetailsResponse) IdentityDetailsResponse {
		if v != nil {
			return *v
		}
		var ret IdentityDetailsResponse
		return ret
	}).(IdentityDetailsResponseOutput)
}

func (o IdentityDetailsResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDetailsResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDetailsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobDetails struct {
	BackupDriveManifest   *bool                       `pulumi:"backupDriveManifest"`
	CancelRequested       *bool                       `pulumi:"cancelRequested"`
	DeliveryPackage       *DeliveryPackageInformation `pulumi:"deliveryPackage"`
	DiagnosticsPath       *string                     `pulumi:"diagnosticsPath"`
	DriveList             []DriveStatus               `pulumi:"driveList"`
	EncryptionKey         *EncryptionKeyDetails       `pulumi:"encryptionKey"`
	Export                *Export                     `pulumi:"export"`
	IncompleteBlobListUri *string                     `pulumi:"incompleteBlobListUri"`
	JobType               *string                     `pulumi:"jobType"`
	LogLevel              *string                     `pulumi:"logLevel"`
	PercentComplete       *float64                    `pulumi:"percentComplete"`
	ProvisioningState     *string                     `pulumi:"provisioningState"`
	ReturnAddress         *ReturnAddress              `pulumi:"returnAddress"`
	ReturnPackage         *PackageInformation         `pulumi:"returnPackage"`
	ReturnShipping        *ReturnShipping             `pulumi:"returnShipping"`
	ShippingInformation   *ShippingInformation        `pulumi:"shippingInformation"`
	State                 *string                     `pulumi:"state"`
	StorageAccountId      *string                     `pulumi:"storageAccountId"`
}





type JobDetailsInput interface {
	pulumi.Input

	ToJobDetailsOutput() JobDetailsOutput
	ToJobDetailsOutputWithContext(context.Context) JobDetailsOutput
}

type JobDetailsArgs struct {
	BackupDriveManifest   pulumi.BoolPtrInput                `pulumi:"backupDriveManifest"`
	CancelRequested       pulumi.BoolPtrInput                `pulumi:"cancelRequested"`
	DeliveryPackage       DeliveryPackageInformationPtrInput `pulumi:"deliveryPackage"`
	DiagnosticsPath       pulumi.StringPtrInput              `pulumi:"diagnosticsPath"`
	DriveList             DriveStatusArrayInput              `pulumi:"driveList"`
	EncryptionKey         EncryptionKeyDetailsPtrInput       `pulumi:"encryptionKey"`
	Export                ExportPtrInput                     `pulumi:"export"`
	IncompleteBlobListUri pulumi.StringPtrInput              `pulumi:"incompleteBlobListUri"`
	JobType               pulumi.StringPtrInput              `pulumi:"jobType"`
	LogLevel              pulumi.StringPtrInput              `pulumi:"logLevel"`
	PercentComplete       pulumi.Float64PtrInput             `pulumi:"percentComplete"`
	ProvisioningState     pulumi.StringPtrInput              `pulumi:"provisioningState"`
	ReturnAddress         ReturnAddressPtrInput              `pulumi:"returnAddress"`
	ReturnPackage         PackageInformationPtrInput         `pulumi:"returnPackage"`
	ReturnShipping        ReturnShippingPtrInput             `pulumi:"returnShipping"`
	ShippingInformation   ShippingInformationPtrInput        `pulumi:"shippingInformation"`
	State                 pulumi.StringPtrInput              `pulumi:"state"`
	StorageAccountId      pulumi.StringPtrInput              `pulumi:"storageAccountId"`
}

func (JobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDetails)(nil)).Elem()
}

func (i JobDetailsArgs) ToJobDetailsOutput() JobDetailsOutput {
	return i.ToJobDetailsOutputWithContext(context.Background())
}

func (i JobDetailsArgs) ToJobDetailsOutputWithContext(ctx context.Context) JobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDetailsOutput)
}

func (i JobDetailsArgs) ToJobDetailsPtrOutput() JobDetailsPtrOutput {
	return i.ToJobDetailsPtrOutputWithContext(context.Background())
}

func (i JobDetailsArgs) ToJobDetailsPtrOutputWithContext(ctx context.Context) JobDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDetailsOutput).ToJobDetailsPtrOutputWithContext(ctx)
}









type JobDetailsPtrInput interface {
	pulumi.Input

	ToJobDetailsPtrOutput() JobDetailsPtrOutput
	ToJobDetailsPtrOutputWithContext(context.Context) JobDetailsPtrOutput
}

type jobDetailsPtrType JobDetailsArgs

func JobDetailsPtr(v *JobDetailsArgs) JobDetailsPtrInput {
	return (*jobDetailsPtrType)(v)
}

func (*jobDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDetails)(nil)).Elem()
}

func (i *jobDetailsPtrType) ToJobDetailsPtrOutput() JobDetailsPtrOutput {
	return i.ToJobDetailsPtrOutputWithContext(context.Background())
}

func (i *jobDetailsPtrType) ToJobDetailsPtrOutputWithContext(ctx context.Context) JobDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDetailsPtrOutput)
}

type JobDetailsOutput struct{ *pulumi.OutputState }

func (JobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDetails)(nil)).Elem()
}

func (o JobDetailsOutput) ToJobDetailsOutput() JobDetailsOutput {
	return o
}

func (o JobDetailsOutput) ToJobDetailsOutputWithContext(ctx context.Context) JobDetailsOutput {
	return o
}

func (o JobDetailsOutput) ToJobDetailsPtrOutput() JobDetailsPtrOutput {
	return o.ToJobDetailsPtrOutputWithContext(context.Background())
}

func (o JobDetailsOutput) ToJobDetailsPtrOutputWithContext(ctx context.Context) JobDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDetails) *JobDetails {
		return &v
	}).(JobDetailsPtrOutput)
}

func (o JobDetailsOutput) BackupDriveManifest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobDetails) *bool { return v.BackupDriveManifest }).(pulumi.BoolPtrOutput)
}

func (o JobDetailsOutput) CancelRequested() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobDetails) *bool { return v.CancelRequested }).(pulumi.BoolPtrOutput)
}

func (o JobDetailsOutput) DeliveryPackage() DeliveryPackageInformationPtrOutput {
	return o.ApplyT(func(v JobDetails) *DeliveryPackageInformation { return v.DeliveryPackage }).(DeliveryPackageInformationPtrOutput)
}

func (o JobDetailsOutput) DiagnosticsPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.DiagnosticsPath }).(pulumi.StringPtrOutput)
}

func (o JobDetailsOutput) DriveList() DriveStatusArrayOutput {
	return o.ApplyT(func(v JobDetails) []DriveStatus { return v.DriveList }).(DriveStatusArrayOutput)
}

func (o JobDetailsOutput) EncryptionKey() EncryptionKeyDetailsPtrOutput {
	return o.ApplyT(func(v JobDetails) *EncryptionKeyDetails { return v.EncryptionKey }).(EncryptionKeyDetailsPtrOutput)
}

func (o JobDetailsOutput) Export() ExportPtrOutput {
	return o.ApplyT(func(v JobDetails) *Export { return v.Export }).(ExportPtrOutput)
}

func (o JobDetailsOutput) IncompleteBlobListUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.IncompleteBlobListUri }).(pulumi.StringPtrOutput)
}

func (o JobDetailsOutput) JobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.JobType }).(pulumi.StringPtrOutput)
}

func (o JobDetailsOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

func (o JobDetailsOutput) PercentComplete() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v JobDetails) *float64 { return v.PercentComplete }).(pulumi.Float64PtrOutput)
}

func (o JobDetailsOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o JobDetailsOutput) ReturnAddress() ReturnAddressPtrOutput {
	return o.ApplyT(func(v JobDetails) *ReturnAddress { return v.ReturnAddress }).(ReturnAddressPtrOutput)
}

func (o JobDetailsOutput) ReturnPackage() PackageInformationPtrOutput {
	return o.ApplyT(func(v JobDetails) *PackageInformation { return v.ReturnPackage }).(PackageInformationPtrOutput)
}

func (o JobDetailsOutput) ReturnShipping() ReturnShippingPtrOutput {
	return o.ApplyT(func(v JobDetails) *ReturnShipping { return v.ReturnShipping }).(ReturnShippingPtrOutput)
}

func (o JobDetailsOutput) ShippingInformation() ShippingInformationPtrOutput {
	return o.ApplyT(func(v JobDetails) *ShippingInformation { return v.ShippingInformation }).(ShippingInformationPtrOutput)
}

func (o JobDetailsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o JobDetailsOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetails) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

type JobDetailsPtrOutput struct{ *pulumi.OutputState }

func (JobDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDetails)(nil)).Elem()
}

func (o JobDetailsPtrOutput) ToJobDetailsPtrOutput() JobDetailsPtrOutput {
	return o
}

func (o JobDetailsPtrOutput) ToJobDetailsPtrOutputWithContext(ctx context.Context) JobDetailsPtrOutput {
	return o
}

func (o JobDetailsPtrOutput) Elem() JobDetailsOutput {
	return o.ApplyT(func(v *JobDetails) JobDetails {
		if v != nil {
			return *v
		}
		var ret JobDetails
		return ret
	}).(JobDetailsOutput)
}

func (o JobDetailsPtrOutput) BackupDriveManifest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobDetails) *bool {
		if v == nil {
			return nil
		}
		return v.BackupDriveManifest
	}).(pulumi.BoolPtrOutput)
}

func (o JobDetailsPtrOutput) CancelRequested() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobDetails) *bool {
		if v == nil {
			return nil
		}
		return v.CancelRequested
	}).(pulumi.BoolPtrOutput)
}

func (o JobDetailsPtrOutput) DeliveryPackage() DeliveryPackageInformationPtrOutput {
	return o.ApplyT(func(v *JobDetails) *DeliveryPackageInformation {
		if v == nil {
			return nil
		}
		return v.DeliveryPackage
	}).(DeliveryPackageInformationPtrOutput)
}

func (o JobDetailsPtrOutput) DiagnosticsPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.DiagnosticsPath
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsPtrOutput) DriveList() DriveStatusArrayOutput {
	return o.ApplyT(func(v *JobDetails) []DriveStatus {
		if v == nil {
			return nil
		}
		return v.DriveList
	}).(DriveStatusArrayOutput)
}

func (o JobDetailsPtrOutput) EncryptionKey() EncryptionKeyDetailsPtrOutput {
	return o.ApplyT(func(v *JobDetails) *EncryptionKeyDetails {
		if v == nil {
			return nil
		}
		return v.EncryptionKey
	}).(EncryptionKeyDetailsPtrOutput)
}

func (o JobDetailsPtrOutput) Export() ExportPtrOutput {
	return o.ApplyT(func(v *JobDetails) *Export {
		if v == nil {
			return nil
		}
		return v.Export
	}).(ExportPtrOutput)
}

func (o JobDetailsPtrOutput) IncompleteBlobListUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.IncompleteBlobListUri
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsPtrOutput) JobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.JobType
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsPtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsPtrOutput) PercentComplete() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *JobDetails) *float64 {
		if v == nil {
			return nil
		}
		return v.PercentComplete
	}).(pulumi.Float64PtrOutput)
}

func (o JobDetailsPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsPtrOutput) ReturnAddress() ReturnAddressPtrOutput {
	return o.ApplyT(func(v *JobDetails) *ReturnAddress {
		if v == nil {
			return nil
		}
		return v.ReturnAddress
	}).(ReturnAddressPtrOutput)
}

func (o JobDetailsPtrOutput) ReturnPackage() PackageInformationPtrOutput {
	return o.ApplyT(func(v *JobDetails) *PackageInformation {
		if v == nil {
			return nil
		}
		return v.ReturnPackage
	}).(PackageInformationPtrOutput)
}

func (o JobDetailsPtrOutput) ReturnShipping() ReturnShippingPtrOutput {
	return o.ApplyT(func(v *JobDetails) *ReturnShipping {
		if v == nil {
			return nil
		}
		return v.ReturnShipping
	}).(ReturnShippingPtrOutput)
}

func (o JobDetailsPtrOutput) ShippingInformation() ShippingInformationPtrOutput {
	return o.ApplyT(func(v *JobDetails) *ShippingInformation {
		if v == nil {
			return nil
		}
		return v.ShippingInformation
	}).(ShippingInformationPtrOutput)
}

func (o JobDetailsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetails) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type JobDetailsResponse struct {
	BackupDriveManifest   *bool                               `pulumi:"backupDriveManifest"`
	CancelRequested       *bool                               `pulumi:"cancelRequested"`
	DeliveryPackage       *DeliveryPackageInformationResponse `pulumi:"deliveryPackage"`
	DiagnosticsPath       *string                             `pulumi:"diagnosticsPath"`
	DriveList             []DriveStatusResponse               `pulumi:"driveList"`
	EncryptionKey         *EncryptionKeyDetailsResponse       `pulumi:"encryptionKey"`
	Export                *ExportResponse                     `pulumi:"export"`
	IncompleteBlobListUri *string                             `pulumi:"incompleteBlobListUri"`
	JobType               *string                             `pulumi:"jobType"`
	LogLevel              *string                             `pulumi:"logLevel"`
	PercentComplete       *float64                            `pulumi:"percentComplete"`
	ProvisioningState     *string                             `pulumi:"provisioningState"`
	ReturnAddress         *ReturnAddressResponse              `pulumi:"returnAddress"`
	ReturnPackage         *PackageInformationResponse         `pulumi:"returnPackage"`
	ReturnShipping        *ReturnShippingResponse             `pulumi:"returnShipping"`
	ShippingInformation   *ShippingInformationResponse        `pulumi:"shippingInformation"`
	State                 *string                             `pulumi:"state"`
	StorageAccountId      *string                             `pulumi:"storageAccountId"`
}





type JobDetailsResponseInput interface {
	pulumi.Input

	ToJobDetailsResponseOutput() JobDetailsResponseOutput
	ToJobDetailsResponseOutputWithContext(context.Context) JobDetailsResponseOutput
}

type JobDetailsResponseArgs struct {
	BackupDriveManifest   pulumi.BoolPtrInput                        `pulumi:"backupDriveManifest"`
	CancelRequested       pulumi.BoolPtrInput                        `pulumi:"cancelRequested"`
	DeliveryPackage       DeliveryPackageInformationResponsePtrInput `pulumi:"deliveryPackage"`
	DiagnosticsPath       pulumi.StringPtrInput                      `pulumi:"diagnosticsPath"`
	DriveList             DriveStatusResponseArrayInput              `pulumi:"driveList"`
	EncryptionKey         EncryptionKeyDetailsResponsePtrInput       `pulumi:"encryptionKey"`
	Export                ExportResponsePtrInput                     `pulumi:"export"`
	IncompleteBlobListUri pulumi.StringPtrInput                      `pulumi:"incompleteBlobListUri"`
	JobType               pulumi.StringPtrInput                      `pulumi:"jobType"`
	LogLevel              pulumi.StringPtrInput                      `pulumi:"logLevel"`
	PercentComplete       pulumi.Float64PtrInput                     `pulumi:"percentComplete"`
	ProvisioningState     pulumi.StringPtrInput                      `pulumi:"provisioningState"`
	ReturnAddress         ReturnAddressResponsePtrInput              `pulumi:"returnAddress"`
	ReturnPackage         PackageInformationResponsePtrInput         `pulumi:"returnPackage"`
	ReturnShipping        ReturnShippingResponsePtrInput             `pulumi:"returnShipping"`
	ShippingInformation   ShippingInformationResponsePtrInput        `pulumi:"shippingInformation"`
	State                 pulumi.StringPtrInput                      `pulumi:"state"`
	StorageAccountId      pulumi.StringPtrInput                      `pulumi:"storageAccountId"`
}

func (JobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDetailsResponse)(nil)).Elem()
}

func (i JobDetailsResponseArgs) ToJobDetailsResponseOutput() JobDetailsResponseOutput {
	return i.ToJobDetailsResponseOutputWithContext(context.Background())
}

func (i JobDetailsResponseArgs) ToJobDetailsResponseOutputWithContext(ctx context.Context) JobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDetailsResponseOutput)
}

func (i JobDetailsResponseArgs) ToJobDetailsResponsePtrOutput() JobDetailsResponsePtrOutput {
	return i.ToJobDetailsResponsePtrOutputWithContext(context.Background())
}

func (i JobDetailsResponseArgs) ToJobDetailsResponsePtrOutputWithContext(ctx context.Context) JobDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDetailsResponseOutput).ToJobDetailsResponsePtrOutputWithContext(ctx)
}









type JobDetailsResponsePtrInput interface {
	pulumi.Input

	ToJobDetailsResponsePtrOutput() JobDetailsResponsePtrOutput
	ToJobDetailsResponsePtrOutputWithContext(context.Context) JobDetailsResponsePtrOutput
}

type jobDetailsResponsePtrType JobDetailsResponseArgs

func JobDetailsResponsePtr(v *JobDetailsResponseArgs) JobDetailsResponsePtrInput {
	return (*jobDetailsResponsePtrType)(v)
}

func (*jobDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDetailsResponse)(nil)).Elem()
}

func (i *jobDetailsResponsePtrType) ToJobDetailsResponsePtrOutput() JobDetailsResponsePtrOutput {
	return i.ToJobDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *jobDetailsResponsePtrType) ToJobDetailsResponsePtrOutputWithContext(ctx context.Context) JobDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDetailsResponsePtrOutput)
}

type JobDetailsResponseOutput struct{ *pulumi.OutputState }

func (JobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDetailsResponse)(nil)).Elem()
}

func (o JobDetailsResponseOutput) ToJobDetailsResponseOutput() JobDetailsResponseOutput {
	return o
}

func (o JobDetailsResponseOutput) ToJobDetailsResponseOutputWithContext(ctx context.Context) JobDetailsResponseOutput {
	return o
}

func (o JobDetailsResponseOutput) ToJobDetailsResponsePtrOutput() JobDetailsResponsePtrOutput {
	return o.ToJobDetailsResponsePtrOutputWithContext(context.Background())
}

func (o JobDetailsResponseOutput) ToJobDetailsResponsePtrOutputWithContext(ctx context.Context) JobDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDetailsResponse) *JobDetailsResponse {
		return &v
	}).(JobDetailsResponsePtrOutput)
}

func (o JobDetailsResponseOutput) BackupDriveManifest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *bool { return v.BackupDriveManifest }).(pulumi.BoolPtrOutput)
}

func (o JobDetailsResponseOutput) CancelRequested() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *bool { return v.CancelRequested }).(pulumi.BoolPtrOutput)
}

func (o JobDetailsResponseOutput) DeliveryPackage() DeliveryPackageInformationResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *DeliveryPackageInformationResponse { return v.DeliveryPackage }).(DeliveryPackageInformationResponsePtrOutput)
}

func (o JobDetailsResponseOutput) DiagnosticsPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.DiagnosticsPath }).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponseOutput) DriveList() DriveStatusResponseArrayOutput {
	return o.ApplyT(func(v JobDetailsResponse) []DriveStatusResponse { return v.DriveList }).(DriveStatusResponseArrayOutput)
}

func (o JobDetailsResponseOutput) EncryptionKey() EncryptionKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *EncryptionKeyDetailsResponse { return v.EncryptionKey }).(EncryptionKeyDetailsResponsePtrOutput)
}

func (o JobDetailsResponseOutput) Export() ExportResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *ExportResponse { return v.Export }).(ExportResponsePtrOutput)
}

func (o JobDetailsResponseOutput) IncompleteBlobListUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.IncompleteBlobListUri }).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponseOutput) JobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.JobType }).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponseOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponseOutput) PercentComplete() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *float64 { return v.PercentComplete }).(pulumi.Float64PtrOutput)
}

func (o JobDetailsResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponseOutput) ReturnAddress() ReturnAddressResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *ReturnAddressResponse { return v.ReturnAddress }).(ReturnAddressResponsePtrOutput)
}

func (o JobDetailsResponseOutput) ReturnPackage() PackageInformationResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *PackageInformationResponse { return v.ReturnPackage }).(PackageInformationResponsePtrOutput)
}

func (o JobDetailsResponseOutput) ReturnShipping() ReturnShippingResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *ReturnShippingResponse { return v.ReturnShipping }).(ReturnShippingResponsePtrOutput)
}

func (o JobDetailsResponseOutput) ShippingInformation() ShippingInformationResponsePtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *ShippingInformationResponse { return v.ShippingInformation }).(ShippingInformationResponsePtrOutput)
}

func (o JobDetailsResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponseOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDetailsResponse) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

type JobDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (JobDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDetailsResponse)(nil)).Elem()
}

func (o JobDetailsResponsePtrOutput) ToJobDetailsResponsePtrOutput() JobDetailsResponsePtrOutput {
	return o
}

func (o JobDetailsResponsePtrOutput) ToJobDetailsResponsePtrOutputWithContext(ctx context.Context) JobDetailsResponsePtrOutput {
	return o
}

func (o JobDetailsResponsePtrOutput) Elem() JobDetailsResponseOutput {
	return o.ApplyT(func(v *JobDetailsResponse) JobDetailsResponse {
		if v != nil {
			return *v
		}
		var ret JobDetailsResponse
		return ret
	}).(JobDetailsResponseOutput)
}

func (o JobDetailsResponsePtrOutput) BackupDriveManifest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BackupDriveManifest
	}).(pulumi.BoolPtrOutput)
}

func (o JobDetailsResponsePtrOutput) CancelRequested() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CancelRequested
	}).(pulumi.BoolPtrOutput)
}

func (o JobDetailsResponsePtrOutput) DeliveryPackage() DeliveryPackageInformationResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *DeliveryPackageInformationResponse {
		if v == nil {
			return nil
		}
		return v.DeliveryPackage
	}).(DeliveryPackageInformationResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) DiagnosticsPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiagnosticsPath
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponsePtrOutput) DriveList() DriveStatusResponseArrayOutput {
	return o.ApplyT(func(v *JobDetailsResponse) []DriveStatusResponse {
		if v == nil {
			return nil
		}
		return v.DriveList
	}).(DriveStatusResponseArrayOutput)
}

func (o JobDetailsResponsePtrOutput) EncryptionKey() EncryptionKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *EncryptionKeyDetailsResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionKey
	}).(EncryptionKeyDetailsResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) Export() ExportResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *ExportResponse {
		if v == nil {
			return nil
		}
		return v.Export
	}).(ExportResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) IncompleteBlobListUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.IncompleteBlobListUri
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponsePtrOutput) JobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.JobType
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponsePtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponsePtrOutput) PercentComplete() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PercentComplete
	}).(pulumi.Float64PtrOutput)
}

func (o JobDetailsResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponsePtrOutput) ReturnAddress() ReturnAddressResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *ReturnAddressResponse {
		if v == nil {
			return nil
		}
		return v.ReturnAddress
	}).(ReturnAddressResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) ReturnPackage() PackageInformationResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *PackageInformationResponse {
		if v == nil {
			return nil
		}
		return v.ReturnPackage
	}).(PackageInformationResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) ReturnShipping() ReturnShippingResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *ReturnShippingResponse {
		if v == nil {
			return nil
		}
		return v.ReturnShipping
	}).(ReturnShippingResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) ShippingInformation() ShippingInformationResponsePtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *ShippingInformationResponse {
		if v == nil {
			return nil
		}
		return v.ShippingInformation
	}).(ShippingInformationResponsePtrOutput)
}

func (o JobDetailsResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o JobDetailsResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type PackageInformation struct {
	CarrierName    string  `pulumi:"carrierName"`
	DriveCount     float64 `pulumi:"driveCount"`
	ShipDate       string  `pulumi:"shipDate"`
	TrackingNumber string  `pulumi:"trackingNumber"`
}





type PackageInformationInput interface {
	pulumi.Input

	ToPackageInformationOutput() PackageInformationOutput
	ToPackageInformationOutputWithContext(context.Context) PackageInformationOutput
}

type PackageInformationArgs struct {
	CarrierName    pulumi.StringInput  `pulumi:"carrierName"`
	DriveCount     pulumi.Float64Input `pulumi:"driveCount"`
	ShipDate       pulumi.StringInput  `pulumi:"shipDate"`
	TrackingNumber pulumi.StringInput  `pulumi:"trackingNumber"`
}

func (PackageInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageInformation)(nil)).Elem()
}

func (i PackageInformationArgs) ToPackageInformationOutput() PackageInformationOutput {
	return i.ToPackageInformationOutputWithContext(context.Background())
}

func (i PackageInformationArgs) ToPackageInformationOutputWithContext(ctx context.Context) PackageInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageInformationOutput)
}

func (i PackageInformationArgs) ToPackageInformationPtrOutput() PackageInformationPtrOutput {
	return i.ToPackageInformationPtrOutputWithContext(context.Background())
}

func (i PackageInformationArgs) ToPackageInformationPtrOutputWithContext(ctx context.Context) PackageInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageInformationOutput).ToPackageInformationPtrOutputWithContext(ctx)
}









type PackageInformationPtrInput interface {
	pulumi.Input

	ToPackageInformationPtrOutput() PackageInformationPtrOutput
	ToPackageInformationPtrOutputWithContext(context.Context) PackageInformationPtrOutput
}

type packageInformationPtrType PackageInformationArgs

func PackageInformationPtr(v *PackageInformationArgs) PackageInformationPtrInput {
	return (*packageInformationPtrType)(v)
}

func (*packageInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageInformation)(nil)).Elem()
}

func (i *packageInformationPtrType) ToPackageInformationPtrOutput() PackageInformationPtrOutput {
	return i.ToPackageInformationPtrOutputWithContext(context.Background())
}

func (i *packageInformationPtrType) ToPackageInformationPtrOutputWithContext(ctx context.Context) PackageInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageInformationPtrOutput)
}

type PackageInformationOutput struct{ *pulumi.OutputState }

func (PackageInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageInformation)(nil)).Elem()
}

func (o PackageInformationOutput) ToPackageInformationOutput() PackageInformationOutput {
	return o
}

func (o PackageInformationOutput) ToPackageInformationOutputWithContext(ctx context.Context) PackageInformationOutput {
	return o
}

func (o PackageInformationOutput) ToPackageInformationPtrOutput() PackageInformationPtrOutput {
	return o.ToPackageInformationPtrOutputWithContext(context.Background())
}

func (o PackageInformationOutput) ToPackageInformationPtrOutputWithContext(ctx context.Context) PackageInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PackageInformation) *PackageInformation {
		return &v
	}).(PackageInformationPtrOutput)
}

func (o PackageInformationOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v PackageInformation) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o PackageInformationOutput) DriveCount() pulumi.Float64Output {
	return o.ApplyT(func(v PackageInformation) float64 { return v.DriveCount }).(pulumi.Float64Output)
}

func (o PackageInformationOutput) ShipDate() pulumi.StringOutput {
	return o.ApplyT(func(v PackageInformation) string { return v.ShipDate }).(pulumi.StringOutput)
}

func (o PackageInformationOutput) TrackingNumber() pulumi.StringOutput {
	return o.ApplyT(func(v PackageInformation) string { return v.TrackingNumber }).(pulumi.StringOutput)
}

type PackageInformationPtrOutput struct{ *pulumi.OutputState }

func (PackageInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageInformation)(nil)).Elem()
}

func (o PackageInformationPtrOutput) ToPackageInformationPtrOutput() PackageInformationPtrOutput {
	return o
}

func (o PackageInformationPtrOutput) ToPackageInformationPtrOutputWithContext(ctx context.Context) PackageInformationPtrOutput {
	return o
}

func (o PackageInformationPtrOutput) Elem() PackageInformationOutput {
	return o.ApplyT(func(v *PackageInformation) PackageInformation {
		if v != nil {
			return *v
		}
		var ret PackageInformation
		return ret
	}).(PackageInformationOutput)
}

func (o PackageInformationPtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageInformation) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

func (o PackageInformationPtrOutput) DriveCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PackageInformation) *float64 {
		if v == nil {
			return nil
		}
		return &v.DriveCount
	}).(pulumi.Float64PtrOutput)
}

func (o PackageInformationPtrOutput) ShipDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageInformation) *string {
		if v == nil {
			return nil
		}
		return &v.ShipDate
	}).(pulumi.StringPtrOutput)
}

func (o PackageInformationPtrOutput) TrackingNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageInformation) *string {
		if v == nil {
			return nil
		}
		return &v.TrackingNumber
	}).(pulumi.StringPtrOutput)
}

type PackageInformationResponse struct {
	CarrierName    string  `pulumi:"carrierName"`
	DriveCount     float64 `pulumi:"driveCount"`
	ShipDate       string  `pulumi:"shipDate"`
	TrackingNumber string  `pulumi:"trackingNumber"`
}





type PackageInformationResponseInput interface {
	pulumi.Input

	ToPackageInformationResponseOutput() PackageInformationResponseOutput
	ToPackageInformationResponseOutputWithContext(context.Context) PackageInformationResponseOutput
}

type PackageInformationResponseArgs struct {
	CarrierName    pulumi.StringInput  `pulumi:"carrierName"`
	DriveCount     pulumi.Float64Input `pulumi:"driveCount"`
	ShipDate       pulumi.StringInput  `pulumi:"shipDate"`
	TrackingNumber pulumi.StringInput  `pulumi:"trackingNumber"`
}

func (PackageInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageInformationResponse)(nil)).Elem()
}

func (i PackageInformationResponseArgs) ToPackageInformationResponseOutput() PackageInformationResponseOutput {
	return i.ToPackageInformationResponseOutputWithContext(context.Background())
}

func (i PackageInformationResponseArgs) ToPackageInformationResponseOutputWithContext(ctx context.Context) PackageInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageInformationResponseOutput)
}

func (i PackageInformationResponseArgs) ToPackageInformationResponsePtrOutput() PackageInformationResponsePtrOutput {
	return i.ToPackageInformationResponsePtrOutputWithContext(context.Background())
}

func (i PackageInformationResponseArgs) ToPackageInformationResponsePtrOutputWithContext(ctx context.Context) PackageInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageInformationResponseOutput).ToPackageInformationResponsePtrOutputWithContext(ctx)
}









type PackageInformationResponsePtrInput interface {
	pulumi.Input

	ToPackageInformationResponsePtrOutput() PackageInformationResponsePtrOutput
	ToPackageInformationResponsePtrOutputWithContext(context.Context) PackageInformationResponsePtrOutput
}

type packageInformationResponsePtrType PackageInformationResponseArgs

func PackageInformationResponsePtr(v *PackageInformationResponseArgs) PackageInformationResponsePtrInput {
	return (*packageInformationResponsePtrType)(v)
}

func (*packageInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageInformationResponse)(nil)).Elem()
}

func (i *packageInformationResponsePtrType) ToPackageInformationResponsePtrOutput() PackageInformationResponsePtrOutput {
	return i.ToPackageInformationResponsePtrOutputWithContext(context.Background())
}

func (i *packageInformationResponsePtrType) ToPackageInformationResponsePtrOutputWithContext(ctx context.Context) PackageInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageInformationResponsePtrOutput)
}

type PackageInformationResponseOutput struct{ *pulumi.OutputState }

func (PackageInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageInformationResponse)(nil)).Elem()
}

func (o PackageInformationResponseOutput) ToPackageInformationResponseOutput() PackageInformationResponseOutput {
	return o
}

func (o PackageInformationResponseOutput) ToPackageInformationResponseOutputWithContext(ctx context.Context) PackageInformationResponseOutput {
	return o
}

func (o PackageInformationResponseOutput) ToPackageInformationResponsePtrOutput() PackageInformationResponsePtrOutput {
	return o.ToPackageInformationResponsePtrOutputWithContext(context.Background())
}

func (o PackageInformationResponseOutput) ToPackageInformationResponsePtrOutputWithContext(ctx context.Context) PackageInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PackageInformationResponse) *PackageInformationResponse {
		return &v
	}).(PackageInformationResponsePtrOutput)
}

func (o PackageInformationResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v PackageInformationResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o PackageInformationResponseOutput) DriveCount() pulumi.Float64Output {
	return o.ApplyT(func(v PackageInformationResponse) float64 { return v.DriveCount }).(pulumi.Float64Output)
}

func (o PackageInformationResponseOutput) ShipDate() pulumi.StringOutput {
	return o.ApplyT(func(v PackageInformationResponse) string { return v.ShipDate }).(pulumi.StringOutput)
}

func (o PackageInformationResponseOutput) TrackingNumber() pulumi.StringOutput {
	return o.ApplyT(func(v PackageInformationResponse) string { return v.TrackingNumber }).(pulumi.StringOutput)
}

type PackageInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (PackageInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageInformationResponse)(nil)).Elem()
}

func (o PackageInformationResponsePtrOutput) ToPackageInformationResponsePtrOutput() PackageInformationResponsePtrOutput {
	return o
}

func (o PackageInformationResponsePtrOutput) ToPackageInformationResponsePtrOutputWithContext(ctx context.Context) PackageInformationResponsePtrOutput {
	return o
}

func (o PackageInformationResponsePtrOutput) Elem() PackageInformationResponseOutput {
	return o.ApplyT(func(v *PackageInformationResponse) PackageInformationResponse {
		if v != nil {
			return *v
		}
		var ret PackageInformationResponse
		return ret
	}).(PackageInformationResponseOutput)
}

func (o PackageInformationResponsePtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

func (o PackageInformationResponsePtrOutput) DriveCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PackageInformationResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.DriveCount
	}).(pulumi.Float64PtrOutput)
}

func (o PackageInformationResponsePtrOutput) ShipDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShipDate
	}).(pulumi.StringPtrOutput)
}

func (o PackageInformationResponsePtrOutput) TrackingNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrackingNumber
	}).(pulumi.StringPtrOutput)
}

type ReturnAddress struct {
	City            string  `pulumi:"city"`
	CountryOrRegion string  `pulumi:"countryOrRegion"`
	Email           string  `pulumi:"email"`
	Phone           string  `pulumi:"phone"`
	PostalCode      string  `pulumi:"postalCode"`
	RecipientName   string  `pulumi:"recipientName"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
}





type ReturnAddressInput interface {
	pulumi.Input

	ToReturnAddressOutput() ReturnAddressOutput
	ToReturnAddressOutputWithContext(context.Context) ReturnAddressOutput
}

type ReturnAddressArgs struct {
	City            pulumi.StringInput    `pulumi:"city"`
	CountryOrRegion pulumi.StringInput    `pulumi:"countryOrRegion"`
	Email           pulumi.StringInput    `pulumi:"email"`
	Phone           pulumi.StringInput    `pulumi:"phone"`
	PostalCode      pulumi.StringInput    `pulumi:"postalCode"`
	RecipientName   pulumi.StringInput    `pulumi:"recipientName"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringInput    `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
}

func (ReturnAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnAddress)(nil)).Elem()
}

func (i ReturnAddressArgs) ToReturnAddressOutput() ReturnAddressOutput {
	return i.ToReturnAddressOutputWithContext(context.Background())
}

func (i ReturnAddressArgs) ToReturnAddressOutputWithContext(ctx context.Context) ReturnAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnAddressOutput)
}

func (i ReturnAddressArgs) ToReturnAddressPtrOutput() ReturnAddressPtrOutput {
	return i.ToReturnAddressPtrOutputWithContext(context.Background())
}

func (i ReturnAddressArgs) ToReturnAddressPtrOutputWithContext(ctx context.Context) ReturnAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnAddressOutput).ToReturnAddressPtrOutputWithContext(ctx)
}









type ReturnAddressPtrInput interface {
	pulumi.Input

	ToReturnAddressPtrOutput() ReturnAddressPtrOutput
	ToReturnAddressPtrOutputWithContext(context.Context) ReturnAddressPtrOutput
}

type returnAddressPtrType ReturnAddressArgs

func ReturnAddressPtr(v *ReturnAddressArgs) ReturnAddressPtrInput {
	return (*returnAddressPtrType)(v)
}

func (*returnAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnAddress)(nil)).Elem()
}

func (i *returnAddressPtrType) ToReturnAddressPtrOutput() ReturnAddressPtrOutput {
	return i.ToReturnAddressPtrOutputWithContext(context.Background())
}

func (i *returnAddressPtrType) ToReturnAddressPtrOutputWithContext(ctx context.Context) ReturnAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnAddressPtrOutput)
}

type ReturnAddressOutput struct{ *pulumi.OutputState }

func (ReturnAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnAddress)(nil)).Elem()
}

func (o ReturnAddressOutput) ToReturnAddressOutput() ReturnAddressOutput {
	return o
}

func (o ReturnAddressOutput) ToReturnAddressOutputWithContext(ctx context.Context) ReturnAddressOutput {
	return o
}

func (o ReturnAddressOutput) ToReturnAddressPtrOutput() ReturnAddressPtrOutput {
	return o.ToReturnAddressPtrOutputWithContext(context.Background())
}

func (o ReturnAddressOutput) ToReturnAddressPtrOutputWithContext(ctx context.Context) ReturnAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReturnAddress) *ReturnAddress {
		return &v
	}).(ReturnAddressPtrOutput)
}

func (o ReturnAddressOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.City }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) CountryOrRegion() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.CountryOrRegion }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.Email }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.PostalCode }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) RecipientName() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.RecipientName }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReturnAddress) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ReturnAddressOutput) StreetAddress1() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddress) string { return v.StreetAddress1 }).(pulumi.StringOutput)
}

func (o ReturnAddressOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReturnAddress) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

type ReturnAddressPtrOutput struct{ *pulumi.OutputState }

func (ReturnAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnAddress)(nil)).Elem()
}

func (o ReturnAddressPtrOutput) ToReturnAddressPtrOutput() ReturnAddressPtrOutput {
	return o
}

func (o ReturnAddressPtrOutput) ToReturnAddressPtrOutputWithContext(ctx context.Context) ReturnAddressPtrOutput {
	return o
}

func (o ReturnAddressPtrOutput) Elem() ReturnAddressOutput {
	return o.ApplyT(func(v *ReturnAddress) ReturnAddress {
		if v != nil {
			return *v
		}
		var ret ReturnAddress
		return ret
	}).(ReturnAddressOutput)
}

func (o ReturnAddressPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.City
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) RecipientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.RecipientName
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return &v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressPtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddress) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

type ReturnAddressResponse struct {
	City            string  `pulumi:"city"`
	CountryOrRegion string  `pulumi:"countryOrRegion"`
	Email           string  `pulumi:"email"`
	Phone           string  `pulumi:"phone"`
	PostalCode      string  `pulumi:"postalCode"`
	RecipientName   string  `pulumi:"recipientName"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
}





type ReturnAddressResponseInput interface {
	pulumi.Input

	ToReturnAddressResponseOutput() ReturnAddressResponseOutput
	ToReturnAddressResponseOutputWithContext(context.Context) ReturnAddressResponseOutput
}

type ReturnAddressResponseArgs struct {
	City            pulumi.StringInput    `pulumi:"city"`
	CountryOrRegion pulumi.StringInput    `pulumi:"countryOrRegion"`
	Email           pulumi.StringInput    `pulumi:"email"`
	Phone           pulumi.StringInput    `pulumi:"phone"`
	PostalCode      pulumi.StringInput    `pulumi:"postalCode"`
	RecipientName   pulumi.StringInput    `pulumi:"recipientName"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringInput    `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
}

func (ReturnAddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnAddressResponse)(nil)).Elem()
}

func (i ReturnAddressResponseArgs) ToReturnAddressResponseOutput() ReturnAddressResponseOutput {
	return i.ToReturnAddressResponseOutputWithContext(context.Background())
}

func (i ReturnAddressResponseArgs) ToReturnAddressResponseOutputWithContext(ctx context.Context) ReturnAddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnAddressResponseOutput)
}

func (i ReturnAddressResponseArgs) ToReturnAddressResponsePtrOutput() ReturnAddressResponsePtrOutput {
	return i.ToReturnAddressResponsePtrOutputWithContext(context.Background())
}

func (i ReturnAddressResponseArgs) ToReturnAddressResponsePtrOutputWithContext(ctx context.Context) ReturnAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnAddressResponseOutput).ToReturnAddressResponsePtrOutputWithContext(ctx)
}









type ReturnAddressResponsePtrInput interface {
	pulumi.Input

	ToReturnAddressResponsePtrOutput() ReturnAddressResponsePtrOutput
	ToReturnAddressResponsePtrOutputWithContext(context.Context) ReturnAddressResponsePtrOutput
}

type returnAddressResponsePtrType ReturnAddressResponseArgs

func ReturnAddressResponsePtr(v *ReturnAddressResponseArgs) ReturnAddressResponsePtrInput {
	return (*returnAddressResponsePtrType)(v)
}

func (*returnAddressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnAddressResponse)(nil)).Elem()
}

func (i *returnAddressResponsePtrType) ToReturnAddressResponsePtrOutput() ReturnAddressResponsePtrOutput {
	return i.ToReturnAddressResponsePtrOutputWithContext(context.Background())
}

func (i *returnAddressResponsePtrType) ToReturnAddressResponsePtrOutputWithContext(ctx context.Context) ReturnAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnAddressResponsePtrOutput)
}

type ReturnAddressResponseOutput struct{ *pulumi.OutputState }

func (ReturnAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnAddressResponse)(nil)).Elem()
}

func (o ReturnAddressResponseOutput) ToReturnAddressResponseOutput() ReturnAddressResponseOutput {
	return o
}

func (o ReturnAddressResponseOutput) ToReturnAddressResponseOutputWithContext(ctx context.Context) ReturnAddressResponseOutput {
	return o
}

func (o ReturnAddressResponseOutput) ToReturnAddressResponsePtrOutput() ReturnAddressResponsePtrOutput {
	return o.ToReturnAddressResponsePtrOutputWithContext(context.Background())
}

func (o ReturnAddressResponseOutput) ToReturnAddressResponsePtrOutputWithContext(ctx context.Context) ReturnAddressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReturnAddressResponse) *ReturnAddressResponse {
		return &v
	}).(ReturnAddressResponsePtrOutput)
}

func (o ReturnAddressResponseOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.City }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) CountryOrRegion() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.CountryOrRegion }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.Email }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.PostalCode }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) RecipientName() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.RecipientName }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReturnAddressResponse) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponseOutput) StreetAddress1() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnAddressResponse) string { return v.StreetAddress1 }).(pulumi.StringOutput)
}

func (o ReturnAddressResponseOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReturnAddressResponse) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

type ReturnAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (ReturnAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnAddressResponse)(nil)).Elem()
}

func (o ReturnAddressResponsePtrOutput) ToReturnAddressResponsePtrOutput() ReturnAddressResponsePtrOutput {
	return o
}

func (o ReturnAddressResponsePtrOutput) ToReturnAddressResponsePtrOutputWithContext(ctx context.Context) ReturnAddressResponsePtrOutput {
	return o
}

func (o ReturnAddressResponsePtrOutput) Elem() ReturnAddressResponseOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) ReturnAddressResponse {
		if v != nil {
			return *v
		}
		var ret ReturnAddressResponse
		return ret
	}).(ReturnAddressResponseOutput)
}

func (o ReturnAddressResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.City
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) RecipientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RecipientName
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ReturnAddressResponsePtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

type ReturnShipping struct {
	CarrierAccountNumber string `pulumi:"carrierAccountNumber"`
	CarrierName          string `pulumi:"carrierName"`
}





type ReturnShippingInput interface {
	pulumi.Input

	ToReturnShippingOutput() ReturnShippingOutput
	ToReturnShippingOutputWithContext(context.Context) ReturnShippingOutput
}

type ReturnShippingArgs struct {
	CarrierAccountNumber pulumi.StringInput `pulumi:"carrierAccountNumber"`
	CarrierName          pulumi.StringInput `pulumi:"carrierName"`
}

func (ReturnShippingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnShipping)(nil)).Elem()
}

func (i ReturnShippingArgs) ToReturnShippingOutput() ReturnShippingOutput {
	return i.ToReturnShippingOutputWithContext(context.Background())
}

func (i ReturnShippingArgs) ToReturnShippingOutputWithContext(ctx context.Context) ReturnShippingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnShippingOutput)
}

func (i ReturnShippingArgs) ToReturnShippingPtrOutput() ReturnShippingPtrOutput {
	return i.ToReturnShippingPtrOutputWithContext(context.Background())
}

func (i ReturnShippingArgs) ToReturnShippingPtrOutputWithContext(ctx context.Context) ReturnShippingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnShippingOutput).ToReturnShippingPtrOutputWithContext(ctx)
}









type ReturnShippingPtrInput interface {
	pulumi.Input

	ToReturnShippingPtrOutput() ReturnShippingPtrOutput
	ToReturnShippingPtrOutputWithContext(context.Context) ReturnShippingPtrOutput
}

type returnShippingPtrType ReturnShippingArgs

func ReturnShippingPtr(v *ReturnShippingArgs) ReturnShippingPtrInput {
	return (*returnShippingPtrType)(v)
}

func (*returnShippingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnShipping)(nil)).Elem()
}

func (i *returnShippingPtrType) ToReturnShippingPtrOutput() ReturnShippingPtrOutput {
	return i.ToReturnShippingPtrOutputWithContext(context.Background())
}

func (i *returnShippingPtrType) ToReturnShippingPtrOutputWithContext(ctx context.Context) ReturnShippingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnShippingPtrOutput)
}

type ReturnShippingOutput struct{ *pulumi.OutputState }

func (ReturnShippingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnShipping)(nil)).Elem()
}

func (o ReturnShippingOutput) ToReturnShippingOutput() ReturnShippingOutput {
	return o
}

func (o ReturnShippingOutput) ToReturnShippingOutputWithContext(ctx context.Context) ReturnShippingOutput {
	return o
}

func (o ReturnShippingOutput) ToReturnShippingPtrOutput() ReturnShippingPtrOutput {
	return o.ToReturnShippingPtrOutputWithContext(context.Background())
}

func (o ReturnShippingOutput) ToReturnShippingPtrOutputWithContext(ctx context.Context) ReturnShippingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReturnShipping) *ReturnShipping {
		return &v
	}).(ReturnShippingPtrOutput)
}

func (o ReturnShippingOutput) CarrierAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnShipping) string { return v.CarrierAccountNumber }).(pulumi.StringOutput)
}

func (o ReturnShippingOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnShipping) string { return v.CarrierName }).(pulumi.StringOutput)
}

type ReturnShippingPtrOutput struct{ *pulumi.OutputState }

func (ReturnShippingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnShipping)(nil)).Elem()
}

func (o ReturnShippingPtrOutput) ToReturnShippingPtrOutput() ReturnShippingPtrOutput {
	return o
}

func (o ReturnShippingPtrOutput) ToReturnShippingPtrOutputWithContext(ctx context.Context) ReturnShippingPtrOutput {
	return o
}

func (o ReturnShippingPtrOutput) Elem() ReturnShippingOutput {
	return o.ApplyT(func(v *ReturnShipping) ReturnShipping {
		if v != nil {
			return *v
		}
		var ret ReturnShipping
		return ret
	}).(ReturnShippingOutput)
}

func (o ReturnShippingPtrOutput) CarrierAccountNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnShipping) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierAccountNumber
	}).(pulumi.StringPtrOutput)
}

func (o ReturnShippingPtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnShipping) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

type ReturnShippingResponse struct {
	CarrierAccountNumber string `pulumi:"carrierAccountNumber"`
	CarrierName          string `pulumi:"carrierName"`
}





type ReturnShippingResponseInput interface {
	pulumi.Input

	ToReturnShippingResponseOutput() ReturnShippingResponseOutput
	ToReturnShippingResponseOutputWithContext(context.Context) ReturnShippingResponseOutput
}

type ReturnShippingResponseArgs struct {
	CarrierAccountNumber pulumi.StringInput `pulumi:"carrierAccountNumber"`
	CarrierName          pulumi.StringInput `pulumi:"carrierName"`
}

func (ReturnShippingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnShippingResponse)(nil)).Elem()
}

func (i ReturnShippingResponseArgs) ToReturnShippingResponseOutput() ReturnShippingResponseOutput {
	return i.ToReturnShippingResponseOutputWithContext(context.Background())
}

func (i ReturnShippingResponseArgs) ToReturnShippingResponseOutputWithContext(ctx context.Context) ReturnShippingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnShippingResponseOutput)
}

func (i ReturnShippingResponseArgs) ToReturnShippingResponsePtrOutput() ReturnShippingResponsePtrOutput {
	return i.ToReturnShippingResponsePtrOutputWithContext(context.Background())
}

func (i ReturnShippingResponseArgs) ToReturnShippingResponsePtrOutputWithContext(ctx context.Context) ReturnShippingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnShippingResponseOutput).ToReturnShippingResponsePtrOutputWithContext(ctx)
}









type ReturnShippingResponsePtrInput interface {
	pulumi.Input

	ToReturnShippingResponsePtrOutput() ReturnShippingResponsePtrOutput
	ToReturnShippingResponsePtrOutputWithContext(context.Context) ReturnShippingResponsePtrOutput
}

type returnShippingResponsePtrType ReturnShippingResponseArgs

func ReturnShippingResponsePtr(v *ReturnShippingResponseArgs) ReturnShippingResponsePtrInput {
	return (*returnShippingResponsePtrType)(v)
}

func (*returnShippingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnShippingResponse)(nil)).Elem()
}

func (i *returnShippingResponsePtrType) ToReturnShippingResponsePtrOutput() ReturnShippingResponsePtrOutput {
	return i.ToReturnShippingResponsePtrOutputWithContext(context.Background())
}

func (i *returnShippingResponsePtrType) ToReturnShippingResponsePtrOutputWithContext(ctx context.Context) ReturnShippingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReturnShippingResponsePtrOutput)
}

type ReturnShippingResponseOutput struct{ *pulumi.OutputState }

func (ReturnShippingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReturnShippingResponse)(nil)).Elem()
}

func (o ReturnShippingResponseOutput) ToReturnShippingResponseOutput() ReturnShippingResponseOutput {
	return o
}

func (o ReturnShippingResponseOutput) ToReturnShippingResponseOutputWithContext(ctx context.Context) ReturnShippingResponseOutput {
	return o
}

func (o ReturnShippingResponseOutput) ToReturnShippingResponsePtrOutput() ReturnShippingResponsePtrOutput {
	return o.ToReturnShippingResponsePtrOutputWithContext(context.Background())
}

func (o ReturnShippingResponseOutput) ToReturnShippingResponsePtrOutputWithContext(ctx context.Context) ReturnShippingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReturnShippingResponse) *ReturnShippingResponse {
		return &v
	}).(ReturnShippingResponsePtrOutput)
}

func (o ReturnShippingResponseOutput) CarrierAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnShippingResponse) string { return v.CarrierAccountNumber }).(pulumi.StringOutput)
}

func (o ReturnShippingResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v ReturnShippingResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

type ReturnShippingResponsePtrOutput struct{ *pulumi.OutputState }

func (ReturnShippingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReturnShippingResponse)(nil)).Elem()
}

func (o ReturnShippingResponsePtrOutput) ToReturnShippingResponsePtrOutput() ReturnShippingResponsePtrOutput {
	return o
}

func (o ReturnShippingResponsePtrOutput) ToReturnShippingResponsePtrOutputWithContext(ctx context.Context) ReturnShippingResponsePtrOutput {
	return o
}

func (o ReturnShippingResponsePtrOutput) Elem() ReturnShippingResponseOutput {
	return o.ApplyT(func(v *ReturnShippingResponse) ReturnShippingResponse {
		if v != nil {
			return *v
		}
		var ret ReturnShippingResponse
		return ret
	}).(ReturnShippingResponseOutput)
}

func (o ReturnShippingResponsePtrOutput) CarrierAccountNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnShippingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierAccountNumber
	}).(pulumi.StringPtrOutput)
}

func (o ReturnShippingResponsePtrOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReturnShippingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CarrierName
	}).(pulumi.StringPtrOutput)
}

type ShippingInformation struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	Phone           *string `pulumi:"phone"`
	PostalCode      *string `pulumi:"postalCode"`
	RecipientName   *string `pulumi:"recipientName"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  *string `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
}





type ShippingInformationInput interface {
	pulumi.Input

	ToShippingInformationOutput() ShippingInformationOutput
	ToShippingInformationOutputWithContext(context.Context) ShippingInformationOutput
}

type ShippingInformationArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	Phone           pulumi.StringPtrInput `pulumi:"phone"`
	PostalCode      pulumi.StringPtrInput `pulumi:"postalCode"`
	RecipientName   pulumi.StringPtrInput `pulumi:"recipientName"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringPtrInput `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
}

func (ShippingInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingInformation)(nil)).Elem()
}

func (i ShippingInformationArgs) ToShippingInformationOutput() ShippingInformationOutput {
	return i.ToShippingInformationOutputWithContext(context.Background())
}

func (i ShippingInformationArgs) ToShippingInformationOutputWithContext(ctx context.Context) ShippingInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingInformationOutput)
}

func (i ShippingInformationArgs) ToShippingInformationPtrOutput() ShippingInformationPtrOutput {
	return i.ToShippingInformationPtrOutputWithContext(context.Background())
}

func (i ShippingInformationArgs) ToShippingInformationPtrOutputWithContext(ctx context.Context) ShippingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingInformationOutput).ToShippingInformationPtrOutputWithContext(ctx)
}









type ShippingInformationPtrInput interface {
	pulumi.Input

	ToShippingInformationPtrOutput() ShippingInformationPtrOutput
	ToShippingInformationPtrOutputWithContext(context.Context) ShippingInformationPtrOutput
}

type shippingInformationPtrType ShippingInformationArgs

func ShippingInformationPtr(v *ShippingInformationArgs) ShippingInformationPtrInput {
	return (*shippingInformationPtrType)(v)
}

func (*shippingInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingInformation)(nil)).Elem()
}

func (i *shippingInformationPtrType) ToShippingInformationPtrOutput() ShippingInformationPtrOutput {
	return i.ToShippingInformationPtrOutputWithContext(context.Background())
}

func (i *shippingInformationPtrType) ToShippingInformationPtrOutputWithContext(ctx context.Context) ShippingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingInformationPtrOutput)
}

type ShippingInformationOutput struct{ *pulumi.OutputState }

func (ShippingInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingInformation)(nil)).Elem()
}

func (o ShippingInformationOutput) ToShippingInformationOutput() ShippingInformationOutput {
	return o
}

func (o ShippingInformationOutput) ToShippingInformationOutputWithContext(ctx context.Context) ShippingInformationOutput {
	return o
}

func (o ShippingInformationOutput) ToShippingInformationPtrOutput() ShippingInformationPtrOutput {
	return o.ToShippingInformationPtrOutputWithContext(context.Background())
}

func (o ShippingInformationOutput) ToShippingInformationPtrOutputWithContext(ctx context.Context) ShippingInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingInformation) *ShippingInformation {
		return &v
	}).(ShippingInformationPtrOutput)
}

func (o ShippingInformationOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) RecipientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.RecipientName }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.StreetAddress1 }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformation) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

type ShippingInformationPtrOutput struct{ *pulumi.OutputState }

func (ShippingInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingInformation)(nil)).Elem()
}

func (o ShippingInformationPtrOutput) ToShippingInformationPtrOutput() ShippingInformationPtrOutput {
	return o
}

func (o ShippingInformationPtrOutput) ToShippingInformationPtrOutputWithContext(ctx context.Context) ShippingInformationPtrOutput {
	return o
}

func (o ShippingInformationPtrOutput) Elem() ShippingInformationOutput {
	return o.ApplyT(func(v *ShippingInformation) ShippingInformation {
		if v != nil {
			return *v
		}
		var ret ShippingInformation
		return ret
	}).(ShippingInformationOutput)
}

func (o ShippingInformationPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) RecipientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.RecipientName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationPtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformation) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

type ShippingInformationResponse struct {
	AdditionalInformation string  `pulumi:"additionalInformation"`
	City                  *string `pulumi:"city"`
	CountryOrRegion       *string `pulumi:"countryOrRegion"`
	Phone                 *string `pulumi:"phone"`
	PostalCode            *string `pulumi:"postalCode"`
	RecipientName         *string `pulumi:"recipientName"`
	StateOrProvince       *string `pulumi:"stateOrProvince"`
	StreetAddress1        *string `pulumi:"streetAddress1"`
	StreetAddress2        *string `pulumi:"streetAddress2"`
}





type ShippingInformationResponseInput interface {
	pulumi.Input

	ToShippingInformationResponseOutput() ShippingInformationResponseOutput
	ToShippingInformationResponseOutputWithContext(context.Context) ShippingInformationResponseOutput
}

type ShippingInformationResponseArgs struct {
	AdditionalInformation pulumi.StringInput    `pulumi:"additionalInformation"`
	City                  pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion       pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	Phone                 pulumi.StringPtrInput `pulumi:"phone"`
	PostalCode            pulumi.StringPtrInput `pulumi:"postalCode"`
	RecipientName         pulumi.StringPtrInput `pulumi:"recipientName"`
	StateOrProvince       pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1        pulumi.StringPtrInput `pulumi:"streetAddress1"`
	StreetAddress2        pulumi.StringPtrInput `pulumi:"streetAddress2"`
}

func (ShippingInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingInformationResponse)(nil)).Elem()
}

func (i ShippingInformationResponseArgs) ToShippingInformationResponseOutput() ShippingInformationResponseOutput {
	return i.ToShippingInformationResponseOutputWithContext(context.Background())
}

func (i ShippingInformationResponseArgs) ToShippingInformationResponseOutputWithContext(ctx context.Context) ShippingInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingInformationResponseOutput)
}

func (i ShippingInformationResponseArgs) ToShippingInformationResponsePtrOutput() ShippingInformationResponsePtrOutput {
	return i.ToShippingInformationResponsePtrOutputWithContext(context.Background())
}

func (i ShippingInformationResponseArgs) ToShippingInformationResponsePtrOutputWithContext(ctx context.Context) ShippingInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingInformationResponseOutput).ToShippingInformationResponsePtrOutputWithContext(ctx)
}









type ShippingInformationResponsePtrInput interface {
	pulumi.Input

	ToShippingInformationResponsePtrOutput() ShippingInformationResponsePtrOutput
	ToShippingInformationResponsePtrOutputWithContext(context.Context) ShippingInformationResponsePtrOutput
}

type shippingInformationResponsePtrType ShippingInformationResponseArgs

func ShippingInformationResponsePtr(v *ShippingInformationResponseArgs) ShippingInformationResponsePtrInput {
	return (*shippingInformationResponsePtrType)(v)
}

func (*shippingInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingInformationResponse)(nil)).Elem()
}

func (i *shippingInformationResponsePtrType) ToShippingInformationResponsePtrOutput() ShippingInformationResponsePtrOutput {
	return i.ToShippingInformationResponsePtrOutputWithContext(context.Background())
}

func (i *shippingInformationResponsePtrType) ToShippingInformationResponsePtrOutputWithContext(ctx context.Context) ShippingInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingInformationResponsePtrOutput)
}

type ShippingInformationResponseOutput struct{ *pulumi.OutputState }

func (ShippingInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingInformationResponse)(nil)).Elem()
}

func (o ShippingInformationResponseOutput) ToShippingInformationResponseOutput() ShippingInformationResponseOutput {
	return o
}

func (o ShippingInformationResponseOutput) ToShippingInformationResponseOutputWithContext(ctx context.Context) ShippingInformationResponseOutput {
	return o
}

func (o ShippingInformationResponseOutput) ToShippingInformationResponsePtrOutput() ShippingInformationResponsePtrOutput {
	return o.ToShippingInformationResponsePtrOutputWithContext(context.Background())
}

func (o ShippingInformationResponseOutput) ToShippingInformationResponsePtrOutputWithContext(ctx context.Context) ShippingInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingInformationResponse) *ShippingInformationResponse {
		return &v
	}).(ShippingInformationResponsePtrOutput)
}

func (o ShippingInformationResponseOutput) AdditionalInformation() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingInformationResponse) string { return v.AdditionalInformation }).(pulumi.StringOutput)
}

func (o ShippingInformationResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.Phone }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) RecipientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.RecipientName }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.StreetAddress1 }).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponseOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingInformationResponse) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

type ShippingInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ShippingInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingInformationResponse)(nil)).Elem()
}

func (o ShippingInformationResponsePtrOutput) ToShippingInformationResponsePtrOutput() ShippingInformationResponsePtrOutput {
	return o
}

func (o ShippingInformationResponsePtrOutput) ToShippingInformationResponsePtrOutputWithContext(ctx context.Context) ShippingInformationResponsePtrOutput {
	return o
}

func (o ShippingInformationResponsePtrOutput) Elem() ShippingInformationResponseOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) ShippingInformationResponse {
		if v != nil {
			return *v
		}
		var ret ShippingInformationResponse
		return ret
	}).(ShippingInformationResponseOutput)
}

func (o ShippingInformationResponsePtrOutput) AdditionalInformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdditionalInformation
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Phone
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) RecipientName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecipientName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ShippingInformationResponsePtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
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





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeliveryPackageInformationOutput{})
	pulumi.RegisterOutputType(DeliveryPackageInformationPtrOutput{})
	pulumi.RegisterOutputType(DeliveryPackageInformationResponseOutput{})
	pulumi.RegisterOutputType(DeliveryPackageInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(DriveBitLockerKeyResponseOutput{})
	pulumi.RegisterOutputType(DriveBitLockerKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(DriveStatusOutput{})
	pulumi.RegisterOutputType(DriveStatusArrayOutput{})
	pulumi.RegisterOutputType(DriveStatusResponseOutput{})
	pulumi.RegisterOutputType(DriveStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(EncryptionKeyDetailsOutput{})
	pulumi.RegisterOutputType(EncryptionKeyDetailsPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyDetailsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportOutput{})
	pulumi.RegisterOutputType(ExportPtrOutput{})
	pulumi.RegisterOutputType(ExportResponseOutput{})
	pulumi.RegisterOutputType(ExportResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityDetailsResponseOutput{})
	pulumi.RegisterOutputType(IdentityDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(JobDetailsOutput{})
	pulumi.RegisterOutputType(JobDetailsPtrOutput{})
	pulumi.RegisterOutputType(JobDetailsResponseOutput{})
	pulumi.RegisterOutputType(JobDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(PackageInformationOutput{})
	pulumi.RegisterOutputType(PackageInformationPtrOutput{})
	pulumi.RegisterOutputType(PackageInformationResponseOutput{})
	pulumi.RegisterOutputType(PackageInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(ReturnAddressOutput{})
	pulumi.RegisterOutputType(ReturnAddressPtrOutput{})
	pulumi.RegisterOutputType(ReturnAddressResponseOutput{})
	pulumi.RegisterOutputType(ReturnAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ReturnShippingOutput{})
	pulumi.RegisterOutputType(ReturnShippingPtrOutput{})
	pulumi.RegisterOutputType(ReturnShippingResponseOutput{})
	pulumi.RegisterOutputType(ReturnShippingResponsePtrOutput{})
	pulumi.RegisterOutputType(ShippingInformationOutput{})
	pulumi.RegisterOutputType(ShippingInformationPtrOutput{})
	pulumi.RegisterOutputType(ShippingInformationResponseOutput{})
	pulumi.RegisterOutputType(ShippingInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
