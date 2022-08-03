# syntax=docker/dockerfile:1

# start react app
FROM node:17.0.0
WORKDIR /app
COPY . .
RUN npm install

EXPOSE 3000

CMD ["npm", "start"]