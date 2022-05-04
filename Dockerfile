FROM ubuntu:20.10
RUN mkdir /app
COPY appdemo /app/
WORKDIR /app
CMD ["/app/appdemo"]
EXPOSE 80