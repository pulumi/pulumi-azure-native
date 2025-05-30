// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// The arm resource definition resource element template details.
    /// </summary>
    public sealed class ArmResourceDefinitionResourceElementTemplateDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource element template type.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ArmResourceDefinitionResourceElementTemplateArgs>? Configuration { get; set; }

        /// <summary>
        /// The depends on profile.
        /// </summary>
        [Input("dependsOnProfile")]
        public Input<Inputs.DependsOnProfileArgs>? DependsOnProfile { get; set; }

        /// <summary>
        /// Name of the resource element template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource element template type.
        /// Expected value is 'ArmResourceDefinition'.
        /// </summary>
        [Input("resourceElementType", required: true)]
        public Input<string> ResourceElementType { get; set; } = null!;

        public ArmResourceDefinitionResourceElementTemplateDetailsArgs()
        {
        }
        public static new ArmResourceDefinitionResourceElementTemplateDetailsArgs Empty => new ArmResourceDefinitionResourceElementTemplateDetailsArgs();
    }
}
