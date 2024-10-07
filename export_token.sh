#!/bin/bash

pwsh -Command ./get_token.ps1

if [ -f ~/.ipam/.token ]; then
    source ~/.ipam/.token
    echo "Token and API endpoint loaded from ~/.ipam/.token"
else
    echo "Token file not found. Please run the PowerShell script to authenticate and cache the token."
fi

# Verify that the environment variables are set
echo "Bash environment variables for provider authentications set:"
env | grep AZURE_IPAM