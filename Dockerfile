# dockerfile for go?
FROM node:14 as AssetBuilder
WORKDIR /app
COPY package.json ./package.json
RUN npm install --silent
# Copy the rest of the project's files over to this image so that we can build it
COPY . ./
RUN ["npm", "run", "build"]
## Create a new image based off the `node:alpine` image, since we want
## to use Go in our production environment.  We specify `/go/src/github.com



FROM golang:1.21.6 as GOBUILDER

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Use alpine to make the image smaller in size
FROM alpine:3.15 AS FINAL

WORKDIR /app
COPY --from=GOBUILDER /app/templates ./templates
COPY --from=GOBUILDER /app/static ./static
COPY --from=AssetBuilder /app/src ./src
COPY --from=AssetBuilder /app/dist ./dist
  
COPY --from=GOBUILDER /app/myapp  .

COPY  .env .

EXPOSE 3500 

ENTRYPOINT [ "./myapp" ]

# multi statge docker file for golang?