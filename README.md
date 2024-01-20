[![Go](https://github.com/earik87/doviz/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/earik87/doviz/actions/workflows/go.yml)

# doviz

This CLI tool gives you the current USD/TRY parity from the terminal. It fetches the parity from doviz.com. 

## Installation

`git clone https://github.com/earik87/doviz`

`cd doviz`

`go install`

`go build`

## Usage

Type `./doviz` in your terminal for help.

## Example

```bash
$ ./doviz usd-try 2570.40
₺77,627.37
```

```bash
$ ./doviz try-eur 49000
€1,487.77
```


## Things to do

- [ ] Add more currencies
- [ ] Add a flag to show historical data