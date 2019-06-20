FROM alpine:edge

RUN apk upgrade --update && \
  apk add build-base openssl openssh-client git go gnupg

RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ARG gpg_secret
ENV GPG_SECRET=$gpg_secret


WORKDIR $GOPATH/src/duks

COPY . .
RUN make deps
RUN make decrypt

CMD ["duks"]
