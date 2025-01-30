FROM node:23 as build-frontend

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./

RUN npm install

COPY frontend/ .

RUN npm run build

FROM golang:1.23 as build-backend

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ .

# disable dynamically linking 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

FROM alpine:3.14

ARG VERSION
ARG GIT_COMMIT
ARG GH_REPO
ENV VERSION=$VERSION
ENV GIT_COMMIT=$GIT_COMMIT
ENV GH_REPO=$GH_REPO


WORKDIR /app

RUN apk --no-cache add git

COPY --from=build-backend /app/main .

COPY --from=build-frontend /app/dist ./public

EXPOSE 8080

CMD ["./main"]