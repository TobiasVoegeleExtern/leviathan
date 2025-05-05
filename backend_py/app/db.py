# app/db.py
from odmantic import AIOEngine
from motor.motor_asyncio import AsyncIOMotorClient
from .config import MONGODB_URI, MONGODB_DATABASE

client = AsyncIOMotorClient(MONGODB_URI)
engine = AIOEngine(client=client, database=MONGODB_DATABASE)

async def get_engine():
    return engine

async def ping_db():
    try:
        await client.admin.command('ping')
        print("MongoDB connection successful!")
    except Exception as e:
        print(f"MongoDB connection failed: {e}")