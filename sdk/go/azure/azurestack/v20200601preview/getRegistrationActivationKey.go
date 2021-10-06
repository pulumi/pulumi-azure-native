


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRegistrationActivationKey(ctx *pulumi.Context, args *GetRegistrationActivationKeyArgs, opts ...pulumi.InvokeOption) (*GetRegistrationActivationKeyResult, error) {
	var rv GetRegistrationActivationKeyResult
	err := ctx.Invoke("azure-native:azurestack/v20200601preview:getRegistrationActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistrationActivationKeyArgs struct {
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type GetRegistrationActivationKeyResult struct {
	ActivationKey *string `pulumi:"activationKey"`
}
