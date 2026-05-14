package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"Привет, мир",
	"Hola, mundo",
	"Bonjour, le monde",
	"Hallo, Welt",
	"Ciao, mondo",
	"Olá, mundo",
}

func main() {
	tiempo := time.Now().UnixNano()
	fmt.Println("Tiempo:", tiempo)
	rand.NewSource(tiempo)
	index := rand.Intn(len(helloList))
	fmt.Println("Lista Hello:", len(helloList))
	fmt.Println("Index:", index)
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))

	}
	return helloList[index], nil
}
