FROM scratch
MAINTAINER Manuel Laufenberg <hello@manuel-laufenberg.de>
COPY dist/ret /
ENTRYPOINT ["/ret"]