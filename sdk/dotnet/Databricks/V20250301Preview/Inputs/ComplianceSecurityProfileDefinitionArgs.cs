// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.V20250301Preview.Inputs
{

    /// <summary>
    /// Status of Compliance Security Profile feature.
    /// </summary>
    public sealed class ComplianceSecurityProfileDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("complianceStandards")]
        private InputList<Union<string, Pulumi.AzureNative.Databricks.V20250301Preview.ComplianceStandard>>? _complianceStandards;

        /// <summary>
        /// Compliance standards associated with the workspace.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Databricks.V20250301Preview.ComplianceStandard>> ComplianceStandards
        {
            get => _complianceStandards ?? (_complianceStandards = new InputList<Union<string, Pulumi.AzureNative.Databricks.V20250301Preview.ComplianceStandard>>());
            set => _complianceStandards = value;
        }

        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.Databricks.V20250301Preview.ComplianceSecurityProfileValue>? Value { get; set; }

        public ComplianceSecurityProfileDefinitionArgs()
        {
        }
        public static new ComplianceSecurityProfileDefinitionArgs Empty => new ComplianceSecurityProfileDefinitionArgs();
    }
}
