// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Confluent.Inputs
{

    /// <summary>
    /// The network associated with this object
    /// </summary>
    public sealed class SCClusterByokEntityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the referred resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// API URL for accessing or modifying the referred object
        /// </summary>
        [Input("related")]
        public Input<string>? Related { get; set; }

        /// <summary>
        /// CRN reference to the referred resource
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        public SCClusterByokEntityArgs()
        {
        }
        public static new SCClusterByokEntityArgs Empty => new SCClusterByokEntityArgs();
    }
}
