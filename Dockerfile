FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /prometheus_exporter .


FROM alpine:latest
COPY --from=build /prometheus_exporter /prometheus_exporter
ENTRYPOINT ["/prometheus_exporter"]
