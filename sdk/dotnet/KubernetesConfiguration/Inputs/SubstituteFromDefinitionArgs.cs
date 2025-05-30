// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KubernetesConfiguration.Inputs
{

    /// <summary>
    /// Array of ConfigMaps/Secrets from which the variables are substituted for this Kustomization.
    /// </summary>
    public sealed class SubstituteFromDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Define whether it is ConfigMap or Secret that holds the variables to be used in substitution.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the ConfigMap/Secret that holds the variables to be used in substitution.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to True to proceed without ConfigMap/Secret, if it is not present.
        /// </summary>
        [Input("optional")]
        public Input<bool>? Optional { get; set; }

        public SubstituteFromDefinitionArgs()
        {
            Optional = false;
        }
        public static new SubstituteFromDefinitionArgs Empty => new SubstituteFromDefinitionArgs();
    }
}
