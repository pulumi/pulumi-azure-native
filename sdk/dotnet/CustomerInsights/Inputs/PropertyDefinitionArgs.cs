// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights.Inputs
{

    /// <summary>
    /// Property definition.
    /// </summary>
    public sealed class PropertyDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Array value separator for properties with isArray set.
        /// </summary>
        [Input("arrayValueSeparator")]
        public Input<string>? ArrayValueSeparator { get; set; }

        [Input("enumValidValues")]
        private InputList<Inputs.ProfileEnumValidValuesFormatArgs>? _enumValidValues;

        /// <summary>
        /// Describes valid values for an enum property.
        /// </summary>
        public InputList<Inputs.ProfileEnumValidValuesFormatArgs> EnumValidValues
        {
            get => _enumValidValues ?? (_enumValidValues = new InputList<Inputs.ProfileEnumValidValuesFormatArgs>());
            set => _enumValidValues = value;
        }

        /// <summary>
        /// Name of the property.
        /// </summary>
        [Input("fieldName", required: true)]
        public Input<string> FieldName { get; set; } = null!;

        /// <summary>
        /// Type of the property.
        /// </summary>
        [Input("fieldType", required: true)]
        public Input<string> FieldType { get; set; } = null!;

        /// <summary>
        /// Indicates if the property is actually an array of the fieldType above on the data api.
        /// </summary>
        [Input("isArray")]
        public Input<bool>? IsArray { get; set; }

        /// <summary>
        /// Whether property is available in graph or not.
        /// </summary>
        [Input("isAvailableInGraph")]
        public Input<bool>? IsAvailableInGraph { get; set; }

        /// <summary>
        /// Indicates if the property is an enum.
        /// </summary>
        [Input("isEnum")]
        public Input<bool>? IsEnum { get; set; }

        /// <summary>
        /// Indicates if the property is an flag enum.
        /// </summary>
        [Input("isFlagEnum")]
        public Input<bool>? IsFlagEnum { get; set; }

        /// <summary>
        /// Whether the property is an Image.
        /// </summary>
        [Input("isImage")]
        public Input<bool>? IsImage { get; set; }

        /// <summary>
        /// Whether the property is a localized string.
        /// </summary>
        [Input("isLocalizedString")]
        public Input<bool>? IsLocalizedString { get; set; }

        /// <summary>
        /// Whether the property is a name or a part of name.
        /// </summary>
        [Input("isName")]
        public Input<bool>? IsName { get; set; }

        /// <summary>
        /// Whether property value is required on instances, IsRequired field only for Interaction. Profile Instance will not check for required field.
        /// </summary>
        [Input("isRequired")]
        public Input<bool>? IsRequired { get; set; }

        /// <summary>
        /// Max length of string. Used only if type is string.
        /// </summary>
        [Input("maxLength")]
        public Input<int>? MaxLength { get; set; }

        /// <summary>
        /// The ID associated with the property.
        /// </summary>
        [Input("propertyId")]
        public Input<string>? PropertyId { get; set; }

        /// <summary>
        /// URL encoded schema.org item prop link for the property.
        /// </summary>
        [Input("schemaItemPropLink")]
        public Input<string>? SchemaItemPropLink { get; set; }

        public PropertyDefinitionArgs()
        {
        }
        public static new PropertyDefinitionArgs Empty => new PropertyDefinitionArgs();
    }
}
