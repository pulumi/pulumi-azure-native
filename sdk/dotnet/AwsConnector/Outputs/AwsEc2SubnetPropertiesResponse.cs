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
    /// Definition of awsEc2Subnet
    /// </summary>
    [OutputType]
    public sealed class AwsEc2SubnetPropertiesResponse
    {
        /// <summary>
        /// Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``. If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
        /// </summary>
        public readonly bool? AssignIpv6AddressOnCreation;
        /// <summary>
        /// The Availability Zone of the subnet. If you update this property, you must also update the ``CidrBlock`` property.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The AZ ID of the subnet.
        /// </summary>
        public readonly string? AvailabilityZoneId;
        /// <summary>
        /// The IPv4 CIDR block assigned to the subnet. If you update this property, we create a new subnet, and then delete the existing one.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. For more information, see [DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-nat64-dns64) in the *User Guide*.
        /// </summary>
        public readonly bool? EnableDns64;
        /// <summary>
        /// Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).
        /// </summary>
        public readonly int? EnableLniAtDeviceIndex;
        /// <summary>
        /// An IPv4 IPAM pool ID for the subnet.
        /// </summary>
        public readonly string? Ipv4IpamPoolId;
        /// <summary>
        /// An IPv4 netmask length for the subnet.
        /// </summary>
        public readonly int? Ipv4NetmaskLength;
        /// <summary>
        /// The IPv6 CIDR block. If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// The IPv6 network ranges for the subnet, in CIDR notation.
        /// </summary>
        public readonly ImmutableArray<string> Ipv6CidrBlocks;
        /// <summary>
        /// An IPv6 IPAM pool ID for the subnet.
        /// </summary>
        public readonly string? Ipv6IpamPoolId;
        /// <summary>
        /// Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.
        /// </summary>
        public readonly bool? Ipv6Native;
        /// <summary>
        /// An IPv6 netmask length for the subnet.
        /// </summary>
        public readonly int? Ipv6NetmaskLength;
        /// <summary>
        /// Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).
        /// </summary>
        public readonly bool? MapPublicIpOnLaunch;
        /// <summary>
        /// Property networkAclAssociationId
        /// </summary>
        public readonly string? NetworkAclAssociationId;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        public readonly string? OutpostArn;
        /// <summary>
        /// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*. Available options:  +  EnableResourceNameDnsAAAARecord (true | false)  +  EnableResourceNameDnsARecord (true | false)  +  HostnameType (ip-name | resource-name)
        /// </summary>
        public readonly Outputs.PrivateDnsNameOptionsOnLaunchModelPropertiesResponse? PrivateDnsNameOptionsOnLaunch;
        /// <summary>
        /// Property subnetId
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Any tags assigned to the subnet.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> Tags;
        /// <summary>
        /// The ID of the VPC the subnet is in. If you update this property, you must also update the ``CidrBlock`` property.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private AwsEc2SubnetPropertiesResponse(
            bool? assignIpv6AddressOnCreation,

            string? availabilityZone,

            string? availabilityZoneId,

            string? cidrBlock,

            bool? enableDns64,

            int? enableLniAtDeviceIndex,

            string? ipv4IpamPoolId,

            int? ipv4NetmaskLength,

            string? ipv6CidrBlock,

            ImmutableArray<string> ipv6CidrBlocks,

            string? ipv6IpamPoolId,

            bool? ipv6Native,

            int? ipv6NetmaskLength,

            bool? mapPublicIpOnLaunch,

            string? networkAclAssociationId,

            string? outpostArn,

            Outputs.PrivateDnsNameOptionsOnLaunchModelPropertiesResponse? privateDnsNameOptionsOnLaunch,

            string? subnetId,

            ImmutableArray<Outputs.TagResponse> tags,

            string? vpcId)
        {
            AssignIpv6AddressOnCreation = assignIpv6AddressOnCreation;
            AvailabilityZone = availabilityZone;
            AvailabilityZoneId = availabilityZoneId;
            CidrBlock = cidrBlock;
            EnableDns64 = enableDns64;
            EnableLniAtDeviceIndex = enableLniAtDeviceIndex;
            Ipv4IpamPoolId = ipv4IpamPoolId;
            Ipv4NetmaskLength = ipv4NetmaskLength;
            Ipv6CidrBlock = ipv6CidrBlock;
            Ipv6CidrBlocks = ipv6CidrBlocks;
            Ipv6IpamPoolId = ipv6IpamPoolId;
            Ipv6Native = ipv6Native;
            Ipv6NetmaskLength = ipv6NetmaskLength;
            MapPublicIpOnLaunch = mapPublicIpOnLaunch;
            NetworkAclAssociationId = networkAclAssociationId;
            OutpostArn = outpostArn;
            PrivateDnsNameOptionsOnLaunch = privateDnsNameOptionsOnLaunch;
            SubnetId = subnetId;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
