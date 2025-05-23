// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of EnclaveOptions
    /// </summary>
    [OutputType]
    public sealed class EnclaveOptionsResponse
    {
        /// <summary>
        /// &lt;p&gt;If this parameter is set to &lt;code&gt;true&lt;/code&gt;, the instance is enabled for Amazon Web Services Nitro Enclaves; otherwise, it is not enabled for Amazon Web Services Nitro Enclaves.&lt;/p&gt;
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private EnclaveOptionsResponse(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
