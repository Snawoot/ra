# ra

Simple sunrise and sunset calculator

This is an offline application, and all calculations performed on your computer, so no Internet access required.

Supported platforms: Windows, MacOS, Linux, \*BSD and other Unix-like.

## Installation

#### Binary download

Pre-built binaries available on [releases](https://github.com/Snawoot/ra/releases/latest) page.

#### From source

Alternatively, you may install ra from source. Run within source directory

```
go install
```

## Usage

Just invoke `ra` with specified `-lat` and `-long` parameters:

```
$ ~/go/bin/ra -lat 37.629562 -long -116.849556
Sunrise	: 2020-06-19 05:25:55 -0700 PDT
Sunset	: 2020-06-19 16:59:59 -0700 PDT

Press ENTER to continue...

```

## Synopsis

```
$ ra - h
Usage of /home/user/go/bin/ra:
  -date string
    	date in YYYY.MM.DD format. Default is current date. (default "2020-06-20")
  -lat float
    	latitude (default NaN)
  -long float
    	longitude (default NaN)
  -nopause
    	don't wait for user to press Enter
```
