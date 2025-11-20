package service

import (
	"context"
	"errors"
	"log"
	"time"

	"backend/internal/domain"
	"backend/internal/repository"

	"github.com/jackc/pgx/v5"
)

type HueGetService struct {
	hueRepo     *repository.HueRepository
	sessionRepo *repository.LoginSessionRepository
	logger      *log.Logger
}

func NewHueGetService(hueRepo *repository.HueRepository, sessionRepo *repository.LoginSessionRepository, logger *log.Logger) *HueGetService {
	if logger == nil {
		logger = log.Default()
	}
	return &HueGetService{hueRepo: hueRepo, sessionRepo: sessionRepo, logger: logger}
}

func (s *HueGetService) GetData(ctx context.Context, session domain.SessionData, recordRange domain.RecordRange) ([]domain.HueRecord, error) {
	loginSession, err := s.sessionRepo.Find(ctx, session.UserID(), session.Token())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logError("session not found", err)
			return nil, domain.ErrInvalidLoginSession
		}
		s.logError("find session", err)
		return nil, err
	}

	if loginSession.IsExpired(time.Now()) {
		s.logError("session expired", domain.ErrExpiredToken)
		if delErr := s.sessionRepo.DeleteByID(ctx, loginSession.ID()); delErr != nil {
			s.logError("cleanup expired session", delErr)
		}
		return nil, domain.ErrExpiredToken
	}

	records, err := s.hueRepo.FindRange(ctx, recordRange)
	if err != nil {
		s.logError("fetch hue records", err)
		return nil, err
	}

	return records, nil
}

func (s *HueGetService) logError(action string, err error) {
	if err == nil {
		return
	}
	s.logger.Printf("[HueGetService] %s: %v", action, err)
}
