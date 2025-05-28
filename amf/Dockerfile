FROM golang:1.24 
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/amf ./cmd/main.go


FROM scratch
COPY --from=0 /bin/amf /bin/amf
CMD ["/bin/amf"]