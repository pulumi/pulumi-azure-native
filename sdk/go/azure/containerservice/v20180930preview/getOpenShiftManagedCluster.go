


package v20180930preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupOpenShiftManagedCluster(ctx *pulumi.Context, args *LookupOpenShiftManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupOpenShiftManagedClusterResult, error) {
	var rv LookupOpenShiftManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20180930preview:getOpenShiftManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupOpenShiftManagedClusterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupOpenShiftManagedClusterResult struct {
	AgentPoolProfiles []OpenShiftManagedClusterAgentPoolProfileResponse `pulumi:"agentPoolProfiles"`
	AuthProfile       *OpenShiftManagedClusterAuthProfileResponse       `pulumi:"authProfile"`
	Fqdn              *string                                           `pulumi:"fqdn"`
	Id                string                                            `pulumi:"id"`
	Location          string                                            `pulumi:"location"`
	MasterPoolProfile *OpenShiftManagedClusterMasterPoolProfileResponse `pulumi:"masterPoolProfile"`
	Name              string                                            `pulumi:"name"`
	NetworkProfile    *NetworkProfileResponse                           `pulumi:"networkProfile"`
	OpenShiftVersion  string                                            `pulumi:"openShiftVersion"`
	Plan              *PurchasePlanResponse                             `pulumi:"plan"`
	ProvisioningState string                                            `pulumi:"provisioningState"`
	PublicHostname    *string                                           `pulumi:"publicHostname"`
	RouterProfiles    []OpenShiftRouterProfileResponse                  `pulumi:"routerProfiles"`
	Tags              map[string]string                                 `pulumi:"tags"`
	Type              string                                            `pulumi:"type"`
}


func (val *LookupOpenShiftManagedClusterResult) Defaults() *LookupOpenShiftManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	return &tmp
}

func LookupOpenShiftManagedClusterOutput(ctx *pulumi.Context, args LookupOpenShiftManagedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupOpenShiftManagedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOpenShiftManagedClusterResult, error) {
			args := v.(LookupOpenShiftManagedClusterArgs)
			r, err := LookupOpenShiftManagedCluster(ctx, &args, opts...)
			var s LookupOpenShiftManagedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOpenShiftManagedClusterResultOutput)
}

type LookupOpenShiftManagedClusterOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupOpenShiftManagedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenShiftManagedClusterArgs)(nil)).Elem()
}


type LookupOpenShiftManagedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupOpenShiftManagedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenShiftManagedClusterResult)(nil)).Elem()
}

func (o LookupOpenShiftManagedClusterResultOutput) ToLookupOpenShiftManagedClusterResultOutput() LookupOpenShiftManagedClusterResultOutput {
	return o
}

func (o LookupOpenShiftManagedClusterResultOutput) ToLookupOpenShiftManagedClusterResultOutputWithContext(ctx context.Context) LookupOpenShiftManagedClusterResultOutput {
	return o
}

func (o LookupOpenShiftManagedClusterResultOutput) AgentPoolProfiles() OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) []OpenShiftManagedClusterAgentPoolProfileResponse {
		return v.AgentPoolProfiles
	}).(OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) AuthProfile() OpenShiftManagedClusterAuthProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) *OpenShiftManagedClusterAuthProfileResponse {
		return v.AuthProfile
	}).(OpenShiftManagedClusterAuthProfileResponsePtrOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) MasterPoolProfile() OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) *OpenShiftManagedClusterMasterPoolProfileResponse {
		return v.MasterPoolProfile
	}).(OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) OpenShiftVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) string { return v.OpenShiftVersion }).(pulumi.StringOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Plan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) *PurchasePlanResponse { return v.Plan }).(PurchasePlanResponsePtrOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) PublicHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) *string { return v.PublicHostname }).(pulumi.StringPtrOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) RouterProfiles() OpenShiftRouterProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) []OpenShiftRouterProfileResponse { return v.RouterProfiles }).(OpenShiftRouterProfileResponseArrayOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOpenShiftManagedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftManagedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOpenShiftManagedClusterResultOutput{})
}
