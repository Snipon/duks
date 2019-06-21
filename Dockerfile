FROM alpine:edge

RUN apk upgrade --update && \
  apk add build-base openssl openssh-client git go gnupg

RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH/src/duks

ARG gpg_secret
ENV GPG_SECRET=$gpg_secret
RUN env

COPY . .
RUN make decrypt
RUN make deps

CMD ["duks"]
