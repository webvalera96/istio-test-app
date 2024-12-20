FROM golang:1.22
WORKDIR /src

COPY . /src

RUN CGO_ENABLED=0 go build -o /bin/cs main.go

FROM fedora:41

COPY --from=0 /bin/cs /bin/cs
ENTRYPOINT ["/bin/cs"]