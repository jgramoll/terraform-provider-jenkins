package main

import (
	"flag"
	"log"
	"os"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func main() {
	jobName := flag.String("job", "", "name of job to import")
	skipEnsure := flag.Bool("skipEnsure", false, "Do not ensure the xml structure")
	outputDir := flag.String("output", "output", "Directory to output the terraform code")
	flag.Parse()

	if *jobName == "" {
		log.Println("[ERROR] Job Name must be provided")
		flag.PrintDefaults()
		os.Exit(128)
	}
	jenkinsClient := initJenkinsClient()
	err := NewJobImportService(jenkinsClient).Import(*jobName, *skipEnsure, *outputDir)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func initJenkinsClient() *client.Client {
	log.Println("[INFO] Initializing jenkins client")

	address := os.Getenv("JENKINS_ADDRESS")
	if address == "" {
		log.Println("[Error] JENKINS_ADDRESS not defined")
	}
	username := os.Getenv("JENKINS_USERNAME")
	if username == "" {
		log.Println("[Error] JENKINS_USERNAME not defined")
	}
	token := os.Getenv("JENKINS_TOKEN")
	if token == "" {
		log.Println("[Error] JENKINS_TOKEN not defined")
	}

	clientConfig := client.Config{
		Address:  address,
		Username: username,
		Token:    token,
	}
	return client.NewClient(clientConfig)
}
