package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shubash/saibaba/moddel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const connectingstring = "mongodb+srv://shubashduttasaibaba:LE19u3bje3SmmDl0@cluster0.ovbu7ak.mongodb.net/"

const dbname = "SaiBABA"
const colname = "user"

var Collection *mongo.Collection

func init() {
	// Load environment variables from .env file

	// Set up client options
	clientOptions := options.Client().ApplyURI(connectingstring)

	// Create a new MongoDB client
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Set the Collection variable to be used throughout the application
	Collection = client.Database("SaiBABA").Collection("user")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to sai baba turst web site")
}
func Singup(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") // Add this line to allow cross-origin requests
    
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user moddel.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    count, err := Collection.CountDocuments(context.Background(), bson.M{"email": user.Email})
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    if count != 0 {
        http.Error(w, "Email is already in use", http.StatusBadRequest)
        return
    }

    num, err := Collection.CountDocuments(context.Background(), bson.M{"phone": user.Phone})
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    if num != 0 {
        http.Error(w, "Phone number is already in use", http.StatusBadRequest)
        return
    }

    password := Hashpassword(user.Password)
    user.Password = password

    insertoneuser(user)

    // Encode and send the user data as the response
    err = json.NewEncoder(w).Encode(user)
    if err != nil {
        log.Println(err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json-x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user moddel.User

	var founduser moddel.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	err1 := Collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&founduser)
	if err1 != nil {
		// log.Println(err1)
		log.Println("email is incorect")
		// json.NewEncoder(w).Encode("email incorrect")
		return

	}
	err2 := bcrypt.CompareHashAndPassword([]byte(founduser.Password), []byte(user.Password))
	if err2 != nil {
		log.Println("password does not match")
		return
	}

	json.NewEncoder(w).Encode(founduser)
	fmt.Println("user is on")
}

func Hashpassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
func insertoneuser(user moddel.User) {
	insert, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted one user in our database", insert.InsertedID)
}
