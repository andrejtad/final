FROM node:12

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json ./
COPY package-lock.json ./
COPY . ./

RUN npm install --silent

CMD ["npm", "start"]
