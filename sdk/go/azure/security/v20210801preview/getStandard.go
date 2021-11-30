


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStandard(ctx *pulumi.Context, args *LookupStandardArgs, opts ...pulumi.InvokeOption) (*LookupStandardResult, error) {
	var rv LookupStandardResult
	err := ctx.Invoke("azure-native:security/v20210801preview:getStandard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStandardArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StandardId        string `pulumi:"standardId"`
}


type LookupStandardResult struct {
	Category     *string                               `pulumi:"category"`
	Components   []StandardComponentPropertiesResponse `pulumi:"components"`
	Description  *string                               `pulumi:"description"`
	DisplayName  *string                               `pulumi:"displayName"`
	Etag         *string                               `pulumi:"etag"`
	Id           string                                `pulumi:"id"`
	Kind         *string                               `pulumi:"kind"`
	Location     *string                               `pulumi:"location"`
	Name         string                                `pulumi:"name"`
	StandardType string                                `pulumi:"standardType"`
	SystemData   SystemDataResponse                    `pulumi:"systemData"`
	Tags         map[string]string                     `pulumi:"tags"`
	Type         string                                `pulumi:"type"`
}
