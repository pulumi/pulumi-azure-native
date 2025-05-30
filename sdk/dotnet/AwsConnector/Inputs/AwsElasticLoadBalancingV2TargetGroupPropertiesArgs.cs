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
    /// Definition of awsElasticLoadBalancingV2TargetGroup
    /// </summary>
    public sealed class AwsElasticLoadBalancingV2TargetGroupPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.
        /// </summary>
        [Input("healthCheckEnabled")]
        public Input<bool>? HealthCheckEnabled { get; set; }

        /// <summary>
        /// The approximate amount of time, in seconds, between health checks of an individual target.
        /// </summary>
        [Input("healthCheckIntervalSeconds")]
        public Input<int>? HealthCheckIntervalSeconds { get; set; }

        /// <summary>
        /// [HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// The port the load balancer uses when performing health checks on targets.
        /// </summary>
        [Input("healthCheckPort")]
        public Input<string>? HealthCheckPort { get; set; }

        /// <summary>
        /// The protocol the load balancer uses when performing health checks on targets.
        /// </summary>
        [Input("healthCheckProtocol")]
        public Input<string>? HealthCheckProtocol { get; set; }

        /// <summary>
        /// The amount of time, in seconds, during which no response from a target means a failed health check.
        /// </summary>
        [Input("healthCheckTimeoutSeconds")]
        public Input<int>? HealthCheckTimeoutSeconds { get; set; }

        /// <summary>
        /// The number of consecutive health checks successes required before considering an unhealthy target healthy.
        /// </summary>
        [Input("healthyThresholdCount")]
        public Input<int>? HealthyThresholdCount { get; set; }

        /// <summary>
        /// The type of IP address used for this target group. The possible values are ipv4 and ipv6.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        [Input("loadBalancerArns")]
        private InputList<string>? _loadBalancerArns;

        /// <summary>
        /// The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.
        /// </summary>
        public InputList<string> LoadBalancerArns
        {
            get => _loadBalancerArns ?? (_loadBalancerArns = new InputList<string>());
            set => _loadBalancerArns = value;
        }

        /// <summary>
        /// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
        /// </summary>
        [Input("matcher")]
        public Input<Inputs.MatcherArgs>? Matcher { get; set; }

        /// <summary>
        /// The name of the target group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port on which the targets receive traffic. This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol to use for routing traffic to the targets.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// [HTTP/HTTPS protocol] The protocol version. The possible values are GRPC, HTTP1, and HTTP2.
        /// </summary>
        [Input("protocolVersion")]
        public Input<string>? ProtocolVersion { get; set; }

        [Input("tags")]
        private InputList<Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags.
        /// </summary>
        public InputList<Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ARN of the Target Group
        /// </summary>
        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        [Input("targetGroupAttributes")]
        private InputList<Inputs.TargetGroupAttributeArgs>? _targetGroupAttributes;

        /// <summary>
        /// The attributes.
        /// </summary>
        public InputList<Inputs.TargetGroupAttributeArgs> TargetGroupAttributes
        {
            get => _targetGroupAttributes ?? (_targetGroupAttributes = new InputList<Inputs.TargetGroupAttributeArgs>());
            set => _targetGroupAttributes = value;
        }

        /// <summary>
        /// The full name of the target group.
        /// </summary>
        [Input("targetGroupFullName")]
        public Input<string>? TargetGroupFullName { get; set; }

        /// <summary>
        /// The name of the target group.
        /// </summary>
        [Input("targetGroupName")]
        public Input<string>? TargetGroupName { get; set; }

        /// <summary>
        /// The type of target that you must specify when registering targets with this target group. You can't specify targets for a target group using more than one target type.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        [Input("targets")]
        private InputList<Inputs.TargetDescriptionArgs>? _targets;

        /// <summary>
        /// The targets.
        /// </summary>
        public InputList<Inputs.TargetDescriptionArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.TargetDescriptionArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The number of consecutive health check failures required before considering a target unhealthy.
        /// </summary>
        [Input("unhealthyThresholdCount")]
        public Input<int>? UnhealthyThresholdCount { get; set; }

        /// <summary>
        /// The identifier of the virtual private cloud (VPC). If the target is a Lambda function, this parameter does not apply.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public AwsElasticLoadBalancingV2TargetGroupPropertiesArgs()
        {
        }
        public static new AwsElasticLoadBalancingV2TargetGroupPropertiesArgs Empty => new AwsElasticLoadBalancingV2TargetGroupPropertiesArgs();
    }
}
