// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    public static class ListLocalRulestackSecurityServices
    {
        /// <summary>
        /// List the security services for rulestack
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListLocalRulestackSecurityServicesResult> InvokeAsync(ListLocalRulestackSecurityServicesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListLocalRulestackSecurityServicesResult>("azure-native:cloudngfw:listLocalRulestackSecurityServices", args ?? new ListLocalRulestackSecurityServicesArgs(), options.WithDefaults());

        /// <summary>
        /// List the security services for rulestack
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListLocalRulestackSecurityServicesResult> Invoke(ListLocalRulestackSecurityServicesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListLocalRulestackSecurityServicesResult>("azure-native:cloudngfw:listLocalRulestackSecurityServices", args ?? new ListLocalRulestackSecurityServicesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List the security services for rulestack
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListLocalRulestackSecurityServicesResult> Invoke(ListLocalRulestackSecurityServicesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListLocalRulestackSecurityServicesResult>("azure-native:cloudngfw:listLocalRulestackSecurityServices", args ?? new ListLocalRulestackSecurityServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListLocalRulestackSecurityServicesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// LocalRulestack resource name
        /// </summary>
        [Input("localRulestackName", required: true)]
        public string LocalRulestackName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("skip")]
        public string? Skip { get; set; }

        [Input("top")]
        public int? Top { get; set; }

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public ListLocalRulestackSecurityServicesArgs()
        {
        }
        public static new ListLocalRulestackSecurityServicesArgs Empty => new ListLocalRulestackSecurityServicesArgs();
    }

    public sealed class ListLocalRulestackSecurityServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// LocalRulestack resource name
        /// </summary>
        [Input("localRulestackName", required: true)]
        public Input<string> LocalRulestackName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("skip")]
        public Input<string>? Skip { get; set; }

        [Input("top")]
        public Input<int>? Top { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListLocalRulestackSecurityServicesInvokeArgs()
        {
        }
        public static new ListLocalRulestackSecurityServicesInvokeArgs Empty => new ListLocalRulestackSecurityServicesInvokeArgs();
    }


    [OutputType]
    public sealed class ListLocalRulestackSecurityServicesResult
    {
        /// <summary>
        /// next link
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// response value
        /// </summary>
        public readonly Outputs.SecurityServicesTypeListResponse Value;

        [OutputConstructor]
        private ListLocalRulestackSecurityServicesResult(
            string? nextLink,

            Outputs.SecurityServicesTypeListResponse value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
