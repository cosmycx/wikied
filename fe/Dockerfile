# -------------------
# step 1 -> build stage
# -------------------
FROM node:latest as builder
WORKDIR /app
COPY . ./
RUN npm install
# needed: from git issues
RUN npm rebuild node-sass
# build app in default build folder
RUN npm run-script build

# -------------------
# step 2 -> run stage
# -------------------
FROM nginx:latest
# copy build folder and assets
COPY --from=builder app/build /usr/share/nginx/html
# file permissions 
RUN chmod 755 -R /usr/share/nginx/html
# exposed port default nginx
EXPOSE 80
