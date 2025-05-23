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
    /// Container App container environment variable.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentVarResponse
    {
        /// <summary>
        /// Environment variable name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Name of the Container App secret from which to pull the environment variable value.
        /// </summary>
        public readonly string? SecretRef;
        /// <summary>
        /// Non-secret environment variable value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private EnvironmentVarResponse(
            string? name,

            string? secretRef,

            string? value)
        {
            Name = name;
            SecretRef = secretRef;
            Value = value;
        }
    }
}
