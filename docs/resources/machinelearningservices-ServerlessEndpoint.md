A Serverless Endpoint requires a Marketplace subscription. You can create one via the [MarketplaceSubscription resource](https://www.pulumi.com/registry/packages/azure-native/api-docs/machinelearningservices/marketplacesubscription/) and then making your endpoint [depend](https://www.pulumi.com/docs/iac/concepts/options/dependson/) on it (adjust for your programming language):

```
marketplace_subscription = azure.machinelearningservices.MarketplaceSubscription(...)

serverless_endpoint=azure.machinelearningservices.v20241001.ServerlessEndpoint(
    ...
    opts=pulumi.ResourceOptions(depends_on=[marketplace_subscription])
```