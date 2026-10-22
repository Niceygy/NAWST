FROM debian:bookworm-slim
# Set working directory in the container
WORKDIR /app/

COPY nawst .

LABEL org.opencontainers.image.description="NAWST"
LABEL org.opencontainers.image.authors="Niceygy (Ava Whale)"

RUN chmod 777 ./nawst

# Expose the application port
EXPOSE 8089
# Run the application
CMD ["./nawst"]   