FROM node:22-alpine

WORKDIR /app

# Если package.json не существует, создать проект
RUN if [ ! -f package.json ]; then \
    npm create vite@latest . -- --template vue; \
  fi

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
