FROM --platform=linux/amd64 golang:alpine
RUN mkdir "/opt/emprepository"
WORKDIR /opt/emprepository
COPY . ./
CMD ["go","run","."]
