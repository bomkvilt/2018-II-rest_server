FROM ubuntu:18.10

ENV PGVER 10
ENV GOVER 1.11
ENV DEBIAN_FRONTEND 'noninteractive'
RUN echo 'Europe/Moscow' > '/etc/timezone'

# install psql
RUN apt-get -y update && apt-get -y install sudo wget software-properties-common
RUN sudo add-apt-repository ppa:longsleep/golang-backports
RUN apt-get -y update && apt-get install -y postgresql-$PGVER golang-go git

ENV GOROOT /usr/lib/go-$GOVER
ENV GOPATH '/opt/go'
ENV GO111MODULE on


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
RUN echo "synchronous_commit = off"        >> /etc/postgresql/$PGVER/main/postgresql.conf
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]
EXPOSE 5432

# run 
USER root
EXPOSE 5000

COPY  . /tmp/go
WORKDIR /tmp/go
RUN go build -mod vendor ./cmd/
CMD service postgresql start && cmd 
# --scheme=http --port=5000 --host=0.0.0.0 