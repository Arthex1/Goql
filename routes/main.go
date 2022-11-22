package routes

import (
	"fmt"
	"io"
	"log"

	"os"

	"github.com/gin-gonic/gin"
)




func Email(g *gin.Context) {		

}

func PostImage(g *gin.Context) {
		
	    file, header , err := g.Request.FormFile("5555555555")
		if err != nil {
			log.Print(err) 
		}
        filename := header.Filename
		
       	fmt.Println(header.Filename)
        out, err := os.Create("./images/"+filename)
		log.Print(header.Size)  
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()
        _, err = io.Copy(out, file)
        if err != nil {
            log.Fatal(err)
        }   
}
