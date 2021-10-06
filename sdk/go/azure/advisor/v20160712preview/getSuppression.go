


package v20160712preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSuppression(ctx *pulumi.Context, args *LookupSuppressionArgs, opts ...pulumi.InvokeOption) (*LookupSuppressionResult, error) {
	var rv LookupSuppressionResult
	err := ctx.Invoke("azure-native:advisor/v20160712preview:getSuppression", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSuppressionArgs struct {
	Name             string `pulumi:"name"`
	RecommendationId string `pulumi:"recommendationId"`
	ResourceUri      string `pulumi:"resourceUri"`
}


type LookupSuppressionResult struct {
	Id            string            `pulumi:"id"`
	Location      *string           `pulumi:"location"`
	Name          string            `pulumi:"name"`
	SuppressionId *string           `pulumi:"suppressionId"`
	Tags          map[string]string `pulumi:"tags"`
	Ttl           *string           `pulumi:"ttl"`
	Type          string            `pulumi:"type"`
}
