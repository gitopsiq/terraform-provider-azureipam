# Step 1: Authenticate with Azure
Connect-AzAccount

# Step 2: Discover the Application ID for the ipam-engine-app
$engineAppName = "ipam-engine-app"
$engineApp = Get-AzADApplication -Filter "DisplayName eq '$engineAppName'"
$engineAppClientId = $engineApp.AppId

# Step 3: Generate the Azure AD token
$accessToken = (Get-AzAccessToken -ResourceUrl "api://$engineAppClientId").Token

# Step 4: Discover the App Service URL that starts with 'ipam'
$appServiceUrl = az webapp list --query "[?starts_with(name, 'ipam')].defaultHostName" -o tsv

# Step 5: Create a folder ~/.ipam if it doesn't exist
$ipamFolderPath = "$HOME/.ipam"
if (!(Test-Path -Path $ipamFolderPath)) {
    New-Item -Path $ipamFolderPath -ItemType Directory
}

# Step 6: Write the token and API endpoint to ~/.ipam/.token
$tokenFilePath = "$ipamFolderPath/.token"
$tokenContent = @"
AZURE_IPAM_TOKEN=$accessToken
AZURE_IPAM_API_ENDPOINT=https://$appServiceUrl
"@
Set-Content -Path $tokenFilePath -Value $tokenContent

# Step 7: Set the environment variables in the current PowerShell session
$env:AZURE_IPAM_TOKEN = $accessToken
$env:AZURE_IPAM_API_ENDPOINT = "https://$appServiceUrl"

# Confirmation message
Write-Host "Token and API endpoint saved to $tokenFilePath and environment variables set."