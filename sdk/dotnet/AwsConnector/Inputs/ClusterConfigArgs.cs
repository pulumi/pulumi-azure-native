// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of ClusterConfig
    /// </summary>
    public sealed class ClusterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;Container for cold storage configuration options.&lt;/p&gt;
        /// </summary>
        [Input("coldStorageOptions")]
        public Input<Inputs.ColdStorageOptionsArgs>? ColdStorageOptions { get; set; }

        /// <summary>
        /// &lt;p&gt;Number of dedicated master nodes in the cluster. This number must be greater than 2 and not 4, otherwise you receive a validation exception.&lt;/p&gt;
        /// </summary>
        [Input("dedicatedMasterCount")]
        public Input<int>? DedicatedMasterCount { get; set; }

        /// <summary>
        /// &lt;p&gt;Indicates whether dedicated master nodes are enabled for the cluster.&lt;code&gt;True&lt;/code&gt; if the cluster will use a dedicated master node.&lt;code&gt;False&lt;/code&gt; if the cluster will not.&lt;/p&gt;
        /// </summary>
        [Input("dedicatedMasterEnabled")]
        public Input<bool>? DedicatedMasterEnabled { get; set; }

        /// <summary>
        /// &lt;p&gt;OpenSearch Service instance type of the dedicated master nodes in the cluster.&lt;/p&gt;
        /// </summary>
        [Input("dedicatedMasterType")]
        public Input<Inputs.OpenSearchPartitionInstanceTypeEnumValueArgs>? DedicatedMasterType { get; set; }

        /// <summary>
        /// &lt;p&gt;Number of data nodes in the cluster. This number must be greater than 1, otherwise you receive a validation exception.&lt;/p&gt;
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// &lt;p&gt;Instance type of data nodes in the cluster.&lt;/p&gt;
        /// </summary>
        [Input("instanceType")]
        public Input<Inputs.OpenSearchPartitionInstanceTypeEnumValueArgs>? InstanceType { get; set; }

        /// <summary>
        /// &lt;p&gt;A boolean that indicates whether a multi-AZ domain is turned on with a standby AZ. For more information, see &lt;a href='https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html'&gt;Configuring a multi-AZ domain in Amazon OpenSearch Service&lt;/a&gt;. &lt;/p&gt;
        /// </summary>
        [Input("multiAZWithStandbyEnabled")]
        public Input<bool>? MultiAZWithStandbyEnabled { get; set; }

        /// <summary>
        /// &lt;p&gt;The number of warm nodes in the cluster.&lt;/p&gt;
        /// </summary>
        [Input("warmCount")]
        public Input<int>? WarmCount { get; set; }

        /// <summary>
        /// &lt;p&gt;Whether to enable warm storage for the cluster.&lt;/p&gt;
        /// </summary>
        [Input("warmEnabled")]
        public Input<bool>? WarmEnabled { get; set; }

        /// <summary>
        /// &lt;p&gt;The instance type for the cluster's warm nodes.&lt;/p&gt;
        /// </summary>
        [Input("warmType")]
        public Input<Inputs.OpenSearchWarmPartitionInstanceTypeEnumValueArgs>? WarmType { get; set; }

        /// <summary>
        /// &lt;p&gt;Container for zone awareness configuration options. Only required if &lt;code&gt;ZoneAwarenessEnabled&lt;/code&gt; is &lt;code&gt;true&lt;/code&gt;.&lt;/p&gt;
        /// </summary>
        [Input("zoneAwarenessConfig")]
        public Input<Inputs.ZoneAwarenessConfigArgs>? ZoneAwarenessConfig { get; set; }

        /// <summary>
        /// &lt;p&gt;Indicates whether multiple Availability Zones are enabled. For more information, see &lt;a href='https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html'&gt;Configuring a multi-AZ domain in Amazon OpenSearch Service&lt;/a&gt;.&lt;/p&gt;
        /// </summary>
        [Input("zoneAwarenessEnabled")]
        public Input<bool>? ZoneAwarenessEnabled { get; set; }

        public ClusterConfigArgs()
        {
        }
        public static new ClusterConfigArgs Empty => new ClusterConfigArgs();
    }
}
