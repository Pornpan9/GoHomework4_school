package todo

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/Pornpan9/test4/database"
	"strconv"
)

func (todo Todo) Update(conn *sql.DB) error{

	query := `
		UPDATE todos SET title=$2, status=$3 WHERE id=$1;
	`;

	stmt, err := conn.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.ID, todo.Title, todo.Status)
	
	return err
}

func UpdateHandler(c *gin.Context)  {
	
	t := Todo{}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))//get param on url
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	t.ID = id

	conn, err := database.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()

	err = t.Update(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}