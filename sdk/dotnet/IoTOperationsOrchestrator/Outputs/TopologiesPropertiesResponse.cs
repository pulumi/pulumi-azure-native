// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsOrchestrator.Outputs
{

    /// <summary>
    /// Defines a desired runtime component.
    /// </summary>
    [OutputType]
    public sealed class TopologiesPropertiesResponse
    {
        /// <summary>
        /// bindings description.
        /// </summary>
        public readonly ImmutableArray<Outputs.BindingPropertiesResponse> Bindings;

        [OutputConstructor]
        private TopologiesPropertiesResponse(ImmutableArray<Outputs.BindingPropertiesResponse> bindings)
        {
            Bindings = bindings;
        }
    }
}
