package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// JWTService gerencia a criação e validação de tokens JWT
type JWTService struct {
	secretKey []byte
	issuer    string
}

// JWTClaims representa as claims de um token JWT
type JWTClaims struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	ExpiresAt int64 `json:"exp"`
	IssuedAt int64 `json:"iat"`
	Issuer string `json:"iss"`
}

// NewJWTService cria uma nova instância do serviço JWT
func NewJWTService(secretKey, issuer string) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
		issuer:    issuer,
	}
}

// GenerateToken cria um novo token JWT para o usuário
func (s *JWTService) GenerateToken(userID, username string, expirationTime time.Duration) (string, error) {
	now := time.Now()
	expires := now.Add(expirationTime).Unix()
	
	// Criar claims
	claims := JWTClaims{
		UserID:    userID,
		Username:  username,
		ExpiresAt: expires,
		IssuedAt:  now.Unix(),
		Issuer:    s.issuer,
	}
	
	// Criar header
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	
	// Codificar header
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	
	// Codificar payload
	payloadJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadJSON)
	
	// Criar assinatura
	signatureInput := headerEncoded + "." + payloadEncoded
	h := hmac.New(sha256.New, s.secretKey)
	h.Write([]byte(signatureInput))
	signature := h.Sum(nil)
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	
	// Montar token
	token := headerEncoded + "." + payloadEncoded + "." + signatureEncoded
	
	return token, nil
}

// ValidateToken verifica se um token JWT é válido
func (s *JWTService) ValidateToken(tokenString string) (*JWTClaims, error) {
	// Dividir token em partes
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("token inválido: formato incorreto")
	}
	
	// Verificar assinatura
	signatureInput := parts[0] + "." + parts[1]
	h := hmac.New(sha256.New, s.secretKey)
	h.Write([]byte(signatureInput))
	expectedSignature := h.Sum(nil)
	
	// Decodificar assinatura do token
	tokenSignature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, fmt.Errorf("token inválido: assinatura mal formatada: %w", err)
	}
	
	// Comparar assinaturas
	if !hmac.Equal(tokenSignature, expectedSignature) {
		return nil, errors.New("token inválido: assinatura inválida")
	}
	
	// Decodificar payload
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("token inválido: payload mal formatado: %w", err)
	}
	
	// Extrair claims
	var claims JWTClaims
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return nil, fmt.Errorf("token inválido: claims mal formatadas: %w", err)
	}
	
	// Verificar expiração
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expirado")
	}
	
	return &claims, nil
}