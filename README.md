# IP-INFO

Get IP address information from commandline.

_This software uses the [ip-api.com](http://ip-api.com) via [goipapi](https://github.com/belovai/goipapi) client package._

## Compile

```bash
go get github.com/belovai/goipapi github.com/pborman/getopt/v2
go build -o bin/ip-info
```

**Copy the `bin/ip-info` file to somewhere into your path. For example:**

```bash
cp bin/ip-info $HOME/.bin/
```

## Usage

```bash
ip-info 140.82.118.4
```

```bash
ip-info 140.82.118.4 172.217.20.14
```

**Working with IPv6 as well**

```bash
ip-info 2606:4700:10::6814:ef
```

**And the pipeline also working**

```bash
echo '140.82.118.4' | ip-info
```

**Each line should contaion exactly one IP address**
```bash
echo '140.82.118.4\n172.217.20.14' > test.txt
cat test.txt | ip-info
```

## Format

Currently there are 2 format available: json and pretty.

```bash
ip-info -f json 172.217.20.14
```
