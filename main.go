package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Define a struct to match the relevant parts of the CISA KEV JSON structure
type CisaKev struct {
	CatalogVersion  string        `json:"catalogVersion"`
	DateReleased    string        `json:"dateReleased"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

type Vulnerability struct {
	CveID              string   `json:"cveID"`
	VendorProject      string   `json:"vendorProject"`
	Product            string   `json:"product"`
	VulnerabilityName  string   `json:"vulnerabilityName"`
	DateAdded          string   `json:"dateAdded"`
	ShortDescription   string   `json:"shortDescription"`
	RequiredAction     string   `json:"requiredAction"`
	DueDate            string   `json:"dueDate"`
	KnownRansomwareCampaignUse string   `json:"knownRansomwareCampaignUse"`
	Notes              string   `json:"notes"`
}

func main() {
	// Define keywords to filter for (customize these based on your interests)
	keywords := []string{"exploit", "vulnerability", "remote code execution", "authentication bypass", "internet-connected assets"}

	fmt.Println("Fetching latest CISA KEV catalog...")
	kevData, err := fetchCisaKevData("https://www.cisa.gov/sites/default/files/feeds/known_exploited_vulnerabilities.json")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}

	fmt.Println("Filtering vulnerabilities...")
	foundVulnerabilities := filterVulnerabilities(kevData.Vulnerabilities, keywords)

	if len(foundVulnerabilities) > 0 {
		fmt.Println("\n--- Found Vulnerabilities Matching Keywords ---")
		for _, vuln := range foundVulnerabilities {
			fmt.Printf("CVE ID: %s\n", vuln.CveID)
			fmt.Printf("Product: %s - %s\n", vuln.VendorProject, vuln.Product)
			fmt.Printf("Description: %s\n", vuln.ShortDescription)
			fmt.Printf("Date Added: %s\n", vuln.DateAdded)
			fmt.Println("---")
		}
		// In a more advanced version, you could implement email or webhook notifications here
		fmt.Printf("\nFound %d matching vulnerabilities.\n", len(foundVulnerabilities))
	} else {
		fmt.Println("\nNo new vulnerabilities found matching the specified keywords.")
	}

	fmt.Println("Vulnerability watch complete at", time.Now().Format(time.RFC3339))
}

// fetchCisaKevData fetches the JSON data from the CISA KEV catalog URL
func fetchCisaKevData(url string) (*CisaKev, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %v", resp.Status)
	}

	var kevData CisaKev
	err = json.NewDecoder(resp.Body).Decode(&kevData)
	if err != nil {
		return nil, err
	}

	return &kevData, nil
}

// filterVulnerabilities iterates through the list of vulnerabilities and checks if their
// description or other relevant fields contain any of the specified keywords (case-insensitive).
func filterVulnerabilities(vulnerabilities []Vulnerability, keywords []string) []Vulnerability {
	var found []Vulnerability
	for _, vuln := range vulnerabilities {
		combinedText := strings.ToLower(vuln.CveID + " " + vuln.VendorProject + " " + vuln.Product + " " + vuln.VulnerabilityName + " " + vuln.ShortDescription)
		for _, keyword := range keywords {
			if strings.Contains(combinedText, strings.ToLower(keyword)) {
				found = append(found, vuln)
				break // Once a keyword is found, no need to check others for this vuln
			}
		}
	}
	return found
}