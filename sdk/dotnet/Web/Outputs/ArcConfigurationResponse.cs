// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    [OutputType]
    public sealed class ArcConfigurationResponse
    {
        public readonly string? ArtifactStorageAccessMode;
        public readonly string? ArtifactStorageClassName;
        public readonly string? ArtifactStorageMountPath;
        public readonly string? ArtifactStorageNodeName;
        public readonly string? ArtifactsStorageType;
        public readonly Outputs.FrontEndConfigurationResponse? FrontEndServiceConfiguration;

        [OutputConstructor]
        private ArcConfigurationResponse(
            string? artifactStorageAccessMode,

            string? artifactStorageClassName,

            string? artifactStorageMountPath,

            string? artifactStorageNodeName,

            string? artifactsStorageType,

            Outputs.FrontEndConfigurationResponse? frontEndServiceConfiguration)
        {
            ArtifactStorageAccessMode = artifactStorageAccessMode;
            ArtifactStorageClassName = artifactStorageClassName;
            ArtifactStorageMountPath = artifactStorageMountPath;
            ArtifactStorageNodeName = artifactStorageNodeName;
            ArtifactsStorageType = artifactsStorageType;
            FrontEndServiceConfiguration = frontEndServiceConfiguration;
        }
    }
}
