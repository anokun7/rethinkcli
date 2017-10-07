FROM golang:onbuild
ENTRYPOINT ["app"]
CMD ["54.174.187.41:28015", "test", "movies"]
