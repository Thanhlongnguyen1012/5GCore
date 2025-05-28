FROM golang:1.24 
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/udm ./cmd/main.go


FROM scratch
COPY --from=0 /bin/udm /bin/udm
CMD ["/bin/udm"]