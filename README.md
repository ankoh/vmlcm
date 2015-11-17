# vmlcm

The **VMware (Fusion) Linked Clones Manager** is a small command line tool written in Go that speeds up the management of linked clones. It is inspired by tools like docker-compose that allow you to describe your environment through a single configuration file.

## Installation

The program can be easily installed with ```go get```. For the setup of your Go environment, please refer to the official [installation instructions](https://golang.org/doc/install).

```bash
# Download and install vmlcm through 'go get'
go get github.com/ankoh/vmlcm

# Add the GOPATH binaries to your path
export PATH=$PATH:$GOPATH/bin/

# Check vmlcm
which vmlcm
```

## Configuration File

```bash
{
  "Vmrun": "/Applications/VMware Fusion.app/Contents/Library/vmrun",
  "TemplatePath": "/Volumes/__/__/build-agents/buildagent-mac-8.vmwarevm/buildagent-mac-8.vmx",
  "ClonesDirectory": "/Volumes/__/__/build-agents/",
  "Prefix": "BuildAgents",
  "Addresses": [
    "08:00:27:__:__:__",
    "08:00:27:__:__:__",
    "08:00:27:__:__:__",
    "08:00:27:__:__:__",
    "08:00:27:__:__:__",
    "08:00:27:__:__:__",
    "08:00:27:__:__:__"
  ]
}
```

## Commands

```bash
# The 'status' command prints the current status of the linked clones
vmlcm -f agents.json status

# The 'verify' command checks if the passed configuration file is valid
vmlcm -f agents.json verify

# The 'use x' command ensures that x linked clones are available.
# (It uses the latest prefixed snapshot and creates one if needed)
# Examples:
# - 'use 3' creates 3 linked clones if 0 exist
# - 'use 2' creates 1 linked clone if 1 exists
# - 'use 2' deletes 1 linked clone if 3 exist
# - 'use 0' deletes all linked clones
vmlcm -f agents.json use 3

# The 'start' command starts all associated linked clones
vmlcm -f agents.json start

# The 'stop' command (force) stops all associated linked clones
# Attention! Force stop == Power off
vmlcm -f agents.json stop

# The 'snapshot' command creates a snapshot from the specified template
vmlcm -f agents.json snapshot

```
