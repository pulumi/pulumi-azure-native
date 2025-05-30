// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Inputs
{

    /// <summary>
    /// Operation parameters details.
    /// </summary>
    public sealed class ParameterContractArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default parameter value.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// Parameter description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("examples")]
        private InputMap<Inputs.ParameterExampleContractArgs>? _examples;

        /// <summary>
        /// Exampled defined for the parameter.
        /// </summary>
        public InputMap<Inputs.ParameterExampleContractArgs> Examples
        {
            get => _examples ?? (_examples = new InputMap<Inputs.ParameterExampleContractArgs>());
            set => _examples = value;
        }

        /// <summary>
        /// Parameter name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies whether parameter is required or not.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// Schema identifier.
        /// </summary>
        [Input("schemaId")]
        public Input<string>? SchemaId { get; set; }

        /// <summary>
        /// Parameter type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Type name defined by the schema.
        /// </summary>
        [Input("typeName")]
        public Input<string>? TypeName { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Parameter values.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ParameterContractArgs()
        {
        }
        public static new ParameterContractArgs Empty => new ParameterContractArgs();
    }
}
