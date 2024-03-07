package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId:             pulumi.String("soapApi"),
			Format:            pulumi.String("wsdl-link"),
			Path:              pulumi.String("currency"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Value:             pulumi.String("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL"),
			WsdlSelector: &apimanagement.ApiCreateOrUpdatePropertiesWsdlSelectorArgs{
				WsdlEndpointName: pulumi.String("CurrencyConvertorSoap"),
				WsdlServiceName:  pulumi.String("CurrencyConvertor"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
