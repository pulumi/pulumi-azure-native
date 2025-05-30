// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of awsElasticBeanstalkConfigurationTemplate
    /// </summary>
    public sealed class AwsElasticBeanstalkConfigurationTemplatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Elastic Beanstalk application to associate with this configuration template.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// An optional description for this configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of an environment whose settings you want to use to create the configuration template. You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        [Input("optionSettings")]
        private InputList<Inputs.ConfigurationOptionSettingArgs>? _optionSettings;

        /// <summary>
        /// Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide.
        /// </summary>
        public InputList<Inputs.ConfigurationOptionSettingArgs> OptionSettings
        {
            get => _optionSettings ?? (_optionSettings = new InputList<Inputs.ConfigurationOptionSettingArgs>());
            set => _optionSettings = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide.
        /// </summary>
        [Input("platformArn")]
        public Input<string>? PlatformArn { get; set; }

        /// <summary>
        /// The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide. You must specify SolutionStackName if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration. Use the ListAvailableSolutionStacks API to obtain a list of available solution stacks.
        /// </summary>
        [Input("solutionStackName")]
        public Input<string>? SolutionStackName { get; set; }

        /// <summary>
        /// An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.Values specified in OptionSettings override any values obtained from the SourceConfiguration.You must specify SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or SolutionStackName.Constraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name.
        /// </summary>
        [Input("sourceConfiguration")]
        public Input<Inputs.SourceConfigurationArgs>? SourceConfiguration { get; set; }

        /// <summary>
        /// The name of the configuration template
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public AwsElasticBeanstalkConfigurationTemplatePropertiesArgs()
        {
        }
        public static new AwsElasticBeanstalkConfigurationTemplatePropertiesArgs Empty => new AwsElasticBeanstalkConfigurationTemplatePropertiesArgs();
    }
}
