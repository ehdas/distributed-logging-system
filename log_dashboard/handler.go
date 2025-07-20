package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"net/http"
)

func logList(c *gin.Context) {

	db, err := sql.Open("postgres", "postgresql://root:secret@postgres:5432/grpc-kafka?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT service_name, level, message, timestamp FROM logs ORDER BY timestamp DESC LIMIT 100")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var logs []map[string]interface{}
	for rows.Next() {
		var service, level, message string
		var ts string
		rows.Scan(&service, &level, &message, &ts)
		logs = append(logs, gin.H{
			"service":   service,
			"level":     level,
			"message":   message,
			"timestamp": ts,
		})
	}

	c.JSON(http.StatusOK, logs)
}
