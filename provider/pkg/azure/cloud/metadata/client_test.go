// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metadata_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"reflect"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud/metadata"
)

var (
	metadata20150101 = `{"galleryEndpoint":"https://gallery.azure.com/","graphEndpoint":"https://graph.windows.net/","portalEndpoint":"https://portal.azure.com/","authentication":{"loginEndpoint":"https://login.microsoftonline.com/","audiences":["https://management.core.windows.net/","https://management.azure.com/"]}}`
	metadata20180101 = `{"galleryEndpoint":"https://gallery.azure.com/","graphEndpoint":"https://graph.windows.net/","portalEndpoint":"https://portal.azure.com/","graphAudience":"https://graph.windows.net/","authentication":{"loginEndpoint":"https://login.microsoftonline.com/","audiences":["https://management.core.windows.net/","https://management.azure.com/"],"tenant":"common","identityProvider":"AAD"},"cloudEndpoint":{"public":{"endpoint":"management.azure.com","locations":["australiacentral","australiacentral2","australiaeast","australiasoutheast","brazilsouth","brazilsoutheast","brazilus","canadacentral","canadaeast","centralindia","centralus","centraluseuap","eastasia","eastus","eastus2","eastus2euap","francecentral","francesouth","germanynorth","germanywestcentral","japaneast","japanwest","jioindiacentral","jioindiawest","koreacentral","koreasouth","northcentralus","northeurope","norwayeast","norwaywest","qatarcentral","southafricanorth","southafricawest","southcentralus","southeastasia","southindia","swedencentral","swedensouth","switzerlandnorth","switzerlandwest","uaecentral","uaenorth","uksouth","ukwest","westcentralus","westeurope","westindia","westus","westus2","westus3","israelcentral","italynorth","malaysiasouth","polandcentral","taiwannorth","taiwannorthwest"]},"chinaCloud":{"endpoint":"management.chinacloudapi.cn","locations":["chinaeast","chinanorth","chinanorth2","chinaeast2","chinanorth3","chinaeast3"]},"usGovCloud":{"endpoint":"management.usgovcloudapi.net","locations":["usgovvirginia","usgoviowa","usdodeast","usdodcentral","usgovtexas","usgovarizona"]},"germanCloud":{"endpoint":"management.microsoftazure.de","locations":["germanycentral","germanynortheast"]}}}`
	metadata20190501 = `[{"portal":"https://portal.azure.com","authentication":{"loginEndpoint":"https://login.microsoftonline.com/","audiences":["https://management.core.windows.net/","https://management.azure.com/"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.azure.net","graphAudience":"https://graph.windows.net/","graph":"https://graph.windows.net/","name":"AzureCloud","suffixes":{"azureDataLakeStoreFileSystem":"azuredatalakestore.net","acrLoginServer":"azurecr.io","sqlServerHostname":"database.windows.net","azureDataLakeAnalyticsCatalogAndJob":"azuredatalakeanalytics.net","keyVaultDns":"vault.azure.net","storage":"core.windows.net","azureFrontDoorEndpointSuffix":"azurefd.net"},"batch":"https://batch.core.windows.net/","resourceManager":"https://management.azure.com/","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","activeDirectoryDataLake":"https://datalake.azure.net/","sqlManagement":"https://management.core.windows.net:8443/","gallery":"https://gallery.azure.com/"},{"portal":"https://portal.azure.cn","authentication":{"loginEndpoint":"https://login.chinacloudapi.cn","audiences":["https://management.core.chinacloudapi.cn","https://management.chinacloudapi.cn"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.chinacloudapi.cn","graphAudience":"https://graph.chinacloudapi.cn","graph":"https://graph.chinacloudapi.cn","name":"AzureChinaCloud","suffixes":{"acrLoginServer":"azurecr.cn","sqlServerHostname":"database.chinacloudapi.cn","keyVaultDns":"vault.azure.cn","storage":"core.chinacloudapi.cn","azureFrontDoorEndpointSuffix":""},"batch":"https://batch.chinacloudapi.cn","resourceManager":"https://management.chinacloudapi.cn","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","sqlManagement":"https://management.core.chinacloudapi.cn:8443","gallery":"https://gallery.chinacloudapi.cn"},{"portal":"https://portal.azure.us","authentication":{"loginEndpoint":"https://login.microsoftonline.us","audiences":["https://management.core.usgovcloudapi.net","https://management.usgovcloudapi.net"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.usgovcloudapi.net","graphAudience":"https://graph.windows.net","graph":"https://graph.windows.net","name":"AzureUSGovernment","suffixes":{"acrLoginServer":"azurecr.us","sqlServerHostname":"database.usgovcloudapi.net","keyVaultDns":"vault.usgovcloudapi.net","storage":"core.usgovcloudapi.net","azureFrontDoorEndpointSuffix":""},"batch":"https://batch.core.usgovcloudapi.net","resourceManager":"https://management.usgovcloudapi.net","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","sqlManagement":"https://management.core.usgovcloudapi.net:8443","gallery":"https://gallery.usgovcloudapi.net"},{"portal":"https://portal.microsoftazure.de","authentication":{"loginEndpoint":"https://login.microsoftonline.de","audiences":["https://management.core.cloudapi.de","https://management.microsoftazure.de"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.cloudapi.de","graphAudience":"https://graph.cloudapi.de","graph":"https://graph.cloudapi.de","name":"AzureGermanCloud","suffixes":{"sqlServerHostname":"database.cloudapi.de","keyVaultDns":"vault.microsoftazure.de","storage":"core.cloudapi.de","azureFrontDoorEndpointSuffix":""},"batch":"https://batch.cloudapi.de","resourceManager":"https://management.microsoftazure.de","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","sqlManagement":"https://management.core.cloudapi.de:8443","gallery":"https://gallery.cloudapi.de"}]`
	metadata20220901 = `{"portal":"https://portal.azure.com","authentication":{"loginEndpoint":"https://login.microsoftonline.com","audiences":["https://management.core.windows.net/","https://management.azure.com/"],"tenant":"common","identityProvider":"AAD"},"media":"https://rest.media.azure.net","graphAudience":"https://graph.windows.net/","graph":"https://graph.windows.net/","name":"AzureCloud","suffixes":{"azureDataLakeStoreFileSystem":"azuredatalakestore.net","acrLoginServer":"azurecr.io","sqlServerHostname":"database.windows.net","azureDataLakeAnalyticsCatalogAndJob":"azuredatalakeanalytics.net","keyVaultDns":"vault.azure.net","storage":"core.windows.net","azureFrontDoorEndpointSuffix":"azurefd.net","storageSyncEndpointSuffix":"afs.azure.net","mhsmDns":"managedhsm.azure.net","mysqlServerEndpoint":"mysql.database.azure.com","postgresqlServerEndpoint":"postgres.database.azure.com","mariadbServerEndpoint":"mariadb.database.azure.com","synapseAnalytics":"dev.azuresynapse.net","attestationEndpoint":"attest.azure.net"},"batch":"https://batch.core.windows.net/","resourceManager":"https://management.azure.com/","vmImageAliasDoc":"https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/arm-compute/quickstart-templates/aliases.json","activeDirectoryDataLake":"https://datalake.azure.net/","sqlManagement":"https://management.core.windows.net:8443/","microsoftGraphResourceId":"https://graph.microsoft.com/","appInsightsResourceId":"https://api.applicationinsights.io","appInsightsTelemetryChannelResourceId":"https://dc.applicationinsights.azure.com/v2/track","attestationResourceId":"https://attest.azure.net","synapseAnalyticsResourceId":"https://dev.azuresynapse.net","logAnalyticsResourceId":"https://api.loganalytics.io","ossrDbmsResourceId":"https://ossrdbms-aad.database.windows.net"}`
)

