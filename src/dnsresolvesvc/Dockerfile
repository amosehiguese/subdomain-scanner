FROM eclipse-temurin:21@sha256:fe90fc98e067d7708811aec14fa60a4b81127d3dc7387aecc4c446c2c30c1277 as builder
WORKDIR /app
COPY ["build.gradle", "gradlew", "./"]
COPY gradle gradle
RUN chmod +x gradlew
RUN ./gradlew downloadRepos

COPY . .
RUN chmod +x gradlew
RUN ./gradlew installDist

FROM eclipse-temurin:21.0.2_13-jre-alpine@sha256:6f78a61a2aa1e6907dda2da3eb791d44ef3d18e36aee1d1bdaa3543bd44cff4b
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app .

EXPOSE 6856
ENTRYPOINT [ "/app/build/install/dnsresolvesvc/bin/dns-resolve-service" ]
