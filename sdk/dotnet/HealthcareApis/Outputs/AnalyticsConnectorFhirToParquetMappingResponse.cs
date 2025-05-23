// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.Outputs
{

    /// <summary>
    /// FHIR Service data mapping configuration for Analytics Connector.
    /// </summary>
    [OutputType]
    public sealed class AnalyticsConnectorFhirToParquetMappingResponse
    {
        /// <summary>
        /// Artifact reference for extension schema.
        /// </summary>
        public readonly string? ExtensionSchemaReference;
        /// <summary>
        /// Artifact reference for filter configurations.
        /// </summary>
        public readonly string? FilterConfigurationReference;
        /// <summary>
        /// Type of data mapping.
        /// Expected value is 'fhirToParquet'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AnalyticsConnectorFhirToParquetMappingResponse(
            string? extensionSchemaReference,

            string? filterConfigurationReference,

            string type)
        {
            ExtensionSchemaReference = extensionSchemaReference;
            FilterConfigurationReference = filterConfigurationReference;
            Type = type;
        }
    }
}
