


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedSubscription(ctx *pulumi.Context, args *LookupLinkedSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupLinkedSubscriptionResult, error) {
	var rv LookupLinkedSubscriptionResult
	err := ctx.Invoke("azure-native:azurestack/v20200601preview:getLinkedSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedSubscriptionArgs struct {
	LinkedSubscriptionName string `pulumi:"linkedSubscriptionName"`
	ResourceGroup          string `pulumi:"resourceGroup"`
}


type LookupLinkedSubscriptionResult struct {
	DeviceConnectionStatus string             `pulumi:"deviceConnectionStatus"`
	DeviceId               string             `pulumi:"deviceId"`
	DeviceLinkState        string             `pulumi:"deviceLinkState"`
	DeviceObjectId         string             `pulumi:"deviceObjectId"`
	Etag                   *string            `pulumi:"etag"`
	Id                     string             `pulumi:"id"`
	Kind                   string             `pulumi:"kind"`
	LastConnectedTime      string             `pulumi:"lastConnectedTime"`
	LinkedSubscriptionId   *string            `pulumi:"linkedSubscriptionId"`
	Location               string             `pulumi:"location"`
	Name                   string             `pulumi:"name"`
	RegistrationResourceId *string            `pulumi:"registrationResourceId"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Tags                   map[string]string  `pulumi:"tags"`
	Type                   string             `pulumi:"type"`
}

func LookupLinkedSubscriptionOutput(ctx *pulumi.Context, args LookupLinkedSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedSubscriptionResult, error) {
			args := v.(LookupLinkedSubscriptionArgs)
			r, err := LookupLinkedSubscription(ctx, &args, opts...)
			var s LookupLinkedSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedSubscriptionResultOutput)
}

type LookupLinkedSubscriptionOutputArgs struct {
	LinkedSubscriptionName pulumi.StringInput `pulumi:"linkedSubscriptionName"`
	ResourceGroup          pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupLinkedSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedSubscriptionArgs)(nil)).Elem()
}


type LookupLinkedSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedSubscriptionResult)(nil)).Elem()
}

func (o LookupLinkedSubscriptionResultOutput) ToLookupLinkedSubscriptionResultOutput() LookupLinkedSubscriptionResultOutput {
	return o
}

func (o LookupLinkedSubscriptionResultOutput) ToLookupLinkedSubscriptionResultOutputWithContext(ctx context.Context) LookupLinkedSubscriptionResultOutput {
	return o
}

func (o LookupLinkedSubscriptionResultOutput) DeviceConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.DeviceConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.DeviceId }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) DeviceLinkState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.DeviceLinkState }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) DeviceObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.DeviceObjectId }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) LastConnectedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.LastConnectedTime }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) LinkedSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) *string { return v.LinkedSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedSubscriptionResultOutput) RegistrationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) *string { return v.RegistrationResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupLinkedSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLinkedSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedSubscriptionResultOutput{})
}
