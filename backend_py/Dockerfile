# Temporär vereinfachter Dockerfile (Single-Stage)
FROM python:3.8-slim

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN pip install --upgrade pip

COPY requirements.txt .

# Installiere direkt in dieses Image
RUN pip install --no-cache-dir -r requirements.txt

COPY ./app /app/app

EXPOSE 8000

# Default command (wird von docker-compose überschrieben)
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]