// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Secret deployment resource id reference.
    /// </summary>
    public sealed class SecretDeploymentResourceReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The resource reference arm id type.
        /// Expected value is 'Secret'.
        /// </summary>
        [Input("idType", required: true)]
        public Input<string> IdType { get; set; } = null!;

        public SecretDeploymentResourceReferenceArgs()
        {
        }
        public static new SecretDeploymentResourceReferenceArgs Empty => new SecretDeploymentResourceReferenceArgs();
    }
}
