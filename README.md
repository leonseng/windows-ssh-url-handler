# SSH URL Handler for Windows

Allows Windows to launch SSH sessions from SSH URLs (e.g.: `ssh://user:password@example.com:123`) via [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/about) (WSL).

Launching SSH sessions in WSL through [Windows Terminal](https://github.com/microsoft/terminal) is also supported.

## Installation

**On WSL:**

1. Download the `ssh-url-handler` binary from [Releases](https://github.com/leonseng/windows-ssh-url-handler/releases) page.

1. Move the `ssh-url-handler` binary into `/usr/local/bin`.

1. Run `chmod +x /usr/local/bin/ssh-url-handler` to make it executable.

**On Windows:**

1. Download the registry file from [Releases](https://github.com/leonseng/windows-ssh-url-handler/releases) page based on your preference:
    - WSL: `wsl-ssh.reg`
    - WSL through Windows Terminal: `wt-wsl-ssh.reg`

1. Double click on the registry file to install the entries which create an association for SSH URLs

You should now be able to click on SSH URLs and have an SSH session launched within WSL.

## Building ssh-url-handler

The `ssh-url-handler` binary can be built from source using [Go](https://golang.org/)
```
go mod download
go build
sudo mv ssh-url-handler /usr/local/bin
```
