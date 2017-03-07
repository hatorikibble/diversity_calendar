FROM golang:onbuild
EXPOSE 6767

HEALTHCHECK --interval=3s --timeout=3s CMD ["/go/src/app/healthchecker/healthchecker", "-port=6767"] || exit 1