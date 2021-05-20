FROM golang:alpine as builder 
ENV GOPROXY=https://goproxy.io
WORKDIR /build  
ADD . /build/
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o app .

FROM scratch
COPY --from=builder /build/app /
CMD ["/app"]
