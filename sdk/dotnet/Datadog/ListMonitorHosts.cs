// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Datadog
{
    public static class ListMonitorHosts
    {
        /// <summary>
        /// Response of a list operation.
        /// 
        /// Uses Azure REST API version 2023-10-20.
        /// 
        /// Other available API versions: 2022-06-01, 2022-08-01, 2023-01-01, 2023-07-07. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datadog [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListMonitorHostsResult> InvokeAsync(ListMonitorHostsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListMonitorHostsResult>("azure-native:datadog:listMonitorHosts", args ?? new ListMonitorHostsArgs(), options.WithDefaults());

        /// <summary>
        /// Response of a list operation.
        /// 
        /// Uses Azure REST API version 2023-10-20.
        /// 
        /// Other available API versions: 2022-06-01, 2022-08-01, 2023-01-01, 2023-07-07. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datadog [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListMonitorHostsResult> Invoke(ListMonitorHostsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListMonitorHostsResult>("azure-native:datadog:listMonitorHosts", args ?? new ListMonitorHostsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Response of a list operation.
        /// 
        /// Uses Azure REST API version 2023-10-20.
        /// 
        /// Other available API versions: 2022-06-01, 2022-08-01, 2023-01-01, 2023-07-07. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datadog [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListMonitorHostsResult> Invoke(ListMonitorHostsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListMonitorHostsResult>("azure-native:datadog:listMonitorHosts", args ?? new ListMonitorHostsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListMonitorHostsArgs : global::Pulumi.InvokeArgs
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

        public ListMonitorHostsArgs()
        {
        }
        public static new ListMonitorHostsArgs Empty => new ListMonitorHostsArgs();
    }

    public sealed class ListMonitorHostsInvokeArgs : global::Pulumi.InvokeArgs
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

        public ListMonitorHostsInvokeArgs()
        {
        }
        public static new ListMonitorHostsInvokeArgs Empty => new ListMonitorHostsInvokeArgs();
    }


    [OutputType]
    public sealed class ListMonitorHostsResult
    {
        /// <summary>
        /// Link to the next set of results, if any.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Results of a list operation.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatadogHostResponse> Value;

        [OutputConstructor]
        private ListMonitorHostsResult(
            string? nextLink,

            ImmutableArray<Outputs.DatadogHostResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
