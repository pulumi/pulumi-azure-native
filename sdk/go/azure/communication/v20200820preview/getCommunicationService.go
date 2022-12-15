


package v20200820preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommunicationService(ctx *pulumi.Context, args *LookupCommunicationServiceArgs, opts ...pulumi.InvokeOption) (*LookupCommunicationServiceResult, error) {
	var rv LookupCommunicationServiceResult
	err := ctx.Invoke("azure-native:communication/v20200820preview:getCommunicationService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommunicationServiceArgs struct {
	CommunicationServiceName string `pulumi:"communicationServiceName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupCommunicationServiceResult struct {
	DataLocation        string            `pulumi:"dataLocation"`
	HostName            string            `pulumi:"hostName"`
	Id                  string            `pulumi:"id"`
	ImmutableResourceId string            `pulumi:"immutableResourceId"`
	Location            *string           `pulumi:"location"`
	Name                string            `pulumi:"name"`
	NotificationHubId   string            `pulumi:"notificationHubId"`
	ProvisioningState   string            `pulumi:"provisioningState"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
	Version             string            `pulumi:"version"`
}

func LookupCommunicationServiceOutput(ctx *pulumi.Context, args LookupCommunicationServiceOutputArgs, opts ...pulumi.InvokeOption) LookupCommunicationServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommunicationServiceResult, error) {
			args := v.(LookupCommunicationServiceArgs)
			r, err := LookupCommunicationService(ctx, &args, opts...)
			var s LookupCommunicationServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommunicationServiceResultOutput)
}

type LookupCommunicationServiceOutputArgs struct {
	CommunicationServiceName pulumi.StringInput `pulumi:"communicationServiceName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommunicationServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommunicationServiceArgs)(nil)).Elem()
}


type LookupCommunicationServiceResultOutput struct{ *pulumi.OutputState }

func (LookupCommunicationServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommunicationServiceResult)(nil)).Elem()
}

func (o LookupCommunicationServiceResultOutput) ToLookupCommunicationServiceResultOutput() LookupCommunicationServiceResultOutput {
	return o
}

func (o LookupCommunicationServiceResultOutput) ToLookupCommunicationServiceResultOutputWithContext(ctx context.Context) LookupCommunicationServiceResultOutput {
	return o
}

func (o LookupCommunicationServiceResultOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.DataLocation }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) ImmutableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.ImmutableResourceId }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCommunicationServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) NotificationHubId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.NotificationHubId }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCommunicationServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCommunicationServiceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationServiceResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommunicationServiceResultOutput{})
}
