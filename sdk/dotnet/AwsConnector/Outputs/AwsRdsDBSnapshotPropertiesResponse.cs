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
    /// Definition of awsRdsDBSnapshot
    /// </summary>
    [OutputType]
    public sealed class AwsRdsDBSnapshotPropertiesResponse
    {
        /// <summary>
        /// &lt;p&gt;Specifies the allocated storage size in gibibytes (GiB).&lt;/p&gt;
        /// </summary>
        public readonly int? AllocatedStorage;
        /// <summary>
        /// &lt;p&gt;Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// &lt;p&gt;Specifies the DB instance identifier of the DB instance this DB snapshot was created from.&lt;/p&gt;
        /// </summary>
        public readonly string? DbInstanceIdentifier;
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) for the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? DbSnapshotArn;
        /// <summary>
        /// &lt;p&gt;Specifies the identifier for the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? DbSnapshotIdentifier;
        /// <summary>
        /// &lt;p&gt;The Oracle system identifier (SID), which is the name of the Oracle database instance that manages your database files. The Oracle SID is also the name of your CDB.&lt;/p&gt;
        /// </summary>
        public readonly string? DbSystemId;
        /// <summary>
        /// &lt;p&gt;The identifier for the source DB instance, which can't be changed and which is unique to an Amazon Web Services Region.&lt;/p&gt;
        /// </summary>
        public readonly string? DbiResourceId;
        /// <summary>
        /// &lt;p&gt;Indicates whether the DB instance has a dedicated log volume (DLV) enabled.&lt;/p&gt;
        /// </summary>
        public readonly bool? DedicatedLogVolume;
        /// <summary>
        /// &lt;p&gt;Indicates whether the DB snapshot is encrypted.&lt;/p&gt;
        /// </summary>
        public readonly bool? Encrypted;
        /// <summary>
        /// &lt;p&gt;Specifies the name of the database engine.&lt;/p&gt;
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// &lt;p&gt;Specifies the version of the database engine.&lt;/p&gt;
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// &lt;p&gt;Indicates whether mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts is enabled.&lt;/p&gt;
        /// </summary>
        public readonly bool? IamDatabaseAuthenticationEnabled;
        /// <summary>
        /// &lt;p&gt;Specifies the time in Coordinated Universal Time (UTC) when the DB instance, from which the snapshot was taken, was created.&lt;/p&gt;
        /// </summary>
        public readonly string? InstanceCreateTime;
        /// <summary>
        /// &lt;p&gt;Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.&lt;/p&gt;
        /// </summary>
        public readonly int? Iops;
        /// <summary>
        /// &lt;p&gt;If &lt;code&gt;Encrypted&lt;/code&gt; is true, the Amazon Web Services KMS key identifier for the encrypted DB snapshot.&lt;/p&gt; &lt;p&gt;The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.&lt;/p&gt;
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// &lt;p&gt;License model information for the restored DB instance.&lt;/p&gt;
        /// </summary>
        public readonly string? LicenseModel;
        /// <summary>
        /// &lt;p&gt;Provides the master username for the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? MasterUsername;
        /// <summary>
        /// &lt;p&gt;Indicates whether the snapshot is of a DB instance using the multi-tenant configuration (TRUE) or the single-tenant configuration (FALSE).&lt;/p&gt;
        /// </summary>
        public readonly bool? MultiTenant;
        /// <summary>
        /// &lt;p&gt;Provides the option group name for the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? OptionGroupName;
        /// <summary>
        /// &lt;p&gt;Specifies the time of the CreateDBSnapshot operation in Coordinated Universal Time (UTC). Doesn't change when the snapshot is copied.&lt;/p&gt;
        /// </summary>
        public readonly string? OriginalSnapshotCreateTime;
        /// <summary>
        /// &lt;p&gt;The percentage of the estimated data that has been transferred.&lt;/p&gt;
        /// </summary>
        public readonly int? PercentProgress;
        /// <summary>
        /// &lt;p&gt;Specifies the port that the database engine was listening on at the time of the snapshot.&lt;/p&gt;
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// &lt;p&gt;The number of CPU cores and the number of threads per core for the DB instance class of the DB instance when the DB snapshot was created.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.ProcessorFeatureResponse> ProcessorFeatures;
        /// <summary>
        /// &lt;p&gt;Specifies when the snapshot was taken in Coordinated Universal Time (UTC). Changes for the copy when the snapshot is copied.&lt;/p&gt;
        /// </summary>
        public readonly string? SnapshotCreateTime;
        /// <summary>
        /// &lt;p&gt;The timestamp of the most recent transaction applied to the database that you're backing up. Thus, if you restore a snapshot, SnapshotDatabaseTime is the most recent transaction in the restored DB instance. In contrast, originalSnapshotCreateTime specifies the system time that the snapshot completed.&lt;/p&gt; &lt;p&gt;If you back up a read replica, you can determine the replica lag by comparing SnapshotDatabaseTime with originalSnapshotCreateTime. For example, if originalSnapshotCreateTime is two hours later than SnapshotDatabaseTime, then the replica lag is two hours.&lt;/p&gt;
        /// </summary>
        public readonly string? SnapshotDatabaseTime;
        /// <summary>
        /// &lt;p&gt;Specifies where manual snapshots are stored: Amazon Web Services Outposts or the Amazon Web Services Region.&lt;/p&gt;
        /// </summary>
        public readonly string? SnapshotTarget;
        /// <summary>
        /// &lt;p&gt;Provides the type of the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? SnapshotType;
        /// <summary>
        /// &lt;p&gt;The DB snapshot Amazon Resource Name (ARN) that the DB snapshot was copied from. It only has a value in the case of a cross-account or cross-Region copy.&lt;/p&gt;
        /// </summary>
        public readonly string? SourceDBSnapshotIdentifier;
        /// <summary>
        /// &lt;p&gt;The Amazon Web Services Region that the DB snapshot was created in or copied from.&lt;/p&gt;
        /// </summary>
        public readonly string? SourceRegion;
        /// <summary>
        /// &lt;p&gt;Specifies the status of this DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// &lt;p&gt;Specifies the storage throughput for the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly int? StorageThroughput;
        /// <summary>
        /// &lt;p&gt;Specifies the storage type associated with DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? StorageType;
        /// <summary>
        /// Property tagList
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> TagList;
        /// <summary>
        /// &lt;p&gt;The ARN from the key store with which to associate the instance for TDE encryption.&lt;/p&gt;
        /// </summary>
        public readonly string? TdeCredentialArn;
        /// <summary>
        /// &lt;p&gt;The time zone of the DB snapshot. In most cases, the &lt;code&gt;Timezone&lt;/code&gt; element is empty. &lt;code&gt;Timezone&lt;/code&gt; content appears only for snapshots taken from Microsoft SQL Server DB instances that were created with a time zone specified.&lt;/p&gt;
        /// </summary>
        public readonly string? Timezone;
        /// <summary>
        /// &lt;p&gt;Provides the VPC ID associated with the DB snapshot.&lt;/p&gt;
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private AwsRdsDBSnapshotPropertiesResponse(
            int? allocatedStorage,

            string? availabilityZone,

            string? dbInstanceIdentifier,

            string? dbSnapshotArn,

            string? dbSnapshotIdentifier,

            string? dbSystemId,

            string? dbiResourceId,

            bool? dedicatedLogVolume,

            bool? encrypted,

            string? engine,

            string? engineVersion,

            bool? iamDatabaseAuthenticationEnabled,

            string? instanceCreateTime,

            int? iops,

            string? kmsKeyId,

            string? licenseModel,

            string? masterUsername,

            bool? multiTenant,

            string? optionGroupName,

            string? originalSnapshotCreateTime,

            int? percentProgress,

            int? port,

            ImmutableArray<Outputs.ProcessorFeatureResponse> processorFeatures,

            string? snapshotCreateTime,

            string? snapshotDatabaseTime,

            string? snapshotTarget,

            string? snapshotType,

            string? sourceDBSnapshotIdentifier,

            string? sourceRegion,

            string? status,

            int? storageThroughput,

            string? storageType,

            ImmutableArray<Outputs.TagResponse> tagList,

            string? tdeCredentialArn,

            string? timezone,

            string? vpcId)
        {
            AllocatedStorage = allocatedStorage;
            AvailabilityZone = availabilityZone;
            DbInstanceIdentifier = dbInstanceIdentifier;
            DbSnapshotArn = dbSnapshotArn;
            DbSnapshotIdentifier = dbSnapshotIdentifier;
            DbSystemId = dbSystemId;
            DbiResourceId = dbiResourceId;
            DedicatedLogVolume = dedicatedLogVolume;
            Encrypted = encrypted;
            Engine = engine;
            EngineVersion = engineVersion;
            IamDatabaseAuthenticationEnabled = iamDatabaseAuthenticationEnabled;
            InstanceCreateTime = instanceCreateTime;
            Iops = iops;
            KmsKeyId = kmsKeyId;
            LicenseModel = licenseModel;
            MasterUsername = masterUsername;
            MultiTenant = multiTenant;
            OptionGroupName = optionGroupName;
            OriginalSnapshotCreateTime = originalSnapshotCreateTime;
            PercentProgress = percentProgress;
            Port = port;
            ProcessorFeatures = processorFeatures;
            SnapshotCreateTime = snapshotCreateTime;
            SnapshotDatabaseTime = snapshotDatabaseTime;
            SnapshotTarget = snapshotTarget;
            SnapshotType = snapshotType;
            SourceDBSnapshotIdentifier = sourceDBSnapshotIdentifier;
            SourceRegion = sourceRegion;
            Status = status;
            StorageThroughput = storageThroughput;
            StorageType = storageType;
            TagList = tagList;
            TdeCredentialArn = tdeCredentialArn;
            Timezone = timezone;
            VpcId = vpcId;
        }
    }
}
