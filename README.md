# TelegramAPIService

It's a mini service in Golang that will receive new message from specified telegram account, and it will send copy of that message with another random message to the sender (chat).
It uses [TDLib](https://core.telegram.org/tdlib/) with Golang wrapper for TDLib: [go-tdlib](https://github.com/zelenin/go-tdlib).

## Installation

### TDLib installation

#### Ubuntu 18-19 / Debian 9

##### Manual compilation

```bash
sudo apt-get update -y
sudo apt-get install -y \
    build-essential \
    ca-certificates \
    ccache \
    cmake \
    git \
    gperf \
    libssl-dev \
    libreadline-dev \
    zlib1g-dev
git clone --depth 1 -b "v1.4.0" "https://github.com/tdlib/td.git" ./tdlib-src
mkdir ./tdlib-src/build
cd ./tdlib-src/build
cmake -DCMAKE_BUILD_TYPE=Release ..
cmake --build .
sudo make install
rm -rf ./../../tdlib-src
```

### Add go-tdlib to packages

```bash
go get github.com/zelenin/go-tdlib/client
```
