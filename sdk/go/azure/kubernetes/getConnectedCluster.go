


package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedCluster(ctx *pulumi.Context, args *LookupConnectedClusterArgs, opts ...pulumi.InvokeOption) (*LookupConnectedClusterResult, error) {
	var rv LookupConnectedClusterResult
	err := ctx.Invoke("azure-native:kubernetes:getConnectedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectedClusterResult struct {
	AgentPublicKeyCertificate                string                           `pulumi:"agentPublicKeyCertificate"`
	AgentVersion                             string                           `pulumi:"agentVersion"`
	ConnectivityStatus                       string                           `pulumi:"connectivityStatus"`
	Distribution                             *string                          `pulumi:"distribution"`
	Id                                       string                           `pulumi:"id"`
	Identity                                 ConnectedClusterIdentityResponse `pulumi:"identity"`
	Infrastructure                           *string                          `pulumi:"infrastructure"`
	KubernetesVersion                        string                           `pulumi:"kubernetesVersion"`
	LastConnectivityTime                     string                           `pulumi:"lastConnectivityTime"`
	Location                                 string                           `pulumi:"location"`
	ManagedIdentityCertificateExpirationTime string                           `pulumi:"managedIdentityCertificateExpirationTime"`
	Name                                     string                           `pulumi:"name"`
	Offering                                 string                           `pulumi:"offering"`
	ProvisioningState                        *string                          `pulumi:"provisioningState"`
	SystemData                               SystemDataResponse               `pulumi:"systemData"`
	Tags                                     map[string]string                `pulumi:"tags"`
	TotalCoreCount                           int                              `pulumi:"totalCoreCount"`
	TotalNodeCount                           int                              `pulumi:"totalNodeCount"`
	Type                                     string                           `pulumi:"type"`
}


func (val *LookupConnectedClusterResult) Defaults() *LookupConnectedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Identity = *tmp.Identity.Defaults()

	return &tmp
}

func LookupConnectedClusterOutput(ctx *pulumi.Context, args LookupConnectedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedClusterResult, error) {
			args := v.(LookupConnectedClusterArgs)
			r, err := LookupConnectedCluster(ctx, &args, opts...)
			var s LookupConnectedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedClusterResultOutput)
}

type LookupConnectedClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedClusterArgs)(nil)).Elem()
}


type LookupConnectedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedClusterResult)(nil)).Elem()
}

func (o LookupConnectedClusterResultOutput) ToLookupConnectedClusterResultOutput() LookupConnectedClusterResultOutput {
	return o
}

func (o LookupConnectedClusterResultOutput) ToLookupConnectedClusterResultOutputWithContext(ctx context.Context) LookupConnectedClusterResultOutput {
	return o
}

func (o LookupConnectedClusterResultOutput) AgentPublicKeyCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.AgentPublicKeyCertificate }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) ConnectivityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.ConnectivityStatus }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) Distribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.Distribution }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) Identity() ConnectedClusterIdentityResponseOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) ConnectedClusterIdentityResponse { return v.Identity }).(ConnectedClusterIdentityResponseOutput)
}

func (o LookupConnectedClusterResultOutput) Infrastructure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.Infrastructure }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedClusterResultOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) LastConnectivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.LastConnectivityTime }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) ManagedIdentityCertificateExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.ManagedIdentityCertificateExpirationTime }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) Offering() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Offering }).(pulumi.StringOutput)
}

func (o LookupConnectedClusterResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectedClusterResultOutput) TotalCoreCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) int { return v.TotalCoreCount }).(pulumi.IntOutput)
}

func (o LookupConnectedClusterResultOutput) TotalNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) int { return v.TotalNodeCount }).(pulumi.IntOutput)
}

func (o LookupConnectedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedClusterResultOutput{})
}
