// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    public static class GetFirewallSupportInfo
    {
        /// <summary>
        /// support info for firewall.
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetFirewallSupportInfoResult> InvokeAsync(GetFirewallSupportInfoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallSupportInfoResult>("azure-native:cloudngfw:getFirewallSupportInfo", args ?? new GetFirewallSupportInfoArgs(), options.WithDefaults());

        /// <summary>
        /// support info for firewall.
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirewallSupportInfoResult> Invoke(GetFirewallSupportInfoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallSupportInfoResult>("azure-native:cloudngfw:getFirewallSupportInfo", args ?? new GetFirewallSupportInfoInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// support info for firewall.
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirewallSupportInfoResult> Invoke(GetFirewallSupportInfoInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallSupportInfoResult>("azure-native:cloudngfw:getFirewallSupportInfo", args ?? new GetFirewallSupportInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallSupportInfoArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// email address on behalf of which this API called
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        /// <summary>
        /// Firewall resource name
        /// </summary>
        [Input("firewallName", required: true)]
        public string FirewallName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFirewallSupportInfoArgs()
        {
        }
        public static new GetFirewallSupportInfoArgs Empty => new GetFirewallSupportInfoArgs();
    }

    public sealed class GetFirewallSupportInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// email address on behalf of which this API called
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Firewall resource name
        /// </summary>
        [Input("firewallName", required: true)]
        public Input<string> FirewallName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFirewallSupportInfoInvokeArgs()
        {
        }
        public static new GetFirewallSupportInfoInvokeArgs Empty => new GetFirewallSupportInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallSupportInfoResult
    {
        /// <summary>
        /// Support account associated with given resource when association type is tenant
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// Support account associated with given resource when association type is billing
        /// </summary>
        public readonly string? AccountIdForBilling;
        /// <summary>
        /// account registered in Customer Support Portal
        /// </summary>
        public readonly string? AccountRegistered;
        /// <summary>
        /// Association Type
        /// </summary>
        public readonly string? AssociationType;
        /// <summary>
        /// Product usage is in free trial period
        /// </summary>
        public readonly string? FreeTrial;
        /// <summary>
        /// Free trial credit remaining
        /// </summary>
        public readonly int? FreeTrialCreditLeft;
        /// <summary>
        /// Free trial days remaining
        /// </summary>
        public readonly int? FreeTrialDaysLeft;
        /// <summary>
        /// URL for paloaltonetworks live community
        /// </summary>
        public readonly string? HelpURL;
        /// <summary>
        /// product Serial associated with given resource
        /// </summary>
        public readonly string? ProductSerial;
        /// <summary>
        /// product SKU associated with given resource
        /// </summary>
        public readonly string? ProductSku;
        /// <summary>
        /// URL for registering product in paloaltonetworks Customer Service Portal
        /// </summary>
        public readonly string? RegisterURL;
        /// <summary>
        /// URL for paloaltonetworks Customer Service Portal
        /// </summary>
        public readonly string? SupportURL;
        /// <summary>
        /// user domain is supported in Customer Support Portal
        /// </summary>
        public readonly string? UserDomainSupported;
        /// <summary>
        /// user registered in Customer Support Portal
        /// </summary>
        public readonly string? UserRegistered;

        [OutputConstructor]
        private GetFirewallSupportInfoResult(
            string? accountId,

            string? accountIdForBilling,

            string? accountRegistered,

            string? associationType,

            string? freeTrial,

            int? freeTrialCreditLeft,

            int? freeTrialDaysLeft,

            string? helpURL,

            string? productSerial,

            string? productSku,

            string? registerURL,

            string? supportURL,

            string? userDomainSupported,

            string? userRegistered)
        {
            AccountId = accountId;
            AccountIdForBilling = accountIdForBilling;
            AccountRegistered = accountRegistered;
            AssociationType = associationType;
            FreeTrial = freeTrial;
            FreeTrialCreditLeft = freeTrialCreditLeft;
            FreeTrialDaysLeft = freeTrialDaysLeft;
            HelpURL = helpURL;
            ProductSerial = productSerial;
            ProductSku = productSku;
            RegisterURL = registerURL;
            SupportURL = supportURL;
            UserDomainSupported = userDomainSupported;
            UserRegistered = userRegistered;
        }
    }
}
