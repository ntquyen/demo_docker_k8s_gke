FROM node:8.7.0-alpine

MAINTAINER quyen@agencyrevolution.com

ENV PORT=80
ENV WORKDIR=/usr/src/app

WORKDIR $WORKDIR

COPY . $WORKDIR

RUN npm install

CMD ["npm", "start"]
