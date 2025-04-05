from pydantic import BaseModel

class UserCreate(BaseModel):
    name: str
    email: str
    phone_number: str
    password: str

class UserResponse(UserCreate):
    id: int

    class Config:
        orm_mode = True