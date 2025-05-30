// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class ServiceManagedResourcesSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings for the service managed cosmosdb account.
        /// </summary>
        [Input("cosmosDb")]
        public Input<Inputs.CosmosDbSettingsArgs>? CosmosDb { get; set; }

        public ServiceManagedResourcesSettingsArgs()
        {
        }
        public static new ServiceManagedResourcesSettingsArgs Empty => new ServiceManagedResourcesSettingsArgs();
    }
}
