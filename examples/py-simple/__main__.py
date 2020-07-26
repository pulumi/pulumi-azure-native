import pulumi

# Create an Azure Resource Group
resource_group = pulumi_azurerm.ResourceGroup('resource_group', name='azurerm-py', location='westus')
