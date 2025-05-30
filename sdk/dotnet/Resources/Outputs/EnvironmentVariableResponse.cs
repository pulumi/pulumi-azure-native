// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources.Outputs
{

    /// <summary>
    /// The environment variable to pass to the script in the container instance.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentVariableResponse
    {
        /// <summary>
        /// The name of the environment variable.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the secure environment variable.
        /// </summary>
        public readonly string? SecureValue;
        /// <summary>
        /// The value of the environment variable.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private EnvironmentVariableResponse(
            string name,

            string? secureValue,

            string? value)
        {
            Name = name;
            SecureValue = secureValue;
            Value = value;
        }
    }
}
