# README

## Introduction

This is the repository for the `rport` program - returns first available tcp port.

## Usage

You can call it directly, for example this will return the first free TCP port from 1024

```bash
rport
```

Or you can specify a min, max and step

```bash
rport -min 10000 -max 20000 -step 50
```

In this case it will search all available ports in steps of 50 from 10000 to 20000.

## Exit Codes

If a free port is found, it is returned to STDOUT and the program exits (0).  If no free port is found, exit code (1) is set with no output returned.
