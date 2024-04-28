FROM node:20.11.1-bullseye

WORKDIR /swagger

RUN npm install -g swagger-cli

COPY . .

CMD ["swagger-cli","bundle", "-o","/swagger/openapi.yml", "-t", "yaml", "/swagger/components/api.yml"]

