import pulumi
from pulumi_azurerm import resources, storage

resource_group = resources.v20200601.ResourceGroup('resource_group', name='azurerm-py', location='westus')

storage_account = storage.v20190601.StorageAccount('sa',
    name='pulumi143pysa',
    resource_group_name=resource_group.name,
    location='westus2',
    sku={
        'name': 'Standard_LRS',
        'tier': 'Standard',
    },
    kind='StorageV2')
