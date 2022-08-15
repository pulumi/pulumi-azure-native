


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerDnsAlias(ctx *pulumi.Context, args *LookupServerDnsAliasArgs, opts ...pulumi.InvokeOption) (*LookupServerDnsAliasResult, error) {
	var rv LookupServerDnsAliasResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getServerDnsAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerDnsAliasArgs struct {
	DnsAliasName      string `pulumi:"dnsAliasName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerDnsAliasResult struct {
	AzureDnsRecord string `pulumi:"azureDnsRecord"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	Type           string `pulumi:"type"`
}

func LookupServerDnsAliasOutput(ctx *pulumi.Context, args LookupServerDnsAliasOutputArgs, opts ...pulumi.InvokeOption) LookupServerDnsAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerDnsAliasResult, error) {
			args := v.(LookupServerDnsAliasArgs)
			r, err := LookupServerDnsAlias(ctx, &args, opts...)
			var s LookupServerDnsAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerDnsAliasResultOutput)
}

type LookupServerDnsAliasOutputArgs struct {
	DnsAliasName      pulumi.StringInput `pulumi:"dnsAliasName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerDnsAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDnsAliasArgs)(nil)).Elem()
}


type LookupServerDnsAliasResultOutput struct{ *pulumi.OutputState }

func (LookupServerDnsAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDnsAliasResult)(nil)).Elem()
}

func (o LookupServerDnsAliasResultOutput) ToLookupServerDnsAliasResultOutput() LookupServerDnsAliasResultOutput {
	return o
}

func (o LookupServerDnsAliasResultOutput) ToLookupServerDnsAliasResultOutputWithContext(ctx context.Context) LookupServerDnsAliasResultOutput {
	return o
}

func (o LookupServerDnsAliasResultOutput) AzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.AzureDnsRecord }).(pulumi.StringOutput)
}

func (o LookupServerDnsAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerDnsAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerDnsAliasResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerDnsAliasResultOutput{})
}
