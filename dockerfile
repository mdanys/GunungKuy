FROM golang:latest

##buat folder APP
RUN mkdir /gunungkuy

##set direktori utama
WORKDIR /gunungkuy

##copy seluruh file ke completedep
ADD . /gunungkuy

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]
