


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLogProfile(ctx *pulumi.Context, args *LookupLogProfileArgs, opts ...pulumi.InvokeOption) (*LookupLogProfileResult, error) {
	var rv LookupLogProfileResult
	err := ctx.Invoke("azure-native:insights:getLogProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLogProfileArgs struct {
	LogProfileName string `pulumi:"logProfileName"`
}


type LookupLogProfileResult struct {
	Categories       []string                `pulumi:"categories"`
	Id               string                  `pulumi:"id"`
	Location         string                  `pulumi:"location"`
	Locations        []string                `pulumi:"locations"`
	Name             string                  `pulumi:"name"`
	RetentionPolicy  RetentionPolicyResponse `pulumi:"retentionPolicy"`
	ServiceBusRuleId *string                 `pulumi:"serviceBusRuleId"`
	StorageAccountId *string                 `pulumi:"storageAccountId"`
	Tags             map[string]string       `pulumi:"tags"`
	Type             string                  `pulumi:"type"`
}
