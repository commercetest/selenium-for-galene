# A Spike for testing Galene using Selenium

Galene currently does not have any browser-based automated tests. This project is a Spike (an experiment) to explore the feasibility and utility for creating various browser-based tests for testing the browser-based functionality of Galene. If it's sufficiently useful it might be integrated into Galene, for now though it's standalone to keep that codebase clear of these experiments.

## Expectations include
This code has various expectations, including the following list of software to be installed and available locally

- <https://github.com/tebeka/selenium> and <https://pkg.go.dev/github.com/tebeka/selenium>
- <https://github.com/jech/galene> will be run locally with a self-signed certificate for SSL connections. This necessitates allowing these certificates in the tests, described in the firefox docs <https://developer.mozilla.org/en-US/docs/Web/WebDriver/Capabilities/acceptInsecureCerts> and here implemented in Go.
- <https://github.com/mozilla/geckodriver> which may already be installed in some operating systems such as Ubuntu when Firefox is installed.
- Selenium Standalone server <https://www.selenium.dev/downloads/>. Download the current release of Selenium Grid and run it locally using java.

## Usage example
Note: this is experimental and subject to breaking changes

Start Selenium Grid in standalone mode
```
java -jar selenium-server-4.26.0.jar standalone
```
Run the automated tests locally once Galene has been started
```
go test -v
```
A firefox instance should start and run briefly and the program should then complete after outputting something similar to:
```
=== RUN   TestBrowserUsingSelenium
--- PASS: TestBrowserUsingSelenium (6.03s)
PASS
ok  	webservertest.com/myserver	6.037s
```

## Further reading
There's a helpful worked step-by-step example of using Selenium in Golang at <https://www.zenrows.com/blog/selenium-golang>

