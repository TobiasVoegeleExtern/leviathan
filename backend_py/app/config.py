# Beispiel für app/config.py oder direkt in db.py
import os
from dotenv import load_dotenv # Nur für lokale .env Dateien, in Docker sind die Variablen gesetzt

# Lade .env nur wenn lokal entwickelt wird (optional)
# load_dotenv()

MONGODB_URI = os.getenv("MONGODB_URI", "mongodb://localhost:27017") # Default für lokale Entwicklung ohne Docker
MONGODB_DATABASE = os.getenv("MONGODB_DATABASE", "default_db")
