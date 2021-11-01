


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistration(ctx *pulumi.Context, args *LookupRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationResult, error) {
	var rv LookupRegistrationResult
	err := ctx.Invoke("azure-native:azurestack/v20170601:getRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationArgs struct {
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type LookupRegistrationResult struct {
	BillingModel *string           `pulumi:"billingModel"`
	CloudId      *string           `pulumi:"cloudId"`
	Etag         *string           `pulumi:"etag"`
	Id           string            `pulumi:"id"`
	Location     string            `pulumi:"location"`
	Name         string            `pulumi:"name"`
	ObjectId     *string           `pulumi:"objectId"`
	Tags         map[string]string `pulumi:"tags"`
	Type         string            `pulumi:"type"`
}
