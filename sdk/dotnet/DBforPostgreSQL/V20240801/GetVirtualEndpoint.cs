// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.V20240801
{
    public static class GetVirtualEndpoint
    {
        /// <summary>
        /// Gets information about a virtual endpoint.
        /// </summary>
        public static Task<GetVirtualEndpointResult> InvokeAsync(GetVirtualEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualEndpointResult>("azure-native:dbforpostgresql/v20240801:getVirtualEndpoint", args ?? new GetVirtualEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a virtual endpoint.
        /// </summary>
        public static Output<GetVirtualEndpointResult> Invoke(GetVirtualEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualEndpointResult>("azure-native:dbforpostgresql/v20240801:getVirtualEndpoint", args ?? new GetVirtualEndpointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a virtual endpoint.
        /// </summary>
        public static Output<GetVirtualEndpointResult> Invoke(GetVirtualEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualEndpointResult>("azure-native:dbforpostgresql/v20240801:getVirtualEndpoint", args ?? new GetVirtualEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual endpoint.
        /// </summary>
        [Input("virtualEndpointName", required: true)]
        public string VirtualEndpointName { get; set; } = null!;

        public GetVirtualEndpointArgs()
        {
        }
        public static new GetVirtualEndpointArgs Empty => new GetVirtualEndpointArgs();
    }

    public sealed class GetVirtualEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual endpoint.
        /// </summary>
        [Input("virtualEndpointName", required: true)]
        public Input<string> VirtualEndpointName { get; set; } = null!;

        public GetVirtualEndpointInvokeArgs()
        {
        }
        public static new GetVirtualEndpointInvokeArgs Empty => new GetVirtualEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualEndpointResult
    {
        /// <summary>
        /// The endpoint type for the virtual endpoint.
        /// </summary>
        public readonly string? EndpointType;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of members for a virtual endpoint
        /// </summary>
        public readonly ImmutableArray<string> Members;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// List of virtual endpoints for a server
        /// </summary>
        public readonly ImmutableArray<string> VirtualEndpoints;

        [OutputConstructor]
        private GetVirtualEndpointResult(
            string? endpointType,

            string id,

            ImmutableArray<string> members,

            string name,

            Outputs.SystemDataResponse systemData,

            string type,

            ImmutableArray<string> virtualEndpoints)
        {
            EndpointType = endpointType;
            Id = id;
            Members = members;
            Name = name;
            SystemData = systemData;
            Type = type;
            VirtualEndpoints = virtualEndpoints;
        }
    }
}
