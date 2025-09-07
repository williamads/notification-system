FROM golang:1.25.1 as stage1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o chatinho

#################################################

FROM scratch

COPY --from=stage1 /app/chatinho /

ENTRYPOINT [ "/chatinho" ]