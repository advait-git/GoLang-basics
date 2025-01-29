package conf

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Debugging: Print loaded environment variables
	// fmt.Println("🔍 Loaded Environment Variables:")
	// fmt.Println("MONGO_URI:", os.Getenv("MONGO_URI"))
	// fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	// fmt.Println("COLLECTION_NAME:", os.Getenv("COLLECTION_NAME"))

	// Ensure MONGO_URI is loaded
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("❌ MONGO_URI is not set or is empty in .env file")
	}

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ MongoDB Connection Error:", err)
	}

	// Ping MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB Ping Failed:", err)
	}

	// Assign connected client to global Client variable
	Client = client
	log.Println("✅ Connected to MongoDB successfully!")
}

// GetCollection returns the MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("❌ MongoDB client is nil. Ensure ConnectDB() is called before accessing collections.")
	}
	dbName := os.Getenv("DB_NAME")
	return Client.Database(dbName).Collection(collectionName)
}
