# ======================
#  GO FIRST STAGE
# ======================

FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git make
USER ${USER}
WORKDIR /usr/src/app
COPY go.mod \
  go.sum ./
RUN go mod download
COPY . ./
ENV GO111MODULE="on" \
  GOARCH="amd64" \
  GOOS="linux" \
  CGO_ENABLED="0"
RUN make goprod
RUN ls

# ======================
#  GO FINAL STAGE
# ======================

FROM scratch
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/.env /usr/src/app/
COPY --from=builder /usr/src/app/cert/* /usr/src/app/cert/
COPY --from=builder /usr/src/app/main /usr/src/app/
EXPOSE 4000
CMD ["./main"]