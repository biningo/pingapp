FROM busybox:1.32.0
EXPOSE 8080
COPY ["./main","./"]
CMD ["./main"]
