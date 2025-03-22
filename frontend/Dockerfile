# Use Node.js as the base image for development
FROM node:18

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Expose the port for Vite
EXPOSE 4200

# Start the application in development mode
CMD ["npm", "run", "dev"]
