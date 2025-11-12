from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def home():
    return {"mensagem": "API Python 3.12 rodando no Docker Alpine!"}

