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
    /// Definition of ReplicationPendingModifiedValues
    /// </summary>
    public sealed class ReplicationPendingModifiedValuesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The amount of storage (in gigabytes) that is allocated for the replication instance.&lt;/p&gt;
        /// </summary>
        [Input("allocatedStorage")]
        public Input<int>? AllocatedStorage { get; set; }

        /// <summary>
        /// &lt;p&gt;The engine version number of the replication instance.&lt;/p&gt;
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// &lt;p&gt; Specifies whether the replication instance is a Multi-AZ deployment. You can't set the &lt;code&gt;AvailabilityZone&lt;/code&gt; parameter if the Multi-AZ parameter is set to &lt;code&gt;true&lt;/code&gt;. &lt;/p&gt;
        /// </summary>
        [Input("multiAZ")]
        public Input<bool>? MultiAZ { get; set; }

        /// <summary>
        /// &lt;p&gt;The type of IP address protocol used by a replication instance, such as IPv4 only or Dual-stack that supports both IPv4 and IPv6 addressing. IPv6 only is not yet supported.&lt;/p&gt;
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// &lt;p&gt;The compute and memory capacity of the replication instance as defined for the specified replication instance class.&lt;/p&gt; &lt;p&gt;For more information on the settings and capacities for the available replication instance classes, see &lt;a href='https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth'&gt; Selecting the right DMS replication instance for your migration&lt;/a&gt;. &lt;/p&gt;
        /// </summary>
        [Input("replicationInstanceClass")]
        public Input<string>? ReplicationInstanceClass { get; set; }

        public ReplicationPendingModifiedValuesArgs()
        {
        }
        public static new ReplicationPendingModifiedValuesArgs Empty => new ReplicationPendingModifiedValuesArgs();
    }
}
