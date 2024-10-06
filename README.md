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

## Examples:

forum.golangbridge.or - please forgive me, your resource help me in devepoling <3

```
go run main.go -domain forum.golangbridge.org/t/net-http-dialcontext/8200 -tag a -attr href
<a href="/"
<a href="/t/net-http-dialcontext/8200"
<a href="/c/getting-help/5"
<a itemprop="url" href='https://forum.golangbridge.org/u/klaymen'
<a itemprop="url" href='https://forum.golangbridge.org/u/Giulio_Iotti'
<a itemprop="url" href='https://forum.golangbridge.org/u/klaymen'
<a itemprop="url" href='https://forum.golangbridge.org/u/calmh'
<a href="https://go.dev/play/p/ulWUzdvOOI0"
<a href="https://go.dev/play/p/ulWUzdvOOI0"
<a itemprop="url" href='https://forum.golangbridge.org/u/klaymen'
<a itemprop="url" href='https://forum.golangbridge.org/u/calmh'
<a itemprop="url" href='https://forum.golangbridge.org/u/system'
<a href='/'
<a href='/categories'
<a href='/guidelines'
<a href='/tos'
<a href='/privacy'
<a href="https://www.discourse.org"
```


```
go run main.go -domain forum.golangbridge.org/t/net-http-dialcontext/8200 -tag head
<head>
    <meta charset="utf-8">
    <title>Net.http: DialContext - Getting Help - Go Forum</title>
    <meta name="description" content="Hi, 
I’m looking for a way to find out the local (ephemeral) TCP port number used in a HTTP client connection; I’m using this to be able to match results from a webcrawler with pcaps. I did find the following way to do i&amp;hellip;">
    <meta name="discourse_theme_id" content="1">
    <meta name="discourse_current_homepage" content="latest">

    <meta name="generator" content="Discourse 3.4.0.beta2-dev - https://github.com/discourse/discourse version 91ac382d83bef9d0c723a4297ab5c68378d512f1">
<link rel="icon" type="image/png" href="https://forum.golangbridge.org/uploads/default/optimized/2X/5/5ad234f7c3dbc68b6cc72d338ce4f2973d9fd81c_2_32x32.png">
<link rel="apple-touch-icon" type="image/png" href="https://forum.golangbridge.org/uploads/default/optimized/2X/5/5ad234f7c3dbc68b6cc72d338ce4f2973d9fd81c_2_180x180.png">
<meta name="theme-color" media="all" content="#E0EBF5">

<meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, user-scalable=yes, viewport-fit=cover">
<link rel="canonical" href="https://forum.golangbridge.org/t/net-http-dialcontext/8200" />

. . .

```
