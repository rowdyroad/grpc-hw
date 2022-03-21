FROM node:current-buster-slim as build
COPY ./web ./app
WORKDIR ./app
RUN npm install && yarn build

FROM nginx as app
COPY --from=build /app/build /usr/share/nginx/html
COPY ./deploy/web/default.conf /etc/nginx/conf.d
