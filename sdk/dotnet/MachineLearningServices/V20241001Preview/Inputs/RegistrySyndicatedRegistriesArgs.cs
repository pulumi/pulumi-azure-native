// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20241001Preview.Inputs
{

    public sealed class RegistrySyndicatedRegistriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Registry id guid of a destination registry that this registry can syndicate to 
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        public RegistrySyndicatedRegistriesArgs()
        {
        }
        public static new RegistrySyndicatedRegistriesArgs Empty => new RegistrySyndicatedRegistriesArgs();
    }
}
