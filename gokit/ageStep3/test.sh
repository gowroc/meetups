echo "POST: {\"yearOfBirth\": 1960} -> http://localhost:8003/age"
echo ""

curl -H "Content-Type: application/json" -X POST -d '{"yearOfBirth": 1960}' http://localhost:8003/age
