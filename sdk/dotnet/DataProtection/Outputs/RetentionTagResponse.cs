// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Outputs
{

    /// <summary>
    /// Retention tag
    /// </summary>
    [OutputType]
    public sealed class RetentionTagResponse
    {
        /// <summary>
        /// Retention Tag version.
        /// </summary>
        public readonly string ETag;
        /// <summary>
        /// Retention Tag version.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Retention Tag Name to relate it to retention rule.
        /// </summary>
        public readonly string TagName;

        [OutputConstructor]
        private RetentionTagResponse(
            string eTag,

            string id,

            string tagName)
        {
            ETag = eTag;
            Id = id;
            TagName = tagName;
        }
    }
}
