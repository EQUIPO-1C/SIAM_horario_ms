FROM golang:1.19.1-bullseye

ENV APP_HOME /go/schedule_ms
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
RUN touch .env
COPY . .

RUN go install
EXPOSE 4000
CMD ["sh","-c","go run main.go"]
