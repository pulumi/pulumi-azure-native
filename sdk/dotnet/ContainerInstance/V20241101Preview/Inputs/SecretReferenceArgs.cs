// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.V20241101Preview.Inputs
{

    /// <summary>
    /// A secret reference
    /// </summary>
    public sealed class SecretReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARM resource id of the managed identity that has access to the secret in the key vault
        /// </summary>
        [Input("identity", required: true)]
        public Input<string> Identity { get; set; } = null!;

        /// <summary>
        /// The identifier of the secret reference
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The URI to the secret in key vault
        /// </summary>
        [Input("secretReferenceUri", required: true)]
        public Input<string> SecretReferenceUri { get; set; } = null!;

        public SecretReferenceArgs()
        {
        }
        public static new SecretReferenceArgs Empty => new SecretReferenceArgs();
    }
}
