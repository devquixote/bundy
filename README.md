# Bundy
Lightweight api for killing processes written in Go.  Intended use is for
coordinated shutdown/cleanup of container-based services as they are no longer
needed.  For instance, an integration test running in one container shutting
down the main process in the container running the service(s) under test.

# Installation
```
curl --location --silent --show-error --fail \
     https://github.com/devquixote/bundy/releases/download/0.0.1/bundy
     > /usr/local/bin/bundy
```

or more specifically for docker containers...

```
curl --location --silent --show-error --fail \
     https://github.com/devquixote/bundy/releases/download/0.0.1/bundy
     > /usr/local/bin/bundy && \
     chmod +x /usr/local/bin/bundy
```

# Usage
```
Usage:
  bundy [OPTIONS]

Application Options:
  -a, --address= Network address to bind to (default: 0.0.0.0)
  -p, --port=    Port to listen on (default: 6666)
  -P, --pids=    Pids to kill (default: 1)
  -s, --signals= Signals to send in order (default: TERM, KILL)
  -d, --delay=   Seconds to wait between signals (default: 20)
  -S, --suicide  If bundy should kill itself after killing the target process (default: true)

Help Options:
  -h, --help     Show this help message
```

# Contributing
1. Fork it
2. Hack it
3. Submit a pull request
4. Update your resume
