FROM node:lts AS frontend
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV CI="true"
RUN corepack enable
COPY web /app
WORKDIR /app

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm build

FROM golang:1.25.3 AS backend
WORKDIR /usr/src/app
COPY go.mod go.sum .
RUN go mod download

COPY . .
COPY --from=frontend /app/dist /web/dist

RUN go build -v -o /usr/local/bin/app main.go

EXPOSE 8000
CMD [ "app" ]
