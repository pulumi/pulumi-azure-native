


package hybridcontainerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProvisionedCluster(ctx *pulumi.Context, args *LookupProvisionedClusterArgs, opts ...pulumi.InvokeOption) (*LookupProvisionedClusterResult, error) {
	var rv LookupProvisionedClusterResult
	err := ctx.Invoke("azure-native:hybridcontainerservice:getProvisionedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupProvisionedClusterArgs struct {
	ProvisionedClustersName string `pulumi:"provisionedClustersName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupProvisionedClusterResult struct {
	ExtendedLocation *ProvisionedClustersResponseResponseExtendedLocation `pulumi:"extendedLocation"`
	Id               string                                               `pulumi:"id"`
	Identity         *ProvisionedClusterIdentityResponse                  `pulumi:"identity"`
	Location         string                                               `pulumi:"location"`
	Name             string                                               `pulumi:"name"`
	Properties       ProvisionedClustersResponsePropertiesResponse        `pulumi:"properties"`
	SystemData       SystemDataResponse                                   `pulumi:"systemData"`
	Tags             map[string]string                                    `pulumi:"tags"`
	Type             string                                               `pulumi:"type"`
}


func (val *LookupProvisionedClusterResult) Defaults() *LookupProvisionedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupProvisionedClusterOutput(ctx *pulumi.Context, args LookupProvisionedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupProvisionedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProvisionedClusterResult, error) {
			args := v.(LookupProvisionedClusterArgs)
			r, err := LookupProvisionedCluster(ctx, &args, opts...)
			var s LookupProvisionedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProvisionedClusterResultOutput)
}

type LookupProvisionedClusterOutputArgs struct {
	ProvisionedClustersName pulumi.StringInput `pulumi:"provisionedClustersName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProvisionedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProvisionedClusterArgs)(nil)).Elem()
}


type LookupProvisionedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupProvisionedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProvisionedClusterResult)(nil)).Elem()
}

func (o LookupProvisionedClusterResultOutput) ToLookupProvisionedClusterResultOutput() LookupProvisionedClusterResultOutput {
	return o
}

func (o LookupProvisionedClusterResultOutput) ToLookupProvisionedClusterResultOutputWithContext(ctx context.Context) LookupProvisionedClusterResultOutput {
	return o
}

func (o LookupProvisionedClusterResultOutput) ExtendedLocation() ProvisionedClustersResponseResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) *ProvisionedClustersResponseResponseExtendedLocation {
		return v.ExtendedLocation
	}).(ProvisionedClustersResponseResponseExtendedLocationPtrOutput)
}

func (o LookupProvisionedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProvisionedClusterResultOutput) Identity() ProvisionedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) *ProvisionedClusterIdentityResponse { return v.Identity }).(ProvisionedClusterIdentityResponsePtrOutput)
}

func (o LookupProvisionedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupProvisionedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProvisionedClusterResultOutput) Properties() ProvisionedClustersResponsePropertiesResponseOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) ProvisionedClustersResponsePropertiesResponse {
		return v.Properties
	}).(ProvisionedClustersResponsePropertiesResponseOutput)
}

func (o LookupProvisionedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupProvisionedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProvisionedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProvisionedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProvisionedClusterResultOutput{})
}
