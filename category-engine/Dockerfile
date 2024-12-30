# Use the official Python image
FROM python:3.9-slim

# Set the working directory
WORKDIR /app
COPY . .

# Install dependencies
RUN pip install gunicorn
RUN pip install -r requirements.txt

# Make the container's port available to the outside
EXPOSE 8080

# Run the application
CMD exec gunicorn --bind 0.0.0.0:$PORT --workers 1 --threads 8 --timeout 0 app:app