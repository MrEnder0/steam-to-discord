set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

default: build-arm

alias b := build
alias ba := build-arm

[doc('Builds the base binary for arm')]
@build-arm:
    $Env:GOOS = "linux"; $Env:GOARCH = "arm64"; $Env:CGO_ENABLED = 0; go build -o bin/ -v

[doc('Builds the binary for the default platform')]
@build:
    go build -o bin/ -v

[doc('Cleans the build artifacts and generated files')]
@clean:
    go clean
    rm bin/
    rm arm_database.dat
    rm database.db
    rm config.toml