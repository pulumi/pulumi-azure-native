// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceNetworking.Inputs
{

    /// <summary>
    /// Association Subnet.
    /// </summary>
    public sealed class AssociationSubnetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Association ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public AssociationSubnetArgs()
        {
        }
        public static new AssociationSubnetArgs Empty => new AssociationSubnetArgs();
    }
}
