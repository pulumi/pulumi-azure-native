// Copyright 2021, Pulumi Corporation.  All rights reserved.

using System.Linq;
using Pulumi;
using Pulumi.AzureNative.ContainerRegistry;
using Pulumi.AzureNative.ContainerRegistry.Inputs;
using Pulumi.AzureNative.Resources;
using Pulumi.AzureNative.Web;
using Pulumi.AzureNative.Web.Inputs;
using Pulumi.Docker;

class MyStack : Stack
{
    public MyStack()
    {
        var resourceGroup = new ResourceGroup("resourceGroup");

        var plan = new AppServicePlan("linux-apps", new AppServicePlanArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Location = resourceGroup.Location,
            Kind = "Linux",
            Reserved = true,
            Sku = new SkuDescriptionArgs
            {
                Name = "B1",
                Tier = "Basic"
            }
        });
        
        //
        // Scenario 1: deploying an image from Docker Hub.
        // The example uses a HelloWorld application written in Go.
        // Image: https://hub.docker.com/r/microsoft/azure-appservices-go-quickstart/
        //
        var imageInDockerHub = "microsoft/azure-appservices-go-quickstart";
        
        var helloApp = new WebApp("hello-app", new WebAppArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Location = plan.Location,
            ServerFarmId = plan.Id,
            SiteConfig = new SiteConfigArgs
            {
                AppSettings = new[] 
                { 
                    new NameValuePairArgs
                    {
                        Name = "WEBSITES_ENABLE_APP_SERVICE_STORAGE",
                        Value = "false"
                    }
                },
                AlwaysOn = true,
                LinuxFxVersion = $"DOCKER|{imageInDockerHub}"
            },
            HttpsOnly = true
        });

        this.HelloEndpoint = Output.Format($"https://{helloApp.DefaultHostName}/hello");
        
        //
        // Scenario 2: deploying a custom image from Azure Container Registry.
        //
        var customImage = "node-app";

        var registry = new Registry("myregistry", new RegistryArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Location = resourceGroup.Location,
            Sku = new SkuArgs { Name = SkuName.Basic },
            AdminUserEnabled = true
        });

        var credentials = Output.Tuple(resourceGroup.Name, registry.Name).Apply(values =>
            ListRegistryCredentials.InvokeAsync(new ListRegistryCredentialsArgs
            {
                ResourceGroupName = values.Item1,
                RegistryName = values.Item2
            }));
        var adminUsername = credentials.Apply(c => c.Username);
        var adminPassword = credentials.Apply(c => Output.CreateSecret(c.Passwords.First().Value));

        var myImage = new Image(customImage, new ImageArgs
        {
            ImageName = Output.Format($"{registry.LoginServer}/{customImage}:v1.0.0"),
            Build = new DockerBuild { Context = $"./{customImage}" },
            Registry = new ImageRegistry
            {
                Server = registry.LoginServer,
                Username = adminUsername,
                Password = adminPassword
            },
        });
        
        var getStartedApp = new WebApp("get-started", new WebAppArgs
        {
            ResourceGroupName = resourceGroup.Name,
            Location = plan.Location,
            ServerFarmId = plan.Id,
            SiteConfig = new SiteConfigArgs
            {
                AppSettings = new[] 
                { 
                    new NameValuePairArgs
                    {
                        Name = "WEBSITES_ENABLE_APP_SERVICE_STORAGE",
                        Value = "false"
                    },
                    new NameValuePairArgs
                    {
                        Name = "DOCKER_REGISTRY_SERVER_URL",
                        Value = Output.Format($"https://{registry.LoginServer}")
                    },
                    new NameValuePairArgs
                    {
                        Name = "DOCKER_REGISTRY_SERVER_USERNAME",
                        Value = adminUsername
                    },
                    new NameValuePairArgs
                    {
                        Name = "DOCKER_REGISTRY_SERVER_PASSWORD",
                        Value = adminPassword
                    },
                    new NameValuePairArgs
                    {
                        Name = "WEBSITES_PORT",
                        Value = "80" // Our custom image exposes port 80. Adjust for your app as needed.
                    }
                },
                AlwaysOn = true,
                LinuxFxVersion = Output.Format($"DOCKER|{myImage.ImageName}")
            },
            HttpsOnly = true
        });
        
        this.GetStartedEndpoint = Output.Format($"https://{getStartedApp.DefaultHostName}");
    }

    [Output] public Output<string> HelloEndpoint { get; set; }
    [Output] public Output<string> GetStartedEndpoint { get; set; }
}
