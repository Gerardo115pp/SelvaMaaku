FROM node:21-alpine3.18 
WORKDIR /app

COPY ./package.json /app
COPY ./package-lock.json /app
COPY ./build /app

RUN which npm ; \
    if [ $? -ne 0 ]; then \
        apk add --no-cache npm; \
    fi

RUN apk add --no-cache \
    tzdata \
    ca-certificates;

RUN npm ci --omit dev

EXPOSE 1366

CMD ["node", "."]

