FROM alpine
# Install ca-certificates to able to call via https
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
COPY ./app /app/

EXPOSE 8080

ENTRYPOINT ["./app"]