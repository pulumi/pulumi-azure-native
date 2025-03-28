// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesConfiguration.V20241101.Outputs
{

    /// <summary>
    /// Parameters to specify which layer to pull from the OCI artifact. By default, the first layer in the artifact is pulled.
    /// </summary>
    [OutputType]
    public sealed class LayerSelectorDefinitionResponse
    {
        /// <summary>
        /// The first layer matching the specified media type will be used.
        /// </summary>
        public readonly string? MediaType;
        /// <summary>
        /// The operation to be performed on the selected layer. The default value is 'extract', but it can be set to 'copy'.
        /// </summary>
        public readonly string? Operation;

        [OutputConstructor]
        private LayerSelectorDefinitionResponse(
            string? mediaType,

            string? operation)
        {
            MediaType = mediaType;
            Operation = operation;
        }
    }
}
