<div align="center">
	<h1>Go-IEEE</h1>
	<a href="https://pkg.go.dev/github.com/pandatix/go-ieee"><img src="https://shields.io/badge/-reference-blue?logo=go&style=for-the-badge" alt="reference"></a>
	<a href="https://goreportcard.com/report/github.com/pandatix/go-ieee"><img src="https://goreportcard.com/badge/github.com/pandatix/go-ieee?style=for-the-badge" alt="go report"></a>
	<a href="https://coveralls.io/github/pandatix/go-ieee?branch=main"><img src="https://img.shields.io/coverallsCoverage/github/pandatix/go-ieee?style=for-the-badge" alt="Coverage Status"></a>
	<a href=""><img src="https://img.shields.io/github/license/pandatix/go-ieee?style=for-the-badge" alt="License"></a>
	<br>
	<a href="https://github.com/pandatix/go-ieee/actions/workflows/ci.yaml"><img src="https://img.shields.io/github/actions/workflow/status/pandatix/go-ieee/ci.yaml?style=for-the-badge&label=CI" alt="CI"></a>
	<a href="https://github.com/pandatix/go-ieee/actions/workflows/codeql-analysis.yaml"><img src="https://img.shields.io/github/actions/workflow/status/pandatix/go-ieee/codeql-analysis.yaml?style=for-the-badge&label=CodeQL" alt="CodeQL"></a>
	<br>
	<a href="https://securityscorecards.dev/viewer/?uri=github.com/pandatix/go-ieee"><img src="https://img.shields.io/ossf-scorecard/github.com/pandatix/go-ieee?label=openssf%20scorecard&style=for-the-badge" alt="OpenSSF Scoreboard"></a>
</div>

Go-IEEE API wraps the IEEExplore REST API, for the following methods/endpoints:
 - POST `search`
 - GET `document/abstract`
 - GET `document/authors`
 - GET `document/citations`
 - GET `document/disclaimer`
 - GET `document/figures`
 - GET `document/footnotes`
 - GET `document/keywords`
 - GET `document/metrics`
 - GET `document/multimedia`
 - GET `document/references`
 - GET `document/similar`

## How to use

Examples use cases could be found in the [examples directory](examples).

The basic idea is to instanciate an `*IEEEClient` and use it to call API endpoints programmatically.
