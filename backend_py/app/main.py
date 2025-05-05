# app/main.py
from fastapi import FastAPI
import strawberry
from strawberry.fastapi import GraphQLRouter
from .db import ping_db
from .schema import Query  # Ensure this points to your GraphQL schema

# Combine schemas
schema = strawberry.Schema(query=Query)

graphql_app = GraphQLRouter(schema)

app = FastAPI()

@app.on_event("startup")
async def startup_event():
    print("Starting up...")
    await ping_db()

app.include_router(graphql_app, prefix="/graphql")

@app.get("/")
def read_root():
    return {"Hello": "Python Backend"}