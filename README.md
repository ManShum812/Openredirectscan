# Overview
This is a simple yet effective Open Redirect Vulnerability Scanner script. It scans URLs for potential open redirect vulnerabilities, specifically checking if URLs redirect to "evil.com." The script is written in Go and can be easily customized for your own needs.

# Features
Concurrency: The script uses goroutines to handle multiple URLs concurrently, making it efficient.

Customizable: Easily modify the script to check for other domains or additional conditions.

Color-coded Output: The script outputs results in color, making it easy to distinguish between vulnerable and non-vulnerable URLs.

# Installation

git clone https://github.com/ManShum812/Openredirectscan.git

cd Openredirectscan

go build main.go

# Running the Scanner
1. Prepare a file (e.g., test.txt) with the URLs you want to scan, one per line.
3. Run the script and pipe the file content to it:

cat test.txt | ./main

# Output
- Vulnerable URLs will be highlighted in red and show a message indicating they redirect to "evil.com".
- Non-vulnerable URLs will be highlighted in green.

Open Redirect Found: https://www.example.com/users/sign_out?redirect_uri=https://evil.com -> evil.com

Not Vulnerable: https://www.abc.com/logout_redirect.do?sysparm_url=https://evil.com

Not Vulnerable: https://www.hello.com/path3/sign_out?redirect_uri=https://evil.com

![Screenshot 2024-06-24 115607](https://github.com/ManShum812/Openredirectscan/assets/43279996/cdbf258e-9174-4738-98bc-c770048d4714)


# Get Involved
Your contributions are welcome! You can help improve this project by opening issues or submitting pull requests. If you have any ideas to enhance the tool, please share them. Together, we can make the web a safer place!
