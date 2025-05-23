// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Conditions required to match a header
    /// </summary>
    [OutputType]
    public sealed class HeaderMatchResponse
    {
        /// <summary>
        /// Exact value of the header
        /// </summary>
        public readonly string? ExactMatch;
        /// <summary>
        /// Name of the header
        /// </summary>
        public readonly string? Header;
        /// <summary>
        /// Prefix value of the header
        /// </summary>
        public readonly string? PrefixMatch;
        /// <summary>
        /// Regex value of the header
        /// </summary>
        public readonly string? RegexMatch;
        /// <summary>
        /// Suffix value of the header
        /// </summary>
        public readonly string? SuffixMatch;

        [OutputConstructor]
        private HeaderMatchResponse(
            string? exactMatch,

            string? header,

            string? prefixMatch,

            string? regexMatch,

            string? suffixMatch)
        {
            ExactMatch = exactMatch;
            Header = header;
            PrefixMatch = prefixMatch;
            RegexMatch = regexMatch;
            SuffixMatch = suffixMatch;
        }
    }
}
