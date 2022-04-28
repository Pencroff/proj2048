# Proj 2048

Experiments with auto-play strategies for 2048 game

## Wasm build

* Windows
  
      SET GOOS=js&&SET GOARCH=wasm&&go build -ldflags="-s -w" -o ./docs/proj2048.wasm github.com/pencroff/proj2048/app

* Linux

      GOOS=js GOARCH=wasm go build -o ./docs/proj2048.wasm github.com/pencroff/proj2048/app
* Pack

        gzip -9 -v -c ./docs/proj2048.wasm > ./docs/proj2048.wasm.gz

## Stats record to DB

* Start Game
  * agent id - human, clockwise_agent, anticlockwise_agent
  * mode - manual, auto_play, simulation
  * game id - 0 - 65535
  * init field
* Move
* Finish game
  * score
  * steps
  * noMoveCounter
  * finish field

#### Refs

* [ebiten + webassembly](https://ebiten.org/documents/webassembly.html)
* [compress webassembly](https://levelup.gitconnected.com/best-practices-for-webassembly-using-golang-1-15-8dfa439827b8)


