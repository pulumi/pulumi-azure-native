using System.Collections.Generic;
using Pulumi.AzureNative.App;

return await Pulumi.Deployment.RunAsync(() =>
{
    var azureDomainVerificationId = GetCustomDomainVerificationId.Invoke().Apply(x => x.Value);

    return new Dictionary<string, object?>
    {
        ["domainVerificationId"] = azureDomainVerificationId
    };
});