


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("azure-native:engagementfabric/v20180901preview:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	AccountName       string `pulumi:"accountName"`
	ChannelName       string `pulumi:"channelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupChannelResult struct {
	ChannelFunctions []string          `pulumi:"channelFunctions"`
	ChannelType      string            `pulumi:"channelType"`
	Credentials      map[string]string `pulumi:"credentials"`
	Id               string            `pulumi:"id"`
	Name             string            `pulumi:"name"`
	Type             string            `pulumi:"type"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ChannelName       pulumi.StringInput `pulumi:"channelName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}


type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ChannelFunctions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []string { return v.ChannelFunctions }).(pulumi.StringArrayOutput)
}

func (o LookupChannelResultOutput) ChannelType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.ChannelType }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupChannelResult) map[string]string { return v.Credentials }).(pulumi.StringMapOutput)
}

func (o LookupChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}
