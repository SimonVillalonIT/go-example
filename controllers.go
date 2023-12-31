package controllers
// getAlbums responds with the list of all albums as JSON.
func getPersons(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, persons)
}

// postAlbums adds an album from JSON received in the request body.
func postPersons(c *gin.Context) {
    var newPerson person

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newPerson); err != nil {
        return
    }

    // Add the new album to the slice.
    persons = append(persons, newPerson)
    c.IndentedJSON(http.StatusCreated, newPerson)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getPersonByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range persons {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
