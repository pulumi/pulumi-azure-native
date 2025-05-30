// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// Describes the parameter of customer managed disk encryption set resource id that can be specified for disk. **Note:** The disk encryption set resource id can only be specified for managed disk. Please refer https://aka.ms/mdssewithcmkoverview for more details.
    /// </summary>
    [OutputType]
    public sealed class DiskEncryptionSetParametersResponse
    {
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private DiskEncryptionSetParametersResponse(string? id)
        {
            Id = id;
        }
    }
}
