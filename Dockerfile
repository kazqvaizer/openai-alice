FROM golang:1.23-alpine

RUN apk add --no-cache libc6-compat 

ENV HOST 0.0.0.0
ENV PORT 8080

RUN mkdir -p /srv
WORKDIR /srv

RUN touch /srv/.env
COPY dist/openai_alice .

EXPOSE 8080

CMD ["./openai_alice"]
