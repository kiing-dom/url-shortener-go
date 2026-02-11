## project plan

### what it does
- shorten urls
- redirect
- track clicks

### constraints
- must run in kubernetes
- horizontal scaling
- stateless api layer

### what is state?
- url mappings
- click counts

&rightarrow; meaning external storage (DB)

### non functional requirements (NFRs)
- health checks
- observability
- graceful shutdown
- config via env variables

## implementation step 1
ignoring kubernetes

set up:
- go module
- router
- db conn
- basic handlers
- example structure:
  ```
    cmd/
    internal/
      handler/
      service/
      repository/
      model/
  ```
separating:
- http layer
- business logic
- persistence

## production shaping
if it crashed in prod what would i want?

- to be able to check health (see status, error logs, uptime)
- graceful shutdown
- health endpoints:
  - /healthz (is process alive?)
  - /readyz (is db connected?)

## containerization
what does it need to run?
- binary
- env variables
- no local files
- external db connection

create:
- multi stage docker build
- small final image
- non-root user if possible

note: if it runs locally within docker, kubernetes becomes easier