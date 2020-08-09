# Build stage
FROM golang:1.13.4-alpine3.10 as codecademy-build
RUN apk update && apk add --no-cache \
    git \
    bash \
    make

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN make build

# Final stage
FROM alpine:3.10 as codecademy-dev
RUN apk update && apk add --no-cache \
    bash
EXPOSE 8080
WORKDIR /
COPY --from=codecademy-build /app /
CMD ["/codecademy"]