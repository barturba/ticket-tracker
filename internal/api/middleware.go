package api

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository"
	"github.com/barturba/ticket-tracker/internal/utils/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userId,omitempty"`
	// Issuer   string `json:"iss,omitempty"`
	// UserRole string `json:"userRole,omitempty"`
	// Name     string `json:"name,omitempty"`
	// Email    string `json:"email,omitempty"`
	// Picture  string `json:"picture,omitempty"`
	// "sub": "0c1120ed-daed-428b-b6fd-6cad6d57d720",
	// "userRole": "user",
	// "userId": "0c1120ed-daed-428b-b6fd-6cad6d57d720",
	// "iat": 1729796927,
	// "exp": 1732388927,
	// "jti": "bc84366a-11d5-4469-a3d7-cfc5471731db"
	// Name     string `json:"name"`
}

func Auth(logger *slog.Logger, db *database.Queries, cfg models.Config) Middleware {
	fmt.Printf("> Auth\n")
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add the "Vary: Authorization" header. This informs caches that the
			// response may vary based on the contents of the Authorization header.
			w.Header().Add("Vary", "Authorization")

			res, err := httputil.DumpRequest(r, true)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print("RAW HTTP:")
			fmt.Print(string(res))

			// Retrieve the value of the Authorization header.
			authorizationHeader := r.Header.Get("Authorization")

			// If the Authorization header is blank, give the request context an
			// anonymous user.
			if authorizationHeader == "" {
				r = ContextSetUser(r, models.AnonymousUser)
				next.ServeHTTP(w, r)
				return
			}

			headerParts := strings.Split(authorizationHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				errors.InvalidAuthenticationTokenResponse(w, r, logger)
				return
			}

			// Get the actual token
			tokenString := headerParts[1]
			claims, err := ValidateToken(logger, cfg, tokenString)
			if err != nil {
				logger.Info("validateToken error", "err", err)
				errors.InvalidAuthenticationTokenResponse(w, r, logger)
				return
			}

			// Parse the id
			id, err := uuid.Parse(claims.UserID)
			if err != nil {
				errors.InvalidAuthenticationTokenResponse(w, r, logger)
				return
			}

			// Retrieve the details of the user associated with the authentication
			// token, again calling the InvalidAuthenticationTokenResponse() helper
			// if no record was found.
			user, err := repository.GetUser(logger, db, r.Context(), id)
			if err != nil {
				errors.ServerErrorResponse(w, r, logger, err)
				return
			}

			// Call the ContextSetUser() helper to add the user to the context.
			r = ContextSetUser(r, &user)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}

// ValidateToken parses and validates the JWT token
func ValidateToken(logger *slog.Logger, cfg models.Config, tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			return nil, fmt.Errorf("token expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

// func parseJWTClaims(logger *slog.Logger, tokenString string) error {
// 	// Parse the token without validating the signature
// 	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
// 	if err != nil {
// 		return fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	// Get the claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return fmt.Errorf("failed to get claims from token")
// 	}

// 	// Loop through all claims
// 	for key, value := range claims {
// 		switch v := value.(type) {
// 		case string:
// 			logger.Info("Claim", "key", key, "type", "(string)", "value", v)
// 		case float64:
// 			logger.Info("Claim", "key", key, "type", "(float64)", "value", v)
// 		case bool:
// 			logger.Info("Claim", "key", key, "type", "(bool)", "value", v)
// 		case []interface{}:
// 			logger.Info("Claim", "key", key, "type", "(array)", "value", v)
// 		case map[string]interface{}:
// 			logger.Info("Claim", "key", key, "type", "(object)", "value", v)
// 		default:
// 			logger.Info("Claim", "key", key, "type", "(unknown type)", "value", v)
// 		}
// 	}

// 	return nil
// }

func RequireActiveUser(logger *slog.Logger, db *database.Queries, cfg models.Config) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use the ContextGetUser() helper to retrieve the user information.
			user := ContextGetUser(r)

			if models.IsAnonymous(user) {
				errors.AuthenticationRequiredResponse(w, r, logger)
				return
			}

			if !user.Active {
				errors.InactiveAccountResponse(w, r, logger)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Log each request to an id
func WithRequestID() Middleware {
	fmt.Printf("> WithRequestID\n")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if there's already an id
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			// Set the request id header
			w.Header().Set("X-Request-ID", requestID)

			// Set the request id in the context
			ctx := ContextSetRequestId(r, requestID)
			next.ServeHTTP(w, ctx)
		})
	}
}

// Custom ResponseWriter to capture status code
type ResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w}
}

func (rw *ResponseWriter) Status() int {
	return rw.status
}

func (rw *ResponseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

// Should use middleware for consistent logging
func Logger(logger *slog.Logger) Middleware {
	fmt.Printf("> Logger\n")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			requestID := GetRequestID(r.Context())

			// Wrap ResponseWriter to capture status code
			ww := NewResponseWriter(w)

			// Add request ID to all log entries
			loggerWithRequestID := logger.With(
				"request_id", requestID,
				"method", r.Method,
				"path", r.URL.Path,
			)

			// Add logger to context
			ctx := context.WithValue(r.Context(), LoggerContextKey, loggerWithRequestID)
			r = r.WithContext(ctx)

			// Process request
			next.ServeHTTP(ww, r)

			// Log completion
			loggerWithRequestID.Info("request completed",
				"status", ww.Status(),
				"duration", time.Since(start),
			)

		})
	}
}
