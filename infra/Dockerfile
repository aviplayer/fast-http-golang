FROM alpine:latest

COPY app app
RUN chmod +x ./app
EXPOSE 8089
CMD ["/bin/sh", "-c", "./app"]
