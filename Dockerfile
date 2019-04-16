FROM golang:1.12.4-stretch as build

RUN apt-get update && \
    apt-get install -y build-essential

WORKDIR /src
RUN wget https://github.com/edenhill/librdkafka/archive/v1.0.0.tar.gz && \
    tar xf v1.0.0.tar.gz && \
    cd librdkafka-1.0.0/ && \
    ./configure --prefix=/usr && \
    make && \
    make install

WORKDIR /app
ADD . /app

RUN go build -o kafka-to-redisearch

FROM debian:stretch

COPY --from=build /usr/lib/librdkafka.a /usr/lib/
COPY --from=build /usr/lib/librdkafka++.so.1 /usr/lib/
COPY --from=build /usr/lib/librdkafka.so /usr/lib/
COPY --from=build /usr/lib/librdkafka++.a /usr/lib/
COPY --from=build /usr/lib/librdkafka++.so /usr/lib/
COPY --from=build /usr/lib/librdkafka.so.1 /usr/lib/
COPY --from=build /usr/include/librdkafka /usr/include/

COPY --from=build /app/kafka-to-redisearch /bin/kafka-to-redisearch
CMD ["/bin/kafka-to-redisearch"]
