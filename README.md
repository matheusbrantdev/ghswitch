# ghswitch

Switch between multiple GitHub accounts on the same machine — SSH keys and Git config handled automatically.

## Install

```bash
curl -sSL https://raw.githubusercontent.com/matheusbrantdev/ghswitch/main/install.sh | sh
```

Restart your terminal after installing.

## Usage

```bash
ghswitch add        # register a new profile
ghswitch use work   # switch to a profile
ghswitch list       # list all profiles
ghswitch update     # edit a profile
ghswitch remove     # remove a profile
ghswitch reset      # delete all profiles
```

## How it works

When you run `ghswitch use <profile>`, it:

1. Updates `~/.ssh/config` to point `github.com` to the correct SSH key
2. Sets `git config --global user.name` and `user.email`
3. Shows the active profile in your terminal prompt

## Requirements

- Linux / WSL
- Git
- SSH key registered on each GitHub account
