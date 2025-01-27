# Port Scanner: A Lightweight TCP Port Scanner in Go

This is a lightweight TCP port scanner written in Go. It allows you to quickly scan all 65,535 TCP ports of a target website or IP address to identify open ports. The tool leverages Go's concurrency capabilities to perform fast and efficient port scanning.

## Features
- Scans all TCP ports (1–65535).
- Utilizes Go's `goroutines` for concurrent scanning.
- Outputs a list of open ports.
- Simple and easy-to-use interface.

## Usage

### Prerequisites
- Go version 1.20 or later installed.

### Installation
1. Clone the repository:
2. cd port_scanner
3. go run main.go

## Example
Enter the website to scan: example.com

80 open
443 open

 

