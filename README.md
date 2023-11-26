# SnappFood Challenge

A simple delivery system

# Quick Start

To initialize and run the application first run command below:

```
make up
```

alternatively you can use:

```
docker-compose up -d
```

Then run the migrations with command below:

```
make migrate
```

Then you can seed dummy data with command below:

```
make seed
```

Finally, You can access the webservice in port 8001

# API Reference


#### Delay
To report a delay on an order
```http  
 POST /api/v1/delay/{orderId}
 ```  
| Parameter | Type     | Description                |  
| :-------- | :------- | :------------------------- |  
| `orderId` | `int` | **Required**. (example:1) |  


**success response:**
```json
{
  "Status": 200,
  "Message": "Sorry for the delay, Our agents will review your report as soon as possible!",
  "Data": null
}
```

#### Assign
To assign a delay report to an agent
```http  
 POST /api/v1/assign
 ```  
| Header | Type     | Description                |  
| :-------- | :------- | :------------------------- |  
| `agent-id` | `int` | **Required**. (example:1) |  

> Note that for simplicity we ignored the authentication and simulated that with a `agent-id` in header

**success response:**
```json
{
  "Status": 200,
  "Message": "Delivery report 12 assigned to you",
  "Data": null
}
```

#### VendorReport
To get a list of vendors with their reports in past week
```http  
 POST /api/v1/vendor/report
 ```  
| Parameter | Type     | Description                |  
| :-------- | :------- | :------------------------- |  
|  | | |  


**success response:**
```json
{
  "Status": 200,
  "Message": "",
  "Data": [
    {
      "ID": 1,
      "Name": "First Vendor",
      "DelayCount": 9
    },
    {
      "ID": 3,
      "Name": "Third Vendor",
      "DelayCount": 3
    }
  ]
}
```
