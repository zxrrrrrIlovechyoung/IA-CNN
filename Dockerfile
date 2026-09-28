FROM alpine:latest

WORKDIR /app

COPY dist/proyectofoco /app/proyectofoco
COPY static /app/static

ENV PORT=8089
EXPOSE 8089

CMD ["/app/proyectofoco"]
