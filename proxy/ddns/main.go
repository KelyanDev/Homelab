package main

// Imports all the necessary libraries
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io"
)

// Here we declare all our constants; Add your ZoneID, apiToken and discordUrl if you want notifications upon changing IP
const (
	ZoneId = ""
	apiToken = ""
	ipfile = "/opt/ddns/latest-ip.txt"
	discordUrl = "" // Optional (discord notifications)
)

// We'll create a custom structure for our records
type Record struct {
	RecordID string
	DnsName string
}

// We'll declare our records here
var records = []Record{
	{
		RecordID: "",
		DnsName: "cloud.example.com",
	},
	{
		RecordID: "",
		DnsName: "www.example.com",
	},
}

// Main function, runs all the other ones depending on the result of the previous one
func main() {
	currentIp, err := getPublicIp()
	if err != nil {
		fmt.Println("Erreur récupération IP:", err)
		return
	}

	lastIp := readLastIp()

	if currentIp == lastIp {
		fmt.Println("Adresse IP inchangée:", currentIp)
		return
	}

	fmt.Println("Nouvelle IP détectée:", currentIp)
	updateError := false

	for _, rec := range records {
		err := updateCloudflare(ZoneId, rec.RecordID, rec.DnsName, currentIp)
		
		if err != nil {
			fmt.Println("Erreur mise à jour de l'IP côté CloudFlare:", err)
			updateError = true
		}
	}

	writeLastIp(currentIp)
	sendDiscordNotification(currentIp, updateError)
}

// This function gets you public IP using the ipify.org api
func getPublicIp() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	return string(ip), err
}

// This function updates all of your previously declared dns records on cloudflare
func updateCloudflare(zoneId, recordId, dnsName, ip string) error {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", zoneId, recordId)

	payload := map[string]interface{} {
		"type": 	"A",
		"name": 	dnsName,
		"content": 	ip,
		"ttl":	1,
		"proxied":	false,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("échec maj: %s", string(b))
	}

	return nil
}

// This function checks the latest known IP of your DDNS
func readLastIp() string {
	data, err := os.ReadFile(ipfile)
	if err != nil {
		return ""
	}
	return string(data)
}

// This function change the known IP adress in case it changed
func writeLastIp(ip string) {
	_ = os.WriteFile(ipfile, []byte(ip), 0644)
}

// This function sends a discord message when your IP changes
func sendDiscordNotification(ip string, status bool) {
	message := map[string]string{}

	if status {
		message = map[string]string{
			"content": fmt.Sprintf("Erreur lors de la mise à jour de l'IP sur CloudFlare: **%s**", ip),
		}
	} else {
		message = map[string]string{
			"content": fmt.Sprintf("IP Publique mise à jour : **%s**", ip),
		}
	}

	body, _ := json.Marshal(message)
	http.Post(discordUrl, "application/json", bytes.NewBuffer(body))
}
