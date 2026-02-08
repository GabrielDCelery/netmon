#!/bin/bash

OWNER="GabrielDCelery"
REPO="netmon"

gh api repos/$OWNER/$REPO/branches/main/protection --method PUT --input .github/settings/branch-protection.json

# Enable Dependabot
# https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#enable-vulnerability-alerts
gh api --method PUT /repos/$OWNER/$REPO/vulnerability-alerts || echo "Dependabot already enabled"

# Enable secret scanning
gh api --method PATCH /repos/$OWNER/$REPO -f 'security_and_analysis.secret_scanning.status=enabled'

# Enable push protection
gh api --method PATCH /repos/$OWNER/$REPO -f 'security_and_analysis.secret_scanning.secret_scanning_push_protection=enabled'
