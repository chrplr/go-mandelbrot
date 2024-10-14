# Go program to draw the Mandelbrot set

## Install Go 

    wget "https://dl.google.com/go/$(curl https://go.dev/VERSION?m=text | head -n1).linux-arm64.tar.gz" -O go.tar.gz
    sudo tar -C /usr/local -xzf go.tar.gz
    
    cat >>.bashrc <<EOF
    export GOPATH=$HOME/go
    export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
    EOF
    
    source .bashrc

## Install dependencies

### Raspberry Pi OS

    sudo apt install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libxxf86vm-dev

## Run

    go run go-mandelbrot.go

