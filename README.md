## Cp8
https://github.com/Chipskein/cp8/assets/47486707/dff47cd7-9e65-481c-9a53-138c67da7a66
#### Description
  Chip8 interpreter/emulator/Virtual Machine written in go
#### KeyMap
| Chip8 Keyboard  |  Converted Keyboard |
|-----------------|---------------------|
| 1 2 3 4         |  1 2 3 C            |
| Q W E R         |  4 5 6 D            |
| A S D F         |  7 8 9 E            |
| Z X C V         |  A 0 B F            |
| Esc             |  Close the window   |

#### ROMS
if you clone this repo using --recursive flag roms folder will be created
with roms found at https://github.com/kripod/chip8-roms

#### Super Chip 48
  **At the moment there is no support for Super chip-48 instructions,there for some roms will not run well**
  
#### Clone
  **Warning**:This repositories uses git submodules then clone using --recursive flag
  
    git clone --recursive https://github.com/Chipskein/cp8.git

  *in case you did forgot of the flag, you can run:*
    
    git submodule update --init
  
#### Build
    go build -o cp8
### Run
    ./cp8 --rom $ROM_FILE_PATH
    
### References
* http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#3.2
* https://en.wikipedia.org/wiki/CHIP-8
* https://tobiasvl.github.io/blog/write-a-chip-8-emulator
