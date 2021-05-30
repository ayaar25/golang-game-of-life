# Game Of Life written in Go

This is a simple implementation of Conway's game of life with the following rules for each cell within Conway's universe
1. Any live cell with two or three live neighbours survives.
2. Any dead cell with three live neighbours becomes a live cell.
3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.

## Requirements
- Go (1.15.3)

## Test
To run the test, please run the following command under main directory
```
go test ./...
```

## Run
Before executing the game, please pick between the following patterns bellow
- Still pattern
  - `block`
  - `boat`
  - `loaf`
- Blink pattern
  - `toad`
  - `pulsar`
  - `pentadecathlon`
- Move pattern
  - `space_filler`
  - `lightweight_spaceship`
  - `gosper_glider_gun`
  - `glider_top_left`
  - `glider_bottom_right`

Once you choose a pattern, execute the following command
```
make <PATTERN_NAME>
```

Example:
```
make pulsar
```

If you want to add new pattern, please add it under `patterns` directory and add the short command in `Makefile`
