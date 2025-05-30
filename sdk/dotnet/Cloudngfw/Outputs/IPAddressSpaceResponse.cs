// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Outputs
{

    /// <summary>
    /// IP Address Space
    /// </summary>
    [OutputType]
    public sealed class IPAddressSpaceResponse
    {
        /// <summary>
        /// Address Space
        /// </summary>
        public readonly string? AddressSpace;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private IPAddressSpaceResponse(
            string? addressSpace,

            string? resourceId)
        {
            AddressSpace = addressSpace;
            ResourceId = resourceId;
        }
    }
}
