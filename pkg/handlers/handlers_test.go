package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/moaabb/product-api/pkg/db"
	"github.com/moaabb/product-api/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetProducts(t *testing.T) {
	mockDB := setupMockDB()
	seedTestProducts(mockDB.DB)

	ph := NewProductHandler(mockDB)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ph.GetProducts(ctx)

	assert.Equal(t, http.StatusOK, w.Code)

	var products []models.Product

	err := json.Unmarshal(w.Body.Bytes(), &products)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(testProducts), len(products))
	for i, p := range products {
		assert.Equal(t, testProducts[i].Name, p.Name)
		assert.Equal(t, testProducts[i].Price, p.Price)
	}
}

func TestXxx(t *testing.T) {
	mockDB := setupMockDB()

	ph := NewProductHandler(mockDB)

	router := gin.Default()
	router.POST("/products", ph.CreateProduct)

	body := []byte(`{"name": "Test Product", "sku": "TP1", "stock": 10}`)
	req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	var createdProduct models.Product
	err = json.Unmarshal(resp.Body.Bytes(), &createdProduct)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}
	assert.Equal(t, "Test Product", createdProduct.Name)
	assert.Equal(t, "TP1", createdProduct.SKU)
	assert.Equal(t, 10, createdProduct.Stock)

	// Check that the product was created in the database
	var retrievedProduct models.Product
	result := mockDB.DB.First(&retrievedProduct, createdProduct.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Test Product", retrievedProduct.Name)
	assert.Equal(t, "TP1", retrievedProduct.SKU)
	assert.Equal(t, 10, retrievedProduct.Stock)

}

func setupMockDB() db.Handler {
	// Create an in-memory SQLite database
	conn, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the database schema (optional)
	err = conn.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to migrate database schema")
	}

	return db.Handler{DB: conn}
}

var testProducts []models.Product

func seedTestProducts(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		price, _ := faker.RandomInt(1, 100)
		product := models.Product{Name: faker.Word(), Price: price[0]}
		db.Create(&product)
		testProducts = append(testProducts, product)
	}
}
