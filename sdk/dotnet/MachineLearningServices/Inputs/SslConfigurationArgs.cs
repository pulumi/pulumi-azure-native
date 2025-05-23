// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// The ssl configuration for scoring
    /// </summary>
    public sealed class SslConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cert data
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// CNAME of the cert
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// Key data
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Leaf domain label of public endpoint
        /// </summary>
        [Input("leafDomainLabel")]
        public Input<string>? LeafDomainLabel { get; set; }

        /// <summary>
        /// Indicates whether to overwrite existing domain label.
        /// </summary>
        [Input("overwriteExistingDomain")]
        public Input<bool>? OverwriteExistingDomain { get; set; }

        /// <summary>
        /// Enable or disable ssl for scoring
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.SslConfigStatus>? Status { get; set; }

        public SslConfigurationArgs()
        {
        }
        public static new SslConfigurationArgs Empty => new SslConfigurationArgs();
    }
}
