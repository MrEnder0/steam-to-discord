set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

default: build

alias b := build
alias c := clean

[doc('Builds the binary for the default platform')]
@build:
    go build -o bin/ -v

[doc('Cleans the build artifacts and generated files')]
@clean:
    go clean
    rm bin/
    rm arm_database.dat
    rm steam_comments.db
    rm config.toml