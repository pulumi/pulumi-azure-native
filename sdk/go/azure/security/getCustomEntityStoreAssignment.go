


package security

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomEntityStoreAssignment(ctx *pulumi.Context, args *LookupCustomEntityStoreAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupCustomEntityStoreAssignmentResult, error) {
	var rv LookupCustomEntityStoreAssignmentResult
	err := ctx.Invoke("azure-native:security:getCustomEntityStoreAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomEntityStoreAssignmentArgs struct {
	CustomEntityStoreAssignmentName string `pulumi:"customEntityStoreAssignmentName"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
}


type LookupCustomEntityStoreAssignmentResult struct {
	EntityStoreDatabaseLink *string `pulumi:"entityStoreDatabaseLink"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	Principal               *string `pulumi:"principal"`
	Type                    string  `pulumi:"type"`
}
