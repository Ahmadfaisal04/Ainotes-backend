package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var supabaseJWKSUrl = "https://asivvleysqlwsdovqboi.supabase.co/auth/v1/keys"

var cachedPublicKeys map[string]string

// fetchSupabasePublicKeys fetches the JWKS and caches the keys
func fetchSupabasePublicKeys() (map[string]string, error) {
	if cachedPublicKeys != nil {
		return cachedPublicKeys, nil
	}
	resp, err := http.Get(supabaseJWKSUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jwks struct {
		Keys []struct {
			Kid string   `json:"kid"`
			Alg string   `json:"alg"`
			Kty string   `json:"kty"`
			N   string   `json:"n"`
			E   string   `json:"e"`
			Use string   `json:"use"`
			X5c []string `json:"x5c"`
		} `json:"keys"`
	}
	if err := json.Unmarshal(body, &jwks); err != nil {
		return nil, err
	}
	keys := make(map[string]string)
	for _, key := range jwks.Keys {
		if len(key.X5c) > 0 {
			keys[key.Kid] = "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}
	cachedPublicKeys = keys
	return keys, nil
}

// SupabaseAuthMiddleware verifies Supabase JWT
func SupabaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		keys, err := fetchSupabasePublicKeys()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch public keys"})
			return
		}
		// Parse token without verifying to get kid
		token, err := jwt.Parse(tokenString, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token header"})
			return
		}
		cert, ok := keys[kid]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unknown key id"})
			return
		}
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		})
		if err != nil || !parsedToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("supabase_claims", parsedToken.Claims)
		c.Next()
	}
}
