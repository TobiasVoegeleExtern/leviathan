# Stage 1: Build the application
FROM node:20-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the SvelteKit application
RUN npm run build

# Stage 2: Run the application in production
FROM node:20-alpine

# Set working directory inside the container
WORKDIR /app

# Copy built application from builder stage
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./
COPY --from=builder /app/node_modules ./node_modules

# Expose port 3000 for the application
EXPOSE 3000

# Run the application
CMD ["node", "build"]