func TestGetMetaData(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	addr := metadataStubServer(ctx)
	defer func() {
		cancel()
	}()

	client := metadata.NewClientWithEndpoint(fmt.Sprintf("http://%s", addr))
	if client == nil {
		t.Fatal("client was nil")
	}

	expected := metadata.MetaData{
		Authentication: metadata.Authentication{
			Audiences: []string{
				"https://management.core.windows.net/",
				"https://management.azure.com/",
			},
			LoginEndpoint:    "https://login.microsoftonline.com",
			IdentityProvider: "AAD",
			Tenant:           "common",
		},
		DnsSuffixes: metadata.DnsSuffixes{
			Attestation:       "attest.azure.net",
			ContainerRegistry: "azurecr.io",
			DataLakeStore:     "azuredatalakestore.net",
			FrontDoor:         "azurefd.net",
			KeyVault:          "vault.azure.net",
			ManagedHSM:        "managedhsm.azure.net",
			MariaDB:           "mariadb.database.azure.com",
			MySql:             "mysql.database.azure.com",
			Postgresql:        "postgres.database.azure.com",
			SqlServer:         "database.windows.net",
			Storage:           "core.windows.net",
			StorageSync:       "afs.azure.net",
			Synapse:           "dev.azuresynapse.net",
		},
		Name: "AzureCloud",
		ResourceIdentifiers: metadata.ResourceIdentifiers{
			Attestation:    "https://attest.azure.net",
			Batch:          "https://batch.core.windows.net",
			DataLake:       "https://datalake.azure.net",
			LogAnalytics:   "https://api.loganalytics.io",
			Media:          "https://rest.media.azure.net",
			MicrosoftGraph: "https://graph.microsoft.com",
			OSSRDBMS:       "https://ossrdbms-aad.database.windows.net",
			Synapse:        "https://dev.azuresynapse.net",
		},
		ResourceManagerEndpoint: "https://management.azure.com/",
	}

	m, err := client.GetMetaData(ctx)
	if err != nil {
		t.Fatalf("error received: %+v", err)
	} else if m == nil {
		t.Fatal("metadata result was nil")
	} else if !reflect.DeepEqual(*m, expected) {
		t.Fatalf("unexpected Metadata{} result.\nexpected:\n\n%#v\n\nreceived:\n\n%#v\n", expected, *m)
	}
}

func metadataStubServer(ctx context.Context) net.Addr {
	handler := http.NewServeMux()

	handler.HandleFunc("/metadata/endpoints", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		apiVersion := q.Get("api-version")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		switch apiVersion {
		case "1.0", "2015-01-01":
			fmt.Fprint(w, metadata20150101)
		case "2018-01-01":
			fmt.Fprint(w, metadata20180101)
		case "2019-05-01":
			fmt.Fprint(w, metadata20190501)
		case "2022-09-01":
			fmt.Fprint(w, metadata20220901)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error":"unrecognized api-version"}`)
			return
		}
	})

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", "localhost:0")
	if err != nil {
		panic(fmt.Sprintf("could not start listener: %s", err))
	}
	server := &http.Server{
		Addr:    listener.Addr().String(),
		Handler: handler,
	}

	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %s\n", err)
		}
		log.Printf("ARM Metadata Stub Service stopped\n")
	}()

	go func() {
		<-ctx.Done()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("failed to gracefully shut down ARM Metadata stub server: %v", err)
		}
	}()

	log.Printf("ARM Metadata Stub Service listening on %s\n", listener.Addr())
	return listener.Addr()
}
