// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// Describes a script action on a running cluster.
    /// </summary>
    [OutputType]
    public sealed class RuntimeScriptActionResponse
    {
        /// <summary>
        /// The application name of the script action, if any.
        /// </summary>
        public readonly string ApplicationName;
        /// <summary>
        /// The name of the script action.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parameters for the script
        /// </summary>
        public readonly string? Parameters;
        /// <summary>
        /// The list of roles where script will be executed.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// The URI to the script.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private RuntimeScriptActionResponse(
            string applicationName,

            string name,

            string? parameters,

            ImmutableArray<string> roles,

            string uri)
        {
            ApplicationName = applicationName;
            Name = name;
            Parameters = parameters;
            Roles = roles;
            Uri = uri;
        }
    }
}
