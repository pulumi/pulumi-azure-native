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
    /// Definition of ReplicationSubnetGroup
    /// </summary>
    public sealed class ReplicationSubnetGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;A description for the replication subnet group.&lt;/p&gt;
        /// </summary>
        [Input("replicationSubnetGroupDescription")]
        public Input<string>? ReplicationSubnetGroupDescription { get; set; }

        /// <summary>
        /// &lt;p&gt;The identifier of the replication instance subnet group.&lt;/p&gt;
        /// </summary>
        [Input("replicationSubnetGroupIdentifier")]
        public Input<string>? ReplicationSubnetGroupIdentifier { get; set; }

        /// <summary>
        /// &lt;p&gt;The status of the subnet group.&lt;/p&gt;
        /// </summary>
        [Input("subnetGroupStatus")]
        public Input<string>? SubnetGroupStatus { get; set; }

        [Input("subnets")]
        private InputList<Inputs.SubnetArgs>? _subnets;

        /// <summary>
        /// &lt;p&gt;The subnets that are in the subnet group.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.SubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.SubnetArgs>());
            set => _subnets = value;
        }

        [Input("supportedNetworkTypes")]
        private InputList<string>? _supportedNetworkTypes;

        /// <summary>
        /// &lt;p&gt;The IP addressing protocol supported by the subnet group. This is used by a replication instance with values such as IPv4 only or Dual-stack that supports both IPv4 and IPv6 addressing. IPv6 only is not yet supported.&lt;/p&gt;
        /// </summary>
        public InputList<string> SupportedNetworkTypes
        {
            get => _supportedNetworkTypes ?? (_supportedNetworkTypes = new InputList<string>());
            set => _supportedNetworkTypes = value;
        }

        /// <summary>
        /// &lt;p&gt;The ID of the VPC.&lt;/p&gt;
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ReplicationSubnetGroupArgs()
        {
        }
        public static new ReplicationSubnetGroupArgs Empty => new ReplicationSubnetGroupArgs();
    }
}
