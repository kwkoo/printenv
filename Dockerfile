FROM scratch

LABEL maintainer="glug71@gmail.com"

COPY bin/printenv /printenv

ENTRYPOINT ["/printenv"]
