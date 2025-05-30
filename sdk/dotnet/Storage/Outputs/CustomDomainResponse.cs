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
    /// The custom domain assigned to this storage account. This can be set via Update.
    /// </summary>
    [OutputType]
    public sealed class CustomDomainResponse
    {
        /// <summary>
        /// Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        /// </summary>
        public readonly bool? UseSubDomainName;

        [OutputConstructor]
        private CustomDomainResponse(
            string name,

            bool? useSubDomainName)
        {
            Name = name;
            UseSubDomainName = useSubDomainName;
        }
    }
}
