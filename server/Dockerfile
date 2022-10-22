FROM golang:1.18

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && apt-get install postgresql-client libpq-dev -y \
                    && apt-get install default-mysql-client -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]