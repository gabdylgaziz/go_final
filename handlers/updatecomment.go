package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a comment...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var comment models.Comment
	json.Unmarshal(body, &comment)

	if result := h.DB.Model(models.Comment{}).Where("id = ?", comment.Id).Update("text", comment.Text); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Updated")

	fmt.Println("comment has been updated")
}
