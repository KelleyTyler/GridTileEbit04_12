# GridTileEbit04_12:

 This is an interative project that is meant to be a follow up to  [GridTileEbitDemo03_17], 
 which I feel has neared the end of it's usefulness as a testbed and has become increasingly hard to maintain as a result. 
 In addition I want to try some other things with the knowledge I've gained from that predecessor.
 

 ---

 ## Dependencies:
  
 - [Golang]
 - [Ebitengine]

 more to come 
 ---

 # Build Instructions
 
 I apologize but I only have a windows machine right now;

 to run: 

 ```
    go run ./app/
 ```

 to build (to bin folder):

 ```
    go build -o bin/GTE_04_03.exe app/main.go
 ```
 to build to webassembly(test)
 ```
    go run github.com/hajimehoshi/wasmserve@latest ./app/

 ```
 to build to WebAssembly (actual file):
 ```
    
 ```
 ---
 [//]: #
 [Golang]: <https://go.dev/>
 [GridTileEbitDemo03_17]: <https://github.com/KelleyTyler/GridTileEbitenDemo03_17.git>
 [Ebitengine]: <https://ebitengine.org/>