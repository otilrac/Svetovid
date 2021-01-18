
<p align="center">
  <a>
    <img alt="svetovid" title="svetovid" src="logo.png" width="450">
  </a>
</p>


<p align="center">
  Post-exploitation scripts and binaries + reverse proxy server
</p>



![Language](https://img.shields.io/badge/Language-Go-purple.svg?longCache=true&style=flat-square)    ![License](https://img.shields.io/badge/License-MIT-green.svg?longCache=true&style=flat-square)   ![Version](https://img.shields.io/badge/Version-1.0-blue.svg?longCache=true&style=flat-square)

## Introduction

This repository gathers useful post-exploitation assets (for Windows and Linux) from various sources.

It comes with a HTTP fileserver (`svetovid.go`) that supports content forwarding with Ngrok tunnel.

Such approach allows downloading stuff from behind NAT.

## Requirements

- https://github.com/redcode-labs/Coldfire
- https://github.com/akamensky/argparse
- https://github.com/atotto/clipboard	
- https://github.com/fatih/color

## Usage
<p align="left">
  <a>
    <img alt="svetovid" title="svetovid" src="usage.png" width="700">
  </a>
</p>

## Directory layout

There are 3 main folders:

- `linux`
- `windows`
- `python`

Scripts under `python` are platform-agnostic. Content inside `linux` and `windows` is further divided into 3 subfolders:

- `x86`     - binaries that can be launched on 32-bit systems
- `x64`     - binaries that can be launched on 64-bit systems
- `multi` - scripts and binaries that can be launched on both 32-bit and 64-bit systems

## License
This software is under [MIT License](https://en.wikipedia.org/wiki/MIT_License)


