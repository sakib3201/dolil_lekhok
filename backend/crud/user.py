from sqlalchemy.orm import Session
from backend.models.user import User
from backend.schemas.user import UserCreate

def create_user(db: Session, user: UserCreate):
    db_user = User(name=user.name, email=user.email, phone_number=user.phone_number, password = user.password)
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user

def get_users(db: Session):
    return db.query(User).all() # to get all db user temporary for test

