FROM agencyrevolution/go:1.8.4-alpine3.6

MAINTAINER quyen@agencyrevolution.com

ENV PORT=80
ENV WORKDIR=$GOPATH/src/app/go-app

WORKDIR $WORKDIR

COPY . $WORKDIR

RUN go get

CMD ["go", "run", "main.go"]

EXPOSE 80
