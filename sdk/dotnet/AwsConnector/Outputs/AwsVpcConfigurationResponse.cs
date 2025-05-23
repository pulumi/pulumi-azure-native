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
    /// Definition of AwsVpcConfiguration
    /// </summary>
    [OutputType]
    public sealed class AwsVpcConfigurationResponse
    {
        /// <summary>
        /// Whether the task's elastic network interface receives a public IP address. The default value is ``DISABLED``.
        /// </summary>
        public readonly string? AssignPublicIp;
        /// <summary>
        /// The IDs of the security groups associated with the task or service. If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per ``AwsVpcConfiguration``.  All specified security groups must be from the same VPC.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// The IDs of the subnets associated with the task or service. There's a limit of 16 subnets that can be specified per ``AwsVpcConfiguration``.  All specified subnets must be from the same VPC.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private AwsVpcConfigurationResponse(
            string? assignPublicIp,

            ImmutableArray<string> securityGroups,

            ImmutableArray<string> subnets)
        {
            AssignPublicIp = assignPublicIp;
            SecurityGroups = securityGroups;
            Subnets = subnets;
        }
    }
}
