package routes

import (
	
	"io"
	"log"
	"path/filepath"

	"os"
	"strings"

	"github.com/gin-gonic/gin"
)




func Email(g *gin.Context) {		

}

func check(path string, name string) bool {
	exists := false 
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
			log.Print(err)
        }
		
		exists = strings.Contains(info.Name(), name) 
		return nil 
        
    })
	return exists 
	
}

func PostImage(g *gin.Context) {
		id := g.Param("userid")
		
		update := func(gs *gin.Context) bool {
			update_query := gs.Query("update")
			if strings.ToLower(update_query) == "true" {
				return true 
			}
			return false 
		}(g) 


		if (!update) {
		  if p := check("./images", id); p {
				
				g.JSON(502, gin.H{
					"error": "Already exists", 
				})
				return 
		  }
		}

	    file, header, err := g.Request.FormFile(id)


		if err != nil {
			log.Print(err) 
			g.JSON(502, gin.H{
				"error": err.Error(),
			})
			return 
		}
		log.Print(header.Size)

        out, err := os.Create("./images/"+id+".png") 

        if err != nil {
            log.Print(err)
			g.JSON(502, gin.H{
				"error": err.Error(),
			})
			return 
        }
        defer out.Close()

        _, err = io.Copy(out, file)


        if err != nil {
            log.Print(err)
			g.JSON(502, gin.H{
				"error": err.Error(),
			})
			return 
    	}


		g.Status(202) 
}
