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
    /// Definition of awsElasticBeanstalkApplication
    /// </summary>
    public sealed class AwsElasticBeanstalkApplicationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// Your description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
        /// </summary>
        [Input("resourceLifecycleConfig")]
        public Input<Inputs.ApplicationResourceLifecycleConfigArgs>? ResourceLifecycleConfig { get; set; }

        public AwsElasticBeanstalkApplicationPropertiesArgs()
        {
        }
        public static new AwsElasticBeanstalkApplicationPropertiesArgs Empty => new AwsElasticBeanstalkApplicationPropertiesArgs();
    }
}
