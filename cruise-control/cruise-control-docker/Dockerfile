FROM eclipse-temurin:11-jdk-focal as cruisecontrol
ARG CRUISE_CONTROL_VERSION=2.5.123
USER root
WORKDIR /
RUN \
    set -xe; \
    apt-get update -qq \
    && apt-get install -qq --no-install-recommends \
    git ca-certificates
RUN \
    set -xe; \
    git clone \
    --branch ${CRUISE_CONTROL_VERSION} \
    --depth 1 \
    https://github.com/linkedin/cruise-control.git
RUN \
    set -xe; \
    cd cruise-control \
    && ./gradlew jar copyDependantLibs \
    && mv -v /cruise-control/cruise-control/build/libs/cruise-control-${CRUISE_CONTROL_VERSION}.jar \
    /cruise-control/cruise-control/build/libs/cruise-control.jar \
    && mv -v /cruise-control/cruise-control/build/dependant-libs/cruise-control-metrics-reporter-${CRUISE_CONTROL_VERSION}.jar \
    /cruise-control/cruise-control/build/dependant-libs/cruise-control-metrics-reporter.jar

FROM node:18.14-buster as cruisecontrol-ui
ARG CRUISE_CONTROL_UI_GIT_REF=v0.4.0
WORKDIR /
RUN \
    set -xe; \
    git clone \
    https://github.com/linkedin/cruise-control-ui.git \
    && cd cruise-control-ui \
    && git checkout ${CRUISE_CONTROL_UI_GIT_REF} \
    && git rev-parse HEAD \
    && npm install \
    && npm run build

FROM eclipse-temurin:11-jre
ENV CRUISE_CONTROL_LIBS="/var/lib/cruise-control-ext-libs/*"
ENV CLASSPATH="${CRUISE_CONTROL_LIBS}"
WORKDIR /opt/cruise-control/
COPY --from=cruisecontrol /cruise-control/cruise-control/build/libs/cruise-control.jar cruise-control/build/libs/cruise-control.jar
COPY --from=cruisecontrol /cruise-control/config config
COPY --from=cruisecontrol /cruise-control/kafka-cruise-control-start.sh .
COPY --from=cruisecontrol /cruise-control/cruise-control/build/dependant-libs cruise-control/build/dependant-libs
COPY --from=cruisecontrol-ui /cruise-control-ui/dist cruise-control-ui/dist
COPY opt/cruise-control .
RUN \
    set -xe; \
    echo "local,localhost,/kafkacruisecontrol" > cruise-control-ui/dist/static/config.csv \
    && chmod +x start.sh
EXPOSE 8090
CMD ["/opt/cruise-control/start.sh"]
