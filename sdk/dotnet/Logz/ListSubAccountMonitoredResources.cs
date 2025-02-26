// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logz
{
    public static class ListSubAccountMonitoredResources
    {
        /// <summary>
        /// Response of a list operation.
        /// Azure REST API version: 2022-01-01-preview.
        /// </summary>
        public static Task<ListSubAccountMonitoredResourcesResult> InvokeAsync(ListSubAccountMonitoredResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSubAccountMonitoredResourcesResult>("azure-native:logz:listSubAccountMonitoredResources", args ?? new ListSubAccountMonitoredResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// Response of a list operation.
        /// Azure REST API version: 2022-01-01-preview.
        /// </summary>
        public static Output<ListSubAccountMonitoredResourcesResult> Invoke(ListSubAccountMonitoredResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSubAccountMonitoredResourcesResult>("azure-native:logz:listSubAccountMonitoredResources", args ?? new ListSubAccountMonitoredResourcesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Response of a list operation.
        /// Azure REST API version: 2022-01-01-preview.
        /// </summary>
        public static Output<ListSubAccountMonitoredResourcesResult> Invoke(ListSubAccountMonitoredResourcesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListSubAccountMonitoredResourcesResult>("azure-native:logz:listSubAccountMonitoredResources", args ?? new ListSubAccountMonitoredResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSubAccountMonitoredResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor resource name
        /// </summary>
        [Input("monitorName", required: true)]
        public string MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Sub Account resource name
        /// </summary>
        [Input("subAccountName", required: true)]
        public string SubAccountName { get; set; } = null!;

        public ListSubAccountMonitoredResourcesArgs()
        {
        }
        public static new ListSubAccountMonitoredResourcesArgs Empty => new ListSubAccountMonitoredResourcesArgs();
    }

    public sealed class ListSubAccountMonitoredResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor resource name
        /// </summary>
        [Input("monitorName", required: true)]
        public Input<string> MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Sub Account resource name
        /// </summary>
        [Input("subAccountName", required: true)]
        public Input<string> SubAccountName { get; set; } = null!;

        public ListSubAccountMonitoredResourcesInvokeArgs()
        {
        }
        public static new ListSubAccountMonitoredResourcesInvokeArgs Empty => new ListSubAccountMonitoredResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class ListSubAccountMonitoredResourcesResult
    {
        /// <summary>
        /// Link to the next set of results, if any.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Results of a list operation.
        /// </summary>
        public readonly ImmutableArray<Outputs.MonitoredResourceResponse> Value;

        [OutputConstructor]
        private ListSubAccountMonitoredResourcesResult(
            string? nextLink,

            ImmutableArray<Outputs.MonitoredResourceResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
