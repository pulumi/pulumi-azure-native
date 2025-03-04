// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforMySQL.V20200701PrivatePreview
{
    public static class GetServer
    {
        /// <summary>
        /// Gets information about a server.
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azure-native:dbformysql/v20200701privatepreview:getServer", args ?? new GetServerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a server.
        /// </summary>
        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerResult>("azure-native:dbformysql/v20200701privatepreview:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a server.
        /// </summary>
        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerResult>("azure-native:dbformysql/v20200701privatepreview:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : global::Pulumi.InvokeArgs
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

        public GetServerArgs()
        {
        }
        public static new GetServerArgs Empty => new GetServerArgs();
    }

    public sealed class GetServerInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetServerInvokeArgs()
        {
        }
        public static new GetServerInvokeArgs Empty => new GetServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// Status showing whether the data encryption is enabled with customer-managed keys.
        /// </summary>
        public readonly string ByokEnforcement;
        /// <summary>
        /// Delegated subnet arguments.
        /// </summary>
        public readonly Outputs.DelegatedSubnetArgumentsResponse? DelegatedSubnetArguments;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string FullyQualifiedDomainName;
        /// <summary>
        /// Enable HA or not for a server.
        /// </summary>
        public readonly string? HaEnabled;
        /// <summary>
        /// The state of a HA server.
        /// </summary>
        public readonly string HaState;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Maintenance window of a server.
        /// </summary>
        public readonly Outputs.MaintenanceWindowResponse? MaintenanceWindow;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string PublicNetworkAccess;
        /// <summary>
        /// The maximum number of replicas that a primary server can have.
        /// </summary>
        public readonly int ReplicaCapacity;
        /// <summary>
        /// The replication role.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// The source MySQL server id.
        /// </summary>
        public readonly string? SourceServerId;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        public readonly string StandbyAvailabilityZone;
        /// <summary>
        /// The state of a server.
        /// </summary>
        public readonly string State;
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
        /// Server version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetServerResult(
            string? administratorLogin,

            string? availabilityZone,

            string byokEnforcement,

            Outputs.DelegatedSubnetArgumentsResponse? delegatedSubnetArguments,

            string earliestRestoreDate,

            string fullyQualifiedDomainName,

            string? haEnabled,

            string haState,

            string id,

            Outputs.IdentityResponse? identity,

            string location,

            Outputs.MaintenanceWindowResponse? maintenanceWindow,

            string name,

            string publicNetworkAccess,

            int replicaCapacity,

            string? replicationRole,

            Outputs.SkuResponse? sku,

            string? sourceServerId,

            string? sslEnforcement,

            string standbyAvailabilityZone,

            string state,

            Outputs.StorageProfileResponse? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            AvailabilityZone = availabilityZone;
            ByokEnforcement = byokEnforcement;
            DelegatedSubnetArguments = delegatedSubnetArguments;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            HaEnabled = haEnabled;
            HaState = haState;
            Id = id;
            Identity = identity;
            Location = location;
            MaintenanceWindow = maintenanceWindow;
            Name = name;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            Sku = sku;
            SourceServerId = sourceServerId;
            SslEnforcement = sslEnforcement;
            StandbyAvailabilityZone = standbyAvailabilityZone;
            State = state;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}
