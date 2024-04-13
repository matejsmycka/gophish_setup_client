# Gophish Setup Client

## Introduction
`gophish_setup_client` is a command-line tool designed to set up Gophish without a graphical user interface. This tool facilitates automation and simplifies the Gophish configuration process.

It will set `password_change_required=false` and API key to the one from the CLI argument.

## Prerequisites
- Golang must be installed on your system to build the executable.

## Installation

Clone the repository and build the executable:

```bash
git clone https://github.com/matejsmycka/gophish_setup_client.git
cd gophish_setup_client
go build
chmod +x gophish_setup_client
```

## Usage

Run the tool using the following command:

```bash
./gophish_setup_client -file <path-to-gophish.db> -api_key=<your-api-key>
```
- Replace <path-to-gophish.db> with the actual path to your gophish.db file.
- Replace <your-api-key> with your actual Gophish API key.

Example:

```bash
./gophish_setup_client -file ../Desktop/phishing/gophish/gophish.db -api_key=12345abcde
```
