# receipt-processor
A small service to calculate the points on a receipt, done for a coding challenge.

## Usage
To use this receipt-processor service, ensure you have Go installed, and  follow the instructions below:

1. **Start the Server:**
   - Run the server by executing the `main.go` file. This will initialize the server and start it on port 8080.
        ```
        go run main.go
        ```

2. **Submit a Receipt:**
   - Use a `POST` request to submit a receipt for processing. The endpoint is:
        ```
        /receipts/process
        ```
   - The request body should contain a JSON object representing the receipt with the following structure:
     ```json
     {
       "retailer": "string",
       "purchaseDate": "YYYY-MM-DD",
       "purchaseTime": "HH:MM",
       "items": [
         {
           "shortDescription": "string",
           "price": "string"
         }
       ],
       "total": "string"
     }
     ```
   - Example using `curl`:
     ```
     curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '
     {
       "retailer": "Costco",
       "purchaseDate": "2020-01-01",
       "purchaseTime": "10:00",
       "items": [
         {
           "shortDescription": "item1",
           "price": "10.00"
         }
       ],
       "total": "10.00"
     }
     '
     ```

3. **Retrieve Receipt Points:**
   - Use a `GET` request to obtain the points for a submitted receipt. The endpoint is below, replace `{id}` with the UUID returned from the submit receipt endpoint. 
        ```
        /receipts/{id}/points
        ```
    - Example using `curl`:
        ```
        curl http://localhost:8080/receipts/{id}/points
        ```

## Considerations
- This Project has no database, so receipts are stored in memory, thus they are not persisted across application runs.
- I'm not super familiar with Go in general so some parts of the code may not be idiomatic Go.
- Since this is a quick single-person project I did not make use of source control the way I would in an enterprise environment.
- It was a fun challenge that served as a way for me to learn Go!