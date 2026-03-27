FROM kprzystalski/projobj-php:latest

WORKDIR /home/student/projobj

USER root

COPY . .

RUN composer install --no-dev --optimize-autoloader

EXPOSE 8000

CMD php bin/console doctrine:migrations:migrate --no-interaction && \
    symfony server:start --allow-http --port=8000 --no-tls
