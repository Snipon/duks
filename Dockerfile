FROM alpine:edge

RUN apk upgrade --update && \
  apk add build-base openssl openssh-client git go

RUN git config --global url."git@github.com":.insteadof "https://github.com/"
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


WORKDIR $GOPATH/src/lavfilm-api

COPY . .
RUN make deps

CMD ["duks"]
