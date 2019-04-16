FROM scratch

LABEL maintainer="glug71@gmail.com"

COPY bin/printenv /printenv
EXPOSE 8080
ENTRYPOINT ["/printenv"]
