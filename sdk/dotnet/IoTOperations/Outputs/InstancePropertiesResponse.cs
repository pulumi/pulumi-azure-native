// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// The properties of the Instance resource.
    /// </summary>
    [OutputType]
    public sealed class InstancePropertiesResponse
    {
        /// <summary>
        /// Detailed description of the Instance.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The reference to the Schema Registry for this AIO Instance.
        /// </summary>
        public readonly Outputs.SchemaRegistryRefResponse SchemaRegistryRef;
        /// <summary>
        /// The Azure IoT Operations version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private InstancePropertiesResponse(
            string? description,

            string provisioningState,

            Outputs.SchemaRegistryRefResponse schemaRegistryRef,

            string version)
        {
            Description = description;
            ProvisioningState = provisioningState;
            SchemaRegistryRef = schemaRegistryRef;
            Version = version;
        }
    }
}
