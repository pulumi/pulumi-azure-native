


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJitRequest(ctx *pulumi.Context, args *LookupJitRequestArgs, opts ...pulumi.InvokeOption) (*LookupJitRequestResult, error) {
	var rv LookupJitRequestResult
	err := ctx.Invoke("azure-native:solutions/v20210701:getJitRequest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJitRequestArgs struct {
	JitRequestName    string `pulumi:"jitRequestName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupJitRequestResult struct {
	ApplicationResourceId    string                             `pulumi:"applicationResourceId"`
	CreatedBy                ApplicationClientDetailsResponse   `pulumi:"createdBy"`
	Id                       string                             `pulumi:"id"`
	JitAuthorizationPolicies []JitAuthorizationPoliciesResponse `pulumi:"jitAuthorizationPolicies"`
	JitRequestState          string                             `pulumi:"jitRequestState"`
	JitSchedulingPolicy      JitSchedulingPolicyResponse        `pulumi:"jitSchedulingPolicy"`
	Location                 *string                            `pulumi:"location"`
	Name                     string                             `pulumi:"name"`
	ProvisioningState        string                             `pulumi:"provisioningState"`
	PublisherTenantId        string                             `pulumi:"publisherTenantId"`
	SystemData               SystemDataResponse                 `pulumi:"systemData"`
	Tags                     map[string]string                  `pulumi:"tags"`
	Type                     string                             `pulumi:"type"`
	UpdatedBy                ApplicationClientDetailsResponse   `pulumi:"updatedBy"`
}
