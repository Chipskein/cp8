## Cp8
https://github.com/Chipskein/cp8/assets/47486707/dff47cd7-9e65-481c-9a53-138c67da7a66
#### Description
  Chip8 interpreter/emulator/Virtual Machine written in go 
  
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
