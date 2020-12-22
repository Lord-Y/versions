ARG DOCKER_IMAGE_NODEJS="mirror.gcr.io/library/node:14-alpine"
FROM ${DOCKER_IMAGE_NODEJS} as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN adduser -D -g '' appuser
RUN rm -rf node_modules .git .env*
RUN apk update && apk upgrade
RUN apk add --no-cache git curl && rm -rf /var/cache/apk/*
RUN npm install
RUN npm run ssr:build

ARG DOCKER_IMAGE_NODEJS="mirror.gcr.io/library/node:14-alpine"
FROM ${DOCKER_IMAGE_NODEJS}
RUN apk update && apk upgrade
RUN apk add --no-cache git curl && rm -rf /var/cache/apk/*
RUN mkdir /app
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/*.config.js ./
RUN npm install --only=production
# The node-prune is an open-source tool for removing unnecessary files from the node_modules folder. Most of the developers forget test files, markdown files, typing files and *.map files in Npm packages. By using node-prune we can safely delete them
# install node-prune (https://github.com/tj/node-prune)
#RUN curl -sfL https://install.goreleaser.com/github.com/tj/node-prune.sh -o /usr/local/bin/node-prune.sh && chmod +x /usr/local/bin/node-prune.sh
RUN npm prune --production
RUN curl -sfL https://install.goreleaser.com/github.com/tj/node-prune.sh | sh -s -- -b /usr/local/bin
RUN /usr/local/bin/node-prune
RUN npm cache clean --force
#RUN for folder in src; do find node_modules/ -type d -name $folder -exec rm -rf {} \;;done; echo "unnecessary folders deleted"
USER appuser
EXPOSE 8080
CMD ./node_modules/@uvue/server/start.js
