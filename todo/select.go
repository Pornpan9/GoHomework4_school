package todo

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/Pornpan9/test4/database"
	"strconv"
)

func (todo Todo) GetAll(conn *sql.DB) ([]Todo, error){
	tt := []Todo{}
	query := "SELECT id,  title, status FROM todos"

	rows, err := conn.Query(query)
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Status); err != nil{
			return nil, err			
		}
		tt = append(tt, t)
	}
	return tt, err
}

func (todo Todo) GetByID(conn *sql.DB) (Todo, error){

	query := `
		SELECT 	id, title, status 
		FROM 	todos 
		where 	id = $1;
	`
	row := conn.QueryRow(query, todo.ID)

	var t Todo
	if err := row.Scan(&t.ID, &t.Title, &t.Status); err != nil{
		return t, err			
	}

	return t, nil
}

func GetHandler(c *gin.Context)  {

	conn, err := database.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()

	t := Todo{}
	tt, err := t.GetAll(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, tt)
}

func GetByIDHandler(c *gin.Context)  {

	t := Todo{}
	
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

	t, err = t.GetByID(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)

}