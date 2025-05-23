// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataLakeAnalytics.Outputs
{

    /// <summary>
    /// Data Lake Store account information.
    /// </summary>
    [OutputType]
    public sealed class DataLakeStoreAccountInformationResponse
    {
        /// <summary>
        /// The resource identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The optional suffix for the Data Lake Store account.
        /// </summary>
        public readonly string Suffix;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DataLakeStoreAccountInformationResponse(
            string id,

            string name,

            string suffix,

            string type)
        {
            Id = id;
            Name = name;
            Suffix = suffix;
            Type = type;
        }
    }
}
