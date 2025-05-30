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
    /// Definition of awsEc2VPCEndpoint
    /// </summary>
    public sealed class AwsEc2VPCEndpointPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property creationTimestamp
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("dnsEntries")]
        private InputList<string>? _dnsEntries;

        /// <summary>
        /// Property dnsEntries
        /// </summary>
        public InputList<string> DnsEntries
        {
            get => _dnsEntries ?? (_dnsEntries = new InputList<string>());
            set => _dnsEntries = value;
        }

        /// <summary>
        /// Property id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("networkInterfaceIds")]
        private InputList<string>? _networkInterfaceIds;

        /// <summary>
        /// Property networkInterfaceIds
        /// </summary>
        public InputList<string> NetworkInterfaceIds
        {
            get => _networkInterfaceIds ?? (_networkInterfaceIds = new InputList<string>());
            set => _networkInterfaceIds = value;
        }

        /// <summary>
        /// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints. For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
        /// </summary>
        [Input("policyDocument")]
        public Input<object>? PolicyDocument { get; set; }

        /// <summary>
        /// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service. To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``. This property is supported only for interface endpoints. Default: ``false``
        /// </summary>
        [Input("privateDnsEnabled")]
        public Input<bool>? PrivateDnsEnabled { get; set; }

        [Input("routeTableIds")]
        private InputList<string>? _routeTableIds;

        /// <summary>
        /// The IDs of the route tables. Routing is supported only for gateway endpoints.
        /// </summary>
        public InputList<string> RouteTableIds
        {
            get => _routeTableIds ?? (_routeTableIds = new InputList<string>());
            set => _routeTableIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The name of the endpoint service.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The type of endpoint. Default: Gateway
        /// </summary>
        [Input("vpcEndpointType")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.VpcEndpointType>? VpcEndpointType { get; set; }

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AwsEc2VPCEndpointPropertiesArgs()
        {
        }
        public static new AwsEc2VPCEndpointPropertiesArgs Empty => new AwsEc2VPCEndpointPropertiesArgs();
    }
}
