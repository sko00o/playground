FROM nvidia/cuda:11.8.0-base-ubuntu22.04

RUN apt-get update && apt-get install -y \
    python3-pip \
    python3-dev \
    libglfw3-dev \
    libgl1-mesa-dev \
    libgles2-mesa-dev \
    libegl1-mesa-dev \
    xvfb \
    && rm -rf /var/lib/apt/lists/*

COPY requirements.txt .
RUN pip3 install --no-cache-dir -r requirements.txt

WORKDIR /app
COPY . .

ENV DISPLAY=:99
CMD ["python3", "main.py"]
