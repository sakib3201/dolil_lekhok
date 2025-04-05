from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from backend.schemas.user import UserCreate, UserResponse
from backend.crud.user import create_user, get_users
from backend.core.database import get_db

router = APIRouter()

@router.post("/", response_model=UserResponse)
def create(user: UserCreate, db: Session = Depends(get_db)):
    return create_user(db, user)

@router.get("/", response_model=list[UserResponse])
def read_all(db: Session = Depends(get_db)):
    return get_users(db)
