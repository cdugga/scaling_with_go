# Stage-2 Production Environment
# Use the nginx 1.12-alpine runtime image for the production environment
FROM nginx:1.12-alpine

ADD ./default.conf /etc/nginx/conf.d/

COPY app-dev.properties /usr/share/nginx/html

# Make port 80 available to the world outside the container
EXPOSE 80

# Run the nginx server
CMD ["nginx", "-g", "daemon off;"]