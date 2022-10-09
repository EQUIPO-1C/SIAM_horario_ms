package configs

import (
    "log"
    "os"
    "fmt"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    user := os.Getenv("DB_USER")
    pswd := os.Getenv("DB_PSWD")
    uriMongo := fmt.Sprintf("mongodb+srv://%s:%s@procesoinscluster.kb4ixgp.mongodb.net/?retryWrites=true&w=majority",user,pswd)
    return uriMongo
}