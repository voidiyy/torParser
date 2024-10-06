# TorParser

TorParser is a powerful tool that allows you to run an HTTP client through the Tor network and parse HTML pages based on the tags and attributes you choose. It is perfect for those who want to anonymously gather data from websites without violating their terms of service.

## Key Features

- Anonymous HTML parsing through the Tor network
- Flexible selection of tags and attributes for information gathering
- User-friendly command-line interface

## Usage

To run TorParser, use the following command:

```bash
sudo apt install tor -y
(linux version)
sudo systemctl enable tor 
go build .
./parse -domain google.com --tag a --attr href
```

##Examples:
