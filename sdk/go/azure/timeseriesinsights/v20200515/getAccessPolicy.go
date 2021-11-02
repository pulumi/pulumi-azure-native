


package v20200515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20200515:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	AccessPolicyName  string `pulumi:"accessPolicyName"`
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccessPolicyResult struct {
	Description       *string  `pulumi:"description"`
	Id                string   `pulumi:"id"`
	Name              string   `pulumi:"name"`
	PrincipalObjectId *string  `pulumi:"principalObjectId"`
	Roles             []string `pulumi:"roles"`
	Type              string   `pulumi:"type"`
}
