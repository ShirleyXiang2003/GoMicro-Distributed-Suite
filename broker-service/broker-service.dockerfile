# Build a tiny docker image
FROM alpine:latest


# Create a directory for the app
RUN mkdir /app


# Copy the built binary from the builder stage to the final image
COPY brokerApp /app


# Specify the command to run when the container starts
CMD [ "/app/brokerApp" ]


# dockerfile => image => container