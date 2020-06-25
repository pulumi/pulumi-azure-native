import pulumi
from pulumi_azurerm import core

# Create an Azure Resource Group
resource_group = core.ResourceGroup('resource_group', name='azurerm-py', location='westus')
