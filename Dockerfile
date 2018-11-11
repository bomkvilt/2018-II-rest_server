FROM ubuntu:18.10

ENV PGVER 10
ENV GOVER 1.10

RUN apt-get -y update
ENV DEBIAN_FRONTEND 'noninteractive'
RUN echo 'Europe/Moscow' > '/etc/timezone' && \
    apt-get install -y postgresql-$PGVER golang-$GOVER

ENV GOROOT /usr/lib/go-$GOVER
ENV GOPATH '/opt/go'

COPY . $GOPATH/src/github.com/bomkvilt/tech-db-app
RUN go build -o '${GOPATH}/bin/forum' 'github.com/bomkvilt/tech-db-app/app/generated/cmd/forum-server'

USER postgres
COPY scheme.sql .
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
    createdb -O docker docker &&\
    psql -d docker -f scheme.sql &&\
    /etc/init.d/postgresql stop
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf
RUN echo "listen_addresses='*'"            >> /etc/postgresql/$PGVER/main/postgresql.conf
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]
EXPOSE 5432

# run 
USER root

EXPOSE 5000

CMD service postgresql start && forum --port 5000