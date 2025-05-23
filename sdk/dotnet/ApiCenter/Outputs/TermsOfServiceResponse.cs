// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiCenter.Outputs
{

    /// <summary>
    /// Terms of service for the API.
    /// </summary>
    [OutputType]
    public sealed class TermsOfServiceResponse
    {
        /// <summary>
        /// URL pointing to the terms of service.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private TermsOfServiceResponse(string url)
        {
            Url = url;
        }
    }
}
