
# go-tfl

Go command line interface for the Transport for London (TFL) API.

## Installation

Clone the repository and build:

```sh
git clone https://github.com/vibe-chung/go-tfl.git
cd go-tfl
go build -o go-tfl
```

## Usage

### Line Commands

#### Get line info
```sh
go-tfl line <line>
# Example: go-tfl line central
```
Fetches information about a specific London Underground line.

#### Get stops for a line
```sh
go-tfl line stops <line>
# Example: go-tfl line stops central
```
Fetches all stop points (stations) for a given line.

#### Get line status
```sh
go-tfl line status <line>
# Example: go-tfl line status central
```
Fetches the current status for a given line.

### Crowding Command

#### Get crowding info for a stop point
```sh
go-tfl crowding <naptan> [--live]
# Example: go-tfl crowding 940GZZLUBND
# Example: go-tfl crowding 940GZZLUBND --live
```
Fetches crowding information for a specific stop point. Use `--live` for live data if available.

## API Reference

- [TFL Line Info](https://api.tfl.gov.uk/Line/{line})
- [TFL Line Stops](https://api.tfl.gov.uk/Line/{line}/StopPoints)
- [TFL Line Status](https://api.tfl.gov.uk/Line/{line}/Status)
- [TFL Crowding](https://api.tfl.gov.uk/StopPoint/{naptan}/Crowding)
- [TFL Live Crowding](https://api.tfl.gov.uk/StopPoint/{naptan}/Crowding/live)
