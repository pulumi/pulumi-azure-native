// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Managed Virtual Network reference type.
    /// </summary>
    [OutputType]
    public sealed class ManagedVirtualNetworkReferenceResponse
    {
        /// <summary>
        /// Reference ManagedVirtualNetwork name.
        /// </summary>
        public readonly string ReferenceName;
        /// <summary>
        /// Managed Virtual Network reference type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ManagedVirtualNetworkReferenceResponse(
            string referenceName,

            string type)
        {
            ReferenceName = referenceName;
            Type = type;
        }
    }
}
