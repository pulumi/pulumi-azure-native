import pulumi
from pulumi_azurerm import resources_v20200601

# Create an Azure Resource Group
resource_group = resources_v20200601.ResourceGroup('resource_group', name='azurerm-py', location='westus')
