// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of awsRedshiftCluster
    /// </summary>
    [OutputType]
    public sealed class AwsRedshiftClusterPropertiesResponse
    {
        /// <summary>
        /// Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True
        /// </summary>
        public readonly bool? AllowVersionUpgrade;
        /// <summary>
        /// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.disabled - Don't use AQUA.auto - Amazon Redshift determines whether to use AQUA.
        /// </summary>
        public readonly string? AquaConfigurationStatus;
        /// <summary>
        /// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1
        /// </summary>
        public readonly int? AutomatedSnapshotRetentionPeriod;
        /// <summary>
        /// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
        /// </summary>
        public readonly bool? AvailabilityZoneRelocation;
        /// <summary>
        /// The availability zone relocation status of the cluster
        /// </summary>
        public readonly string? AvailabilityZoneRelocationStatus;
        /// <summary>
        /// A boolean value indicating whether the resize operation is using the classic resize process. If you don't provide this parameter or set the value to false , the resize type is elastic.
        /// </summary>
        public readonly bool? Classic;
        /// <summary>
        /// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
        /// </summary>
        public readonly string? ClusterIdentifier;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster namespace.
        /// </summary>
        public readonly string? ClusterNamespaceArn;
        /// <summary>
        /// The name of the parameter group to be associated with this cluster.
        /// </summary>
        public readonly string? ClusterParameterGroupName;
        /// <summary>
        /// A list of security groups to be associated with this cluster.
        /// </summary>
        public readonly ImmutableArray<string> ClusterSecurityGroups;
        /// <summary>
        /// The name of a cluster subnet group to be associated with this cluster.
        /// </summary>
        public readonly string? ClusterSubnetGroupName;
        /// <summary>
        /// The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
        /// </summary>
        public readonly string? ClusterType;
        /// <summary>
        /// The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
        /// </summary>
        public readonly string? ClusterVersion;
        /// <summary>
        /// The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.
        /// </summary>
        public readonly string? DbName;
        /// <summary>
        /// A boolean indicating whether to enable the deferred maintenance window.
        /// </summary>
        public readonly bool? DeferMaintenance;
        /// <summary>
        /// An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time. The duration must be 45 days or less.
        /// </summary>
        public readonly int? DeferMaintenanceDuration;
        /// <summary>
        /// A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.
        /// </summary>
        public readonly string? DeferMaintenanceEndTime;
        /// <summary>
        /// A unique identifier for the deferred maintenance window.
        /// </summary>
        public readonly string? DeferMaintenanceIdentifier;
        /// <summary>
        /// A timestamp indicating the start time for the deferred maintenance window.
        /// </summary>
        public readonly string? DeferMaintenanceStartTime;
        /// <summary>
        /// The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
        /// </summary>
        public readonly string? DestinationRegion;
        /// <summary>
        /// The Elastic IP (EIP) address for the cluster.
        /// </summary>
        public readonly string? ElasticIp;
        /// <summary>
        /// If true, the data in the cluster is encrypted at rest.
        /// </summary>
        public readonly bool? Encrypted;
        /// <summary>
        /// Property endpoint
        /// </summary>
        public readonly Outputs.RedshiftClusterEndpointResponse? Endpoint;
        /// <summary>
        /// An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.If this option is true , enhanced VPC routing is enabled.Default: false
        /// </summary>
        public readonly bool? EnhancedVpcRouting;
        /// <summary>
        /// Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM
        /// </summary>
        public readonly string? HsmClientCertificateIdentifier;
        /// <summary>
        /// Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
        /// </summary>
        public readonly string? HsmConfigurationIdentifier;
        /// <summary>
        /// A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request
        /// </summary>
        public readonly ImmutableArray<string> IamRoles;
        /// <summary>
        /// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// Property loggingProperties
        /// </summary>
        public readonly Outputs.LoggingPropertiesResponse? LoggingProperties;
        /// <summary>
        /// The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
        /// </summary>
        public readonly string? MaintenanceTrackName;
        /// <summary>
        /// A boolean indicating if the redshift cluster's admin user credentials is managed by Redshift or not. You can't use MasterUserPassword if ManageMasterPassword is true. If ManageMasterPassword is false or not set, Amazon Redshift uses MasterUserPassword for the admin user account's password.
        /// </summary>
        public readonly bool? ManageMasterPassword;
        /// <summary>
        /// The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.The value must be either -1 or an integer between 1 and 3,653.
        /// </summary>
        public readonly int? ManualSnapshotRetentionPeriod;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the cluster's admin user credentials secret.
        /// </summary>
        public readonly string? MasterPasswordSecretArn;
        /// <summary>
        /// The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret.
        /// </summary>
        public readonly string? MasterPasswordSecretKmsKeyId;
        /// <summary>
        /// The password associated with the master user account for the cluster that is being created. You can't use MasterUserPassword if ManageMasterPassword is true. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.
        /// </summary>
        public readonly string? MasterUserPassword;
        /// <summary>
        /// The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.
        /// </summary>
        public readonly string? MasterUsername;
        /// <summary>
        /// A boolean indicating if the redshift cluster is multi-az or not. If you don't provide this parameter or set the value to false, the redshift cluster will be single-az.
        /// </summary>
        public readonly bool? MultiAZ;
        /// <summary>
        /// The namespace resource policy document that will be attached to a Redshift cluster.
        /// </summary>
        public readonly object? NamespaceResourcePolicy;
        /// <summary>
        /// The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge | ra3.16xlarge
        /// </summary>
        public readonly string? NodeType;
        /// <summary>
        /// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.
        /// </summary>
        public readonly int? NumberOfNodes;
        /// <summary>
        /// Property ownerAccount
        /// </summary>
        public readonly string? OwnerAccount;
        /// <summary>
        /// The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The weekly time range (in UTC) during which automated cluster maintenance can occur.
        /// </summary>
        public readonly string? PreferredMaintenanceWindow;
        /// <summary>
        /// If true, the cluster can be accessed from a public network.
        /// </summary>
        public readonly bool? PubliclyAccessible;
        /// <summary>
        /// The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs
        /// </summary>
        public readonly string? ResourceAction;
        /// <summary>
        /// The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
        /// </summary>
        public readonly string? RevisionTarget;
        /// <summary>
        /// A boolean indicating if we want to rotate Encryption Keys.
        /// </summary>
        public readonly bool? RotateEncryptionKey;
        /// <summary>
        /// The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
        /// </summary>
        public readonly string? SnapshotClusterIdentifier;
        /// <summary>
        /// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
        /// </summary>
        public readonly string? SnapshotCopyGrantName;
        /// <summary>
        /// Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
        /// </summary>
        public readonly bool? SnapshotCopyManual;
        /// <summary>
        /// The number of days to retain automated snapshots in the destination region after they are copied from the source region.  Default is 7.  Constraints: Must be at least 1 and no more than 35.
        /// </summary>
        public readonly int? SnapshotCopyRetentionPeriod;
        /// <summary>
        /// The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
        /// </summary>
        public readonly string? SnapshotIdentifier;
        /// <summary>
        /// The list of tags for the cluster parameter group.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> Tags;
        /// <summary>
        /// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private AwsRedshiftClusterPropertiesResponse(
            bool? allowVersionUpgrade,

            string? aquaConfigurationStatus,

            int? automatedSnapshotRetentionPeriod,

            string? availabilityZone,

            bool? availabilityZoneRelocation,

            string? availabilityZoneRelocationStatus,

            bool? classic,

            string? clusterIdentifier,

            string? clusterNamespaceArn,

            string? clusterParameterGroupName,

            ImmutableArray<string> clusterSecurityGroups,

            string? clusterSubnetGroupName,

            string? clusterType,

            string? clusterVersion,

            string? dbName,

            bool? deferMaintenance,

            int? deferMaintenanceDuration,

            string? deferMaintenanceEndTime,

            string? deferMaintenanceIdentifier,

            string? deferMaintenanceStartTime,

            string? destinationRegion,

            string? elasticIp,

            bool? encrypted,

            Outputs.RedshiftClusterEndpointResponse? endpoint,

            bool? enhancedVpcRouting,

            string? hsmClientCertificateIdentifier,

            string? hsmConfigurationIdentifier,

            ImmutableArray<string> iamRoles,

            string? kmsKeyId,

            Outputs.LoggingPropertiesResponse? loggingProperties,

            string? maintenanceTrackName,

            bool? manageMasterPassword,

            int? manualSnapshotRetentionPeriod,

            string? masterPasswordSecretArn,

            string? masterPasswordSecretKmsKeyId,

            string? masterUserPassword,

            string? masterUsername,

            bool? multiAZ,

            object? namespaceResourcePolicy,

            string? nodeType,

            int? numberOfNodes,

            string? ownerAccount,

            int? port,

            string? preferredMaintenanceWindow,

            bool? publiclyAccessible,

            string? resourceAction,

            string? revisionTarget,

            bool? rotateEncryptionKey,

            string? snapshotClusterIdentifier,

            string? snapshotCopyGrantName,

            bool? snapshotCopyManual,

            int? snapshotCopyRetentionPeriod,

            string? snapshotIdentifier,

            ImmutableArray<Outputs.TagResponse> tags,

            ImmutableArray<string> vpcSecurityGroupIds)
        {
            AllowVersionUpgrade = allowVersionUpgrade;
            AquaConfigurationStatus = aquaConfigurationStatus;
            AutomatedSnapshotRetentionPeriod = automatedSnapshotRetentionPeriod;
            AvailabilityZone = availabilityZone;
            AvailabilityZoneRelocation = availabilityZoneRelocation;
            AvailabilityZoneRelocationStatus = availabilityZoneRelocationStatus;
            Classic = classic;
            ClusterIdentifier = clusterIdentifier;
            ClusterNamespaceArn = clusterNamespaceArn;
            ClusterParameterGroupName = clusterParameterGroupName;
            ClusterSecurityGroups = clusterSecurityGroups;
            ClusterSubnetGroupName = clusterSubnetGroupName;
            ClusterType = clusterType;
            ClusterVersion = clusterVersion;
            DbName = dbName;
            DeferMaintenance = deferMaintenance;
            DeferMaintenanceDuration = deferMaintenanceDuration;
            DeferMaintenanceEndTime = deferMaintenanceEndTime;
            DeferMaintenanceIdentifier = deferMaintenanceIdentifier;
            DeferMaintenanceStartTime = deferMaintenanceStartTime;
            DestinationRegion = destinationRegion;
            ElasticIp = elasticIp;
            Encrypted = encrypted;
            Endpoint = endpoint;
            EnhancedVpcRouting = enhancedVpcRouting;
            HsmClientCertificateIdentifier = hsmClientCertificateIdentifier;
            HsmConfigurationIdentifier = hsmConfigurationIdentifier;
            IamRoles = iamRoles;
            KmsKeyId = kmsKeyId;
            LoggingProperties = loggingProperties;
            MaintenanceTrackName = maintenanceTrackName;
            ManageMasterPassword = manageMasterPassword;
            ManualSnapshotRetentionPeriod = manualSnapshotRetentionPeriod;
            MasterPasswordSecretArn = masterPasswordSecretArn;
            MasterPasswordSecretKmsKeyId = masterPasswordSecretKmsKeyId;
            MasterUserPassword = masterUserPassword;
            MasterUsername = masterUsername;
            MultiAZ = multiAZ;
            NamespaceResourcePolicy = namespaceResourcePolicy;
            NodeType = nodeType;
            NumberOfNodes = numberOfNodes;
            OwnerAccount = ownerAccount;
            Port = port;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            PubliclyAccessible = publiclyAccessible;
            ResourceAction = resourceAction;
            RevisionTarget = revisionTarget;
            RotateEncryptionKey = rotateEncryptionKey;
            SnapshotClusterIdentifier = snapshotClusterIdentifier;
            SnapshotCopyGrantName = snapshotCopyGrantName;
            SnapshotCopyManual = snapshotCopyManual;
            SnapshotCopyRetentionPeriod = snapshotCopyRetentionPeriod;
            SnapshotIdentifier = snapshotIdentifier;
            Tags = tags;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
        }
    }
}
