FROM node:alpine AS builder

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm ci

ENTRYPOINT ["npm", "run", "dev"]