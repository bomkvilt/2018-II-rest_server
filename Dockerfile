FROM ubuntu:18.10

ENV PGVER 10
ENV GOVER 1.11

# install psql
RUN apt-get -y update && apt-get -y install sudo wget software-properties-common
RUN sudo add-apt-repository ppa:longsleep/golang-backports
RUN sudo apt-get update && apt-get -y install golang-go
ENV DEBIAN_FRONTEND 'noninteractive'
RUN echo 'Europe/Moscow' > '/etc/timezone' && apt-get install -y postgresql-$PGVER golang-$GOVER git

ENV GOROOT /usr/lib/go-$GOVER
ENV GOPATH '/opt/go'

RUN go get 'github.com/bomkvilt/tech-db-app' && go build -o '${GOPATH}/bin/forum' 'github.com/bomkvilt/tech-db-app/app/generated/cmd/forum-server'

# stup psql
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