function Get-GoDependencies() {
    Write-Output "Grabbing go dependencies..."
    go get -t -v -d ./...
}

function Get-WebDependencies() {
    Write-Output "Grabbing webclient dependencies..."
    Set-Location $rootdir\webclient
    yarn
    Set-Location $rootdir
}

function Get-LocalCertificates() {
    if (-not (Test-Path env:GOROOT)) {
        Write-Output "GOROOT is not defined. Set GOROOT to continue..."
        EXIT 1
    }

    if (-not (Test-Path "$rootdir\certs")) {
        New-Item -ItemType Directory "$rootdir\certs"
    }

    Set-Location $rootdir\certs

    Write-Output "Generating local certificates..."
    go run $env:GOROOT/src/crypto/tls/generate_cert.go --host localhost

    Set-Location $rootdir
}

function Start-BuildWebClient() {
    Write-Output "Building web client..."
    Set-Location $rootdir\webclient
    yarn run build
    Set-Location $rootdir
}

function Start-BuildServer($environmentFile) {
    Write-Output "Building server..."

    $envFile = "-X main.envFile=$environmentFile"

    go build -ldflags "$envFile" -o Hacksite server/cmd/server.go
}

function Start-RunWebTests() {
    Write-Output "Running webclient tests..."
    Set-Location $rootdir\webclient
    yarn run test
    Set-Location $rootdir
}

function Start-RunGoTests() {
    Write-Output "Running go tests..."

    go test -v ./server/pkg/...
}

$env_file = "environments/dev.env.json"

if (Test-Path env:environment_file) {
    $env_file = Get-Item env:environment_file
}

$command = "default"
$rootdir = (Get-Item $PSScriptRoot).parent.FullName

Write-Output "Args: $args"

if ($args.Length -eq 1) {
    $command = $args[0]
}

# Get-GoDependencies
# Get-WebDependencies
# Get-LocalCertificates

Write-Output "Command: $command"

if ($command -eq "default") {
    Start-BuildWebClient
    Start-BuildServer
    EXIT $LASTEXITCODE
}

if ($command -eq "test") {
    Start-RunWebTests
    Start-RunGoTests
    EXIT $LASTEXITCODE
}

if ($command -eq "buildServer") {
    Start-BuildServer $env_file
    EXIT $LASTEXITCODE
}

if ($command -eq "buildWeb") {
    Start-BuildWebClient
    EXIT $LASTEXITCODE
}

if ($command -eq "testServer") {
    Start-RunGoTests
    EXIT $LASTEXITCODE
}

if ($command -eq "testClient") {
    Start-RunWebTests
    EXIT $LASTEXITCODE
}
