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
    /// Definition of VpcSecurityGroupMembership
    /// </summary>
    public sealed class VpcSecurityGroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The status of the VPC security group.&lt;/p&gt;
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// &lt;p&gt;The VPC security group ID.&lt;/p&gt;
        /// </summary>
        [Input("vpcSecurityGroupId")]
        public Input<string>? VpcSecurityGroupId { get; set; }

        public VpcSecurityGroupMembershipArgs()
        {
        }
        public static new VpcSecurityGroupMembershipArgs Empty => new VpcSecurityGroupMembershipArgs();
    }
}
