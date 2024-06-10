FROM golang:1.20-alpine

# Install necessary packages using apk
RUN apk update && \
    apk add --no-cache \
    build-base \
    imagemagick-dev \
    imagemagick \
    libjpeg-turbo-dev \
    libpng-dev \
    libwebp-dev \
    pkgconfig

WORKDIR /app
COPY . .

# Create the uploads directory
RUN mkdir -p /app/uploads

RUN go build -o main .

# Expose port 8080
EXPOSE 8080

CMD ["./main"]
