FROM node:20.12.2-alpine@sha256:ec0c413b1d84f3f7f67ec986ba885930c57b5318d2eb3abc6960ee05d4f2eb28 as base
FROM base as builder
WORKDIR /usr/src/app
COPY package.json ./
RUN npm install -g yarn
RUN yarn install --only=production

FROM base
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/node_modules ./node_modules
COPY . .
EXPOSE 50051
ENTRYPOINT [ "yarn", "run", "server" ]
