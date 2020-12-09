FROM ubuntu
RUN mkdir /app
ADD ./adisasson-api-linux-amd64 /app/
WORKDIR /app
EXPOSE 8080 8080
CMD ["/app/adisasson-api-linux-amd64"]
