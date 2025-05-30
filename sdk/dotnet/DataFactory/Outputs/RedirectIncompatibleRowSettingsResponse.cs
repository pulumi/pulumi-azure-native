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
    /// Redirect incompatible row settings
    /// </summary>
    [OutputType]
    public sealed class RedirectIncompatibleRowSettingsResponse
    {
        /// <summary>
        /// Name of the Azure Storage, Storage SAS, or Azure Data Lake Store linked service used for redirecting incompatible row. Must be specified if redirectIncompatibleRowSettings is specified. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object LinkedServiceName;
        /// <summary>
        /// The path for storing the redirect incompatible row data. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Path;

        [OutputConstructor]
        private RedirectIncompatibleRowSettingsResponse(
            object linkedServiceName,

            object? path)
        {
            LinkedServiceName = linkedServiceName;
            Path = path;
        }
    }
}
