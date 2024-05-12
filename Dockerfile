FROM golang:1.22-alpine3.19 as build-stage

WORKDIR /projects/bebastukar-be

COPY . .

RUN go mod download

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o bebastukar ./cmd/api/

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /projects/bebastukar-be/web/ /web/
COPY --from=build-stage /projects/bebastukar-be/docs/ /docs/
COPY --from=build-stage /projects/bebastukar-be/config.yaml /config.yaml
COPY --from=build-stage /projects/bebastukar-be/bebastukar /bebastukar

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./bebastukar"]
