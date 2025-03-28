// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesConfiguration.V20241101.Outputs
{

    /// <summary>
    /// The source reference for the OCIRepository object.
    /// </summary>
    [OutputType]
    public sealed class OCIRepositoryRefDefinitionResponse
    {
        /// <summary>
        /// The image digest to pull from OCI repository, the value should be in the format ‘sha256:’. This takes precedence over semver.
        /// </summary>
        public readonly string? Digest;
        /// <summary>
        /// The semver range used to match against OCI repository tags. This takes precedence over tag.
        /// </summary>
        public readonly string? Semver;
        /// <summary>
        /// The OCI repository image tag name to pull. This defaults to 'latest'.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private OCIRepositoryRefDefinitionResponse(
            string? digest,

            string? semver,

            string? tag)
        {
            Digest = digest;
            Semver = semver;
            Tag = tag;
        }
    }
}
