package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)


type Movie struct {
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'

}
type Direrctor{

}