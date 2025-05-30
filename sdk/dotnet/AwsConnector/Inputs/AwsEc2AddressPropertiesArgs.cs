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
    /// Definition of awsEc2Address
    /// </summary>
    public sealed class AwsEc2AddressPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The ID representing the allocation of the address.&lt;/p&gt;
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID representing the association of the address with an instance.&lt;/p&gt;
        /// </summary>
        [Input("associationId")]
        public Input<string>? AssociationId { get; set; }

        /// <summary>
        /// &lt;p&gt;The carrier IP address associated. This option is only available for network interfaces which reside in a subnet in a Wavelength Zone (for example an EC2 instance). &lt;/p&gt;
        /// </summary>
        [Input("carrierIp")]
        public Input<string>? CarrierIp { get; set; }

        /// <summary>
        /// &lt;p&gt;The customer-owned IP address.&lt;/p&gt;
        /// </summary>
        [Input("customerOwnedIp")]
        public Input<string>? CustomerOwnedIp { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID of the customer-owned address pool.&lt;/p&gt;
        /// </summary>
        [Input("customerOwnedIpv4Pool")]
        public Input<string>? CustomerOwnedIpv4Pool { get; set; }

        /// <summary>
        /// &lt;p&gt;The network (&lt;code&gt;vpc&lt;/code&gt;).&lt;/p&gt;
        /// </summary>
        [Input("domain")]
        public Input<Inputs.DomainTypeEnumValueArgs>? Domain { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID of the instance that the address is associated with (if any).&lt;/p&gt;
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// &lt;p&gt;The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses.&lt;/p&gt;
        /// </summary>
        [Input("networkBorderGroup")]
        public Input<string>? NetworkBorderGroup { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID of the network interface.&lt;/p&gt;
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID of the Amazon Web Services account that owns the network interface.&lt;/p&gt;
        /// </summary>
        [Input("networkInterfaceOwnerId")]
        public Input<string>? NetworkInterfaceOwnerId { get; set; }

        /// <summary>
        /// &lt;p&gt;The private IP address associated with the Elastic IP address.&lt;/p&gt;
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// &lt;p&gt;The Elastic IP address.&lt;/p&gt;
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID of an address pool.&lt;/p&gt;
        /// </summary>
        [Input("publicIpv4Pool")]
        public Input<string>? PublicIpv4Pool { get; set; }

        [Input("tags")]
        private InputList<Inputs.TagArgs>? _tags;

        /// <summary>
        /// &lt;p&gt;Any tags assigned to the Elastic IP address.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagArgs>());
            set => _tags = value;
        }

        public AwsEc2AddressPropertiesArgs()
        {
        }
        public static new AwsEc2AddressPropertiesArgs Empty => new AwsEc2AddressPropertiesArgs();
    }
}
