FROM google/cloud-sdk:latest
WORKDIR /app
COPY . .
RUN chmod +x start.sh
EXPOSE 8000
CMD ["./start.sh"]