FROM docker.io/library/nginx:1.23.3

COPY default.conf /etc/nginx/conf.d/default.conf
COPY index.html script.js /usr/share/nginx/html/
COPY index.html favicon.ico /usr/share/nginx/html/

ARG SERVER_ADDR=localhost:1323

RUN sed 's#^\(const SERVER_ADDR\).*$#\1 = "'$SERVER_ADDR'"#' -i /usr/share/nginx/html/script.js

EXPOSE 8888
