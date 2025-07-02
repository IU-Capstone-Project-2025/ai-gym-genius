FROM debian:buster-slim AS base
WORKDIR /app

RUN apt-get update && apt-get install -y git curl unzip && rm -rf /var/lib/apt/lists/*

FROM base AS build
ARG FLUTTER_VERSION=stable
ARG FLUTTER_DIR=/usr/local/flutter

RUN git clone --depth 1 --branch "$FLUTTER_VERSION" https://github.com/flutter/flutter.git "$FLUTTER_DIR"
ENV PATH="$PATH:$FLUTTER_DIR/bin"
RUN flutter doctor

FROM build AS dependencies
COPY pubspec.yaml ./
RUN flutter pub get
COPY . .

FROM dependencies AS compile
ARG API_URL

RUN BUILD_VERSION=$(date +%Y%m%d-%H%M) && \
    flutter build web ${API_URL:+"--dart-define=API_URL=$API_URL"} && \
    sed -i "s/{{flutter_service_worker_version}}/$BUILD_VERSION/" /app/build/web/index.html && \
    sed -i "s/main.dart.js/main.dart.js?v=$BUILD_VERSION/" /app/build/web/index.html && \
    echo "<script>window.buildVersion = '$BUILD_VERSION';</script>" >> /app/build/web/index.html

FROM nginx:alpine AS production

ARG PORT=80
EXPOSE $PORT

RUN touch /var/run/nginx.pid && chown -R nginx:nginx /var/cache/nginx/ \
/etc/nginx/conf.d /usr/share/nginx/html /var/run/nginx.pid /var/log/nginx && \
sed -i "s/listen  .*/listen $PORT;/g" /etc/nginx/conf.d/default.conf

COPY --from=compile --chown=nginx:nginx /app/build/web/ /usr/share/nginx/html/
USER nginx

CMD ["nginx", "-g", "daemon off;"]
