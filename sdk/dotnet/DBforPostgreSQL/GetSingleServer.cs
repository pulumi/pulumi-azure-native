// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL
{
    public static class GetSingleServer
    {
        /// <summary>
        /// Gets information about a server.
        /// 
        /// Uses Azure REST API version 2017-12-01.
        /// </summary>
        public static Task<GetSingleServerResult> InvokeAsync(GetSingleServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSingleServerResult>("azure-native:dbforpostgresql:getSingleServer", args ?? new GetSingleServerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a server.
        /// 
        /// Uses Azure REST API version 2017-12-01.
        /// </summary>
        public static Output<GetSingleServerResult> Invoke(GetSingleServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSingleServerResult>("azure-native:dbforpostgresql:getSingleServer", args ?? new GetSingleServerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a server.
        /// 
        /// Uses Azure REST API version 2017-12-01.
        /// </summary>
        public static Output<GetSingleServerResult> Invoke(GetSingleServerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSingleServerResult>("azure-native:dbforpostgresql:getSingleServer", args ?? new GetSingleServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSingleServerArgs : global::Pulumi.InvokeArgs
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

        public GetSingleServerArgs()
        {
        }
        public static new GetSingleServerArgs Empty => new GetSingleServerArgs();
    }

    public sealed class GetSingleServerInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSingleServerInvokeArgs()
        {
        }
        public static new GetSingleServerInvokeArgs Empty => new GetSingleServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetSingleServerResult
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Status showing whether the server data encryption is enabled with customer-managed keys.
        /// </summary>
        public readonly string ByokEnforcement;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string? EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string? FullyQualifiedDomainName;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponse? Identity;
        /// <summary>
        /// Status showing whether the server enabled infrastructure encryption.
        /// </summary>
        public readonly string? InfrastructureEncryption;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The master server id of a replica server.
        /// </summary>
        public readonly string? MasterServerId;
        /// <summary>
        /// Enforce a minimal Tls version for the server.
        /// </summary>
        public readonly string? MinimalTlsVersion;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of private endpoint connections on a server
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The maximum number of replicas that a master server can have.
        /// </summary>
        public readonly int? ReplicaCapacity;
        /// <summary>
        /// The replication role of the server.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        public readonly Outputs.SingleServerSkuResponse? Sku;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        public readonly Outputs.StorageProfileResponse? StorageProfile;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A state of a server that is visible to user.
        /// </summary>
        public readonly string? UserVisibleState;
        /// <summary>
        /// Server version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetSingleServerResult(
            string? administratorLogin,

            string azureApiVersion,

            string byokEnforcement,

            string? earliestRestoreDate,

            string? fullyQualifiedDomainName,

            string id,

            Outputs.ResourceIdentityResponse? identity,

            string? infrastructureEncryption,

            string location,

            string? masterServerId,

            string? minimalTlsVersion,

            string name,

            ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponse> privateEndpointConnections,

            string? publicNetworkAccess,

            int? replicaCapacity,

            string? replicationRole,

            Outputs.SingleServerSkuResponse? sku,

            string? sslEnforcement,

            Outputs.StorageProfileResponse? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? userVisibleState,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            AzureApiVersion = azureApiVersion;
            ByokEnforcement = byokEnforcement;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            Id = id;
            Identity = identity;
            InfrastructureEncryption = infrastructureEncryption;
            Location = location;
            MasterServerId = masterServerId;
            MinimalTlsVersion = minimalTlsVersion;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            Sku = sku;
            SslEnforcement = sslEnforcement;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            UserVisibleState = userVisibleState;
            Version = version;
        }
    }
}
