// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry
{
    public static class GetConnectedRegistry
    {
        /// <summary>
        /// Gets the properties of the connected registry.
        /// 
        /// Uses Azure REST API version 2023-01-01-preview.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-08-01-preview, 2023-11-01-preview, 2024-11-01-preview.
        /// </summary>
        public static Task<GetConnectedRegistryResult> InvokeAsync(GetConnectedRegistryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectedRegistryResult>("azure-native:containerregistry:getConnectedRegistry", args ?? new GetConnectedRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the connected registry.
        /// 
        /// Uses Azure REST API version 2023-01-01-preview.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-08-01-preview, 2023-11-01-preview, 2024-11-01-preview.
        /// </summary>
        public static Output<GetConnectedRegistryResult> Invoke(GetConnectedRegistryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectedRegistryResult>("azure-native:containerregistry:getConnectedRegistry", args ?? new GetConnectedRegistryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the connected registry.
        /// 
        /// Uses Azure REST API version 2023-01-01-preview.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-08-01-preview, 2023-11-01-preview, 2024-11-01-preview.
        /// </summary>
        public static Output<GetConnectedRegistryResult> Invoke(GetConnectedRegistryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectedRegistryResult>("azure-native:containerregistry:getConnectedRegistry", args ?? new GetConnectedRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectedRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the connected registry.
        /// </summary>
        [Input("connectedRegistryName", required: true)]
        public string ConnectedRegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectedRegistryArgs()
        {
        }
        public static new GetConnectedRegistryArgs Empty => new GetConnectedRegistryArgs();
    }

    public sealed class GetConnectedRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the connected registry.
        /// </summary>
        [Input("connectedRegistryName", required: true)]
        public Input<string> ConnectedRegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetConnectedRegistryInvokeArgs()
        {
        }
        public static new GetConnectedRegistryInvokeArgs Empty => new GetConnectedRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectedRegistryResult
    {
        /// <summary>
        /// The activation properties of the connected registry.
        /// </summary>
        public readonly Outputs.ActivationPropertiesResponse Activation;
        /// <summary>
        /// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
        /// </summary>
        public readonly ImmutableArray<string> ClientTokenIds;
        /// <summary>
        /// The current connection state of the connected registry.
        /// </summary>
        public readonly string ConnectionState;
        /// <summary>
        /// The resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last activity time of the connected registry.
        /// </summary>
        public readonly string LastActivityTime;
        /// <summary>
        /// The logging properties of the connected registry.
        /// </summary>
        public readonly Outputs.LoggingPropertiesResponse? Logging;
        /// <summary>
        /// The login server properties of the connected registry.
        /// </summary>
        public readonly Outputs.LoginServerPropertiesResponse? LoginServer;
        /// <summary>
        /// The mode of the connected registry resource that indicates the permissions of the registry.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of notifications subscription information for the connected registry.
        /// </summary>
        public readonly ImmutableArray<string> NotificationsList;
        /// <summary>
        /// The parent of the connected registry.
        /// </summary>
        public readonly Outputs.ParentPropertiesResponse Parent;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The list of current statuses of the connected registry.
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusDetailPropertiesResponse> StatusDetails;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The current version of ACR runtime on the connected registry.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetConnectedRegistryResult(
            Outputs.ActivationPropertiesResponse activation,

            ImmutableArray<string> clientTokenIds,

            string connectionState,

            string id,

            string lastActivityTime,

            Outputs.LoggingPropertiesResponse? logging,

            Outputs.LoginServerPropertiesResponse? loginServer,

            string mode,

            string name,

            ImmutableArray<string> notificationsList,

            Outputs.ParentPropertiesResponse parent,

            string provisioningState,

            ImmutableArray<Outputs.StatusDetailPropertiesResponse> statusDetails,

            Outputs.SystemDataResponse systemData,

            string type,

            string version)
        {
            Activation = activation;
            ClientTokenIds = clientTokenIds;
            ConnectionState = connectionState;
            Id = id;
            LastActivityTime = lastActivityTime;
            Logging = logging;
            LoginServer = loginServer;
            Mode = mode;
            Name = name;
            NotificationsList = notificationsList;
            Parent = parent;
            ProvisioningState = provisioningState;
            StatusDetails = statusDetails;
            SystemData = systemData;
            Type = type;
            Version = version;
        }
    }
}
