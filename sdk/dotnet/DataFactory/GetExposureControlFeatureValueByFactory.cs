// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory
{
    public static class GetExposureControlFeatureValueByFactory
    {
        /// <summary>
        /// Get exposure control feature for specific factory.
        /// 
        /// Uses Azure REST API version 2018-06-01.
        /// </summary>
        public static Task<GetExposureControlFeatureValueByFactoryResult> InvokeAsync(GetExposureControlFeatureValueByFactoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExposureControlFeatureValueByFactoryResult>("azure-native:datafactory:getExposureControlFeatureValueByFactory", args ?? new GetExposureControlFeatureValueByFactoryArgs(), options.WithDefaults());

        /// <summary>
        /// Get exposure control feature for specific factory.
        /// 
        /// Uses Azure REST API version 2018-06-01.
        /// </summary>
        public static Output<GetExposureControlFeatureValueByFactoryResult> Invoke(GetExposureControlFeatureValueByFactoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExposureControlFeatureValueByFactoryResult>("azure-native:datafactory:getExposureControlFeatureValueByFactory", args ?? new GetExposureControlFeatureValueByFactoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get exposure control feature for specific factory.
        /// 
        /// Uses Azure REST API version 2018-06-01.
        /// </summary>
        public static Output<GetExposureControlFeatureValueByFactoryResult> Invoke(GetExposureControlFeatureValueByFactoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExposureControlFeatureValueByFactoryResult>("azure-native:datafactory:getExposureControlFeatureValueByFactory", args ?? new GetExposureControlFeatureValueByFactoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExposureControlFeatureValueByFactoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The factory name.
        /// </summary>
        [Input("factoryName", required: true)]
        public string FactoryName { get; set; } = null!;

        /// <summary>
        /// The feature name.
        /// </summary>
        [Input("featureName")]
        public string? FeatureName { get; set; }

        /// <summary>
        /// The feature type.
        /// </summary>
        [Input("featureType")]
        public string? FeatureType { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetExposureControlFeatureValueByFactoryArgs()
        {
        }
        public static new GetExposureControlFeatureValueByFactoryArgs Empty => new GetExposureControlFeatureValueByFactoryArgs();
    }

    public sealed class GetExposureControlFeatureValueByFactoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The factory name.
        /// </summary>
        [Input("factoryName", required: true)]
        public Input<string> FactoryName { get; set; } = null!;

        /// <summary>
        /// The feature name.
        /// </summary>
        [Input("featureName")]
        public Input<string>? FeatureName { get; set; }

        /// <summary>
        /// The feature type.
        /// </summary>
        [Input("featureType")]
        public Input<string>? FeatureType { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetExposureControlFeatureValueByFactoryInvokeArgs()
        {
        }
        public static new GetExposureControlFeatureValueByFactoryInvokeArgs Empty => new GetExposureControlFeatureValueByFactoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetExposureControlFeatureValueByFactoryResult
    {
        /// <summary>
        /// The feature name.
        /// </summary>
        public readonly string FeatureName;
        /// <summary>
        /// The feature value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetExposureControlFeatureValueByFactoryResult(
            string featureName,

            string value)
        {
            FeatureName = featureName;
            Value = value;
        }
    }
}
