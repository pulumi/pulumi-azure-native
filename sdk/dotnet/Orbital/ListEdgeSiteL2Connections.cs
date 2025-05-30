// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Orbital
{
    public static class ListEdgeSiteL2Connections
    {
        /// <summary>
        /// Returns a list of L2 Connections attached to an edge site.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListEdgeSiteL2ConnectionsResult> InvokeAsync(ListEdgeSiteL2ConnectionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListEdgeSiteL2ConnectionsResult>("azure-native:orbital:listEdgeSiteL2Connections", args ?? new ListEdgeSiteL2ConnectionsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of L2 Connections attached to an edge site.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListEdgeSiteL2ConnectionsResult> Invoke(ListEdgeSiteL2ConnectionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListEdgeSiteL2ConnectionsResult>("azure-native:orbital:listEdgeSiteL2Connections", args ?? new ListEdgeSiteL2ConnectionsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of L2 Connections attached to an edge site.
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native orbital [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListEdgeSiteL2ConnectionsResult> Invoke(ListEdgeSiteL2ConnectionsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListEdgeSiteL2ConnectionsResult>("azure-native:orbital:listEdgeSiteL2Connections", args ?? new ListEdgeSiteL2ConnectionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListEdgeSiteL2ConnectionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Edge site name.
        /// </summary>
        [Input("edgeSiteName", required: true)]
        public string EdgeSiteName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListEdgeSiteL2ConnectionsArgs()
        {
        }
        public static new ListEdgeSiteL2ConnectionsArgs Empty => new ListEdgeSiteL2ConnectionsArgs();
    }

    public sealed class ListEdgeSiteL2ConnectionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Edge site name.
        /// </summary>
        [Input("edgeSiteName", required: true)]
        public Input<string> EdgeSiteName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListEdgeSiteL2ConnectionsInvokeArgs()
        {
        }
        public static new ListEdgeSiteL2ConnectionsInvokeArgs Empty => new ListEdgeSiteL2ConnectionsInvokeArgs();
    }


    [OutputType]
    public sealed class ListEdgeSiteL2ConnectionsResult
    {
        /// <summary>
        /// The URL to get the next set of results.
        /// </summary>
        public readonly string NextLink;
        /// <summary>
        /// A list of Azure Resource IDs.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdListResultResponseValue> Value;

        [OutputConstructor]
        private ListEdgeSiteL2ConnectionsResult(
            string nextLink,

            ImmutableArray<Outputs.ResourceIdListResultResponseValue> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
