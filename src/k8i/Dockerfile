
ARG GO_VERSION=1.23.0

FROM golang:${GO_VERSION}-alpine AS build
RUN apk add --no-cache git
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /main ./main.go
 

FROM gcr.io/distroless/static AS final
LABEL maintainer="luis"
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /main /main
ENTRYPOINT ["/main"]