# SSH URL Handler for Windows

Allows Windows to launch SSH sessions from SSH URLs (e.g.: `ssh://user:password@example.com:123`) via [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/about) (WSL).

Launching SSH sessions through WSL in [Windows Terminal](https://github.com/microsoft/terminal) is also supported.

## Instructions

**On WSL:**

1. Download the latest `ssh-url-handler` binary from [Releases](https://github.com/leonseng/windows-ssh-url-handler/releases)
1. Move the `ssh-url-handler` binary into `/usr/local/bin`

**On Windows:**

1. Download [ssh.reg](./ssh.reg) onto your Windows environment.
1. If you use WSL through [Windows Terminal](https://github.com/microsoft/terminal), update the launch command in [ssh.reg](./ssh.reg) to

    `@="wt wsl $(ssh-url-handler %1)"`
1. Double click on [ssh.reg](./ssh.reg) to install the registry entries which create an association for SSH URLs.

You should now be able to click on SSH URLs and have an SSH session launched within WSL.

## Building ssh-url-handler

The `ssh-url-handler` binary can be built from source using [Go](https://golang.org/)
```
go mod download
go build
sudo mv ssh-url-handler /usr/local/bin
```
