package main

import (
	"flag"
	"log"
	"os"

	"github.com/jgramoll/terraform-provider-jenkins/client"
)

func main() {
	jobName := flag.String("job", "", "name of job to import")
	flag.Parse()

	if *jobName == "" {
		log.Println("[ERROR] Job Name must be provided")
		flag.PrintDefaults()
		return
	}

	c := initClient()
	js := client.JobService{Client: c}
	t, err := js.GetJobs(*jobName)
	if err != nil {
		log.Println(err.Error())
		return
	}
	println("here", t)
}

func initClient() *client.Client {
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
