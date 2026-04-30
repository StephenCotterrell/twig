# twig

Terminal WireGuard Interface

## Overview

Terminal WireGuard Interface for managing local tunnel profiles and inspecting live connection state.

`twig` is built for quickly switching between local WireGuard configs and checking which tunnels are active without remembering exact config names.

The current version focuses on a small, daily-usable workflow: discover local WireGuard profiles, inspect tunnel state, and run basic up/down actions.

## Features

- Discovers WireGuard profiles from `/etc/wireguard`.
- Shows active and inactive profiles.
- Displays live `wg show` output.
- Allows selecting a profile from the terminal UI.
- Brings a selected profile up.
- Brings a selected profile down.
- Brings all active profiles down.
- Shows status feedback after up/down actions.

## Requirements

- Go, for building from source
- WireGuard tools installed locally: `wg` and `wg-quick`
- WireGuard config files in `/etc/wireguard`
- Permission to run `wg-quick up` and `wg-quick down`

`twig` needs permission to run `wg-quick up` and `wg-quick down`. The expected workflow is to build the binary first, then run the built binary with `sudo`.

## Usage

Build and run with `sudo`:

```sh
go build ./cmd/twig
sudo ./twig
```

Running directly with `go run` is not recommended because the program needs elevated permissions. Build the binary first, then run that binary with `sudo`.

For development, you can still run parts of the code and tests without `sudo`:

```sh
go test ./...
```

## Keybindings

- `j` / `down`: move cursor down
- `k` / `up`: move cursor up
- `enter` / `space`: select profile
- `esc`: clear selection
- `u`: bring selected profile up
- `d`: bring selected profile down
- `ctrl+d`: bring all active profiles down
- `q` / `ctrl+c`: quit

## Current Limitations

- Profiles are currently discovered from `/etc/wireguard` only.
- `wg-quick` must be available on the system path.
- `twig` currently expects to be run as a built binary with `sudo`.
- The UI does not yet include a confirmation prompt before disconnecting tunnels.
- The app does not yet check public IP, DNS state, or route-leak protection.
- Error handling is focused on common local command failures, not exhaustive diagnostics.

## Verification

The core flow has been manually verified with local WireGuard configs:

- Launch the app.
- View discovered profiles.
- Select a profile.
- Bring the selected profile up.
- Bring the selected profile down.
- Bring all active profiles down.
- Quit the app.

## Feature Roadmap

- Add a confirmation prompt before disconnecting active tunnels.
- Add route and DNS checks to better answer "am I protected right now?"
- Show current public IP.
- Display handshake age and transfer stats in a friendlier format.
- Add a logs or diagnostics view for failed `wg` and `wg-quick` commands.
- Support configurable WireGuard profile directories.
- Add favorites or friendly names for commonly used profiles.
- Improve packaging and installation.
- Add more tests around parsing, profile discovery, and command result behavior.
