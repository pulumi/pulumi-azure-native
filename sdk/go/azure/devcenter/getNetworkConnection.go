


package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkConnection(ctx *pulumi.Context, args *LookupNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkConnectionResult, error) {
	var rv LookupNetworkConnectionResult
	err := ctx.Invoke("azure-native:devcenter:getNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkConnectionArgs struct {
	NetworkConnectionName string `pulumi:"networkConnectionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNetworkConnectionResult struct {
	DomainJoinType              string             `pulumi:"domainJoinType"`
	DomainName                  *string            `pulumi:"domainName"`
	DomainPassword              *string            `pulumi:"domainPassword"`
	DomainUsername              *string            `pulumi:"domainUsername"`
	HealthCheckStatus           string             `pulumi:"healthCheckStatus"`
	Id                          string             `pulumi:"id"`
	Location                    string             `pulumi:"location"`
	Name                        string             `pulumi:"name"`
	NetworkingResourceGroupName *string            `pulumi:"networkingResourceGroupName"`
	OrganizationUnit            *string            `pulumi:"organizationUnit"`
	ProvisioningState           string             `pulumi:"provisioningState"`
	SubnetId                    string             `pulumi:"subnetId"`
	SystemData                  SystemDataResponse `pulumi:"systemData"`
	Tags                        map[string]string  `pulumi:"tags"`
	Type                        string             `pulumi:"type"`
}

func LookupNetworkConnectionOutput(ctx *pulumi.Context, args LookupNetworkConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkConnectionResult, error) {
			args := v.(LookupNetworkConnectionArgs)
			r, err := LookupNetworkConnection(ctx, &args, opts...)
			var s LookupNetworkConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkConnectionResultOutput)
}

type LookupNetworkConnectionOutputArgs struct {
	NetworkConnectionName pulumi.StringInput `pulumi:"networkConnectionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkConnectionArgs)(nil)).Elem()
}


type LookupNetworkConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkConnectionResult)(nil)).Elem()
}

func (o LookupNetworkConnectionResultOutput) ToLookupNetworkConnectionResultOutput() LookupNetworkConnectionResultOutput {
	return o
}

func (o LookupNetworkConnectionResultOutput) ToLookupNetworkConnectionResultOutputWithContext(ctx context.Context) LookupNetworkConnectionResultOutput {
	return o
}

func (o LookupNetworkConnectionResultOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.DomainJoinType }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkConnectionResultOutput) DomainPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.DomainPassword }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkConnectionResultOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.DomainUsername }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkConnectionResultOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) NetworkingResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.NetworkingResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkConnectionResultOutput) OrganizationUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.OrganizationUnit }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupNetworkConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNetworkConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkConnectionResultOutput{})
}
