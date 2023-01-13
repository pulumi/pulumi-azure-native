


package v20220904

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOpenShiftCluster(ctx *pulumi.Context, args *LookupOpenShiftClusterArgs, opts ...pulumi.InvokeOption) (*LookupOpenShiftClusterResult, error) {
	var rv LookupOpenShiftClusterResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20220904:getOpenShiftCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenShiftClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupOpenShiftClusterResult struct {
	ApiserverProfile        *APIServerProfileResponse        `pulumi:"apiserverProfile"`
	ClusterProfile          *ClusterProfileResponse          `pulumi:"clusterProfile"`
	ConsoleProfile          *ConsoleProfileResponse          `pulumi:"consoleProfile"`
	Id                      string                           `pulumi:"id"`
	IngressProfiles         []IngressProfileResponse         `pulumi:"ingressProfiles"`
	Location                string                           `pulumi:"location"`
	MasterProfile           *MasterProfileResponse           `pulumi:"masterProfile"`
	Name                    string                           `pulumi:"name"`
	NetworkProfile          *NetworkProfileResponse          `pulumi:"networkProfile"`
	ProvisioningState       *string                          `pulumi:"provisioningState"`
	ServicePrincipalProfile *ServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	SystemData              SystemDataResponse               `pulumi:"systemData"`
	Tags                    map[string]string                `pulumi:"tags"`
	Type                    string                           `pulumi:"type"`
	WorkerProfiles          []WorkerProfileResponse          `pulumi:"workerProfiles"`
}

func LookupOpenShiftClusterOutput(ctx *pulumi.Context, args LookupOpenShiftClusterOutputArgs, opts ...pulumi.InvokeOption) LookupOpenShiftClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOpenShiftClusterResult, error) {
			args := v.(LookupOpenShiftClusterArgs)
			r, err := LookupOpenShiftCluster(ctx, &args, opts...)
			var s LookupOpenShiftClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOpenShiftClusterResultOutput)
}

type LookupOpenShiftClusterOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupOpenShiftClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenShiftClusterArgs)(nil)).Elem()
}


type LookupOpenShiftClusterResultOutput struct{ *pulumi.OutputState }

func (LookupOpenShiftClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenShiftClusterResult)(nil)).Elem()
}

func (o LookupOpenShiftClusterResultOutput) ToLookupOpenShiftClusterResultOutput() LookupOpenShiftClusterResultOutput {
	return o
}

func (o LookupOpenShiftClusterResultOutput) ToLookupOpenShiftClusterResultOutputWithContext(ctx context.Context) LookupOpenShiftClusterResultOutput {
	return o
}

func (o LookupOpenShiftClusterResultOutput) ApiserverProfile() APIServerProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *APIServerProfileResponse { return v.ApiserverProfile }).(APIServerProfileResponsePtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) ClusterProfile() ClusterProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *ClusterProfileResponse { return v.ClusterProfile }).(ClusterProfileResponsePtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) ConsoleProfile() ConsoleProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *ConsoleProfileResponse { return v.ConsoleProfile }).(ConsoleProfileResponsePtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOpenShiftClusterResultOutput) IngressProfiles() IngressProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) []IngressProfileResponse { return v.IngressProfiles }).(IngressProfileResponseArrayOutput)
}

func (o LookupOpenShiftClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOpenShiftClusterResultOutput) MasterProfile() MasterProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *MasterProfileResponse { return v.MasterProfile }).(MasterProfileResponsePtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOpenShiftClusterResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) ServicePrincipalProfile() ServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *ServicePrincipalProfileResponse {
		return v.ServicePrincipalProfile
	}).(ServicePrincipalProfileResponsePtrOutput)
}

func (o LookupOpenShiftClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOpenShiftClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOpenShiftClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupOpenShiftClusterResultOutput) WorkerProfiles() WorkerProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) []WorkerProfileResponse { return v.WorkerProfiles }).(WorkerProfileResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOpenShiftClusterResultOutput{})
}
