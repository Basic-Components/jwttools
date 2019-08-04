ASSETS="bin"
GOARCHS=("386" "amd64")
GOOSS=("linux" "darwin" "windows")
export GO111MODULE="on"
# Set the GOPROXY environment variable
export GOPROXY="https://goproxy.io"

case $(uname) in
Darwin)
    case $(uname -m) in
    x86_64)
        cmd="mac"
        ;;
    *)
        cmd="mac32"
        ;;
    esac
    ;;
*)
    case $(uname -m) in
    x86_64)
        cmd="linux64"
        ;;
    *)
        cmd="linux32"
        ;;
    esac
    ;;
esac

cmd="mac"
name="jwtrpc"
if test $# -eq 0; then
    cmd="mac"
elif test $# -eq 1; then
    cmd=$1
elif test $# -eq 2; then
    cmd=$1
    name=$2
else
    echo "args too much"
    exit 0
fi

if ! test -d $ASSETS; then
    mkdir $ASSETS
    protoc -I schema schema/jwtrpcdeclare.proto --go_out=plugins=grpc:jwtrpcdeclare
fi

case $cmd in
all)
    for goarch in ${GOARCHS[@]}; do
        for goos in ${GOOSS[@]}; do
            export GOARCH=$goarch
            export GOOS=$goos
            target="$ASSETS/$GOOS-$GOARCH"
            echo "---------$target----------------"
            if ! test -d $target; then
                mkdir $target
            fi
            case $goos in
            windows)
                go build -o $target/$name.exe server/main.go
                ;;
            *)
                go build -o $target/$name server/main.go
                ;;
            esac
        done
    done
    ;;
win32)
    export GOARCH="386"
    export GOOS="windows"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name.exe server/main.go
    ;;
win64)
    export GOARCH="amd64"
    export GOOS="windows"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name.exe server/main.go
    ;;
mac)
    export GOARCH="amd64"
    export GOOS="darwin"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name server/main.go
    ;;
mac32)
    export GOARCH="386"
    export GOOS="darwin"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name server/main.go
    ;;
linux32)
    export GOARCH="386"
    export GOOS="linux"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name server/main.go
    ;;
linux64)
    export GOARCH="amd64"
    export GOOS="linux"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name server/main.go
    ;;
linuxarm)
    export GOARCH="arm"
    export GOOS="linux"
    target="$ASSETS/$GOOS-$GOARCH"
    if ! test -d $target; then
        mkdir $target
    fi
    go build -o $target/$name server/main.go
    ;;
*)
    echo "unknown cmd $cmd"
    ;;
esac
