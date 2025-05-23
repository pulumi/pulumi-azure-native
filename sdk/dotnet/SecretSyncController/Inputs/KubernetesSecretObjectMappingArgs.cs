// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecretSyncController.Inputs
{

    /// <summary>
    /// Properties defining the mapping between a cloud secret store object and a Kubernetes Secret.
    /// </summary>
    public sealed class KubernetesSecretObjectMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SourcePath is the identifier for the secret data as defined by the external secret provider. This is the key or path to the secret in the provider's system, which gets mounted to a specific path in the pod. The value should match the name of the secret as specified in the SecretProviderClass's objects array.
        /// </summary>
        [Input("sourcePath", required: true)]
        public Input<string> SourcePath { get; set; } = null!;

        /// <summary>
        /// TargetKey is the key in the Kubernetes secret's data field where the secret value will be stored. This key is used to reference the secret data within Kubernetes, and it should be unique within the secret.
        /// </summary>
        [Input("targetKey", required: true)]
        public Input<string> TargetKey { get; set; } = null!;

        public KubernetesSecretObjectMappingArgs()
        {
        }
        public static new KubernetesSecretObjectMappingArgs Empty => new KubernetesSecretObjectMappingArgs();
    }
}
