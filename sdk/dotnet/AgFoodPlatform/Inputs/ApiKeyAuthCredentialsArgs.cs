// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AgFoodPlatform.Inputs
{

    /// <summary>
    /// ApiKeyAuthCredentials class for ApiKey based Auth.
    /// </summary>
    public sealed class ApiKeyAuthCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Properties of the key vault.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<Inputs.KeyVaultPropertiesArgs> ApiKey { get; set; } = null!;

        /// <summary>
        /// Enum for different types of AuthCredentials supported.
        /// Expected value is 'ApiKeyAuthCredentials'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        public ApiKeyAuthCredentialsArgs()
        {
        }
        public static new ApiKeyAuthCredentialsArgs Empty => new ApiKeyAuthCredentialsArgs();
    }
}
