FROM golang:1.24 
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/smf ./cmd/main.go


FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /bin/smf /bin/smf
CMD ["/bin/smf"]