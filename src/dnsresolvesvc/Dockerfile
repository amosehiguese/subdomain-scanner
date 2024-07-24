FROM eclipse-temurin:8u412-b08-jdk as builder
WORKDIR /app
COPY ["build.gradle", "gradlew", "./"]
COPY gradle gradle
RUN chmod +x gradlew
RUN ./gradlew downloadRepos

COPY . .
RUN chmod +x gradlew
RUN ./gradlew installDist

FROM eclipse-temurin:8u412-b08-jre-alpine
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app .

EXPOSE 50053
ENTRYPOINT [ "/app/build/install/dnsresolvesvc/bin/dns-resolve-service" ]
