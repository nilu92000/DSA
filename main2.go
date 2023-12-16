type Payload struct {
	ToSort [][]int `json:"to_sort"`
}
{
	"to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
  }
  curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]}' http://127.0.0.1:8000/process-single
  curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]}' http://127.0.0.1:8000/process-concurrent
  