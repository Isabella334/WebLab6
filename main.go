package main // Ejecutable independiente

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type serie struct{
	ID string
	Title string
	Episodes int
	Rating float32

}

func main() {
	
}