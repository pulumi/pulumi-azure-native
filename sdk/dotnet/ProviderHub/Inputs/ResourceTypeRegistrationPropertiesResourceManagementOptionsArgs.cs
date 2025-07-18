// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    /// <summary>
    /// Resource management options.
    /// </summary>
    public sealed class ResourceTypeRegistrationPropertiesResourceManagementOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Batch provisioning support.
        /// </summary>
        [Input("batchProvisioningSupport")]
        public Input<Inputs.ResourceTypeRegistrationPropertiesBatchProvisioningSupportArgs>? BatchProvisioningSupport { get; set; }

        [Input("deleteDependencies")]
        private InputList<Inputs.DeleteDependencyArgs>? _deleteDependencies;

        /// <summary>
        /// Delete dependencies.
        /// </summary>
        public InputList<Inputs.DeleteDependencyArgs> DeleteDependencies
        {
            get => _deleteDependencies ?? (_deleteDependencies = new InputList<Inputs.DeleteDependencyArgs>());
            set => _deleteDependencies = value;
        }

        /// <summary>
        /// Nested provisioning support.
        /// </summary>
        [Input("nestedProvisioningSupport")]
        public Input<Inputs.ResourceTypeRegistrationPropertiesNestedProvisioningSupportArgs>? NestedProvisioningSupport { get; set; }

        public ResourceTypeRegistrationPropertiesResourceManagementOptionsArgs()
        {
        }
        public static new ResourceTypeRegistrationPropertiesResourceManagementOptionsArgs Empty => new ResourceTypeRegistrationPropertiesResourceManagementOptionsArgs();
    }
}
