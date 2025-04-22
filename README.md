# Simple Vulnerability Watcher (Go)

This is a small Go program that fetches the latest Known Exploited Vulnerabilities (KEV) catalog from the Cybersecurity and Infrastructure Security Agency (CISA) and filters the list based on relevant keywords.

## Purpose

This project was created as a demonstration of my interest in vulnerability research and my ability to automate tasks related to threat intelligence. It's specifically tailored to showcase skills relevant to the Vulnerability Researcher role at runZero.

By monitoring the CISA KEV catalog for specific terms, this script can help proactively identify actively exploited vulnerabilities that might be of interest for further analysis and understanding of attack vectors.

## How to Run

1.  **Ensure Go is installed** on your system.
2.  **Save the `main.go` file** in a directory (e.g., `vulnwatcher`).
3.  **Open a terminal or command prompt** and navigate to that directory.
4.  **Run the script** using the command:
    ```bash
    go run main.go
    ```

## Keywords Used

The script currently filters the CISA KEV catalog for the following keywords:

* `vulnerability discovery`
* `exploit`
* `attack vector`
* `internet-connected devices`
* `zero-day`

These keywords were chosen to align with the focus of the Vulnerability Researcher role at runZero, particularly in understanding how vulnerabilities are found, how they are exploited, the methods attackers use, the types of assets involved, and the importance of emerging threats.

## Next Steps and Potential Enhancements

This is a basic version of a vulnerability watcher. Potential future enhancements could include:

* Implementing different notification mechanisms (e.g., email, Slack webhook).
* Adding more sophisticated filtering options (e.g., by vendor, product, severity).
* Storing the vulnerability data in a database for more complex querying and analysis.
* Fetching data from other vulnerability sources (e.g., NIST NVD).
* Developing more advanced analysis of vulnerability details.

## Skills Demonstrated

* **Programming in Go:** Demonstrates basic Go syntax, networking (`net/http`), and JSON parsing (`encoding/json`).
* **Automation:** Shows the ability to automate the process of monitoring security information.
* **Understanding of Vulnerability Concepts:** Utilizes relevant keywords related to vulnerability research and threat intelligence.
* **Git and GitHub:** Demonstrates familiarity with version control and sharing code on GitHub.

## Contact

[Your Name] - [Your Email Address (Optional)] - [Your LinkedIn Profile URL (Optional)]