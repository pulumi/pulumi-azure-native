


package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerCommunicationLink(ctx *pulumi.Context, args *LookupServerCommunicationLinkArgs, opts ...pulumi.InvokeOption) (*LookupServerCommunicationLinkResult, error) {
	var rv LookupServerCommunicationLinkResult
	err := ctx.Invoke("azure-native:sql:getServerCommunicationLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerCommunicationLinkArgs struct {
	CommunicationLinkName string `pulumi:"communicationLinkName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ServerName            string `pulumi:"serverName"`
}


type LookupServerCommunicationLinkResult struct {
	Id            string `pulumi:"id"`
	Kind          string `pulumi:"kind"`
	Location      string `pulumi:"location"`
	Name          string `pulumi:"name"`
	PartnerServer string `pulumi:"partnerServer"`
	State         string `pulumi:"state"`
	Type          string `pulumi:"type"`
}

func LookupServerCommunicationLinkOutput(ctx *pulumi.Context, args LookupServerCommunicationLinkOutputArgs, opts ...pulumi.InvokeOption) LookupServerCommunicationLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerCommunicationLinkResult, error) {
			args := v.(LookupServerCommunicationLinkArgs)
			r, err := LookupServerCommunicationLink(ctx, &args, opts...)
			var s LookupServerCommunicationLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerCommunicationLinkResultOutput)
}

type LookupServerCommunicationLinkOutputArgs struct {
	CommunicationLinkName pulumi.StringInput `pulumi:"communicationLinkName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName            pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerCommunicationLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCommunicationLinkArgs)(nil)).Elem()
}


type LookupServerCommunicationLinkResultOutput struct{ *pulumi.OutputState }

func (LookupServerCommunicationLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCommunicationLinkResult)(nil)).Elem()
}

func (o LookupServerCommunicationLinkResultOutput) ToLookupServerCommunicationLinkResultOutput() LookupServerCommunicationLinkResultOutput {
	return o
}

func (o LookupServerCommunicationLinkResultOutput) ToLookupServerCommunicationLinkResultOutputWithContext(ctx context.Context) LookupServerCommunicationLinkResultOutput {
	return o
}

func (o LookupServerCommunicationLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerCommunicationLinkResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupServerCommunicationLinkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerCommunicationLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerCommunicationLinkResultOutput) PartnerServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.PartnerServer }).(pulumi.StringOutput)
}

func (o LookupServerCommunicationLinkResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupServerCommunicationLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerCommunicationLinkResultOutput{})
}
