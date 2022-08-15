


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFailoverGroup(ctx *pulumi.Context, args *LookupFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupFailoverGroupResult, error) {
	var rv LookupFailoverGroupResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFailoverGroupArgs struct {
	FailoverGroupName string `pulumi:"failoverGroupName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupFailoverGroupResult struct {
	Databases         []string                               `pulumi:"databases"`
	Id                string                                 `pulumi:"id"`
	Location          string                                 `pulumi:"location"`
	Name              string                                 `pulumi:"name"`
	PartnerServers    []PartnerInfoResponse                  `pulumi:"partnerServers"`
	ReadOnlyEndpoint  *FailoverGroupReadOnlyEndpointResponse `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint FailoverGroupReadWriteEndpointResponse `pulumi:"readWriteEndpoint"`
	ReplicationRole   string                                 `pulumi:"replicationRole"`
	ReplicationState  string                                 `pulumi:"replicationState"`
	Tags              map[string]string                      `pulumi:"tags"`
	Type              string                                 `pulumi:"type"`
}

func LookupFailoverGroupOutput(ctx *pulumi.Context, args LookupFailoverGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFailoverGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFailoverGroupResult, error) {
			args := v.(LookupFailoverGroupArgs)
			r, err := LookupFailoverGroup(ctx, &args, opts...)
			var s LookupFailoverGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFailoverGroupResultOutput)
}

type LookupFailoverGroupOutputArgs struct {
	FailoverGroupName pulumi.StringInput `pulumi:"failoverGroupName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupFailoverGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverGroupArgs)(nil)).Elem()
}


type LookupFailoverGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFailoverGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverGroupResult)(nil)).Elem()
}

func (o LookupFailoverGroupResultOutput) ToLookupFailoverGroupResultOutput() LookupFailoverGroupResultOutput {
	return o
}

func (o LookupFailoverGroupResultOutput) ToLookupFailoverGroupResultOutputWithContext(ctx context.Context) LookupFailoverGroupResultOutput {
	return o
}

func (o LookupFailoverGroupResultOutput) Databases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) []string { return v.Databases }).(pulumi.StringArrayOutput)
}

func (o LookupFailoverGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFailoverGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupFailoverGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFailoverGroupResultOutput) PartnerServers() PartnerInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) []PartnerInfoResponse { return v.PartnerServers }).(PartnerInfoResponseArrayOutput)
}

func (o LookupFailoverGroupResultOutput) ReadOnlyEndpoint() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) *FailoverGroupReadOnlyEndpointResponse { return v.ReadOnlyEndpoint }).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o LookupFailoverGroupResultOutput) ReadWriteEndpoint() FailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) FailoverGroupReadWriteEndpointResponse { return v.ReadWriteEndpoint }).(FailoverGroupReadWriteEndpointResponseOutput)
}

func (o LookupFailoverGroupResultOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

func (o LookupFailoverGroupResultOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.ReplicationState }).(pulumi.StringOutput)
}

func (o LookupFailoverGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFailoverGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFailoverGroupResultOutput{})
}
