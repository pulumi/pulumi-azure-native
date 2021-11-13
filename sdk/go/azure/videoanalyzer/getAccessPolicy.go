


package videoanalyzer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("azure-native:videoanalyzer:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	AccessPolicyName  string `pulumi:"accessPolicyName"`
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccessPolicyResult struct {
	Authentication *JwtAuthenticationResponse `pulumi:"authentication"`
	Id             string                     `pulumi:"id"`
	Name           string                     `pulumi:"name"`
	Role           *string                    `pulumi:"role"`
	SystemData     SystemDataResponse         `pulumi:"systemData"`
	Type           string                     `pulumi:"type"`
}
