package auth

import (
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"

	"github.com/medvedev-v/radiocontest/internal/model"
)

// UserRepository определяет методы для работы с пользователями, необходимые middleware.
type UserRepository interface {
    GetByID(id int64) (*model.User, error)
}

// SessionRepository определяет методы для работы с сессиями, необходимые middleware.
type SessionRepository interface {
    GetActiveByToken(token string) (*model.Session, error)
}

// AuthMiddleware возвращает Gin-обработчик, проверяющий валидность токена.
func AuthMiddleware(userRepo UserRepository, sessionRepo SessionRepository) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Извлечь токен из заголовка Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
            return
        }

        // Ожидаем формат "Bearer <token>"
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
            return
        }
        token := parts[1]

        // 2. Найти активную сессию по токену
        session, err := sessionRepo.GetActiveByToken(token)
        if err != nil {
            // В реальности можно логировать ошибку, но клиенту говорим только 401
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
            return
        }

        // 3. Проверить, не истекла ли сессия (дополнительная проверка, хотя репозиторий уже должен фильтровать)
        if session.ExpiresAt.Before(time.Now()) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
            return
        }

        // 4. Получить пользователя по user_id из сессии
        user, err := userRepo.GetByID(session.UserID)
        if err != nil {
            // Если пользователь не найден (возможно, удалён), сессия невалидна
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
            return
        }

        // 5. Сохранить пользователя и сессию в контексте для последующих хендлеров
        c.Set("user", user)
        c.Set("session", session)

        // 6. Продолжить обработку
        c.Next()
    }
}

// GetUserFromContext возвращает пользователя из контекста Gin.
// Используйте в хендлерах после применения AuthMiddleware.
func GetUserFromContext(c *gin.Context) (*model.User, bool) {
    user, exists := c.Get("user")
    if !exists {
        return nil, false
    }
    u, ok := user.(*model.User)
    return u, ok
}

// GetSessionFromContext возвращает сессию из контекста Gin.
func GetSessionFromContext(c *gin.Context) (*model.Session, bool) {
    sess, exists := c.Get("session")
    if !exists {
        return nil, false
    }
    s, ok := sess.(*model.Session)
    return s, ok
}
