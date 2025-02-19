<div align="center">
  <a href=".">
    <img src="./assets/logo.png" alt="mess logo" height="200"/>
  </a>
  
  <br>
  
  <a href="./LICENSE">
    <img src="https://img.shields.io/github/license/rakarchive/mess?style=for-the-badge">
  </a>
  <a href="https://github.com/raklaptudirm/mess/actions/workflows/ci.yml">
    <img src="https://img.shields.io/github/actions/workflow/status/rakarchive/mess/ci.yml?style=for-the-badge">
  </a>
<a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fraklaptudirm%2Fmess?ref=badge_shield" alt="FOSSA Status"><img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fraklaptudirm%2Fmess.svg?type=shield"/></a>
  <br>
  <a href="https://github.com/raklaptudirm/mess/releases">
    <img src="https://img.shields.io/github/v/release/rakarchive/mess?style=for-the-badge">
  </a>
</div>


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fraklaptudirm%2Fmess.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fraklaptudirm%2Fmess?ref=badge_large)

# Overview

Mess is an open source, UCI-compliant chess engine. Mess is not a complete chess program
and requires an UCI-compliant chess graphical user interface(e.g. XBoard with PolyGlot,
Scid, Cute Chess, eboard, Arena, Sigma Chess, Shredder, Chess Partner or Fritz) to be
used comfortably.

Mess's code structure is extremely modular and thus it may be used as a library for
developing chess engines in go. The [`./pkg`](./pkg) directory will contain all the
packages that are available for use publicly.

# Mess on Lichess

> The MessChessEngine Bot is currently hosted on my laptop, and therefore might not
always be online. The most reliable way to play against mess is to download the engine
and use it with an UCI-compliant graphical user interface.

- 🖥️ [Watch Mess Play on Lichess](https://lichess.org/@/MessChessEngine/tv)
- 🎮 [Challange Mess on Lichess](https://lichess.org/@/MessChessEngine)

# Installation

### Prebuilt Binaries

Prebuilt binaries for mess can be found in the releases section of this repository.
Binaries are only provided for release versions of Mess. Binaries have been provided for
all the major operating systems, including Windows, Linux, and Darwin.

A list of all of Mess's releases can be found [here](https://github.com/raklaptudirm/mess/releases).

### Prerequisites

The following need to be installed before you can install Mess:
- [The Go Programming Language](https://go.dev/dl/)
- [Git](https://git-scm.com/downloads) (for building from source)

### Install Globally

```bash
go install laptudirm.com/x/mess@latest
```

### Build from source
```bash
# downloard mess (other methods also work)
git clone https://github.com/rakarchive/mess.git
cd mess

# building mess binary
make EXE=path # creates binary in path
```

# License

Mess is licensed under the [Apache 2.0 License](./LICENSE).