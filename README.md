# preflight-netpath

a preflight check for a simple tcp network connection.

## Build

```bash
make
```

## Install

NOTE: you will need `curl`, `bash`, and `jq` installed for the install script to work. It will attempt to install the binary in `/usr/local/bin` and will require `sudo` access. You can override the install directory by setting the `INSTALL_DIR` environment variable.

```bash
curl -sSL https://raw.githubusercontent.com/robertlestak/preflight-netpath/main/scripts/install.sh | bash
```

## Usage

```bash
Usage of preflight-netpath:
  -config string
        path to config file
  -endpoint string
        endpoint to test in the form of <host>:<port>
  -log-level string
        log level (default "info")
  -timeout duration
        timeout in seconds (default 5s)
```

## Example

```bash
preflight-netpath -endpoint google.com:443
```

### Docker example

```bash
docker run --rm robertlestak/preflight-netpath \
      -endpoint google.com:443
```

## Config file

You can also use a config file rather than cli args.

```yaml
endpoint: google.com:443
timeout: 5s
```

```bash
preflight-netpath -config config.yaml
```