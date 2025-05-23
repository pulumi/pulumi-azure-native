// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL
{
    /// <summary>
    /// Represents a cluster.
    /// 
    /// Uses Azure REST API version 2023-03-02-preview.
    /// 
    /// Other available API versions: 2022-11-08. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:dbforpostgresql:ServerGroupCluster")]
    public partial class ServerGroupCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the cluster was created using AAD authentication.
        /// </summary>
        [Output("aadAuthEnabled")]
        public Output<string> AadAuthEnabled { get; private set; } = null!;

        /// <summary>
        /// The administrator's login name of the servers in the cluster.
        /// </summary>
        [Output("administratorLogin")]
        public Output<string> AdministratorLogin { get; private set; } = null!;

        /// <summary>
        /// Authentication configuration of a cluster.
        /// </summary>
        [Output("authConfig")]
        public Output<Outputs.ServerGroupClusterAuthConfigResponse?> AuthConfig { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The Citus extension version on all cluster servers.
        /// </summary>
        [Output("citusVersion")]
        public Output<string?> CitusVersion { get; private set; } = null!;

        /// <summary>
        /// If public access is enabled on coordinator.
        /// </summary>
        [Output("coordinatorEnablePublicIpAccess")]
        public Output<bool?> CoordinatorEnablePublicIpAccess { get; private set; } = null!;

        /// <summary>
        /// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
        /// </summary>
        [Output("coordinatorServerEdition")]
        public Output<string?> CoordinatorServerEdition { get; private set; } = null!;

        /// <summary>
        /// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Output("coordinatorStorageQuotaInMb")]
        public Output<int?> CoordinatorStorageQuotaInMb { get; private set; } = null!;

        /// <summary>
        /// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Output("coordinatorVCores")]
        public Output<int?> CoordinatorVCores { get; private set; } = null!;

        /// <summary>
        /// The data encryption properties of a cluster.
        /// </summary>
        [Output("dataEncryption")]
        public Output<Outputs.ServerGroupClusterDataEncryptionResponse?> DataEncryption { get; private set; } = null!;

        /// <summary>
        /// The database name of the cluster. Only one database per cluster is supported.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The earliest restore point time (ISO8601 format) for the cluster.
        /// </summary>
        [Output("earliestRestoreTime")]
        public Output<string> EarliestRestoreTime { get; private set; } = null!;

        /// <summary>
        /// If cluster backup is stored in another Azure region in addition to the copy of the backup stored in the cluster's region. Enabled only at the time of cluster creation.
        /// </summary>
        [Output("enableGeoBackup")]
        public Output<bool?> EnableGeoBackup { get; private set; } = null!;

        /// <summary>
        /// If high availability (HA) is enabled or not for the cluster.
        /// </summary>
        [Output("enableHa")]
        public Output<bool?> EnableHa { get; private set; } = null!;

        /// <summary>
        /// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
        /// </summary>
        [Output("enableShardsOnCoordinator")]
        public Output<bool?> EnableShardsOnCoordinator { get; private set; } = null!;

        /// <summary>
        /// Describes the identity of the cluster.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityPropertiesResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Maintenance window of a cluster.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.ServerGroupClusterMaintenanceWindowResponse?> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
        /// </summary>
        [Output("nodeCount")]
        public Output<int?> NodeCount { get; private set; } = null!;

        /// <summary>
        /// If public access is enabled on worker nodes.
        /// </summary>
        [Output("nodeEnablePublicIpAccess")]
        public Output<bool?> NodeEnablePublicIpAccess { get; private set; } = null!;

        /// <summary>
        /// The edition of a node server (default: MemoryOptimized).
        /// </summary>
        [Output("nodeServerEdition")]
        public Output<string?> NodeServerEdition { get; private set; } = null!;

        /// <summary>
        /// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Output("nodeStorageQuotaInMb")]
        public Output<int?> NodeStorageQuotaInMb { get; private set; } = null!;

        /// <summary>
        /// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Output("nodeVCores")]
        public Output<int?> NodeVCores { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the cluster was created with a password or using AAD authentication.
        /// </summary>
        [Output("passwordEnabled")]
        public Output<string> PasswordEnabled { get; private set; } = null!;

        /// <summary>
        /// Date and time in UTC (ISO8601 format) for cluster restore.
        /// </summary>
        [Output("pointInTimeUTC")]
        public Output<string?> PointInTimeUTC { get; private set; } = null!;

        /// <summary>
        /// The major PostgreSQL version on all cluster servers.
        /// </summary>
        [Output("postgresqlVersion")]
        public Output<string?> PostgresqlVersion { get; private set; } = null!;

        /// <summary>
        /// Preferred primary availability zone (AZ) for all cluster servers.
        /// </summary>
        [Output("preferredPrimaryZone")]
        public Output<string?> PreferredPrimaryZone { get; private set; } = null!;

        /// <summary>
        /// The private endpoint connections for a cluster.
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.SimplePrivateEndpointConnectionResponse>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the cluster
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The array of read replica clusters.
        /// </summary>
        [Output("readReplicas")]
        public Output<ImmutableArray<string>> ReadReplicas { get; private set; } = null!;

        /// <summary>
        /// The list of server names in the cluster
        /// </summary>
        [Output("serverNames")]
        public Output<ImmutableArray<Outputs.ServerNameItemResponse>> ServerNames { get; private set; } = null!;

        /// <summary>
        /// The Azure region of source cluster for read replica clusters.
        /// </summary>
        [Output("sourceLocation")]
        public Output<string?> SourceLocation { get; private set; } = null!;

        /// <summary>
        /// The resource id of source cluster for read replica clusters.
        /// </summary>
        [Output("sourceResourceId")]
        public Output<string?> SourceResourceId { get; private set; } = null!;

        /// <summary>
        /// A state of a cluster/server that is visible to user.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerGroupCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerGroupCluster(string name, ServerGroupClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-native:dbforpostgresql:ServerGroupCluster", name, args ?? new ServerGroupClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerGroupCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:dbforpostgresql:ServerGroupCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql/v20201005privatepreview:ServerGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql/v20201005privatepreview:ServerGroupCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql/v20221108:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql/v20221108:ServerGroupCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql/v20230302preview:Cluster" },
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql/v20230302preview:ServerGroupCluster" },
                    new global::Pulumi.Alias { Type = "azure-native:dbforpostgresql:Cluster" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerGroupCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerGroupCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerGroupCluster(name, id, options);
        }
    }

    public sealed class ServerGroupClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password of the administrator login. Required for creation.
        /// </summary>
        [Input("administratorLoginPassword")]
        public Input<string>? AdministratorLoginPassword { get; set; }

        /// <summary>
        /// Authentication configuration of a cluster.
        /// </summary>
        [Input("authConfig")]
        public Input<Inputs.ServerGroupClusterAuthConfigArgs>? AuthConfig { get; set; }

        /// <summary>
        /// The Citus extension version on all cluster servers.
        /// </summary>
        [Input("citusVersion")]
        public Input<string>? CitusVersion { get; set; }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// If public access is enabled on coordinator.
        /// </summary>
        [Input("coordinatorEnablePublicIpAccess")]
        public Input<bool>? CoordinatorEnablePublicIpAccess { get; set; }

        /// <summary>
        /// The edition of a coordinator server (default: GeneralPurpose). Required for creation.
        /// </summary>
        [Input("coordinatorServerEdition")]
        public Input<string>? CoordinatorServerEdition { get; set; }

        /// <summary>
        /// The storage of a server in MB. Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Input("coordinatorStorageQuotaInMb")]
        public Input<int>? CoordinatorStorageQuotaInMb { get; set; }

        /// <summary>
        /// The vCores count of a server (max: 96). Required for creation. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Input("coordinatorVCores")]
        public Input<int>? CoordinatorVCores { get; set; }

        /// <summary>
        /// The data encryption properties of a cluster.
        /// </summary>
        [Input("dataEncryption")]
        public Input<Inputs.ServerGroupClusterDataEncryptionArgs>? DataEncryption { get; set; }

        /// <summary>
        /// The database name of the cluster. Only one database per cluster is supported.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// If cluster backup is stored in another Azure region in addition to the copy of the backup stored in the cluster's region. Enabled only at the time of cluster creation.
        /// </summary>
        [Input("enableGeoBackup")]
        public Input<bool>? EnableGeoBackup { get; set; }

        /// <summary>
        /// If high availability (HA) is enabled or not for the cluster.
        /// </summary>
        [Input("enableHa")]
        public Input<bool>? EnableHa { get; set; }

        /// <summary>
        /// If distributed tables are placed on coordinator or not. Should be set to 'true' on single node clusters. Requires shard rebalancing after value is changed.
        /// </summary>
        [Input("enableShardsOnCoordinator")]
        public Input<bool>? EnableShardsOnCoordinator { get; set; }

        /// <summary>
        /// Describes the identity of the cluster.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityPropertiesArgs>? Identity { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Maintenance window of a cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.ServerGroupClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Worker node count of the cluster. When node count is 0, it represents a single node configuration with the ability to create distributed tables on that node. 2 or more worker nodes represent multi-node configuration. Node count value cannot be 1. Required for creation.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// If public access is enabled on worker nodes.
        /// </summary>
        [Input("nodeEnablePublicIpAccess")]
        public Input<bool>? NodeEnablePublicIpAccess { get; set; }

        /// <summary>
        /// The edition of a node server (default: MemoryOptimized).
        /// </summary>
        [Input("nodeServerEdition")]
        public Input<string>? NodeServerEdition { get; set; }

        /// <summary>
        /// The storage in MB on each worker node. See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Input("nodeStorageQuotaInMb")]
        public Input<int>? NodeStorageQuotaInMb { get; set; }

        /// <summary>
        /// The compute in vCores on each worker node (max: 104). See https://learn.microsoft.com/azure/cosmos-db/postgresql/resources-compute for more information.
        /// </summary>
        [Input("nodeVCores")]
        public Input<int>? NodeVCores { get; set; }

        /// <summary>
        /// Date and time in UTC (ISO8601 format) for cluster restore.
        /// </summary>
        [Input("pointInTimeUTC")]
        public Input<string>? PointInTimeUTC { get; set; }

        /// <summary>
        /// The major PostgreSQL version on all cluster servers.
        /// </summary>
        [Input("postgresqlVersion")]
        public Input<string>? PostgresqlVersion { get; set; }

        /// <summary>
        /// Preferred primary availability zone (AZ) for all cluster servers.
        /// </summary>
        [Input("preferredPrimaryZone")]
        public Input<string>? PreferredPrimaryZone { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Azure region of source cluster for read replica clusters.
        /// </summary>
        [Input("sourceLocation")]
        public Input<string>? SourceLocation { get; set; }

        /// <summary>
        /// The resource id of source cluster for read replica clusters.
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerGroupClusterArgs()
        {
        }
        public static new ServerGroupClusterArgs Empty => new ServerGroupClusterArgs();
    }
}
