package main

import (
	"fmt"
	"log"
	"net/http"
	//"go.uber.org/zap"
	gopivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/shinji62/generate-opsman-status/opsmanager"
	"github.com/shinji62/generate-opsman-status/pivnet"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	opsmanagerEndpoint     = kingpin.Flag("opsman-endpoint", "OpsMan Endpoint").OverrideDefaultFromEnvar("OPSMAN_ENDPOINT").Required().String()
	clientIDOpsManagager   = kingpin.Flag("client-id-opsman", "Client ID.").OverrideDefaultFromEnvar("OPSMAN_CLIENT_ID").Required().String()
	clientSecretOpsManager = kingpin.Flag("client-secret-opsman", "Client secret.").OverrideDefaultFromEnvar("OPSMAN_CLIENT_SECRET").Required().String()
	pivnetToken            = kingpin.Flag("pivnet-api-token", "Pivnet API token").OverrideDefaultFromEnvar("PIVNET_TOKEN").Required().String()
	skipSSLValidation      = kingpin.Flag("skip-ssl-validation", "Please don't").Default("false").OverrideDefaultFromEnvar("SKIP_SSL_VALIDATION").Bool()
	pathProf               = kingpin.Flag("path-output", "Set the Path to Opsman file /path/opsman.json by default result").Default("result").OverrideDefaultFromEnvar("PATH_PROF").String()
)

func main() {

	kingpin.Parse()

	//SetUp HTTPClient
	hClientOpsMan := http.DefaultClient

	//opsMan Client
	opsClient := opsmanager.NewOpsManagerApi(hClientOpsMan, *opsmanagerEndpoint, *clientIDOpsManagager, *clientSecretOpsManager, *skipSSLValidation)
	if !opsClient.CloseSessions() {
		log.Fatal("Could not logout user from OpsManager")
	}
	fmt.Println("All users have been disconnected from OpsManager")
	opsClient = opsmanager.NewOpsManagerApi(hClientOpsMan, *opsmanagerEndpoint, *clientIDOpsManagager, *clientSecretOpsManager, *skipSSLValidation)

	boshDiag, err := opsClient.GetDiagnosticReport()
	if err != nil {
		log.Fatalf("Could not get Diagnostic Report from OpsManager ", err)
	}

	if !opsClient.CloseSessions() {
		fmt.Println("Could not logout user from OpsManager, after operation")
	}
	fmt.Println("Users have been disconnected from OpsManager")

	//Pivnet Client

	config := gopivnet.ClientConfig{
		Host:      gopivnet.DefaultHost,
		Token:     *pivnetToken,
		UserAgent: opsmanager.UserAgent,
	}
	client := pivnet.NewPivnetClient(config, boshDiag)
	client.CreateInfoPCF()
	client.GenerateJson("./result.json")

	//release, err := client.ReleasesForProductSlug("elastic-runtime")
	// 	release, err := client.GetRelease("elastic-runtime", "1.8.27")
	// 	//products, err := client.G
	// 	//product, _ := client.Products.Get("elastic-runtime")
	// 	fmt.Printf("Err: %v", err)
	// 	//fmt.Printf("products: %v", products)
	// 	fmt.Printf("releases: %v", release.ReleaseNotesURL)
}
