from fastapi import FastAPI
from backend.api.v1 import users
from backend.core.database import Base, engine

# Create tables (or use Alembic for migrations)
Base.metadata.create_all(bind=engine)

app = FastAPI()

app.include_router(users.router, prefix="/api/v1/users", tags=["Users"])
