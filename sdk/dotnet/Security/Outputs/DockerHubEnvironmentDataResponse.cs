// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// The Docker Hub connector environment data
    /// </summary>
    [OutputType]
    public sealed class DockerHubEnvironmentDataResponse
    {
        /// <summary>
        /// The Docker Hub organization authentication details
        /// </summary>
        public readonly Outputs.AccessTokenAuthenticationResponse? Authentication;
        /// <summary>
        /// The type of the environment data.
        /// Expected value is 'DockerHubOrganization'.
        /// </summary>
        public readonly string EnvironmentType;
        /// <summary>
        /// Scan interval in hours (value should be between 1-hour to 24-hours)
        /// </summary>
        public readonly double? ScanInterval;

        [OutputConstructor]
        private DockerHubEnvironmentDataResponse(
            Outputs.AccessTokenAuthenticationResponse? authentication,

            string environmentType,

            double? scanInterval)
        {
            Authentication = authentication;
            EnvironmentType = environmentType;
            ScanInterval = scanInterval;
        }
    }
}
