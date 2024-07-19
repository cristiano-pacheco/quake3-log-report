# Quake 3 Log Report

![CI](https://github.com/cristiano-pacheco/quake3-log-report/actions/workflows/go.yml/badge.svg)

This Go program reads and parses a Quake3 log file, grouping kill events by match and calculating statistics such as total kills, unique players, kills per player, and kills by type.

The requirements for it is on [this file](docs/requirements.md).

## Stack Requirements
- Go 1.22.5+

## Usage

Command to show the game ranking:
```bash
go run ./cmd/cli/main.go -file=quake3.log -outputType=ranking
```

It will show an output like this:
```json
{
        "game_1": {
                "total_kills": 5,
                "players": [
                        "Isgalamido",
                        "Mocinha"
                ],
                "kills": {
                        "Isgalamido": 4,
                        "Mocinha": 1
                }
        },
        "game_2": {
                "total_kills": 12,
                "players": [
                        "Isgalamido",
                        "Mocinha"
                ],
                "kills": {
                        "Isgalamido": 4,
                        "Mocinha": 8
                }
        },
        "game_3": {
                "total_kills": 4,
                "players": [
                        "Isgalamido",
                        "Mocinha"
                ],
                "kills": {
                        "Isgalamido": 0,
                        "Mocinha": 0
                }
        },
        "game_4": {
                "total_kills": 2,
                "players": [
                        "Isgalamido",
                        "Mocinha"
                ],
                "kills": {
                        "Isgalamido": -1,
                        "Mocinha": -1
                }
        }
}
```

Command to show a report of deaths grouped by the death types:
```bash
go run ./cmd/cli/main.go -file=quake3.log -outputType=report
```

It will show an output like this:
```json
{
        "game_1": {
                "kills_by_means": {
                        "MOD_ROCKET_SPLASH": 5
                }
        },
        "game_2": {
                "kills_by_means": {
                        "MOD_ROCKET_SPLASH": 12
                }
        },
        "game_3": {
                "kills_by_means": {
                        "MOD_ROCKET_SPLASH": 2,
                        "MOD_TRIGGER_HURT": 2
                }
        },
        "game_4": {
                "kills_by_means": {
                        "MOD_TRIGGER_HURT": 2
                }
        }
}
```

## Docker image build

```bash
docker build -t quake3logreport .
```

Running the executable from the docker image:
```bash
docker run -it --rm quake3logreport ./main -file=quake3.log -outputType=ranking
```

## Useful commands

Lint:

```bash
make lint
```

Vulnerability check:

```bash
make vuln-check
```

Tests:
```bash
make test-only
```

All in one:

```bash
make test
```
