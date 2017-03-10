FROM iron/base
EXPOSE 6767

# export CGO_ENABLED=0 



ADD diversity_calendar /
ADD healthchecker/healthchecker /

HEALTHCHECK --interval=3s --timeout=3s CMD ["./healthchecker", "-port=6767"] || exit 1

ENTRYPOINT ["./diversity_calendar"]