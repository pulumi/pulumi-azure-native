// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// The credentials used to login to the image repository that has access to the specified image
    /// </summary>
    [OutputType]
    public sealed class VmImageRepositoryCredentialsResponse
    {
        /// <summary>
        /// Password for accessing image repository
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Username for accessing image repository
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private VmImageRepositoryCredentialsResponse(
            string password,

            string username)
        {
            Password = password;
            Username = username;
        }
    }
}
