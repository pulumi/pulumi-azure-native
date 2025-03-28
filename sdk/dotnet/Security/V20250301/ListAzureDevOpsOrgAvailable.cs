// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.V20250301
{
    public static class ListAzureDevOpsOrgAvailable
    {
        /// <summary>
        /// List of RP resources which supports pagination.
        /// </summary>
        public static Task<ListAzureDevOpsOrgAvailableResult> InvokeAsync(ListAzureDevOpsOrgAvailableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListAzureDevOpsOrgAvailableResult>("azure-native:security/v20250301:listAzureDevOpsOrgAvailable", args ?? new ListAzureDevOpsOrgAvailableArgs(), options.WithDefaults());

        /// <summary>
        /// List of RP resources which supports pagination.
        /// </summary>
        public static Output<ListAzureDevOpsOrgAvailableResult> Invoke(ListAzureDevOpsOrgAvailableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListAzureDevOpsOrgAvailableResult>("azure-native:security/v20250301:listAzureDevOpsOrgAvailable", args ?? new ListAzureDevOpsOrgAvailableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List of RP resources which supports pagination.
        /// </summary>
        public static Output<ListAzureDevOpsOrgAvailableResult> Invoke(ListAzureDevOpsOrgAvailableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListAzureDevOpsOrgAvailableResult>("azure-native:security/v20250301:listAzureDevOpsOrgAvailable", args ?? new ListAzureDevOpsOrgAvailableInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListAzureDevOpsOrgAvailableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The security connector name.
        /// </summary>
        [Input("securityConnectorName", required: true)]
        public string SecurityConnectorName { get; set; } = null!;

        public ListAzureDevOpsOrgAvailableArgs()
        {
        }
        public static new ListAzureDevOpsOrgAvailableArgs Empty => new ListAzureDevOpsOrgAvailableArgs();
    }

    public sealed class ListAzureDevOpsOrgAvailableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The security connector name.
        /// </summary>
        [Input("securityConnectorName", required: true)]
        public Input<string> SecurityConnectorName { get; set; } = null!;

        public ListAzureDevOpsOrgAvailableInvokeArgs()
        {
        }
        public static new ListAzureDevOpsOrgAvailableInvokeArgs Empty => new ListAzureDevOpsOrgAvailableInvokeArgs();
    }


    [OutputType]
    public sealed class ListAzureDevOpsOrgAvailableResult
    {
        /// <summary>
        /// Gets or sets next link to scroll over the results.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Gets or sets list of resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureDevOpsOrgResponse> Value;

        [OutputConstructor]
        private ListAzureDevOpsOrgAvailableResult(
            string? nextLink,

            ImmutableArray<Outputs.AzureDevOpsOrgResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
