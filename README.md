# SSH URL Handler for Windows

Allows Windows to launch SSH sessions from SSH URLs (e.g.: `ssh://user:password@example.com:123`) via [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/about) (WSL).

Launching SSH sessions through WSL in [Windows Terminal](https://github.com/microsoft/terminal) is also supported.

## Instructions

Build the `ssh-url-handler` binary in WSL using [Go](https://golang.org/)
```
go build
sudo mv ssh-url-handler /usr/local/bin
```

> If you use WSL through [Windows Terminal](https://github.com/microsoft/terminal), update the launch command in [ssh.reg](./ssh.reg) to
>
> `@="wt wsl $(ssh-url-handler %1)"`

Execute [ssh.reg](./ssh.reg) on your Windows desktop to install the registry entries which create an association for SSH URLs.

You should now be able to click on SSH URLs and have an SSH session launched within WSL.
