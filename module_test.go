package module

import (
	"fmt"
	"os"
	"testing"
)

var privatekey = ""
var publickey = ""
var encode = ""
var dbname = "GIS"
var collname = "geojson"

func TestGeneratePaseto(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("Private Key: " + privateKey)
	fmt.Println("Public Key: " + publicKey)
}

func TestEncode(t *testing.T) {
	name := "Test Nama"
	username := "Test Username"
	role := "Test Role"

	tokenstring, err := Encode(name, username, role, privatekey)
	fmt.Println("error : ", err)
	fmt.Println("token : ", tokenstring)
}

func TestDecode(t *testing.T) {
	pay, err := Decode(publickey, encode)
	name := DecodeGetName(publickey, encode)
	username := DecodeGetUsername(publickey, encode)
	role := DecodeGetRole(publickey, encode)

	fmt.Println("name :", name)
	fmt.Println("username :", username)
	fmt.Println("role :", role)
	fmt.Println("err : ", err)
	fmt.Println("payload : ", pay)
}

func TestGetAllUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	datagedung := GetAllUser(mconn, "user")
	fmt.Println(datagedung)
}
func TestCobaUsernameExists(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	var user User
	user.Username = "megah123"
	datagedung := UsernameExists(mconn, "user", user)
	fmt.Println(datagedung)
}

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			107.59438892887096, -6.9928869227076405,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{107.59438892887096, -6.9928869227076405},
				{107.5943782000349, -6.993685599063639},
				{107.59544571922281, -6.993648327530735},
				{107.59438892887096, -6.9928869227076405},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", dbname, collname)
	coordinates := Point{
		Coordinates: []float64{
			107.59438892887096, -6.9928869227076405,
		},
	}
	datagedung := Near(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			107.59438892887096, -6.9928869227076405,
		},
	}
	datagedung := NearSphere(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polyline{
		Coordinates: [][]float64{
			{107.59438892887096, -6.9928869227076405},
			{107.5943782000349, -6.993685599063639},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestFindUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	datagedung := FindUser(mconn, collname, User{Username: "megah"})
	tokenstring, tokenerr := Encode(datagedung.Name, datagedung.Username, datagedung.Role, os.Getenv("privatekey"))
	if tokenerr != nil {
		fmt.Println("Gagal encode token: " + tokenerr.Error())
	}
	fmt.Println(tokenstring)
}

func TestGetAllBangunan(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	x := GetAllBangunan(mconn, collname)
	fmt.Println(x)
}

func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	var user User
	user.Name = "test"
	InsertUser(mconn, "test", user)
}
