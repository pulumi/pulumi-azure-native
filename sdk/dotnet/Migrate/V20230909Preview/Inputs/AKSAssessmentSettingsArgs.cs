// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.V20230909Preview.Inputs
{

    /// <summary>
    /// Data model of AKS Assessment Settings.
    /// </summary>
    public sealed class AKSAssessmentSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets azure location.
        /// </summary>
        [Input("azureLocation", required: true)]
        public Input<string> AzureLocation { get; set; } = null!;

        /// <summary>
        /// Gets or sets azure VM category.
        /// </summary>
        [Input("category", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.AzureVmCategory> Category { get; set; } = null!;

        /// <summary>
        /// Gets or sets consolidation type.
        /// </summary>
        [Input("consolidation", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.ConsolidationType> Consolidation { get; set; } = null!;

        /// <summary>
        /// Gets or sets currency.
        /// </summary>
        [Input("currency", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.AzureCurrency> Currency { get; set; } = null!;

        /// <summary>
        /// Gets or sets discount percentage.
        /// </summary>
        [Input("discountPercentage")]
        public Input<double>? DiscountPercentage { get; set; }

        /// <summary>
        /// Gets or sets environment type.
        /// </summary>
        [Input("environmentType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.AzureEnvironmentType> EnvironmentType { get; set; } = null!;

        /// <summary>
        /// Gets or sets licensing program.
        /// </summary>
        [Input("licensingProgram", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.LicensingProgram> LicensingProgram { get; set; } = null!;

        /// <summary>
        /// Gets or sets performance data settings.
        /// </summary>
        [Input("performanceData")]
        public Input<Inputs.PerfDataSettingsArgs>? PerformanceData { get; set; }

        /// <summary>
        /// Gets or sets pricing tier.
        /// </summary>
        [Input("pricingTier", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.PricingTier> PricingTier { get; set; } = null!;

        /// <summary>
        /// Gets or sets savings options.
        /// </summary>
        [Input("savingsOptions", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.SavingsOptions> SavingsOptions { get; set; } = null!;

        /// <summary>
        /// Gets or sets scaling factor.
        /// </summary>
        [Input("scalingFactor")]
        public Input<double>? ScalingFactor { get; set; }

        /// <summary>
        /// Gets or sets sizing criteria.
        /// </summary>
        [Input("sizingCriteria", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20230909Preview.AssessmentSizingCriterion> SizingCriteria { get; set; } = null!;

        public AKSAssessmentSettingsArgs()
        {
        }
        public static new AKSAssessmentSettingsArgs Empty => new AKSAssessmentSettingsArgs();
    }
}
