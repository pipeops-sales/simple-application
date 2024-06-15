# Use the official Nginx image as the base image
FROM nginx:alpine

# Copy the index.html file to the default Nginx HTML location
COPY . /usr/share/nginx/html/

# Expose port 80 to the outside world
EXPOSE 80

# Start Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]
