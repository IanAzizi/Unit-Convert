# 🌐 Unit Converter (Go + Gin)

A simple **unit conversion web app** built with Go and Gin framework.  
This project allows users to convert between various units of **length**, **weight**, and **temperature** — quickly and accurately.

> 🔗 Project source on roadmap.sh: [Unit Converter](https://roadmap.sh/projects/unit-converter)

---

## 🧩 Features

✅ Convert between multiple measurement units:  
- **Length:** millimeter, centimeter, meter, kilometer, inch, foot, yard, mile  
- **Weight:** milligram, gram, kilogram, ounce, pound  
- **Temperature:** Celsius, Fahrenheit, Kelvin  

✅ Simple RESTful API endpoint for conversion  
✅ Modular code structure (`service`, `handler`, `router`)  
✅ JSON-based request/response  
✅ Easily extendable for new unit types  

---

## 🚀 API Example

### **POST** `/api/convert`

**Request Body:**
```json
{
  "from": "Gram",
  "to": "Kilogram",
  "value": 2000
}
responce

{
  "from": "Gram",
  "to": "Kilogram",
  "value": 2000,
  "result": 2
}

Project Structure
unit-convert/
│
├── api/
│   ├── handler/
│   │   └── unit_handler.go
│   └── router/
│       └── router.go
│
├── service/
│   └── unit.go
│
├── go.mod
├── go.sum
├── main.go
└── README.md


Setup & Run
git clone https://github.com/YOUR_USERNAME/Unit-Convert.git
cd Unit-Convert

install depandacy
go mod tidy

run server 
go run main.go


test api
curl -X POST http://localhost:8080/api/convert \
-H "Content-Type: application/json" \
-d '{"from":"Meter","to":"Kilometer","value":1000}'
