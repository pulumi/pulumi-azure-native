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
    /// Definition of ProjectEnvironment
    /// </summary>
    public sealed class ProjectEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The ARN of the Amazon S3 bucket, path prefix, and object key that contains the PEM-encoded certificate for the build project. For more information, see &lt;a href='https://docs.aws.amazon.com/codebuild/latest/userguide/create-project-cli.html#cli.environment.certificate'&gt;certificate&lt;/a&gt; in the &lt;i&gt;CodeBuild User Guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// &lt;p&gt;Information about the compute resources the build project uses. Available values include:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_GENERAL1_SMALL&lt;/code&gt;: Use up to 3 GB memory and 2 vCPUs for builds.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_GENERAL1_MEDIUM&lt;/code&gt;: Use up to 7 GB memory and 4 vCPUs for builds.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_GENERAL1_LARGE&lt;/code&gt;: Use up to 16 GB memory and 8 vCPUs for builds, depending on your environment type.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_GENERAL1_XLARGE&lt;/code&gt;: Use up to 70 GB memory and 36 vCPUs for builds, depending on your environment type.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_GENERAL1_2XLARGE&lt;/code&gt;: Use up to 145 GB memory, 72 vCPUs, and 824 GB of SSD storage for builds. This compute type supports Docker images up to 100 GB uncompressed.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_LAMBDA_1GB&lt;/code&gt;: Use up to 1 GB memory for builds. Only available for environment type &lt;code&gt;LINUX_LAMBDA_CONTAINER&lt;/code&gt; and &lt;code&gt;ARM_LAMBDA_CONTAINER&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_LAMBDA_2GB&lt;/code&gt;: Use up to 2 GB memory for builds. Only available for environment type &lt;code&gt;LINUX_LAMBDA_CONTAINER&lt;/code&gt; and &lt;code&gt;ARM_LAMBDA_CONTAINER&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_LAMBDA_4GB&lt;/code&gt;: Use up to 4 GB memory for builds. Only available for environment type &lt;code&gt;LINUX_LAMBDA_CONTAINER&lt;/code&gt; and &lt;code&gt;ARM_LAMBDA_CONTAINER&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_LAMBDA_8GB&lt;/code&gt;: Use up to 8 GB memory for builds. Only available for environment type &lt;code&gt;LINUX_LAMBDA_CONTAINER&lt;/code&gt; and &lt;code&gt;ARM_LAMBDA_CONTAINER&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;BUILD_LAMBDA_10GB&lt;/code&gt;: Use up to 10 GB memory for builds. Only available for environment type &lt;code&gt;LINUX_LAMBDA_CONTAINER&lt;/code&gt; and &lt;code&gt;ARM_LAMBDA_CONTAINER&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt; If you use &lt;code&gt;BUILD_GENERAL1_SMALL&lt;/code&gt;: &lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; For environment type &lt;code&gt;LINUX_CONTAINER&lt;/code&gt;, you can use up to 3 GB memory and 2 vCPUs for builds. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; For environment type &lt;code&gt;LINUX_GPU_CONTAINER&lt;/code&gt;, you can use up to 16 GB memory, 4 vCPUs, and 1 NVIDIA A10G Tensor Core GPU for builds.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; For environment type &lt;code&gt;ARM_CONTAINER&lt;/code&gt;, you can use up to 4 GB memory and 2 vCPUs on ARM-based processors for builds.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt; If you use &lt;code&gt;BUILD_GENERAL1_LARGE&lt;/code&gt;: &lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; For environment type &lt;code&gt;LINUX_CONTAINER&lt;/code&gt;, you can use up to 15 GB memory and 8 vCPUs for builds. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; For environment type &lt;code&gt;LINUX_GPU_CONTAINER&lt;/code&gt;, you can use up to 255 GB memory, 32 vCPUs, and 4 NVIDIA Tesla V100 GPUs for builds.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; For environment type &lt;code&gt;ARM_CONTAINER&lt;/code&gt;, you can use up to 16 GB memory and 8 vCPUs on ARM-based processors for builds.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;note&gt; &lt;p&gt;If you're using compute fleets during project creation, &lt;code&gt;computeType&lt;/code&gt; will be ignored.&lt;/p&gt; &lt;/note&gt; &lt;p&gt;For more information, see &lt;a href='https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html'&gt;Build Environment Compute Types&lt;/a&gt; in the &lt;i&gt;CodeBuild User Guide.&lt;/i&gt; &lt;/p&gt;
        /// </summary>
        [Input("computeType")]
        public Input<Inputs.ComputeTypeEnumValueArgs>? ComputeType { get; set; }

        [Input("environmentVariables")]
        private InputList<Inputs.EnvironmentVariableArgs>? _environmentVariables;

        /// <summary>
        /// &lt;p&gt;A set of environment variables to make available to builds for this build project.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.EnvironmentVariableArgs> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputList<Inputs.EnvironmentVariableArgs>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// &lt;p&gt;A ProjectFleet object to use for this build project.&lt;/p&gt;
        /// </summary>
        [Input("fleet")]
        public Input<Inputs.ProjectFleetArgs>? Fleet { get; set; }

        /// <summary>
        /// &lt;p&gt;The image tag or image digest that identifies the Docker image to use for this build project. Use the following formats:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;For an image tag: &lt;code&gt;&amp;lt;registry&amp;gt;/&amp;lt;repository&amp;gt;:&amp;lt;tag&amp;gt;&lt;/code&gt;. For example, in the Docker repository that CodeBuild uses to manage its Docker images, this would be &lt;code&gt;aws/codebuild/standard:4.0&lt;/code&gt;. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;For an image digest: &lt;code&gt;&amp;lt;registry&amp;gt;/&amp;lt;repository&amp;gt;@&amp;lt;digest&amp;gt;&lt;/code&gt;. For example, to specify an image with the digest 'sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf,' use &lt;code&gt;&amp;lt;registry&amp;gt;/&amp;lt;repository&amp;gt;@sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt;For more information, see &lt;a href='https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html'&gt;Docker images provided by CodeBuild&lt;/a&gt; in the &lt;i&gt;CodeBuild user guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// &lt;p&gt; The type of credentials CodeBuild uses to pull images in your build. There are two valid values: &lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CODEBUILD&lt;/code&gt; specifies that CodeBuild uses its own credentials. This requires that you modify your ECR repository policy to trust CodeBuild service principal. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;SERVICE_ROLE&lt;/code&gt; specifies that CodeBuild uses your build project's service role. &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt; When you use a cross-account or private registry image, you must use SERVICE_ROLE credentials. When you use an CodeBuild curated image, you must use CODEBUILD credentials. &lt;/p&gt;
        /// </summary>
        [Input("imagePullCredentialsType")]
        public Input<Inputs.ImagePullCredentialsTypeEnumValueArgs>? ImagePullCredentialsType { get; set; }

        /// <summary>
        /// &lt;p&gt;Enables running the Docker daemon inside a Docker container. Set to true only if the build project is used to build Docker images. Otherwise, a build that attempts to interact with the Docker daemon fails. The default setting is &lt;code&gt;false&lt;/code&gt;.&lt;/p&gt; &lt;p&gt;You can initialize the Docker daemon during the install phase of your build by adding one of the following sets of commands to the install phase of your buildspec file:&lt;/p&gt; &lt;p&gt;If the operating system's base image is Ubuntu Linux:&lt;/p&gt; &lt;p&gt; &lt;code&gt;- nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay&amp;amp;&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;- timeout 15 sh -c 'until docker info; do echo .; sleep 1; done'&lt;/code&gt; &lt;/p&gt; &lt;p&gt;If the operating system's base image is Alpine Linux and the previous command does not work, add the &lt;code&gt;-t&lt;/code&gt; argument to &lt;code&gt;timeout&lt;/code&gt;:&lt;/p&gt; &lt;p&gt; &lt;code&gt;- nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay&amp;amp;&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;- timeout -t 15 sh -c 'until docker info; do echo .; sleep 1; done'&lt;/code&gt; &lt;/p&gt;
        /// </summary>
        [Input("privilegedMode")]
        public Input<bool>? PrivilegedMode { get; set; }

        /// <summary>
        /// &lt;p&gt; The credentials for access to a private registry.&lt;/p&gt;
        /// </summary>
        [Input("registryCredential")]
        public Input<Inputs.RegistryCredentialArgs>? RegistryCredential { get; set; }

        /// <summary>
        /// &lt;p&gt;The type of build environment to use for related builds.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;The environment type &lt;code&gt;ARM_CONTAINER&lt;/code&gt; is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), Asia Pacific (Mumbai), Asia Pacific (Tokyo), Asia Pacific (Sydney), and EU (Frankfurt).&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;The environment type &lt;code&gt;LINUX_CONTAINER&lt;/code&gt; is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Canada (Central), EU (Ireland), EU (London), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific (Singapore), Asia Pacific (Sydney), China (Beijing), and China (Ningxia).&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;The environment type &lt;code&gt;LINUX_GPU_CONTAINER&lt;/code&gt; is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Canada (Central), EU (Ireland), EU (London), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific (Singapore), Asia Pacific (Sydney) , China (Beijing), and China (Ningxia).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;The environment types &lt;code&gt;ARM_LAMBDA_CONTAINER&lt;/code&gt; and &lt;code&gt;LINUX_LAMBDA_CONTAINER&lt;/code&gt; are available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Asia Pacific (Mumbai), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo), EU (Frankfurt), EU (Ireland), and South America (São Paulo).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;The environment types &lt;code&gt;WINDOWS_CONTAINER&lt;/code&gt; and &lt;code&gt;WINDOWS_SERVER_2019_CONTAINER&lt;/code&gt; are available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), and EU (Ireland).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;note&gt; &lt;p&gt;If you're using compute fleets during project creation, &lt;code&gt;type&lt;/code&gt; will be ignored.&lt;/p&gt; &lt;/note&gt; &lt;p&gt;For more information, see &lt;a href='https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html'&gt;Build environment compute types&lt;/a&gt; in the &lt;i&gt;CodeBuild user guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        [Input("type")]
        public Input<Inputs.EnvironmentTypeEnumValueArgs>? Type { get; set; }

        public ProjectEnvironmentArgs()
        {
        }
        public static new ProjectEnvironmentArgs Empty => new ProjectEnvironmentArgs();
    }
}
