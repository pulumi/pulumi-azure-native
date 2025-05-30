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
    /// Definition of awsEmrCluster
    /// </summary>
    [OutputType]
    public sealed class AwsEmrClusterPropertiesResponse
    {
        /// <summary>
        /// &lt;p&gt;The applications installed on this cluster.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationResponse> Applications;
        /// <summary>
        /// &lt;p&gt;An IAM role for automatic scaling policies. The default role is &lt;code&gt;EMR_AutoScaling_DefaultRole&lt;/code&gt;. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate Amazon EC2 instances in an instance group.&lt;/p&gt;
        /// </summary>
        public readonly string? AutoScalingRole;
        /// <summary>
        /// &lt;p&gt;Specifies whether the cluster should terminate after completing all steps.&lt;/p&gt;
        /// </summary>
        public readonly bool? AutoTerminate;
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name of the cluster.&lt;/p&gt;
        /// </summary>
        public readonly string? ClusterArn;
        /// <summary>
        /// &lt;p&gt;Applies only to Amazon EMR releases 4.x and later. The list of configurations that are supplied to the Amazon EMR cluster.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigurationResponse> Configurations;
        /// <summary>
        /// &lt;p&gt;Available only in Amazon EMR releases 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.&lt;/p&gt;
        /// </summary>
        public readonly string? CustomAmiId;
        /// <summary>
        /// &lt;p&gt;The IOPS, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance. Available in Amazon EMR releases 6.15.0 and later.&lt;/p&gt;
        /// </summary>
        public readonly int? EbsRootVolumeIops;
        /// <summary>
        /// &lt;p&gt;The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance. Available in Amazon EMR releases 4.x and later.&lt;/p&gt;
        /// </summary>
        public readonly int? EbsRootVolumeSize;
        /// <summary>
        /// &lt;p&gt;The throughput, in MiB/s, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance. Available in Amazon EMR releases 6.15.0 and later.&lt;/p&gt;
        /// </summary>
        public readonly int? EbsRootVolumeThroughput;
        /// <summary>
        /// &lt;p&gt;Provides information about the Amazon EC2 instances in a cluster grouped by category. For example, key name, subnet ID, IAM instance profile, and so on.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.Ec2InstanceAttributesResponse? Ec2InstanceAttributes;
        /// <summary>
        /// &lt;p&gt;The unique identifier for the cluster.&lt;/p&gt;
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// &lt;note&gt; &lt;p&gt;The instance fleet configuration is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x versions.&lt;/p&gt; &lt;/note&gt; &lt;p&gt;The instance group configuration of the cluster. A value of &lt;code&gt;INSTANCE_GROUP&lt;/code&gt; indicates a uniform instance group configuration. A value of &lt;code&gt;INSTANCE_FLEET&lt;/code&gt; indicates an instance fleets configuration.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.InstanceCollectionTypeEnumValueResponse? InstanceCollectionType;
        /// <summary>
        /// &lt;p&gt;Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration. For more information see &lt;a href='https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html'&gt;Use Kerberos Authentication&lt;/a&gt; in the &lt;i&gt;Amazon EMR Management Guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.KerberosAttributesResponse? KerberosAttributes;
        /// <summary>
        /// &lt;p&gt; The KMS key used for encrypting log files. This attribute is only available with Amazon EMR 5.30.0 and later, excluding Amazon EMR 6.0.0. &lt;/p&gt;
        /// </summary>
        public readonly string? LogEncryptionKmsKeyId;
        /// <summary>
        /// &lt;p&gt;The path to the Amazon S3 location where logs for this cluster are stored.&lt;/p&gt;
        /// </summary>
        public readonly string? LogUri;
        /// <summary>
        /// &lt;p&gt;The DNS name of the master node. If the cluster is on a private subnet, this is the private DNS name. On a public subnet, this is the public DNS name.&lt;/p&gt;
        /// </summary>
        public readonly string? MasterPublicDnsName;
        /// <summary>
        /// &lt;p&gt;The name of the cluster. This parameter can't contain the characters &amp;lt;, &amp;gt;, $, |, or ` (backtick).&lt;/p&gt;
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// &lt;p&gt;An approximation of the cost of the cluster, represented in m1.small/hours. This value is incremented one time for every hour an m1.small instance runs. Larger instances are weighted more, so an Amazon EC2 instance that is roughly four times more expensive would result in the normalized instance hours being incremented by four. This result is only an approximation and does not reflect the actual billing rate.&lt;/p&gt;
        /// </summary>
        public readonly int? NormalizedInstanceHours;
        /// <summary>
        /// &lt;p&gt;The Amazon Linux release specified in a cluster launch RunJobFlow request. If no Amazon Linux release was specified, the default Amazon Linux release is shown in the response.&lt;/p&gt;
        /// </summary>
        public readonly string? OsReleaseLabel;
        /// <summary>
        /// &lt;p&gt; The Amazon Resource Name (ARN) of the Outpost where the cluster is launched. &lt;/p&gt;
        /// </summary>
        public readonly string? OutpostArn;
        /// <summary>
        /// &lt;p&gt;Placement group configured for an Amazon EMR cluster.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.PlacementGroupConfigResponse> PlacementGroups;
        /// <summary>
        /// &lt;p&gt;The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster. Release labels are in the form &lt;code&gt;emr-x.x.x&lt;/code&gt;, where x.x.x is an Amazon EMR release version such as &lt;code&gt;emr-5.14.0&lt;/code&gt;. For more information about Amazon EMR release versions and included application versions and features, see &lt;a href='https://docs.aws.amazon.com/emr/latest/ReleaseGuide/'&gt;https://docs.aws.amazon.com/emr/latest/ReleaseGuide/&lt;/a&gt;. The release label applies only to Amazon EMR releases version 4.0 and later. Earlier versions use &lt;code&gt;AmiVersion&lt;/code&gt;.&lt;/p&gt;
        /// </summary>
        public readonly string? ReleaseLabel;
        /// <summary>
        /// &lt;p&gt;Applies only when &lt;code&gt;CustomAmiID&lt;/code&gt; is used. Specifies the type of updates that the Amazon Linux AMI package repositories apply when an instance boots using the AMI.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.RepoUpgradeOnBootEnumValueResponse? RepoUpgradeOnBoot;
        /// <summary>
        /// &lt;p&gt;The AMI version requested for this cluster.&lt;/p&gt;
        /// </summary>
        public readonly string? RequestedAmiVersion;
        /// <summary>
        /// &lt;p&gt;The AMI version running on this cluster.&lt;/p&gt;
        /// </summary>
        public readonly string? RunningAmiVersion;
        /// <summary>
        /// &lt;p&gt;The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized. &lt;code&gt;TERMINATE_AT_INSTANCE_HOUR&lt;/code&gt; indicates that Amazon EMR terminates nodes at the instance-hour boundary, regardless of when the request to terminate the instance was submitted. This option is only available with Amazon EMR 5.1.0 and later and is the default for clusters created using that version. &lt;code&gt;TERMINATE_AT_TASK_COMPLETION&lt;/code&gt; indicates that Amazon EMR adds nodes to a deny list and drains tasks from nodes before terminating the Amazon EC2 instances, regardless of the instance-hour boundary. With either behavior, Amazon EMR removes the least active nodes first and blocks instance termination if it could lead to HDFS corruption. &lt;code&gt;TERMINATE_AT_TASK_COMPLETION&lt;/code&gt; is available only in Amazon EMR releases 4.1.0 and later, and is the default for versions of Amazon EMR earlier than 5.1.0.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.ScaleDownBehaviorEnumValueResponse? ScaleDownBehavior;
        /// <summary>
        /// &lt;p&gt;The name of the security configuration applied to the cluster.&lt;/p&gt;
        /// </summary>
        public readonly string? SecurityConfiguration;
        /// <summary>
        /// &lt;p&gt;The IAM role that Amazon EMR assumes in order to access Amazon Web Services resources on your behalf.&lt;/p&gt;
        /// </summary>
        public readonly string? ServiceRole;
        /// <summary>
        /// &lt;p&gt;The current status details about the cluster.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.EmrClusterStatusResponse? Status;
        /// <summary>
        /// &lt;p&gt;Specifies the number of steps that can be executed concurrently.&lt;/p&gt;
        /// </summary>
        public readonly int? StepConcurrencyLevel;
        /// <summary>
        /// &lt;p&gt;A list of tags associated with a cluster.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> Tags;
        /// <summary>
        /// &lt;p&gt;Indicates whether Amazon EMR will lock the cluster to prevent the Amazon EC2 instances from being terminated by an API call or user intervention, or in the event of a cluster error.&lt;/p&gt;
        /// </summary>
        public readonly bool? TerminationProtected;
        /// <summary>
        /// &lt;p&gt;Indicates whether Amazon EMR should gracefully replace Amazon EC2 core instances that have degraded within the cluster.&lt;/p&gt;
        /// </summary>
        public readonly bool? UnhealthyNodeReplacement;
        /// <summary>
        /// &lt;p&gt;Indicates whether the cluster is visible to IAM principals in the Amazon Web Services account associated with the cluster. When &lt;code&gt;true&lt;/code&gt;, IAM principals in the Amazon Web Services account can perform Amazon EMR cluster actions on the cluster that their IAM policies allow. When &lt;code&gt;false&lt;/code&gt;, only the IAM principal that created the cluster and the Amazon Web Services account root user can perform Amazon EMR actions, regardless of IAM permissions policies attached to other IAM principals.&lt;/p&gt; &lt;p&gt;The default value is &lt;code&gt;true&lt;/code&gt; if a value is not provided when creating a cluster using the Amazon EMR API &lt;a&gt;RunJobFlow&lt;/a&gt; command, the CLI &lt;a href='https://docs.aws.amazon.com/cli/latest/reference/emr/create-cluster.html'&gt;create-cluster&lt;/a&gt; command, or the Amazon Web Services Management Console.&lt;/p&gt;
        /// </summary>
        public readonly bool? VisibleToAllUsers;

        [OutputConstructor]
        private AwsEmrClusterPropertiesResponse(
            ImmutableArray<Outputs.ApplicationResponse> applications,

            string? autoScalingRole,

            bool? autoTerminate,

            string? clusterArn,

            ImmutableArray<Outputs.ConfigurationResponse> configurations,

            string? customAmiId,

            int? ebsRootVolumeIops,

            int? ebsRootVolumeSize,

            int? ebsRootVolumeThroughput,

            Outputs.Ec2InstanceAttributesResponse? ec2InstanceAttributes,

            string? id,

            Outputs.InstanceCollectionTypeEnumValueResponse? instanceCollectionType,

            Outputs.KerberosAttributesResponse? kerberosAttributes,

            string? logEncryptionKmsKeyId,

            string? logUri,

            string? masterPublicDnsName,

            string? name,

            int? normalizedInstanceHours,

            string? osReleaseLabel,

            string? outpostArn,

            ImmutableArray<Outputs.PlacementGroupConfigResponse> placementGroups,

            string? releaseLabel,

            Outputs.RepoUpgradeOnBootEnumValueResponse? repoUpgradeOnBoot,

            string? requestedAmiVersion,

            string? runningAmiVersion,

            Outputs.ScaleDownBehaviorEnumValueResponse? scaleDownBehavior,

            string? securityConfiguration,

            string? serviceRole,

            Outputs.EmrClusterStatusResponse? status,

            int? stepConcurrencyLevel,

            ImmutableArray<Outputs.TagResponse> tags,

            bool? terminationProtected,

            bool? unhealthyNodeReplacement,

            bool? visibleToAllUsers)
        {
            Applications = applications;
            AutoScalingRole = autoScalingRole;
            AutoTerminate = autoTerminate;
            ClusterArn = clusterArn;
            Configurations = configurations;
            CustomAmiId = customAmiId;
            EbsRootVolumeIops = ebsRootVolumeIops;
            EbsRootVolumeSize = ebsRootVolumeSize;
            EbsRootVolumeThroughput = ebsRootVolumeThroughput;
            Ec2InstanceAttributes = ec2InstanceAttributes;
            Id = id;
            InstanceCollectionType = instanceCollectionType;
            KerberosAttributes = kerberosAttributes;
            LogEncryptionKmsKeyId = logEncryptionKmsKeyId;
            LogUri = logUri;
            MasterPublicDnsName = masterPublicDnsName;
            Name = name;
            NormalizedInstanceHours = normalizedInstanceHours;
            OsReleaseLabel = osReleaseLabel;
            OutpostArn = outpostArn;
            PlacementGroups = placementGroups;
            ReleaseLabel = releaseLabel;
            RepoUpgradeOnBoot = repoUpgradeOnBoot;
            RequestedAmiVersion = requestedAmiVersion;
            RunningAmiVersion = runningAmiVersion;
            ScaleDownBehavior = scaleDownBehavior;
            SecurityConfiguration = securityConfiguration;
            ServiceRole = serviceRole;
            Status = status;
            StepConcurrencyLevel = stepConcurrencyLevel;
            Tags = tags;
            TerminationProtected = terminationProtected;
            UnhealthyNodeReplacement = unhealthyNodeReplacement;
            VisibleToAllUsers = visibleToAllUsers;
        }
    }
}
