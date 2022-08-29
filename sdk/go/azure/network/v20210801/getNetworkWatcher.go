


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkWatcher(ctx *pulumi.Context, args *LookupNetworkWatcherArgs, opts ...pulumi.InvokeOption) (*LookupNetworkWatcherResult, error) {
	var rv LookupNetworkWatcherResult
	err := ctx.Invoke("azure-native:network/v20210801:getNetworkWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkWatcherArgs struct {
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupNetworkWatcherResult struct {
	Etag              string            `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

func LookupNetworkWatcherOutput(ctx *pulumi.Context, args LookupNetworkWatcherOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkWatcherResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkWatcherResult, error) {
			args := v.(LookupNetworkWatcherArgs)
			r, err := LookupNetworkWatcher(ctx, &args, opts...)
			var s LookupNetworkWatcherResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkWatcherResultOutput)
}

type LookupNetworkWatcherOutputArgs struct {
	NetworkWatcherName pulumi.StringInput `pulumi:"networkWatcherName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkWatcherOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkWatcherArgs)(nil)).Elem()
}


type LookupNetworkWatcherResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkWatcherResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkWatcherResult)(nil)).Elem()
}

func (o LookupNetworkWatcherResultOutput) ToLookupNetworkWatcherResultOutput() LookupNetworkWatcherResultOutput {
	return o
}

func (o LookupNetworkWatcherResultOutput) ToLookupNetworkWatcherResultOutputWithContext(ctx context.Context) LookupNetworkWatcherResultOutput {
	return o
}

func (o LookupNetworkWatcherResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkWatcherResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkWatcherResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkWatcherResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkWatcherResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkWatcherResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkWatcherResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkWatcherResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkWatcherResultOutput{})
}
