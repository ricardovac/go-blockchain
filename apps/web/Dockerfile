FROM node:lts-alpine as build-stage

WORKDIR /app

ARG VITE_BASE_PATH
ENV VITE_BASE_PATH=${VITE_BASE_PATH}
ARG VITE_SERVER_PORT
ENV VITE_SERVER_PORT=${VITE_SERVER_PORT}

RUN corepack enable

COPY apps/web/package*.json /app/apps/web/

RUN cd /app/apps/web && pnpm install --no-frozen-lockfile

COPY apps/web /app/apps/web

RUN cd /app/apps/web && pnpm run build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/apps/web/dist /usr/share/nginx/html
EXPOSE 8888
CMD ["nginx", "-g", "daemon off;"]