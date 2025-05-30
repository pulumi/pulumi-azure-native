// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    /// <summary>
    /// Multichannel setting. Applies to Premium FileStorage only.
    /// </summary>
    [OutputType]
    public sealed class MultichannelResponse
    {
        /// <summary>
        /// Indicates whether multichannel is enabled
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private MultichannelResponse(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
