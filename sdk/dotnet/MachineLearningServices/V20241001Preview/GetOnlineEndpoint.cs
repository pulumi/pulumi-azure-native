// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20241001Preview
{
    public static class GetOnlineEndpoint
    {
        public static Task<GetOnlineEndpointResult> InvokeAsync(GetOnlineEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOnlineEndpointResult>("azure-native:machinelearningservices/v20241001preview:getOnlineEndpoint", args ?? new GetOnlineEndpointArgs(), options.WithDefaults());

        public static Output<GetOnlineEndpointResult> Invoke(GetOnlineEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOnlineEndpointResult>("azure-native:machinelearningservices/v20241001preview:getOnlineEndpoint", args ?? new GetOnlineEndpointInvokeArgs(), options.WithDefaults());

        public static Output<GetOnlineEndpointResult> Invoke(GetOnlineEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOnlineEndpointResult>("azure-native:machinelearningservices/v20241001preview:getOnlineEndpoint", args ?? new GetOnlineEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOnlineEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Online Endpoint name.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetOnlineEndpointArgs()
        {
        }
        public static new GetOnlineEndpointArgs Empty => new GetOnlineEndpointArgs();
    }

    public sealed class GetOnlineEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Online Endpoint name.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetOnlineEndpointInvokeArgs()
        {
        }
        public static new GetOnlineEndpointInvokeArgs Empty => new GetOnlineEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetOnlineEndpointResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// [Required] Additional attributes of the entity.
        /// </summary>
        public readonly Outputs.OnlineEndpointResponse OnlineEndpointProperties;
        /// <summary>
        /// Sku details required for ARM contract for Autoscaling.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetOnlineEndpointResult(
            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string? kind,

            string location,

            string name,

            Outputs.OnlineEndpointResponse onlineEndpointProperties,

            Outputs.SkuResponse? sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Identity = identity;
            Kind = kind;
            Location = location;
            Name = name;
            OnlineEndpointProperties = onlineEndpointProperties;
            Sku = sku;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
