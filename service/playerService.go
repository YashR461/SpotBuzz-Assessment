package service

import (
	"strings"
	"strconv"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	model "yash_rastogi/SpotBuzz-Assessment/src/model"
	mysql "yash_rastogi/SpotBuzz-Assessment/persistence"
	log "github.com/sirupsen/logrus"
	"database/sql"
)

// AddPlayer fetches all player by using a POST request.
func AddPlayer(c *gin.Context) {
	var newPlayer model.Player
	// Call BindJSON to bind the received JSON to newPlayer
	if err := c.BindJSON(&newPlayer); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new player into the database
	_, err := mysql.DB.Exec("INSERT INTO players (Name, Country, Score) VALUES (?, ?, ?)",
	newPlayer.Name, newPlayer.Country, newPlayer.Score)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add player to database"})
		return
	}

	// Returns the message
	c.JSON(http.StatusOK, gin.H{"message": "Player added successfully"})
}

// UpdatePlayerByID updates the name and score attribute of a player by ID using a PUT request.
func UpdatePlayerByID(c *gin.Context) {
	// Parse the playerID value from the URL parameter
	playerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	var updatedPlayer model.Player
	// Call BindJSON to bind the received JSON to newPlayer
	if err := c.ShouldBindJSON(&updatedPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Check for additional fields in the JSON request
	if updatedPlayer.Country != "" || updatedPlayer.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only name and score can be updated"})
		return
	}

	// Create an update query based on the fields provided in the request
	updateQuery := "UPDATE players SET"
	updateFields := []string{}

	if updatedPlayer.Name != "" {
		updateFields = append(updateFields, "Name = '"+updatedPlayer.Name+"'")
	}

	if updatedPlayer.Score != 0 {
		updateFields = append(updateFields, "Score = "+strconv.Itoa(updatedPlayer.Score))
	}

	// Construct the final query
	updateQuery += " " + strings.Join(updateFields, ", ") + " WHERE ID = " + strconv.Itoa(playerID)

	// Execute the update query
	result, err := mysql.DB.Exec(updateQuery)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Returns the message
	c.JSON(http.StatusOK, gin.H{"message": "Player updated successfully"})
}


// DeletePlayerByID deletes a player by ID using a DELETE request.
func DeletePlayerByID(c *gin.Context) {

	// Parse the playerID value from the URL parameter
	playerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	// Query the database to retrieve the player with the specified ID
	result, err := mysql.DB.Exec("DELETE FROM players WHERE ID = ?", playerID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Returns the message
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}

// GetPlayers fetches all players by using a GET request.
func GetPlayers(c *gin.Context) {
	// Query the database to in order to retrieve the players in descending order according to the score
	rows, err := mysql.DB.Query("SELECT * FROM players order by score desc")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var players []model.Player

	for rows.Next() {

		var player model.Player

		// Scan the retrieved row data into the player instance
		if err := rows.Scan(&player.ID, &player.Name, &player.Country, &player.Score); err != nil {
			log.Fatal(err)
			return
		}
		//Appends player(new player) to players 
		players = append(players, player)
	}

	// Returns all the players data in JSON format
	c.IndentedJSON(http.StatusOK, players)
}

// GetPlayerByRank fetches a player by their rank using a GET request.
func GetPlayerByRank(c *gin.Context) {

	// Parse the rank value from the URL parameter
	rankVal, err := strconv.Atoi(c.Param("val"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rank value"})
		return
	}

	var player model.Player

	// Query the database to retrieve the player with the specified rank
	row := mysql.DB.QueryRow("SELECT * FROM players ORDER BY Score DESC LIMIT 1 OFFSET ?", rankVal-1)

	// Scan the retrieved row data into the player instance
	err = row.Scan(&player.ID, &player.Name, &player.Country, &player.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve player by rank from database"})
		return
	}

	// Return the retrieved player data in JSON format
	c.IndentedJSON(http.StatusOK, player)
}

// GetRandomPlayer fetches a random player using a GET request.
func GetRandomPlayer(c *gin.Context) {

	var player model.Player

	// Query the database to retrieve a random player
	row := mysql.DB.QueryRow("SELECT * FROM players ORDER BY RAND() LIMIT 1")

	// Scan the retrieved row data into the player instance
	err := row.Scan(&player.ID, &player.Name, &player.Country, &player.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No players found"})
			return
		}

		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve random player from database"})
		return
	}

	// Return a random player data in JSON format
	c.JSON(http.StatusOK, player)
}
